package main

import (
	"github.com/spf13/cobra"
	"strings"
)

var CmdTrim = &cobra.Command{
	Use:     "trim",
	Aliases: []string{"t"},
	Short:   "逐行去除前后的空白字符（空格，tab等空白字符）",
	Run: func(cmd *cobra.Command, args []string) {
		TrimAction(cmd, true, true)
	},
}

var CmdTrimLeft = &cobra.Command{
	Use:     "trimleft",
	Aliases: []string{"tl"},
	Short:   "逐行去除左侧的空白字符（空格，tab等空白字符）",
	Run: func(cmd *cobra.Command, args []string) {
		TrimAction(cmd, true, false)
	},
}

var CmdTrimRight = &cobra.Command{
	Use:     "trimright",
	Aliases: []string{"tr"},
	Short:   "逐行去除右侧的空白字符（空格，tab等空白字符）",
	Run: func(cmd *cobra.Command, args []string) {
		TrimAction(cmd, false, true)
	},
}

func TrimAction(cmd *cobra.Command, trimStart bool, trimEnd bool) {
	LineAction(cmd, func(line string) string {
		if trimStart {
			line = strings.TrimLeftFunc(line, func(r rune) bool {
				return r == ' ' || r == '\t'
			})
		}
		if trimEnd {
			line = strings.TrimRightFunc(line, func(r rune) bool {
				return r == ' ' || r == '\t'
			})
		}
		return line
	})
}

func initCmdTrim() {
	// 去除
	CmdTrim.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdTrim.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源，如果为true，则不需要指定file参数")
	CmdTrim.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	// 去除左侧的
	CmdTrimLeft.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdTrimLeft.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源，如果为true，则不需要指定file参数")
	CmdTrimLeft.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	// 去除右侧
	CmdTrimRight.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdTrimRight.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源，如果为true，则不需要指定file参数")
	CmdTrimRight.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
}
