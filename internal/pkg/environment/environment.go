package environment

import (
	"github.com/spf13/viper"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"text/tabwriter"
	"net/url"
	"net"
)

type Environment struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	UserCert string `json:"user_cert"`
	HostCert string `json:"host_cert"`
	V        int    `json:"__v"`
}


func GetAllEnvs() {
	queryUrl := viper.GetString("server.base-url") + "/environment"
	_, err := url.ParseRequestURI(queryUrl)
	if err != nil {
		errorHandler(err)
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorHandler(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		log.Fatal("Error: Invalid server response")
	}
	var allEnvironments []Environment
	if err := json.NewDecoder(resp.Body).Decode(&allEnvironments); err != nil {
		log.Println("Invalid JSON Response: ", err)
	}

	printEnvs(allEnvironments)
}

func GetEnvById(envId string) {
	//need custom error on bad envId

	queryUrl := viper.GetString("server.base-url") + "/environment/" + envId
	_, err := url.ParseRequestURI(queryUrl)
	if err != nil {
		errorHandler(err)
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorHandler(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		log.Fatal("Error: Invalid server response")
	}
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

func errorHandler(err error) {
	if ue, ok := err.(*url.Error); ok {
		//handle connection refused error
		switch uet := ue.Err.(type) {
			//if oe, ok := ue.Err.(*net.OpError); ok {
		case *net.OpError:
				switch oet := uet.Err.(type) {
				case *os.SyscallError:
					//if se, ok := oe.Err.(*os.SyscallError); ok {
						if oet.Err.Error() == "connection refused" {
							log.Fatalf("Error: Connection refused when attempting to connect to scureshell server at %s", uet.Addr.(*net.TCPAddr))
						}
					//}
			}
		default:
			if ue.Op == "parse" {
				log.Fatalf("Error: Invalid URL provided for scureshell server - \"%s\"", ue.URL)
			}
		}
	}
}
