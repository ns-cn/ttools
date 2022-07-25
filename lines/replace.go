package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var CmdReplace = &cobra.Command{
	Use:     "replace",
	Aliases: []string{"r"},
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
			var number string
			if keepOriginal {
				number = fmt.Sprintf(numberFormat, original)
			} else {
				number = fmt.Sprintf(numberFormat, index)
			}
			index++
			return fmt.Sprintf("%s%s", strings.ReplaceAll(box, LINE, number), line)
		})
	},
}
