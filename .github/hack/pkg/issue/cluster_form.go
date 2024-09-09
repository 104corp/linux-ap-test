package issue

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"hack/clusterMap"
	"hack/pkg/config"
)

type configYAML struct {
	Cluster map[string]string `yaml:"cluster"`
}

func SetCluster(apRequest *[]map[int]string, env string) error {
	fmt.Printf("the env: %s\n", env)
	cluster := (*apRequest)[config.ClusterLocation][config.ClusterLocation]
	fmt.Printf("the cluster: %s\n", cluster)
	var yamlConfig configYAML
	data, err := embedFS.GetYamlFile(env)()
	if err != nil {
		errors := fmt.Errorf("failed to read config file: %v", err)
		return errors
	}

	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		errors := fmt.Errorf("error: %v", err)
		return errors
	}
	fmt.Printf("the yaml data:\n %s \n", data)
	fmt.Printf("the yaml Config: %+v \n", yamlConfig)
	(*apRequest)[config.ClusterLocation][config.ClusterLocation] = yamlConfig.Cluster[cluster]
	return nil
}
