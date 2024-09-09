package schema

import (
	"embed"
)

const (
	CheckLinuxAPOwner    = "check-linux-ap-owner.schema.json"
	CheckLinuxAPTeam     = "check-linux-ap-team.schema.json"
	CheckLinuxAPApNum    = "check-linux-ap-apnum.schema.json"
	CheckLinuxAPHost     = "check-linux-ap-host.schema.json"
	CheckLinuxAPJava     = "check-linux-ap-java.schema.json"
	CheckLinuxAPRepo     = "check-linux-ap-repo.schema.json"
	CheckLinuxAPBranch   = "check-linux-ap-branch.schema.json"
	CheckLinuxAPEnv      = "check-linux-ap-env.schema.json"
	CheckLinuxAPLocation = "check-linux-ap-location.schema.json"
	CheckLinuxAPTime     = "check-linux-ap-time.schema.json"
	CheckLinuxAPDrRun    = "check-linux-ap-dr-run.schema.json"
)

var (
	//go:embed check-linux-ap-owner.schema.json
	//go:embed check-linux-ap-team.schema.json
	//go:embed check-linux-ap-apnum.schema.json
	//go:embed check-linux-ap-host.schema.json
	//go:embed check-linux-ap-java.schema.json
	//go:embed check-linux-ap-repo.schema.json
	//go:embed check-linux-ap-branch.schema.json
	//go:embed check-linux-ap-env.schema.json
	//go:embed check-linux-ap-location.schema.json
	//go:embed check-linux-ap-time.schema.json
	//go:embed check-linux-ap-dr-run.schema.json
	EmbedFiles embed.FS
)

func GetSchemaPath(checkSchemaItem int) string {
	schemaMap := []string{
		CheckLinuxAPOwner,
		CheckLinuxAPTeam,
		CheckLinuxAPApNum,
		CheckLinuxAPHost,
		CheckLinuxAPJava,
		CheckLinuxAPRepo,
		CheckLinuxAPBranch,
		CheckLinuxAPEnv,
		CheckLinuxAPLocation,
		CheckLinuxAPTime,
		CheckLinuxAPDrRun,
	}
	return schemaMap[checkSchemaItem]
}
