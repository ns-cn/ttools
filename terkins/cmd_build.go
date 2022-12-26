package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdBuild = &cobra.Command{
	Use:   "build",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("----------")
		reader := bufio.NewReader(os.Stdin)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				return
			}
			fmt.Println(string(line))
		}
	},
}
