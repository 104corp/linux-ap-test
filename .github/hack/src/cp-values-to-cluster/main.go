package main

import (
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v2"
	"hack/embedFS"
	"hack/pkg/cmd"
	"hack/pkg/config"
	"hack/pkg/issue"
	"log"
)

const controller = config.ConfigCluster

func getValues(yamlConfig configYAML, locateConfig locationConfig, input issue.ClusterValue) map[string]string {
	//fmt.Println(string(config.Apps["104se-apps"])) // 輸出: cluster cluster.values.yaml
	vars := map[string]string{
		"JavaVersion": yamlConfig.Java[input.Java],
		"Team":        input.NameSpace,
		"Time":        input.Time,
		"APNum":       input.Name,
		"Location":    locateConfig.Location[input.Location],
		"DrRun":       input.Suspend,
	}
	fmt.Printf("CP-to-cluster the Java input is %s, the mapping version is %s ", input.Java, yamlConfig.Java[input.Java])
	fmt.Printf("CP-to-cluster the name input is %s", input.NameSpace)
	fmt.Printf("CP-to-cluster the Time input is %s", input.Time)
	fmt.Printf("CP-to-cluster the APNum input is %s", input.Name)
	fmt.Printf("CP-to-cluster the cluster input is %s, the location input is %s", input.Location, locateConfig.Location[input.Location])
	fmt.Printf("CP-to-cluster the suspend input is %s", input.Suspend)
	return vars
}

func initInput() issue.ClusterValue {
	var input issue.ClusterValue
	input.Java = githubactions.GetInput("java")
	if input.Java == "" {
		githubactions.Fatalf("missing input: java")
	}
	input.Time = githubactions.GetInput("time")
	if input.Time == "" {
		githubactions.Fatalf("missing input: time")
	}
	input.Name = githubactions.GetInput("apnum")
	if input.Name == "" {
		githubactions.Fatalf("missing input: apnum")
	}
	input.Team = githubactions.GetInput("team")
	if input.Team == "" {
		githubactions.Fatalf("missing input: team")
	}
	input.Env = githubactions.GetInput("env")
	if input.Env == "" {
		githubactions.Fatalf("missing input: env")
	}
	input.Location = githubactions.GetInput("location")
	if input.Location == "" {
		githubactions.Fatalf("missing input: cluster")
	}
	input.Suspend = githubactions.GetInput("suspend")
	if input.Suspend == "" {
		githubactions.Fatalf("missing input: suspend")
	}
	return input
}

type configYAML struct {
	Java map[string]string `yaml:"versionMap"`
}

type locationConfig struct {
	Location map[string]string `yaml:"location"`
}

func main() {
	input := initInput()

	input.NameSpace = config.GetNameSpace(input.Env, input.Team, input.Name)

	// 讀取 yaml 文件
	var yamlConfig configYAML
	var locateConfig locationConfig

	data, err := embedFS.GetYamlFile(controller)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	fmt.Printf("the Namespace is %s", input.NameSpace)
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	data2, err := embedFS.GetYamlFile(config.ConfigLocatipn)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data2, &locateConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = cmd.RunPWDCmd()
	if err != nil {
		println(err.Error())
	}

	// 設定環境變數
	if err = config.SetEnvVars(getValues(yamlConfig, locateConfig, input), controller); err != nil {
		println(err.Error())
	}
	// 執行命令 envsubst 生成 output.yaml 文件
	if err = cmd.RunCommandToFile(controller)(); err != nil {
		println(err.Error())
	}
}
