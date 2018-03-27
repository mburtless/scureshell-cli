package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	RootCmd.AddCommand(reqCmd)
}

var reqCmd = &cobra.Command {
	Use:	"request",
	Short:	"Manage requests",
	Long:	"Manage requests",
}
