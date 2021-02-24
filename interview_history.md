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

