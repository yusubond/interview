## 记录

- 快速排序
- 按照要求打印堆栈信息，程序文件在[format_gdblog.go](code/format_gdblog.go)

```bash
输入为文件，其文件内容和格式为
Thread:
#0  0x0000000000c8dc41 in func7
#1  0x00000000004674c5 in func8
#2  0x00000000004585d8 in func9
#3  0x000000000046bac4 in a_very_long_function_name_across_lines_10
#4  0x000000000046ba51 in another_very_long_function_name_across_more_lines_11
#5  0x000000000046ba02 in func12


Thread:
#0  0x0000000000c8dc41 in func13
#1  0x00000000004674c5 in func14
#2  0x00000000004585d8 in func15
#3  0x000000000046ba02 in func16

要求输出为：
func7 <- func8 <- func9 <- a_very_long_function_name_across_lines_10 <- another_very_long_function_name_across_more_lines_11 <- func12 
func13 <- func14 <- func15 <- func16
```

- 动态规划

给定一个int数组A，数组中元素互不重复，给定一个数x，求所有求和能得到x的数字组合，组合中的元素来自A，可重复使用 

- 非递归层序遍历二叉树，3月5日 腾讯



## 问题2 

> 给定一个int数组A，数组中元素互不重复，给定一个数x，求所有求和能得到x的数字组合，组合中的元素来自A，可重复使用

解析：

这道题是 leetcode 第 39题，组合总和，可用回溯算法解决。

原题目是：

> 给你一个 **无重复元素** 的整数数组 `candidates` 和一个目标整数 `target` ，找出 `candidates` 中可以使数字和为目标数 `target` 的 所有 **不同组合** ，并以列表形式返回。你可以按 **任意顺序** 返回这些组合。
>
> `candidates` 中的 **同一个** 数字可以 **无限制重复被选取** 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
>
> 对于给定的输入，保证和为 `target` 的不同组合数少于 `150` 个。

解题算法如下：

```go
```

