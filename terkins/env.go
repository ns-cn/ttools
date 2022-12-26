package main

import "github.com/ns-cn/goter"

const (
	TERKINS_HOST      = "TERKINS_HOST"
	TERKINS_USER      = "TERKINS_USER"
	TERKINS_PASS      = "TERKINS_PASS"
	TERKINS_ENCRYPTED = "TERKINS_ENCRYPTED"
)

var (
	// effected setting
	host      = ""
	user      = ""
	pass      = ""
	encrypted = ""
	debug     = ""

	flagHost    = goter.CmdStringFlag{P: &host, Name: "host", Shorthand: "H", Value: "", Usage: "host of jenkins, can use ENV TERKINS_HOST instead"}
	flagUser    = goter.CmdStringFlag{P: &user, Name: "user", Shorthand: "U", Value: "", Usage: "user of jenkins, can use ENV TERKINS_USER instead"}
	flagPass    = goter.CmdStringFlag{P: &pass, Name: "password", Shorthand: "P", Value: "", Usage: "password of jenkins, can use ENV TERKINS_PASS instead"}
	flagEncrypt = goter.CmdStringFlag{P: &encrypted, Name: "encrypted", Shorthand: "E", Value: "", Usage: "password encrypted: Y/N , can use TERKINS_ENCRYPTED instead"}
	flagDebug   = goter.CmdStringFlag{P: &debug, Name: "debug", Shorthand: "D", Value: "N", Usage: "debug: Y/N"}

	isEncrypted = true
	// env from system
	envHost      = ""
	envUser      = ""
	envPass      = ""
	envEncrypted = ""
)

// 构建使用的参数
var (
	infoToBuild     = ""
	isInfoToBuild   = true
	flagInfoToBuild = goter.CmdStringFlag{P: &infoToBuild, Name: "info", Shorthand: "I", Value: "", Usage: "info each job to build or not"}
)
