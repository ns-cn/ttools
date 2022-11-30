package ttools

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func VersionCommand(tools, version string) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "打印当前版本号",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%s version: v%s(%s/%s)", tools, version, runtime.GOOS, runtime.GOARCH))
		},
	}
}
