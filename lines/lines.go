package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var (
	LINE      = "#number"
	file      = ""
	format    = "%d"
	prefix    = ""
	skipEmpty = false
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
		if file == "" {
			cmd.Help()
			return
		}
		file, err := os.OpenFile(file, os.O_RDWR, 0666)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Open file error!%s\n", err.Error()))
			return
		}
		defer file.Close()
		if err != nil {
			panic(err)
		}
		buf := bufio.NewReader(file)
		var index = 0
		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("Read file error!", err)
					return
				}
			}
			if strings.TrimSpace(line) == "" && skipEmpty {
				continue
			}
			number := fmt.Sprintf(format, index)
			fmt.Printf("%s%s", strings.ReplaceAll(prefix, LINE, number), line)
			index++
		}
	},
}

func main() {
	root.Flags().StringVarP(&file, "file", "F", "", "[*]目标文件")
	root.Flags().StringVarP(&prefix, "prefix", "P", "#number", "行前置格式化内容，其中#number为占位符标示行号, \n"+
		"例如\"#number|\"表示使用下划线分割行号和正文")
	root.Flags().StringVarP(&format, "number", "N", "%d", "行号的格式化风格，例如%4d，则格式化为4位，%-4d则4位居左")
	root.Flags().BoolVarP(&skipEmpty, "skipEmpty", "S", false, "是否跳过空白行，例如true表示跳过空白行")
	root.Execute()
}
