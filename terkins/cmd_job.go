package main

import (
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
)

var CmdJob = &cobra.Command{
	Use:   "job",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		initSetting()
		parseInfo()
		goter.Required(host, func(u string) bool { return u != "" }, "run without host", func() { _ = cmd.Help() })
		goter.Required(user, func(u string) bool { return u != "" }, "run without username", func() { _ = cmd.Help() })
		goter.Required(pass, func(u string) bool { return u != "" }, "run without password", func() { _ = cmd.Help() })
	},
}
