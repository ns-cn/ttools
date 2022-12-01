package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func updateAll(repositories []repository) error {
	for _, r := range repositories {
		updateRepository(r)
	}
	return nil
}

func updateRepository(repository repository) {
	dir := repository.dir
	if svnUpdate(dir) && !fullThrough {
		return
	}
	for _, subRepository := range repository.subRepositories() {
		updateRepository(subRepository)
	}
}

func svnUpdate(dir string) bool {
	stat, err := os.Stat(filepath.Join(dir, DIR_SVN))
	if err != nil || !stat.IsDir() {
		return false
	}
	fmt.Printf("svn update in:%s\n", dir)
	command := exec.Command("svn", "update")
	command.Dir = dir
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	return err == nil
}
