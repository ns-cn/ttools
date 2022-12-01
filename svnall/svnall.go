package main

import (
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

	DIR_SVN = ".svn" // svn目录
)

var (
	depth       int  // 遍历深度
	exclude     bool // 排除环境变量内容,不排除默认所有
	fullThrough bool // 是否做整体穿透，即便找到.svn也继续往下
)

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
			help(cmd)
			return
		}
		updateAll(repositories)
	},
}

func main() {
	root.AddCommand(ttools.VersionCommand("svnall", VERSION))
	root.AddCommand(update)
	update.Flags().IntVarP(&depth, "depth", "d", -1, "the depth searching in dir(不指定则使用环境变量：SVNALL_DEPTH，默认值2)")
	update.Flags().BoolVarP(&exclude, "exclude", "e", false, "是否排除环境变量配置仓库,默认不排除（环境变量：SVNALL_REPOSITORIES）")
	update.Flags().BoolVarP(&fullThrough, "through", "t", false, "是否已经找到.svn继续往下查找，默认不继续往下")
	// 数据源
	_ = root.Execute()
}
