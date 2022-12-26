package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"os"
)

var session *gojenkins.Jenkins

func getSession() *gojenkins.Jenkins {
	if session != nil {
		return session
	}
	if debug == "Y" {
		fmt.Printf("terkins://%s@%s(using password: %s)\n", host, user, pass)
	}
	jenkins := gojenkins.CreateJenkins(nil, host, user, pass)
	_, err := jenkins.Init(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	session = jenkins
	return jenkins
}
