package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mburtless/scureshell-cli/internal/pkg/sign"
)

var SignReqID string
var SignUserID string
var SignValidity string
var SignPrincipal string
var SignComment string

func init() {
	RootCmd.AddCommand(signCmd)
	signCmd.Flags().StringVarP(&SignReqID, "reqid", "r", "", "ID of request to be signed (required)")
	signCmd.Flags().StringVarP(&SignUserID, "userid", "u", "", "ID of user who submitted request (required)")
	signCmd.Flags().StringVarP(&SignValidity, "validity", "v", "", "Validity period of signed certificate")
	signCmd.Flags().StringVarP(&SignPrincipal, "principal", "p", "", "Principal name associated with signed certificate")
	signCmd.Flags().StringVarP(&SignComment, "comment", "c", "", "Comment to be included with signed certificate (IS THIS USED?)")
	signCmd.MarkFlagRequired("reqid")
	signCmd.MarkFlagRequired("userid")
}

var signCmd = &cobra.Command {
	Use:	"sign [file]",
	Short:	"Request signed certificates",
	Long:	"Request signed public keys from SSH CA for approved requests",
	Args:	cobra.ExactArgs(1),
	Run:	func(cmd *cobra.Command, args []string) {
		signParams := sign.SignParams{
			PubKeyFilename: args[0],
			ReqID: SignReqID,
			UserID: SignUserID,
			Validity: SignValidity,
			Principal: SignPrincipal,
			Comment: SignComment,
		}
		//fmt.Println("Pub key filename: ", pubKeyFilename)
		sign.Request(&signParams)
	},
}
