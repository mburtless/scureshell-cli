package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/environment"
)

var EnvId string

func init() {
	envCmd.AddCommand(envLs)
	envLs.Flags().StringVarP(&EnvId, "id", "i", "", "Filter by environment ID")
}

var envLs = &cobra.Command {
	Use:	"ls",
	Short:	"List environments",
	Long:	"List environments",
	Run:	func(cmd *cobra.Command, args []string) {
		//fmt.Println("dummy env")
		if EnvId != "" {
			environment.GetEnvById(EnvId)
		} else {
			environment.GetAllEnvs()
		}
	},

}
