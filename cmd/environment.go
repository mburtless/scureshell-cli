package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/environment"
)

func init() {
	RootCmd.AddCommand(envCmd)
}

var envCmd = &cobra.Command {
	Use:	"environment",
	Short:	"Manage environments",
	Long:	"Commands:\n\n" +
			"ls\tList environments\n" +
			"rm\tDelete environment\n" +
			"create\tCreate environment\n" +
			"update\tUpdate environment",
	Run:	func(cmd *cobra.Command, args []string) {
		//fmt.Println("dummy env")
		environment.GetAllEnvs()
	},
}
