package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var CmdSkipEmpty = &cobra.Command{
	Use:     "skipempty",
	Aliases: []string{"se"},
	Short:   "去除空白行（仅包含空格或tab等空白字符）",
	Run: func(cmd *cobra.Command, args []string) {
		LineAction(cmd, func(line string) {
			if strings.TrimSpace(line) == "" {
				return
			}
			fmt.Print(line)
		})
	},
}

func initSkipEmpty() {
	CmdSkipEmpty.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
}
