package main

import (
	"github.com/ns-cn/goter"
)

func main() {
	root := goter.Root("terkins", "tool to operate jenkins with terminal", VERSION)
	root.AddCommand(cmdBindAll(CmdJobs))
	root.AddCommand(cmdBindAll(CmdJob))
	root.AddCommand(cmdBindAll(CmdBuild))
	root.AddCommand(cmdBind(CmdEncrypted, FlagPass, FlagUser, FlagEncrypt))
	// 数据源
	_ = root.Execute()
}
