package request

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"text/tabwriter"
	"bytes"
	"github.com/mburtless/scureshell-cli/internal/pkg/validationHelper"
	"github.com/mburtless/scureshell-cli/internal/pkg/errorHandler"
)

type PostRes struct {
	Message       string  `json:"message"`
	Request      RequestRes `json:"request"`
}

type RequestRes struct {
	Status        []string `json:"status"`
	ID            string   `json:"_id"`
	EnvironmentID string   `json:"environment_id"`
	UserID        string   `json:"user_id"`
	V             int      `json:"__v"`
}

type RequestReq struct {
	EnvironmentID string   `json:"environment_id"`
	UserID        string   `json:"user_id"`
}

func CreateReq(reqUserID string, reqEnvID string) {
	queryUrl := viper.GetString("server.base-url") + "/request"
	_, err := validationHelper.Url(queryUrl)
	if err != nil {
		errorHandler.Handle(err)
	}

	requestStr := RequestReq{EnvironmentID: reqEnvID, UserID: reqUserID}
	requestJson, err := json.Marshal(requestStr)
	if err != nil {
		log.Fatal("JSON err: ", err)
	}

	req, err := http.NewRequest("POST", queryUrl, bytes.NewBuffer(requestJson))
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
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		log.Fatal("Error: Invalid server response")
	}

	var postRes PostRes
	if err := json.NewDecoder(resp.Body).Decode(&postRes); err != nil {
		log.Println("Invalid JSON Response ", err)
	}

	if postRes.Message == "Request added" {
		printReqs([]RequestRes{postRes.Request})
	} else {
		log.Fatal("Error: Could not create request")
	}
}

func GetAllReqs() {
	queryUrl := viper.GetString("server.base-url") + "/request"

	_, err := validationHelper.Url(queryUrl)
	if err != nil {
		errorHandler.Handle(err)
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorHandler.Handle(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		log.Fatal("Error: Invalid server response")
	}
	var allRequests []RequestRes
	if err := json.NewDecoder(resp.Body).Decode(&allRequests); err != nil {
		log.Println("Invalid JSON Response ", err)
	}

	printReqs(allRequests)
}

func GetReqById(reqId string) {
	queryUrl := viper.GetString("server.base-url") + "/request/" + reqId

	_, err := validationHelper.Url(queryUrl)
	if err != nil {
		errorHandler.Handle(err)
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorHandler.Handle(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		log.Fatal("Error: Invalid server response")
	}
	var thisReq RequestRes
	//var allRequests []RequestRes
	if err := json.NewDecoder(resp.Body).Decode(&thisReq); err != nil {
		log.Println("Invalid JSON Response ", err)
	}
	//allRequests = append(allRequests, thisReq)
	printReqs([]RequestRes{thisReq})

}

func printReqs(allRequests []RequestRes) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "Request ID\tEnvironment ID\tUser ID\tStatus\t")
	for _, req := range allRequests {
		reqString := req.ID + "\t" + req.EnvironmentID + "\t" + req.UserID + "\t" + req.Status[0] + "\t"
		fmt.Fprintln(w, reqString)
	}
	fmt.Fprintln(w)
	w.Flush()

}
