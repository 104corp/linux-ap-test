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
	"reflect"
)

const controller = config.Configpvc

func getValues(input issue.StorageTemplate, volValue VolumeTemplate, enenvValue map[string]string, conYaml configYAML) map[string]string {
	vars := map[string]string{
		"Env":     enenvValue[input.Location],
		"Team":    input.Team,
		"Name":    input.Name,
		"Volume":  volValue.Volume,
		"Bu":      conYaml.Apps[input.Repo],
		"Server":  volValue.Server,
		"Path":    volValue.ServerPath,
		"Linuxap": input.Host,
	}
	return vars
}

func initInput() issue.StorageTemplate {
	var input issue.StorageTemplate
	input.Repo = githubactions.GetInput("repo")
	if input.Repo == "" {
		githubactions.Fatalf("missing input: repo")
	}
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
	input.Host = githubactions.GetInput("host")
	if input.Host == "" {
		githubactions.Fatalf("missing input: host")
	}
	return input
}

type VolumeTemplate struct {
	Volume     string `yaml:"volume"`
	Server     string `yaml:"server"`
	ServerPath string `yaml:"server_path"`
}

type configTemplate struct {
	NFS struct {
		Lab  []VolumeTemplate `yaml:"lab"`
		Stg  []VolumeTemplate `yaml:"stg"`
		Prod []VolumeTemplate `yaml:"prod"`
		Dr   []VolumeTemplate `yaml:"dr"`
	} `yaml:"nfs"`
}

type configYAML struct {
	Apps map[string]string `yaml:"apps"`
}

func main() {
	input := initInput()
	var yamlConfig configTemplate
	var conYAML configYAML
	data, err := embedFS.GetYamlFile(controller)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err = yaml.Unmarshal(data, &yamlConfig); err != nil {
		log.Fatalf("error: %v", err)
	}

	data2, err := embedFS.GetYamlFile(config.ConfigGlobal)()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	if err = yaml.Unmarshal(data2, &conYAML); err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Storage: %+v\n", conYAML)

	envValue := map[string]string{
		"lab":   "lab",
		"stg":   "stg",
		"prod":  "prod",
		"gke31": "dr",
	}
	fieldName := capitalize(envValue[input.Location])
	field := reflect.ValueOf(yamlConfig.NFS).FieldByName(fieldName)
	if field.IsValid() && field.Kind() == reflect.Slice {
		for i := 0; i < field.Len(); i++ {
			value := field.Index(i)
			// 設定環境變數
			if err := config.SetEnvVars(getValues(input, value.Interface().(VolumeTemplate), envValue, conYAML), controller); err != nil {
				println(err.Error())
			}

			// 執行命令 envsubst 生成 output.yaml 文件
			if err := cmd.RunCommandToFile(controller)(); err != nil {
				println(err.Error())
			}
			outputFileName := fmt.Sprintf("%s.yaml", value.Interface().(VolumeTemplate).Volume)
			fmt.Printf("命令已成功執行並生成新的 yaml 文件: %s\n", outputFileName)
		}
	} else {
		fmt.Printf("Field %s not found or is not a slice\n", envValue[input.Location])
	}

}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-'a'+'A') + s[1:]
}
