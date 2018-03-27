package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	reqCmd.AddCommand(reqUpdate)
}

var reqUpdate = &cobra.Command {
	Use:	"update",
	Short:	"Update request",
	Long:	"Update request",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("NOT IMPLEMENTED")
	},
}
