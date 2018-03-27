package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	reqCmd.AddCommand(reqCreate)
}

var reqCreate = &cobra.Command {
	Use:	"create",
	Short:	"Create request",
	Long:	"Create request",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("NOT IMPLEMENTED")
	},
}
