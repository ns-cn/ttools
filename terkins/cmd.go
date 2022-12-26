package main

import "github.com/spf13/cobra"

type cmdStringFlag struct {
	p               *string
	name, shorthand string
	value           string
	usage           string
}

var (
	FlagHost    = cmdStringFlag{p: &host, name: "host", shorthand: "H", value: "", usage: "host of jenkins, can use ENV TERKINS_HOST instead"}
	FlagUser    = cmdStringFlag{p: &user, name: "user", shorthand: "U", value: "", usage: "user of jenkins, can use ENV TERKINS_USER instead"}
	FlagPass    = cmdStringFlag{p: &pass, name: "password", shorthand: "P", value: "", usage: "password of jenkins, can use ENV TERKINS_PASS instead"}
	FlagEncrypt = cmdStringFlag{p: &encrypted, name: "encrypted", shorthand: "E", value: "", usage: "password encrypted: Y/N , can use TERKINS_ENCRYPTED instead"}
)

func cmdBindAll(cmd *cobra.Command) *cobra.Command {
	cmdBind(cmd, FlagHost, FlagUser, FlagPass, FlagEncrypt)
	return cmd
}

func cmdBind(cmd *cobra.Command, flags ...cmdStringFlag) *cobra.Command {
	for _, flag := range flags {
		cmd.Flags().StringVarP(flag.p, flag.name, flag.shorthand, flag.value, flag.usage)
	}
	return cmd
}
