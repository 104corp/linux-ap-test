package main

import (
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"hack/pkg/cmd"
	"hack/pkg/config"
	"hack/pkg/issue"
)

const controller = config.ConfigNamespace

func getValues(input issue.NameSpaceTempalte) map[string]string {
	vars := map[string]string{
		"Team":     input.Team,
		"Name":     input.Name,
		"Location": input.Location,
	}
	return vars
}

func initInput() issue.NameSpaceTempalte {
	var input issue.NameSpaceTempalte
	input.Team = githubactions.GetInput("team")
	if input.Team == "" {
		githubactions.Fatalf("missing input: team")
	}
	input.Name = fmt.Sprintf("ap-%s", githubactions.GetInput("apnum"))
	if input.Name == "" {
		githubactions.Fatalf("missing input: apnum")
	}
	input.Location = githubactions.GetInput("location")
	if input.Location == "" {
		githubactions.Fatalf("missing input: location")
	}
	//	if input.Location == "gke31" {
	//		input.Location = "prod,gke31"
	//	}
	return input
}

func main() {
	input := initInput()

	// 設定環境變數
	if err := config.SetEnvVars(getValues(input), controller); err != nil {
		println(err.Error())
	}

	// 執行命令 envsubst 生成 output.yaml 文件
	if err := cmd.RunCommandToFile(controller)(); err != nil {
		println(err.Error())
	}

	fmt.Println("命令已成功執行並生成新的 yaml 文件: output.yaml")
}
