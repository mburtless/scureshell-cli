package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/request"
)

var ReqUserId string
var ReqEnvId string

func init() {
	reqCmd.AddCommand(reqCreate)
	reqCreate.Flags().StringVarP(&ReqUserId, "user", "u", "", "Username (required)")
	reqCreate.Flags().StringVarP(&ReqEnvId, "envid", "e", "", "ID of environment used to sign request (required)")
	reqCreate.MarkFlagRequired("envid")
	reqCreate.MarkFlagRequired("user")
}

var reqCreate = &cobra.Command {
	Use:	"create",
	Short:	"Create request",
	Long:	"Create request",
	Run:	func(cmd *cobra.Command, args []string) {
		//fmt.Println("NOT IMPLEMENTED")
		if(ReqUserId != "" && ReqEnvId != "") {
			request.CreateReq(ReqUserId, ReqEnvId)
		}
	},
}
