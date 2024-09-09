package embedFS

import (
	"embed"
	"fmt"
)

const (
	DescribeYaml    = "teamDescribe.yaml"
	JavaVersionYaml = "javaVersion.yaml"
	LocationnYaml   = "locationMap.yaml"
	StorageMapYaml  = "storageMap.yaml"
	FunctionNum     = 4
)

var (
	//go:embed teamDescribe.yaml
	//go:embed javaVersion.yaml
	//go:embed locationMap.yaml
	//go:embed storageMap.yaml
	ConfigYAML embed.FS
)

func ConfigDescribe() ([]byte, error) {
	return ConfigYAML.ReadFile(DescribeYaml)
}

func ConfigJavaVersion() ([]byte, error) {
	return ConfigYAML.ReadFile(JavaVersionYaml)
}

func ConfigLocationMap() ([]byte, error) {
	return ConfigYAML.ReadFile(LocationnYaml)
}

func COnfigStorageMap() ([]byte, error) {
	return ConfigYAML.ReadFile(StorageMapYaml)
}

// 使用 func array 來定義
type FuncType func() ([]byte, error)

var embedReadFile = []FuncType{
	ConfigDescribe,    // Global Setting
	ConfigJavaVersion, // Cluster Setting
	ConfigLocationMap, // Location change
	COnfigStorageMap,  // Storage Setting
}

func GetYamlFile(controller int) FuncType {
	if controller >= 0 && controller < FunctionNum {
		return embedReadFile[controller]
	}
	// 如果索引超出範圍，返回一個錯誤函式
	return func() ([]byte, error) {
		return nil, fmt.Errorf("invalid controller index")
	}
}
