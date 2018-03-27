package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	RootCmd.AddCommand(envCmd)
}

var envCmd = &cobra.Command {
	Use:	"environment",
	Short:	"Manage environments",
	Long:	"Manage environments",
}
