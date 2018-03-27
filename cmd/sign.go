package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(signCmd)
}

var signCmd = &cobra.Command {
	Use:	"sign",
	Short:	"Request signed certificates",
	Long:	"Request signed public keys from SSH CA for approved requests",
}
