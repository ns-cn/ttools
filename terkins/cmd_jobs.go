package main

import "github.com/spf13/cobra"

var CmdJobs = &cobra.Command{
	Use:   "jobs",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
