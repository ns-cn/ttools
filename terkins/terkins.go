package main

import (
	"github.com/ns-cn/goter"
)

func main() {
	root := goter.NewRootCmd("terkins", "tool to operate jenkins with terminal", VERSION)
	root.AddCommand(CmdJobs.Bind(flagHost, flagUser, flagPass, flagEncrypt))
	root.AddCommand(CmdJob.Bind(flagHost, flagUser, flagPass, flagEncrypt))
	root.AddCommand(CmdBuild.Bind(flagHost, flagUser, flagPass, flagEncrypt, flagInfoToBuild))
	root.AddCommand(CmdEncrypted.Bind(flagPass, flagUser))
	// 数据源
	_ = root.Execute()
}
