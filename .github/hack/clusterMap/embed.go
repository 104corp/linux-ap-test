package embedFS

import (
	"embed"
)

const (
	labYaml   = "labCluster.yaml"
	stgYaml   = "stgCluster.yaml"
	prodYaml  = "prodCluster.yaml"
	drYaml    = "drCluster.yaml"
	configNum = 4
)

var (
	//go:embed labCluster.yaml
	//go:embed stgCluster.yaml
	//go:embed prodCluster.yaml
	//go:embed drCluster.yaml
	ConfigYAML embed.FS
)

func ConfigLabYaml() ([]byte, error) {
	return ConfigYAML.ReadFile(labYaml)
}

func ConfigStgYaml() ([]byte, error) {
	return ConfigYAML.ReadFile(stgYaml)
}

func ConfigProdYaml() ([]byte, error) {
	return ConfigYAML.ReadFile(prodYaml)
}

func ConfigDrYaml() ([]byte, error) {
	return ConfigYAML.ReadFile(drYaml)
}

// 使用 func array 來定義
type FuncType func() ([]byte, error)

var embedReadFile = map[string]FuncType{
	"lab":  ConfigLabYaml,  // lab cluster config
	"stg":  ConfigStgYaml,  // stg cluster config
	"prod": ConfigProdYaml, // prod cluster config
	"dr":   ConfigDrYaml,   // dr cluster config
}

func GetYamlFile(controller string) FuncType {
	return embedReadFile[controller]
}
