package issue

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func AddComment(issue Issue, comment string) error {
	if _, err := exec.LookPath(ghCmd); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(ghCmd, "issue", "comment", issue.HtmlURL, "--body", comment)
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to send comment to issue: %w", err)
	}
	return nil
}
