package main

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

	isEncrypted = true

	// env from system
	envHost      = ""
	envUser      = ""
	envPass      = ""
	envEncrypted = ""
)
