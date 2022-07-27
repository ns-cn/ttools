package main

import (
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var (
	replaceFrom         = ""    // 目标替换字符串
	replaceTo           = ""    // 替换为目标字符串
	replaceUsingRegular = false // 是否是正则表达式形式的替换
)

var CmdReplace = &cobra.Command{
	Use:     "replace",
	Aliases: []string{"r"},
	Short:   "为文本增加前缀",
	Long: `逐行为文本增加前缀;
lines replace [-F {filepath}| -P] [-p {prefix-numberFormat}] [-N {lineIndex-numberFormat}] [-os]
`,
	Run: func(cmd *cobra.Command, args []string) {
		var reg *regexp.Regexp
		var err error
		if replaceUsingRegular {
			reg, err = regexp.Compile(replaceFrom)
			if err != nil {
				handleErrWithTips("正则表达式错误", err)
			}
		}
		LineAction(cmd, func(line Line) string {
			if replaceUsingRegular {
				return reg.ReplaceAllString(line.value, replaceTo)
			} else {
				return strings.ReplaceAll(line.value, replaceFrom, replaceTo)
			}
		})
	},
}

func initReplace() {
	CmdReplace.Flags().StringVarP(&replaceFrom, "from", "f", "", "需要替换的字符串部分")
	CmdReplace.Flags().StringVarP(&replaceTo, "to", "t", "", "替换为目标字符串")
	CmdReplace.Flags().BoolVarP(&replaceUsingRegular, "regular", "r", false, "是否是正则形式的替换")
}
