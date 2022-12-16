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
	jenkins := gojenkins.CreateJenkins(nil, "http://172.20.21.1:8080/", "tangyujun", "2V@3EYw^vXQS")
	_, err := jenkins.Init(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	session = jenkins
	return jenkins
}
