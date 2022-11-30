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
	workdir, err := os.Getwd()
	if err != nil {
		return
	}
	unparsedRepositories := make([]string, 0)
	if len(args) > 0 {
		unparsedRepositories = append(unparsedRepositories, args...)
	} else {
		envRepositories := os.Getenv(ENV_REPOSITORIES)
		strings.Split(envRepositories, "#")
	}
	// 从参数中读取
	for _, unparsedRepository := range unparsedRepositories {
		repository, err := arg2Repo(workdir, unparsedRepository)
		if err > 0 {
			wrongRepositories = append(wrongRepositories, unparsedRepository)
		} else {
			repositories = append(repositories, repository)
		}
	}
	if len(wrongRepositories) != 0 {
		err = fmt.Errorf("wrong repository format: %v\n", wrongRepositories)
	}
	return
}

/*
arg2Repo: 将特定的参数转换为对应的仓库地址信息
格式要求:仓库地址路径[#寻址深度]
*/
func arg2Repo(workdir, arg string) (repository, int) {
	// 还原环境
	defer func(dir string) {
		_ = os.Chdir(dir)
	}(workdir)
	splited := strings.Split(arg, SPLITOR_DEPTH)
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
		err = os.Chdir(dir)
		if err != nil {
			return repository{}, ERR_ACCESS
		}
		return repository{dir: splited[0], depth: simpleDepth}, NO_ERR
	}
}
