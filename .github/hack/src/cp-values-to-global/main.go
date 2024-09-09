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

const controller = config.ConfigGlobal

func getValues(yamlConfig configYAML, javaConfig configJava, input issue.GlobalValue) map[string]string {
	//fmt.Println(string(config.Apps["104se-apps"])) // 輸出: cluster cluster.values.yaml
	fmt.Printf("java 版本是: %s ,對應版本是: %s \n", input.Java, javaConfig.AES[input.Java])
	vars := map[string]string{
		"Business":    yamlConfig.Apps[input.Repo], // 使用 repo 變數取代 "104se-apps"
		"Description": config.Description,          // 使用常量 Description 取代 "linux ap cronjob"
		"Owner":       input.Owner,                 // 使用 user 變數取代 "deep.huang"
		"Repository":  input.Repo,                  // 使用 repo 變數
		"AES":         javaConfig.AES[input.Java],  // 使用 app 變數
	}
	return vars
}

func initInput() issue.GlobalValue {
	var input issue.GlobalValue
	input.Repo = githubactions.GetInput("repo")
	if input.Repo == "" {
		githubactions.Fatalf("missing input: issue")
	}
	input.Owner = githubactions.GetInput("owner")
	if input.Owner == "" {
		githubactions.Fatalf("missing input: owner")
	}
	input.App = githubactions.GetInput("app")
	if input.App == "" {
		githubactions.Fatalf("missing input: app")
	}
	input.Java = githubactions.GetInput("java")
	if input.Java == "" {
		githubactions.Fatalf("missing input: java")
	}
	return input
}

type configYAML struct {
	Apps map[string]string `yaml:"apps"`
}
type configJava struct {
	AES map[string]string `yaml:"aes"`
}

func main() {

	input := initInput()
	// 讀取 yaml 文件
	var yamlConfig configYAML
	var javaConfig configJava

	data, err := embedFS.GetYamlFile(controller)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err = yaml.Unmarshal(data, &yamlConfig); err != nil {
		log.Fatalf("error: %v", err)
	}

	data2, err := embedFS.GetYamlFile(config.ConfigCluster)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err = yaml.Unmarshal(data2, &javaConfig); err != nil {
		log.Fatalf("error: %v", err)
	}

	if err = cmd.RunPWDCmd(); err != nil {
		println(err.Error())
	}
	fmt.Printf("the yamlConfig: %v \n the yamlConfig.app: %s", yamlConfig, yamlConfig.Apps["k8s-104dtt-apps"])
	// 設置環境變數
	if err = config.SetEnvVars(getValues(yamlConfig, javaConfig, input), controller); err != nil {
		println(err.Error())
	}

	// 執行命令 envsubst 生成 output.yaml 文件
	if err = cmd.RunCommandToFile(controller)(); err != nil {
		println(err.Error())
	}

	fmt.Println("命令已成功執行並生成新的 yaml 文件: output.yaml")
}
