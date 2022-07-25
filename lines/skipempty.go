package main

import (
	"github.com/spf13/cobra"
	"strings"
)

var CmdSkipEmpty = &cobra.Command{
	Use:     "skipempty",
	Aliases: []string{"se"},
	Short:   "去除空白行（仅包含空格或tab等空白字符）",
	Run: func(cmd *cobra.Command, args []string) {
		LineAction(cmd, func(line string) string {
			if strings.TrimSpace(line) == "" {
				return ""
			}
			return line
		})
	},
}

func initSkipEmpty() {
	CmdSkipEmpty.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdSkipEmpty.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源")
	CmdSkipEmpty.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
}
