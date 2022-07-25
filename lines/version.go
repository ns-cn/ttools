package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdVersion = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "打印当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("v.%s", VERSION))
	},
}
