package main

import (
	"github.com/spf13/cobra"
)

const (
	VERSION = "1.0"
)

var root = &cobra.Command{
	Use:   "cuttlefish",
	Short: "八爪鱼，一个任务分发的",
	Long:  `八爪鱼，一个任务分发的`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	root.AddCommand(CmdVersion) // 打印版本号
	// 数据源
	_ = root.Execute()
}
