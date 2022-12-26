package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
)

var CmdEncrypted = goter.Command{Cmd: &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		ReadSetting()
		goter.Required(user, func(u string) bool { return u != "" }, "run without username", nil)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter password for %s: ", user)
		inputPass, _ := reader.ReadString('\n')
		encodedPass := hex.EncodeToString(goter.AesEncryptCBC([]byte(inputPass), goter.GetKey(user)))
		fmt.Printf("加密后的密码: %s\n", encodedPass)
		if envUser == user && isEncrypted && envPass != encodedPass {
			fmt.Printf("系统参数配置了%s的加密密码，但与刚输入的密码不匹配，请检查\n", user)
		}
	},
}}
