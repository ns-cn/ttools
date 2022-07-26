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
		LineAction(cmd, func(line string) string {
			if replaceUsingRegular {
				return reg.ReplaceAllString(line, replaceTo)
			} else {
				return strings.ReplaceAll(line, replaceFrom, replaceTo)
			}
		})
	},
}

func initReplace() {
	CmdReplace.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdReplace.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false,
		"是否从粘贴板读取数据作为格式化数据的数据源，如果为true，则不需要指定file参数")
	CmdReplace.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	CmdReplace.Flags().StringVarP(&replaceFrom, "from", "f", "", "需要替换的字符串部分")
	CmdReplace.Flags().StringVarP(&replaceTo, "to", "t", "", "替换为目标字符串")
	CmdReplace.Flags().BoolVarP(&replaceUsingRegular, "regular", "r", false, "是否是正则形式的替换")
}
