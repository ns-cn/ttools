# ttools
tiny tools

# lines
用于单行的处理，目前只支持行首的处理
1. 支持```-P```(或```--prefix```)为每行增加前缀，支持#number占位符标识行号
2. 支持```-N```(或```--number```)格式化行号，例如%4d为4位长度
3. 支持```-S```(或```--skipEmpty```)跳过空白行

#### 使用实例
```shell
# 处理mine.txt文件，为行首增加居左显示3位宽度的行号，并使用|分割行号和正文
> lines.exe -F mine.txt -P "#number|" -N "%-3d" -S false
# 重定向输出到target.txt，方便后续编辑读取等操作
> lines.exe -F mine.txt -P "#number|" -N "%-3d" -S false >> target.txt
```

#### 后续支持计划
- [ ] 支持从管道读取
- [ ] 支持行尾编辑操作
- [ ] 支持应用内级别的文件写入
- [ ] 支持多文件同时处理并写入文件（支持文件位置指定）
