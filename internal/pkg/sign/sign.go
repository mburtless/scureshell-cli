package sign

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
	//"os"
	//"fmt"
	//"text/tabwriter"
	"bytes"
	"github.com/mburtless/scureshell-cli/internal/pkg/validationHelper"
	"github.com/mburtless/scureshell-cli/internal/pkg/errorHandler"
)

type SignParams struct {
	PubKeyFilename	string
	ReqID			string
	UserID			string
	Validity		string
	Principal		string
	Comment			string
}

type SignReq struct {
	PublicKey string `json:"public_key,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Validity  string `json:"validity,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Principal string `json:"principal,omitempty"`
}

type PostRes struct {
	Status int `json:"status"`
	Data   []struct {
		Signedkey string `json:"signedkey"`
	} `json:"data"`
	Message  string `json:"message,omitempty"`
}

func Request(params *SignParams) {
	queryUrl := viper.GetString("server.base-url") + "/sign"
	_, err := validationHelper.Url(queryUrl)
	if err != nil {
		errorHandler.Handle(err)
	}

	_, err = validationHelper.File(params.PubKeyFilename)
	if err != nil {
		log.Fatalf("Error: Public key file %s could not be found", params.PubKeyFilename)
	}

	//Read public key file to string
	pubKeyBytes, err := ioutil.ReadFile(params.PubKeyFilename)
	if err != nil {
		log.Fatalf("Error: Public key file %s could not be read", params.PubKeyFilename)
	}
	pubKey := string(pubKeyBytes)

	signStr := SignReq{
		PublicKey: pubKey,
		RequestID: params.ReqID,
		UserID: params.UserID,
		Validity: params.Validity,
		Comment: params.Comment,
		Principal: params.Principal,
	}

	signJson, err := json.Marshal(signStr)
	if err != nil {
		log.Fatal("JSON err: ", err)
	}

	//log.Printf("%#v", signJson)
	req, err := http.NewRequest("POST", queryUrl, bytes.NewBuffer(signJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorHandler.Handle(err)
	}

	defer resp.Body.Close()
	if (resp.StatusCode > 299 || resp.StatusCode < 200) && (resp.StatusCode != 400) {
		log.Fatal("Error: Invalid server response")
	}

	var postRes PostRes
	if err := json.NewDecoder(resp.Body).Decode(&postRes); err != nil {
		log.Println("Invalid JSON Response ", err)
	}

	if postRes.Status == 200 {
		//printReqs([]RequestRes{postRes.Request})
		log.Printf("%v", postRes.Data[0].Signedkey)
	} else {
		log.Fatalf("Error: Public key could not be signed\n%s", postRes.Message)
	}
}
