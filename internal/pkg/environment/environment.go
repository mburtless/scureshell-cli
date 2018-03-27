package environment

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"text/tabwriter"
)

type Environment struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	UserCert string `json:"user_cert"`
	HostCert string `json:"host_cert"`
	V        int    `json:"__v"`
}


func GetAllEnvs() {
	url := viper.GetString("server.base-url") + "/environment"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	var allEnvironments []Environment
	if err := json.NewDecoder(resp.Body).Decode(&allEnvironments); err != nil {
		log.Println("Invalid JSON Response: ", err)
	}

	printEnvs(allEnvironments)
}

func GetEnvById(envId string) {
	//fmt.Println("NOT IMPLEMENTED")
	url := viper.GetString("server.base-url") + "/environment/" + envId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	var env Environment
	var allEnvironments []Environment
	if err := json.NewDecoder(resp.Body).Decode(&env); err != nil {
		log.Println("Invalid JSON Response: ", err)
	}
	allEnvironments = append(allEnvironments, env)
	printEnvs(allEnvironments)

}

func printEnvs(allEnvironments []Environment) {
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
