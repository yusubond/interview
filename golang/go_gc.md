# Go 垃圾回收（GC）

> Go 语言中 GC 机制演进过程中经历过哪些变化？分别是什么？

Go 语言 GC 机制主要经历过三次重大变革，分别是：

- go v1.3 之前的**标记清除（Mark and Sweep）法**
- Go v1.5 的**三色并发标记法**。"强弱三色不变式"、插入屏障、删除屏障。
- Go v1.8 的**混合写屏障机制**。



**v1.3 标记清除法**

标记清除算法的基本流程：先启动 STW，使程序暂停，执行标记，执行数据回收，停止 STW，恢复程序。

这个方法的缺点是整个 GC 过程被包裹在 STW 时间内，影响程序性能。

v1.3 做了简单的优化，将「停止 STW」的步骤提前到标记执行结束，来缩短 STW 时间对程序性能的影响。

但无论如何优化，STW 的时间对程序的影响都很大。



**v1.5 三色并发标记法**

为了解决原始标记清除算法带来的长时间 STW，多数现代的追踪式垃圾收集器都会实现三色标记算法的变种以缩短 STW 的时间。

三色标记算法将程序中的对象分为白色、黑色和灰色三类：

- 白色对象—潜在的垃圾，其内存可能会被垃圾收集器回收
- 黑色对象—活跃的对象，包括不存在任何引用外部指针的对象以及根对象可达的对象
- 灰色对象—活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象



**三色标记的标记阶段**



1. 初始状态下所有的对象都是白色的
2. 从根节点开始遍历，把遍历到的对象变成灰色对象【只遍历一层】
3. 遍历灰色对象，将灰色对象引用的对象变成灰色对象，并把遍历过的灰色对象标记为黑色对象【保证该对象和被该对象引用的对象都不被回收】
4. 重复步骤3，直到灰色对象全部变成黑色

当三色标记阶段结束之后，应用程序的堆中就不存在任何灰色对象，我们只能看到黑色的存活对象和白色的垃圾对象。垃圾收集器可以回收这些白色垃圾。



因为用户程序可能在标记执行的过程中修改对象指针，所以三色标记算法本身是不可并发或者增量执行的，它仍然需要 STW。想要并发或者增量地标记对象需要屏障技术。



**屏障技术**

内存屏障技术是一种屏障指令，它可以让 CPU 或者编译器在执行内存相关操作时遵循特定的约束，目前多数的现代处理器都会乱序执行指令以最大化性能，但该技术能够保证内存操作的顺序性，在内存屏障钱执行的操作一定会先于内存屏障后执行的操作。

想要并发或增量的标记算法中的正确性，我们需要达成以下两种三色不变性中的一种：

- 强三色不变性—黑色对象不会指向白色对象，只会指向灰色对象或黑色对象
- 弱三色不变性—黑色对象指向的白色对象必须包含一条从灰色对象经由多个白色对象的可达路径



垃圾收集器的屏障技术更像是一个钩子方法，它是在用户程序读取对象、创建对象以及更新对象指针时执行一段代码，根据操作类型的不同，我们可以将它们分成读屏障和写屏障两种。因为读屏障需要在读操作中加入一段代码，对用户程序的性能影响很大，所以编程语言往往都会采用写屏障保证三色不变性。

Go 语言中使用两种写屏障技术，分别是 Dijkstra 提出的插入写屏障和 Yussa 提出的删除写屏障。

**插入写屏障**

Dijkstra 在 1978 年提出了插入写屏障，通过如下所示的写屏障，用户程序和垃圾收集器可以在交替工作的情况下保证程序执行的正确性：

```go
writePointer(slot, ptr):
    shade(ptr)
    *slot = ptr
```

上述插入写屏障的伪代码非常好理解，每当执行类似 `*slot = ptr` 的表达式时，我们会执行上述写屏障通过 `shade` 函数尝试改变指针的颜色。如果 `ptr` 指针是白色的，那么该函数会将该对象设置成灰色，其他情况则保持不变。



**v1.8 混合写屏障机制**