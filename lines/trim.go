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
	LineAction(cmd, func(line Line) string {
		var result = line.value
		if trimStart {
			result = strings.TrimLeftFunc(result, func(r rune) bool {
				return r == ' ' || r == '\t'
			})
		}
		if trimEnd {
			result = strings.TrimRightFunc(result, func(r rune) bool {
				return r == ' ' || r == '\t'
			})
		}
		return result
	})
}
