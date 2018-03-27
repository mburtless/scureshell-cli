package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	envCmd.AddCommand(envCreate)
}

var envCreate = &cobra.Command {
	Use:	"create",
	Short:	"Create environment",
	Long:	"Create environment",
	Run:	func(cmd *cobra.Command, args []string) {
		log.Print("NOT IMPLEMENTED")
	},
}
