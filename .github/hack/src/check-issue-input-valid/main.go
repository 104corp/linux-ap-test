package main

import (
	"encoding/json"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"hack/pkg/config"
	"hack/pkg/issue"
	"hack/pkg/validator"
	"os"
)

func main() {
	input := githubactions.GetInput("issue")
	if input == "" {
		githubactions.Fatalf("missing input: issue")
	}

	var isu issue.Issue
	if err := json.Unmarshal([]byte(input), &isu); err != nil {
		githubactions.Fatalf("issue comment unmarshal err: %s", err)
		os.Exit(1)
	}

	var apRequest []map[int]string
	if err := issue.ParseLinuxApForm(&apRequest, &isu); err != nil {
		_ = issue.RemoveLabel(isu, issue.LabelValid)
		githubactions.Fatalf("failed to parse body: %s", err)
		os.Exit(2)
	}

	if err := validator.ValidateRequestForm(apRequest); err != nil {
		_ = issue.RemoveLabel(isu, issue.LabelValid)
		if errors := issue.AddComment(isu, err.Error()); errors != nil {
			githubactions.Fatalf("%s", errors)
		}
		os.Exit(3)
	}
	_ = issue.AddLabel(isu, issue.LabelValid)
	_ = issue.AddLabel(isu, issue.LabelNeedsApprove)
	fmt.Printf("the input is %s:", apRequest[config.Environment][config.Environment])
	_ = issue.AddLabel(isu, issue.GetEnvLabel(apRequest[config.Environment][config.Environment]))
}
func printRequest(apRequest []map[int]string) {
	for _, values := range apRequest {
		for key, value := range values {
			fmt.Printf("the key: %s, the value: %s \n", config.OutputItems[key], value)
		}
	}
}
