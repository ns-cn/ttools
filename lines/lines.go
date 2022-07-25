package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

const (
	// LINE 占位符
	LINE    = "#number"
	VERSION = "1.02"
)

var (
	// 数据源
	filePath      = ""
	fromClipboard = false // 是否从粘贴板读取字符数据，优先级第一
	toClipboard   = false // 是否将结果写入到粘贴板而不是命令行标准输出
)

var root = &cobra.Command{
	Short: "针对文本的逐行读取小工具",
	Long: `
针对文件内文本的逐行读取及简易编辑操作
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
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
	root.AddCommand(CmdVersion)   // 打印版本号
	// 数据源
	_ = root.Execute()
}

func InputAction(cmd *cobra.Command, action func(*bufio.Reader)) {
	var reader *bufio.Reader
	if fromClipboard {
		clipboardString, err := clipboard.ReadAll()
		if err != nil {
			handleErrWithTips("读取粘贴板失败", err)
			return
		}
		reader = bufio.NewReader(strings.NewReader(clipboardString))
	} else if filePath == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			handleErrWithTips("open file error!", err)
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

/*LineAction
行处理函数，需要传递单行文本处理函数供统筹处理
*/
func LineAction(cmd *cobra.Command, action func(line string) string) {
	InputAction(cmd, func(reader *bufio.Reader) {
		result := strings.Builder{}
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					actionResult := action(line)
					result.WriteString(actionResult)
					break
				} else {
					handleErrWithTips("Read file error!", err)
					return
				}
			}
			actionResult := action(line)
			result.WriteString(actionResult)
		}
		if toClipboard {
			err := clipboard.WriteAll(result.String())
			if err != nil {
				_, _ = os.Stderr.WriteString(err.Error())
			}
		} else {
			fmt.Println(result.String())
		}
	})
}
