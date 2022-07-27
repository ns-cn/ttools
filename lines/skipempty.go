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
		LineAction(cmd, func(line Line) string {
			if strings.TrimSpace(line.value) == "" {
				return ""
			}
			return line.value
		})
	},
}
