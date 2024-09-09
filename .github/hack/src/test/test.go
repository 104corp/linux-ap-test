package main

import (
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v2"
	"hack/clusterMap"
	"hack/pkg/config"
	"hack/pkg/issue"
	"hack/pkg/validator"
	"log"
)

const (
	Owner           = 0
	Team            = 1
	ApNo            = 2
	LinuxapHost     = 3
	JavaVersion     = 4
	TeamRepo        = 5
	RepoBranch      = 6
	Environment     = 7
	ClusterLocation = 8
	CrontabTime     = 9
)

type configYAML struct {
	Cluster map[string]string `yaml:"cluster"`
}

func main() {
	formList := []map[int]string{
		{
			Owner: "deep.huang",
		},
		{
			Team: "devops",
		},
		{
			ApNo: "222",
		},
		{
			LinuxapHost: "linux-u1801",
		},
		{
			JavaVersion: "8",
		},
		{
			TeamRepo: "team-repo",
		},
		{
			RepoBranch: "master",
		},
		{
			Environment: "dr",
		},
		{
			ClusterLocation: "on-premise",
		},
		{
			CrontabTime: "0 0 * * *",
		},
	}
	if err := validator.ValidateRequestForm(formList); err != nil {
		fmt.Println(err.Error())
	}

	var yamlConfig configYAML
	controller := "lab"
	data, err := embedFS.GetYamlFile(controller)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	println("")
	if err := issue.SetCluster(&formList, formList[config.Environment][config.Environment]); err != nil {
		githubactions.Fatalf("failed to set cluster: %s", err)
	}
	//fmt.Printf("the data: %s \n", data)
	//fmt.Printf("the yamlConfig: %+v \n", yamlConfig)
	fmt.Printf("the value: %s \n", formList[config.ClusterLocation][ClusterLocation])

}
