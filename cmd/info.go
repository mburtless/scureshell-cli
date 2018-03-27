package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command {
	Use:	"info",
	Short:	"Display configuration information",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Printf("Scureshell Server: %s\n", viper.GetString("server.base-url"))
		fmt.Printf("Client SSH Directory: %s\n", viper.GetString("client.ssh-cert-dir"))
	},
}
