package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	reqCmd.AddCommand(reqRm)
}

var reqRm = &cobra.Command {
	Use:	"remove",
	Short:	"Delete request",
	Long:	"Delete request",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("NOT IMPLEMENTED")
	},
}
