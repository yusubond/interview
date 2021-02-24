# Golang

1、Map的底层实现？



2、Map的扩容是怎么实现的？

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



