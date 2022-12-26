package main

import (
	"github.com/ns-cn/goter"
	"os"
	"strings"
)

func ReadSetting() {
	envHost, _ = os.LookupEnv(TERKINS_HOST)
	envUser, _ = os.LookupEnv(TERKINS_USER)
	envPass, _ = os.LookupEnv(TERKINS_PASS)
	envEncrypted, _ = os.LookupEnv(TERKINS_ENCRYPTED)
	if host == "" {
		host = envHost
	}
	if user == "" {
		user = envUser
	}
	if pass == "" {
		pass = envPass
	}
	if encrypted == "" {
		encrypted = envEncrypted
	}
	isEncrypted = strings.ToUpper(encrypted) == "Y" || envEncrypted == ""
	if isEncrypted {
		pass = string(goter.AesEncryptCBC([]byte(pass), goter.GetKey(user)))
	}
}
