package main

import (
	"github.com/ns-cn/goter"
)

func main() {
	root := goter.Root("terkins", "tool to operate jenkins with terminal", VERSION)
	root.AddCommand(cmdBindAll(CmdJobs))
	root.AddCommand(cmdBindAll(CmdJob))
	cmdBindPass(CmdEncrypted)
	cmdBindUser(CmdEncrypted)
	cmdBindEncrypted(CmdEncrypted)
	root.AddCommand(CmdEncrypted)
	// 数据源
	_ = root.Execute()
}
