package main

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"hack/pkg/config"
	"hack/pkg/issue"
)

func main() {
	input := githubactions.GetInput("issue")
	if input == "" {
		githubactions.Fatalf("missing input: issue")
	}

	var isu issue.Issue
	if err := json.Unmarshal([]byte(input), &isu); err != nil {
		githubactions.Fatalf("issue comment unmarshal err: %s", err)
	}

	var apRequest []map[int]string
	if err := issue.ParseLinuxApForm(&apRequest, &isu); err != nil {
		githubactions.Fatalf("failed to parse body: %s", err)
	}
	if err := issue.SetCluster(&apRequest, apRequest[config.Environment][config.Environment]); err != nil {
		githubactions.Fatalf("failed to set cluster: %s", err)
	}

	githubactions.Infof("%s", "The linux ap resources has been generated")
	githubactions.SetOutput("issue_url", isu.HtmlURL)
	for _, values := range apRequest {
		for key, value := range values {
			if key == config.Environment {
				envMap := map[string]string{
					"lab":  "lab",
					"stg":  "stg",
					"prod": "prod",
					"dr":   "prod",
				}
				value = envMap[value]
			}
			if key == config.DrRun {
				envMap := map[string]string{
					"yes": "false",
					"no":  "true",
				}
				value = envMap[value]
			}
			githubactions.SetOutput(config.OutputItems[key], value)
			fmt.Printf("the %x is %s ", key, value)
		}
	}
	githubactions.SetOutput("namespace", config.GetNameSpace(apRequest[config.Environment][config.Environment], apRequest[config.Team][config.Team], apRequest[config.ApNo][config.ApNo]))
}
