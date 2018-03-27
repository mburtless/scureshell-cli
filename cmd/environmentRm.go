package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	envCmd.AddCommand(envRm)
}

var envRm = &cobra.Command {
	Use:	"rm",
	Short:	"Delete environment",
	Long:	"Delete environment",
	Run:	func(cmd *cobra.Command, args []string) {
		log.Print("NOT IMPLEMENTED")
	},
}
