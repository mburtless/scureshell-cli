package environment

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"bytes"
	"text/tabwriter"
	"github.com/mburtless/scureshell-cli/internal/pkg/validationHelper"
	"github.com/mburtless/scureshell-cli/internal/pkg/errorHandler"
)

type PostRes struct {
	Message       string         `json:"message"`
	Environment   EnvironmentRes `json:"environment"`
}

type EnvironmentRes struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	UserCert string `json:"user_cert"`
	HostCert string `json:"host_cert"`
	V        int    `json:"__v"`
}

type EnvironmentReq struct {
	Name     string `json:"name"`
	UserCert string `json:"user_cert"`
	HostCert string `json:"host_cert"`
}

func CreateEnv(envName string, envUserCert string, envHostCert string) {
	queryUrl := viper.GetString("server.base-url") + "/environment"
	_, err := validationHelper.Url(queryUrl)
	if err != nil {
		errorHandler.Handle(err)
	}

	environmentStr := EnvironmentReq{Name: envName, UserCert: envUserCert, HostCert: envHostCert}
	environmentJson, err := json.Marshal(environmentStr)
	if err != nil {
		log.Fatal("JSON err: ", err)
	}

	req, err := http.NewRequest("POST", queryUrl, bytes.NewBuffer(environmentJson))
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

	if postRes.Message == "Environment added" {
		printEnvs([]EnvironmentRes{postRes.Environment})
	} else {
		log.Fatalf("Error: Could not create environment\n%s", postRes.Message)
	}
}

func GetAllEnvs() {
	queryUrl := viper.GetString("server.base-url") + "/environment"

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
	var allEnvironments []EnvironmentRes
	if err := json.NewDecoder(resp.Body).Decode(&allEnvironments); err != nil {
		log.Println("Invalid JSON Response: ", err)
	}

	printEnvs(allEnvironments)
}

func GetEnvById(envId string) {
	//need custom error on bad envId
	/*_, err := validationHelper.environmentId(envId)
	if err != nil {
		errorHandler.Handle(err)
	}*/

	queryUrl := viper.GetString("server.base-url") + "/environment/" + envId

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
	var env EnvironmentRes
	var allEnvironments []EnvironmentRes
	if err := json.NewDecoder(resp.Body).Decode(&env); err != nil {
		log.Println("Invalid JSON Response: ", err)
	}
	allEnvironments = append(allEnvironments, env)
	printEnvs(allEnvironments)

}

func printEnvs(allEnvironments []EnvironmentRes) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(w, "Environment ID\tName\tUser CA Certificate\tHost CA Certificate\t")
	for _, env := range allEnvironments {
		envString := env.ID + "\t" + env.Name + "\t" + env.UserCert + "\t" + env.HostCert + "\t"
		fmt.Fprintln(w, envString)
	}
	fmt.Fprintln(w)
	w.Flush()

}
