package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const (
	// LINE 占位符
	LINE = "#number"
)

var (
	// 数据源
	filePath = ""
)

var root = &cobra.Command{
	Short: "针对文本的逐行读取小工具",
	Long: `
针对文件内文本的逐行读取及简易编辑操作
1. 编辑行前缀（行号占位符）
2. 编辑行后缀（暂未实现）
3. 支持跳过空白行（空格、tab等空白字符）
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	initCmdTrim()
	initSkipEmpty()
	initPrefix()

	root.AddCommand(CmdPrefix)    // 前缀
	root.AddCommand(CmdSuffix)    // 后缀
	root.AddCommand(CmdTrim)      // 去除先后的空白字符
	root.AddCommand(CmdTrimLeft)  // 去除左侧的空白字符
	root.AddCommand(CmdTrimRight) // 去除右侧的空白字符
	root.AddCommand(CmdSkipEmpty) // 跳过空白字符行
	// 数据源
	root.Execute()
}

func FileAction(cmd *cobra.Command, action func(*bufio.Reader)) {
	var reader *bufio.Reader
	if filePath == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Open file error!%s\n", err.Error()))
			return
		}
		defer file.Close()
		if err != nil {
			panic(err)
		}
		reader = bufio.NewReader(file)
	}
	action(reader)
}

func LineAction(cmd *cobra.Command, action func(line string)) {
	FileAction(cmd, func(reader *bufio.Reader) {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					action(line)
					break
				} else {
					fmt.Println("Read file error!", err)
					return
				}
			}
			action(line)
		}
	})
}
