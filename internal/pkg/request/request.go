package request

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"text/tabwriter"
	"github.com/mburtless/scureshell-cli/internal/pkg/validationHelper"
	"github.com/mburtless/scureshell-cli/internal/pkg/errorHandler"
)

type Request struct {
	Status        []string `json:"status"`
	ID            string   `json:"_id"`
	EnvironmentID string   `json:"environment_id"`
	UserID        string   `json:"user_id"`
	V             int      `json:"__v"`
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
	var allRequests []Request
	if err := json.NewDecoder(resp.Body).Decode(&allRequests); err != nil {
		log.Println("Invalid JSON Response ", err)
	}

	printReqs(allRequests)
}

func GetReqById(reqId string) {
	queryUrl := viper.GetString("server.base-url") + "/request" + reqId

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
	var thisReq Request
	var allRequests []Request
	if err := json.NewDecoder(resp.Body).Decode(&thisReq); err != nil {
		log.Println("Invalid JSON Response ", err)
	}
	allRequests = append(allRequests, thisReq)
	printReqs(allRequests)

}

func printReqs(allRequests []Request) {
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
