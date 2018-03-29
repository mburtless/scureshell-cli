package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/environment"
)
var EnvName string
var EnvUserCert string
var EnvHostCert string

func init() {
	envCmd.AddCommand(envCreate)
	envCreate.Flags().StringVarP(&EnvName, "name", "n", "", "Environment Name (required)")
	envCreate.Flags().StringVarP(&EnvUserCert, "usercert", "u", "", "CA certificate name for user key signing requests (required)")
	envCreate.Flags().StringVarP(&EnvHostCert, "servercert", "s", "", "CA certificate name for server key signing requests (required)")
	envCreate.MarkFlagRequired("name")
	envCreate.MarkFlagRequired("usercert")
	envCreate.MarkFlagRequired("servercert")
}

var envCreate = &cobra.Command {
	Use:	"create",
	Short:	"Create environment",
	Long:	"Create environment",
	Run:	func(cmd *cobra.Command, args []string) {
		if(EnvName != "" && EnvUserCert != "" && EnvHostCert != "") {
			environment.CreateEnv(EnvName, EnvUserCert, EnvHostCert)
		}
	},
}
