package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// Base cmd, without args
var RootCmd = &cobra.Command{
	Use:	"scureshell",
	Short:	"A client for the scureshell SSH CA",
	Long:	`Request, sign and recieve SSH certificates signed by an SSH CA directly from the command line`,
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*func init() {
	cobra.OnInitialize(initConfig)
}*/
