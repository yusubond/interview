# Golang

> 1、Map的底层实现？



> 2、Map的扩容是怎么实现的？

使用哈希表的目的就是要快速查找目标key，然而，随着向map中添加的key越来越多，key发生碰撞的概率也越来越大。bucket中的8个cell会被越塞越满，查找，插入，删除key的效率也会越来越低。

最理想的情况是一个bucket只装一个key，这样就能达到O(1)的效率，但这样空间消耗太大，用空间换时间的代价太高。

Go语言采用一个bucket里装载8个key，定位到某个bucket后，还需要再定位到具体的key，这实际上又用了时间换空间。

当然这样做要有一个度，不然所有的key都落在了同一个bucket里，直接退化成了链表，各种操作的效率直接降为O(n)，是不行的。

因此，通过一个指标来衡量前面描述的情况，这就是装载因子。Go源码里这样定义装载因子：

```
loadFactor := count / (2^B)
```

count就是map的元素个数，2^B表示bucket数量。

参考资料：

map扩容机制：https://github.com/yusubond/Crius/blob/master/std/73e254a888a9fbebaaed6875fc1c1a3f.md



> 3、为什么遍历MAP是无序的？

1）bucket序号会发生变化，当扩容(B+1)发生后，rehase bucket序号会变化

2）每次遍历并不是从0号bucket开始，而是随机值序号，而且从这个bucket的随机序号的cell开始



> 4、goroutine是怎么工作的？

goroutine 可以看做对 thread的抽象，更加轻量级，可单独执行。

goroutine 和 thread 的区别，可以从三个角度分析，分别是内存消耗，创建和销毁，切换。

**内存消耗**：创建一个 goroutine 仅需 2KB 内存，如果栈空间不够用，会自动扩容；创建一个 thread 需要 1MB 内存，还需要一个称为 "a guard page" 的区域用于和其他线程栈空间的隔离。

**创建和销毁**：thread 是内核级，跟 OS 打交道，创建，销毁有巨大消耗；goroutine 用户级，由 goroutine 管理，创建，销毁消耗小

**切换**：thread 切换需要保存各种寄存器，切换消耗1000-1500纳秒，1 纳秒平均12-18条指令；goroutine 切换只保存三个寄存器，Program Counter，Stack Pointer 和 BP (psesudo SP)



> 5、什么是Golang的M:N模型

go runtime负责 goroutine 的整个生命周期。Runtime在启动时，创建 M 个线程，然后创建 N 个 goroutine，这些 goroutine 都会依附在 M 个线程上执行，这就是 M:N 模型。

同一时刻，一个线程只能运行一个 goroutine；当某个 goroutine 发生阻塞时，run time 会把当前 goroutine 调度走，让其他 goroutine来执行，不让一个线程闲着。



> 6、Golang GMP调度模型

前置知识：

OS调度

线程是指"按顺序执行指令"，也是 CPU 调度的实体。线程有三个状态，`Waiting`, `Runnable` 和 `Executing`。

线程能做的事情分两种，一种是计算型，占用CPU资源，一种是 IO 型，获取外部资源。

线程切换就是操作系统用一个 `Runnable` 的线程，将 CPU 上正在运行的处于 `Executing` 状态的线程换下来的过程。换下来的线程变为 `Runnable` (计算型) 或者 `Waiting` (IO型)。

对于OS调度而言，最重要的是不要让一个 CPU 核心闲着，尽量让每个 CPU 核心都有任务可做。

