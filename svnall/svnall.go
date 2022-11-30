package main

import (
	"fmt"
	"github.com/ns-cn/ttools"
	"github.com/spf13/cobra"
	"os"
)

const (
	ENV_REPOSITORIES   = "SVNALL_REPOSITORIES"
	ENV_DEPTH          = "SVNALL_DEPTH"
	DEFAULT_DEPTH      = 2 // 默认的遍历深度
	SPLITOR_DEPTH      = "#"
	SPLITOR_REPOSITORY = ":"
	NO_ERR             = 0 // 无异常
	ERR_FORMAT_ARG     = 1 // 参数格式
	ERR_ACCESS         = 2 // 路径不合法,或无访问权限
)

var (
	depth int
)

type repository struct {
	dir   string // 目录位置
	depth int    // 查询深度
}

var root = &cobra.Command{
	Use:   "svnall",
	Short: "tool to update multi svn repository update ",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var update = &cobra.Command{
	Use:   "update",
	Short: "tool to update multi svn repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repositories, err := initEnv(args)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		fmt.Println(repositories)
	},
}

func main() {
	root.AddCommand(ttools.VersionCommand("svnall", VERSION))
	root.AddCommand(update)
	root.Flags().IntVarP(&depth, "depth", "d", -1, "the depth searching in dir")
	// 数据源
	_ = root.Execute()
}
