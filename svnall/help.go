package main

import "github.com/spf13/cobra"

func help(cmd *cobra.Command) {
	_ = cmd.Help()
}
