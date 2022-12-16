package main

import "github.com/spf13/cobra"

var CmdBuild = cobra.Command{
	Use:   "build",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
