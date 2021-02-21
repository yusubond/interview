# Tofino可编程芯片

1、Parsers解析器的原理

P4解析器是基于有限状态机的思想来设计的。

P4语言中的Parser是指**数据包到达Switch的时候，并不能马上进行Match-Action匹配，需要 parser解析器进行处理加工成MA单元能够匹配的程序。**

解析器中解析的过程可以被一个解析图(parser graph)所表示，解析图中所表示的某一个状态(或者说，在P4语言中的某一个解析函数)看做是一个状态节点，每一个状态转换等同于跨越状态节点之间的边界。

解析器根据数据包的第一个字节进行调控，它保存了一个指向该数据包包头中特殊单位字节的指针(current offset)。

当解析器开始对首部实例进行提取操作时，它根据作为extract函数参数的首部实例的格式进行提取，将数据包的数据更新到该首部实例中，同时更新该数据包的解析表示。在提取操作结束之后，解析器保存的指针(current offset)移向下一个要处理的位置，等同于进行一次状态转移。

参考资料：

https://www.sdnlab.com/18021.html

2、mirror的底层原理



3、meter的如何实现的

rfc 2698

https://blog.csdn.net/changchangaaa/article/details/108165170



4、PRE(Packet Replication Engine)的工作内容

PRE称为数据包复制引擎，也是TM(Traffic Manager)的一部分。主要负责多播，copy_to_cpu。



5、Mirror的工作原理

Mirror是指创建一个新的，完整的数据包的副本，**原始报文和mirrored的报文副本是被独立处理的**。P4支持两种类型的Mirror:

- Ingress Mirror

  Ingress mirror发生在ingress deparse中，当执行`pkt.emit()`的时候即发生mirror。

- Engress Mirror

  Engress mirror发生在engress deparse中，当执行`pkt.emit()`的时候即发生mirror。

所以，**mirror实际上是创建了一个新的报文副本给到PRE，并且不计入数据包的统计中，当高负载时，mirror报文会被第一时间丢弃掉**。



6、Resubmit & Recirculate

Resubmit是指将Ingress Deparser的数据包重新提交到Ingress Parser部分，并可以携带额外的metadata。

报文仅可被resubmit一次。

Recirculate是指将数据包重定向到同一个(或者不同Pipeline)的端口，怎么理解呢？

就是同一个Pepeline内Recirculate通过CPU68口完成回环；不同的Pepeline可指定端口，例如

```bash
action recirculate_odd_pipeline() {
	ig_tm_md.ucast_egress_port = ig_intr_md.ingress_port | 0x80;
}
```



