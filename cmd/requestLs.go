package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/request"
)

var ReqId string

func init() {
	reqCmd.AddCommand(reqLs)
	reqLs.Flags().StringVarP(&ReqId, "id", "i", "", "Filter by request ID")
}

var reqLs = &cobra.Command {
	Use:	"ls",
	Short:	"List requests",
	Long:	"List requests",
	Run:	func(cmd *cobra.Command, args []string) {
		if ReqId != "" {
			request.GetReqById(ReqId)
		} else {
			request.GetAllReqs()
		}
	},
}
