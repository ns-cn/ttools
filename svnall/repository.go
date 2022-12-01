package main

import (
	"os"
	"path/filepath"
)

type repository struct {
	dir   string // 目录位置
	depth int    // 查询深度
}

func (r repository) subRepositories() []repository {
	subDirs, _ := os.ReadDir(r.dir)
	subRepositories := make([]repository, 0)
	for _, subDir := range subDirs {
		if subDir == nil || !subDir.IsDir() || subDir.Name() == DIR_SVN {
			continue
		}
		subRepository := r.stepIn(subDir.Name())
		if subRepository != nil {
			subRepositories = append(subRepositories, *subRepository)
		}
	}
	return subRepositories
}

func (r repository) stepIn(subDir string) *repository {
	if depth <= 0 {
		return nil
	}
	return &repository{dir: filepath.Join(r.dir, subDir), depth: r.depth - 1}
}
