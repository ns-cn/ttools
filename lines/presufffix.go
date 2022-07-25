package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	// 导出格式的处理
	numberFormat = "%d"
	box          = ""

	// 范围的处理

	// 空行的处理
	skipEmpty    = false // 是否跳过空白字符行
	keepOriginal = true  // 非空白字符行保持原有行号还是连续编号
)

var CmdPrefix = &cobra.Command{
	Use:     "prefix",
	Aliases: []string{"pf"},
	Short:   "为文本增加前缀",
	Long: `逐行为文本增加前缀;
lines prefix [-F {filepath}| -P] [-p {prefix-numberFormat}] [-N {lineIndex-numberFormat}] [-os]
`,
	Run: func(cmd *cobra.Command, args []string) {
		var index = 0
		var original = 0
		LineAction(cmd, func(line string) string {
			original++
			if strings.TrimSpace(line) == "" && skipEmpty {
				return ""
			}
			index++
			var number string
			if keepOriginal {
				number = fmt.Sprintf(numberFormat, original)
			} else {
				number = fmt.Sprintf(numberFormat, index)
			}
			return fmt.Sprintf("%s%s", strings.ReplaceAll(box, LINE, number), line)
		})
	},
}

var CmdSuffix = &cobra.Command{
	Use:     "suffix",
	Aliases: []string{"sf"},
	Short:   "为文本增加后缀",
	Long: `逐行为文本增加后缀;
lines suffix [-F {filepath}| -P] [- {box}] [-n {number-numberFormat}] [-os]`,
	Run: func(cmd *cobra.Command, args []string) {
		var index = 0
		var original = 0
		LineAction(cmd, func(line string) string {
			original++
			if strings.TrimSpace(line) == "" && skipEmpty {
				return ""
			}
			index++
			var number string
			if keepOriginal {
				number = fmt.Sprintf(numberFormat, original)
			} else {
				number = fmt.Sprintf(numberFormat, index)
			}
			return fmt.Sprintf("%s%s\n", line[:len(line)-1], strings.ReplaceAll(box, LINE, number))
		})
	},
}

func initPrefix() {
	CmdPrefix.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdPrefix.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源")
	CmdPrefix.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	CmdSuffix.Flags().StringVarP(&filePath, "file", "F", "", "目标文件, 不指定则从管道中读取")
	CmdSuffix.Flags().BoolVarP(&fromClipboard, "fromClipboard", "C", false, "是否从粘贴板读取数据作为格式化数据的数据源")
	CmdSuffix.Flags().BoolVarP(&toClipboard, "toClipboard", "c", false, "是否将处理结果粘贴到粘贴板（默认输出到标准输出）")
	// ----------------前缀----------------
	// 格式数据
	CmdPrefix.Flags().StringVarP(&box, "box", "b", "", "目标位置填写内容，其中#number为占位符标示行号, \n"+
		"例如\"#number|\"表示使用下划线分割行号和正文")
	CmdPrefix.Flags().StringVarP(&numberFormat, "number", "n", "%d", "行号的格式化风格，例如%4d，则格式化为4位，%-4d则4位居左")
	// 空白行处理
	CmdPrefix.Flags().BoolVarP(&skipEmpty, "skipEmpty", "s", false, "是否跳过空白行，例如true表示跳过空白行")
	CmdPrefix.Flags().BoolVarP(&keepOriginal, "keepOriginal", "o", true, "非空白字符行保持原有行号还是连续编号，为false则针对显示字符行进行连续边行")

	// ----------------后缀----------------
	// 格式数据
	CmdSuffix.Flags().StringVarP(&box, "box", "b", "", "目标位置填写内容，其中#number为占位符标示行号, \n"+
		"例如\"#number|\"表示使用下划线分割行号和正文")
	CmdSuffix.Flags().StringVarP(&numberFormat, "number", "n", "%d", "行号的格式化风格，例如%4d，则格式化为4位，%-4d则4位居左")
	// 空白行处理
	CmdSuffix.Flags().BoolVarP(&skipEmpty, "skipEmpty", "s", false, "是否跳过空白行，例如true表示跳过空白行")
	CmdSuffix.Flags().BoolVarP(&keepOriginal, "keepOriginal", "o", true, "非空白字符行保持原有行号还是连续编号，为false则针对显示字符行进行连续边行")
}
