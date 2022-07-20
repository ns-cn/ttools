# ttools
tiny tools: 小工具集合

# lines
用于文本文件的逐行的处理，支持增加前缀、后缀、去除空白字符、跳过空白行

#### 使用实例
```shell
# 增加文件前缀: 实例为为mine.txt中非空白行增加行号显示并显示原始文件行号
> lines.exe prefix -F mine.txt -c "第#number行. " -n "%3d" -os
> lines.exe prefix --file mine.txt --content "第#number行. " --number "%3d" --skipEmpty --keepOriginal
# 增加文件后缀
> lines.exe suffix -F mine.txt -c "第#number行. " -n "%3d" -os
# 去除左侧的空白字符
> lines.exe trimleft -F mine.txt
# 去除右侧的空白字符
> lines.exe trimright -F mine.txt
# 去除左右两侧的空白字符
> lines.exe trim -F mine.txt
# 跳过空白行
> lines.exe skipempty -F mine.txt

# 管道连续操作
> lines.exe trim -F mine.txt | lines.exe skipempty | lines.exe prefix -c "第#number行. " >> output.txt
```

> 完整文档参见：```lines.exe --help```

#### 后续支持计划
- [x] 支持从管道读取
- [x] 支持行尾编辑操作
- [ ] 支持行内替换操作
- [ ] 支持指定行范围操作
- [ ] 支持应用内级别的文件写入
- [ ] 支持多文件同时处理并写入文件（支持文件位置指定）
- [ ] 简单GUI