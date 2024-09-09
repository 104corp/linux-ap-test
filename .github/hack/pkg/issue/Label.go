package issue

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	ghCmd             = "gh"
	LabelMigrateAP    = "migrate-team"            // color: blue
	LabelCreateAP     = "create-team-application" // color: blue
	LabelValid        = "valid"                   // color: green
	LabelNeedsApprove = "needs-approve"           // color: red
	LabelApproved     = "approved"                // color: green
	LabelDev          = "lab-cluster"             // color: blue
	LabelStg          = "stg-cluster"             // color: blue
	LabelProd         = "prod-cluster"            // color: blue
	LabelDR           = "gke31-cluster"           // color: blue
	LabelError        = "error"                   // color: red
)

type Label struct {
	Name string `json:"name"`
}

func HasLabel(labels []Label, name string) bool {
	for _, label := range labels {
		if label.Name == name {
			return true
		}
	}
	return false
}

func AddLabel(issue Issue, label string) error {
	if _, err := exec.LookPath(ghCmd); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(ghCmd, "issue", "edit", issue.HtmlURL, "--add-label", label)
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add label to issue: %w", err)
	}
	return nil
}

func RemoveLabel(issue Issue, label string) error {
	if !HasLabel(issue.Labels, label) {
		return nil
	}

	if _, err := exec.LookPath(ghCmd); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(ghCmd, "issue", "edit", issue.HtmlURL, "--remove-label", label)
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove label from issue: %w", err)
	}
	return nil
}

func GetEnvLabel(env string) string {
	fmt.Printf("the cluster is: %s", env)
	if env == "lab" {
		return LabelDev
	}
	if env == "stg" {
		return LabelStg
	}
	if env == "prod" {
		return LabelProd
	}
	if env == "dr" {
		return LabelDR
	}
	return LabelError

}
