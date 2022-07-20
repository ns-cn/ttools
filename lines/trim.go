package main

import (
	"fmt"
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
	LineAction(cmd, func(line string) {
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
		fmt.Print(line)
	})
}

func initCmdTrim() {
	CmdTrim.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdTrimLeft.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdTrimRight.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
}
