package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	envCmd.AddCommand(envUpdate)
}

var envUpdate = &cobra.Command {
	Use:	"update",
	Short:	"Update environment",
	Long:	"Update environment",
	Run:	func(cmd *cobra.Command, args []string) {
		log.Print("NOT IMPLEMENTED")
	},
}
