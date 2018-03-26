package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(reqCmd)
}

var reqCmd = &cobra.Command {
	Use:	"request",
	Short:	"Manage requests",
	Long:	"Commands:\n\n" +
			"ls\tList requests\n" +
			"rm\tDelete request\n" +
			"create\tCreate request\n" +
			"update\tUpdate request",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("dummy req")
	},
}
