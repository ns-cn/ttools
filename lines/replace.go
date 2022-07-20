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
lines prefix [-F {filepath}| -P] [-p {prefix-format}] [-N {lineIndex-format}] [-os]
`,
	Run: func(cmd *cobra.Command, args []string) {
		var index = 0
		var original = 0
		LineAction(cmd, func(line string) {
			original++
			if strings.TrimSpace(line) == "" && skipEmpty {
				return
			}
			var number string
			if keepOriginal {
				number = fmt.Sprintf(format, original)
			} else {
				number = fmt.Sprintf(format, index)
			}
			fmt.Printf("%s%s", strings.ReplaceAll(content, LINE, number), line)
			index++
		})
	},
}
