package main

import "os"

// 异常处理模块
func handleErr(err error) {
	handleErrString(err.Error())
}
func handleErrWithTips(tip string, err error) {
	handleErrString(tip, err.Error())
}
func handleErrString(err ...string) {
	for _, errStr := range err {
		_, _ = os.Stderr.WriteString(errStr)
	}
	_, _ = os.Stderr.WriteString("\n")
	//panic(err)
}
