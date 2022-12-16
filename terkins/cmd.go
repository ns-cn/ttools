package main

import "github.com/spf13/cobra"

func cmdBindAll(cmd *cobra.Command) *cobra.Command {
	cmdBindHost(cmd)
	cmdBindUser(cmd)
	cmdBindPass(cmd)
	cmdBindEncrypted(cmd)
	return cmd
}
func cmdBindHost(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&host, "host", "H", "", "host of jenkins, can use ENV TERKINS_HOST instead")
	return cmd
}
func cmdBindUser(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&user, "user", "U", "", "user of jenkins, can use ENV TERKINS_USER instead")
	return cmd
}
func cmdBindPass(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&pass, "password", "P", "", "password of jenkins, can use ENV TERKINS_PASS instead")
	return cmd
}
func cmdBindEncrypted(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&encrypted, "encrypted", "E", "", "password encrypted: Y/N , can use TERKINS_ENCRYPTED instead")
	return cmd
}
