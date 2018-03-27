package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/environment"
)

var EnvId string

func init() {
	RootCmd.AddCommand(envCmd)
	envCmd.Flags().StringVarP(&EnvId, "id", "i", "", "Environment ID")
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
		if EnvId != "" {
			environment.GetEnvById(EnvId)
		} else {
			environment.GetAllEnvs()
		}
	},
}
