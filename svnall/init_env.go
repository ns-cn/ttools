package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func initEnv(args []string) (repositories []repository, err error) {
	// no user input or wrong user input
	if depth < 0 {
		envDepth := os.Getenv(ENV_DEPTH)
		parsedDepth, err := strconv.Atoi(envDepth)
		if err != nil || parsedDepth < 0 {
			depth = DEFAULT_DEPTH
		} else {
			depth = parsedDepth
		}
	}
	repositories = make([]repository, 0)
	wrongRepositories := make([]string, 0)
	if err != nil {
		return
	}
	unparsedRepositories := make([]string, 0)
	if len(args) > 0 {
		unparsedRepositories = append(unparsedRepositories, args...)
	}
	if !exclude {
		envRepositories := os.Getenv(ENV_REPOSITORIES)
		for _, unparsedRepository := range strings.Split(envRepositories, SPLITOR_REPOSITORY) {
			if unparsedRepository != "" {
				unparsedRepositories = append(unparsedRepositories, unparsedRepository)
			}
		}
	}
	if unparsedRepositories == nil || len(unparsedRepositories) == 0 {
		err = fmt.Errorf("尚未指定更新仓库地址,可选参数或环境变量方式\n")
		return
	}
	// 从参数中读取
	for _, unparsedRepository := range unparsedRepositories {
		repository, err := parseRepository(unparsedRepository)
		if err > 0 {
			wrongRepositories = append(wrongRepositories, unparsedRepository)
		} else {
			repositories = append(repositories, repository)
		}
	}
	if len(wrongRepositories) != 0 {
		err = fmt.Errorf("wrong repository: %v\n", wrongRepositories)
	}
	return
}

/*
arg2Repo: 将特定的参数转换为对应的仓库地址信息
格式要求:仓库地址路径[#寻址深度]
*/
func parseRepository(unparsedRepository string) (repository, int) {
	splited := strings.Split(unparsedRepository, SPLITOR_DEPTH)
	if len(splited) > 2 || splited[0] == "" {
		return repository{}, ERR_FORMAT_ARG
	} else {
		var simpleDepth = depth
		var err error
		if len(splited) == 2 && splited[1] != "" {
			simpleDepth, err = strconv.Atoi(splited[1])
			if err != nil {
				return repository{}, ERR_FORMAT_ARG
			}
		}
		dir := splited[0]
		if !isDirExists(dir) {
			return repository{}, ERR_ACCESS
		}
		return repository{dir: dir, depth: simpleDepth}, NO_ERR
	}
}

func isDirExists(path string) bool {
	state, err := os.Stat(path)
	return err == nil && state.IsDir()
}
