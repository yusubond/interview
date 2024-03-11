## GO五个问题

[toc]

## 1、golang 中 make 和 new 的区别。

`new`函数的原型：

```go
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
 
func new(Type) *Type
```

new函数用来分配内存，第一个参数是一种类型，而不是一个值，返回值是一个指向该类型零值位置的指针。

也就是说 new 函数有两个作用：

- 分配内存，并不初始化内存，只是将其设置为该类型的零值。
- 返回一个指向该内存的指针。



`make`函数的原型是：

```go
//The make built-in function allocates and initializes an object
//of type slice, map, or chan (only). Like new, the first argument is
// a type, not a value. Unlike new, make's return type is the same as
// the type of its argument, not a pointer to it.
 
func make(t Type, size ...IntegerType) Type
```

`make`函数只能为`slice`,`map`和`chan`类型分配内存和初始化一个对象。

第一个参数是一种类型，而不是一个值。

返回值是就是传入的类型，而不是指针。

**因为这三种类型都是引用类型，所以没必要返回他们的指针。**

所以，`make`函数的作用有：

- 分配内存，并初始化一个类型对象。
- 返回该类型。



在编译期间的类型检查阶段，Go 语言会将`make`关键字中的`OMAKE`节点根据参数类型的不同转换为`OMAKESLICE`、`OMAKEMAP`和`OMAKECHAN`三种不同类型的节点，进而调用不同的运行时函数初始化相应的数据结构。



`new`函数创建的变量，如果变量会逃逸到堆上，就会在堆上申请内存；如果变量不需要在当前作用域外生存，那么就不需要初始化在堆上。



## 2、range 的时候地址会发生变化吗？

```go
for a, b := range c
```

在上面的代码中，a和b在内存里只会存一份，每次循环遍历的时候，都是以值覆盖的方式赋值为a和b，a和b 的内存地址不会变。

正是因为这个，如果要在 for 中创建协程，需要创建临时变量，不能直接将a和b传递给协程。



