package config

import (
	"fmt"
	"os"
)

// issue 表單 項目 包含 偷尾
const (
	first           = 0 // 這是第一個索引
	Owner           = 0
	Team            = 1
	ApNo            = 2
	LinuxapHost     = 3
	JavaVersion     = 4
	TeamRepo        = 5
	RepoBranch      = 6
	Environment     = 7
	ClusterLocation = 8
	CrontabTime     = 9
	DrRun           = 10
	end             = 11 // 這是最後一個索引
)

// OutputItems issue 表單輸出 workflow 項目
var OutputItems = []string{
	"owner",
	"team",
	"apnum",
	"host",
	"java",
	"repo",
	"branch",
	"env",
	"location",
	"time",
	"suspend",
}

// GetNameSpace 獲取環境變數值
func GetNameSpace(env, team, name string) string {
	EnvMap := map[string]string{
		"lab":  "l",
		"stg":  "s",
		"prod": "p",
		"dr":   "p",
	}
	envValue := fmt.Sprintf("%s-%s-ap-%s", EnvMap[env], team, name)
	return envValue
}

func GetFormListValue() map[int]string {
	FormList := map[int]string{
		Owner:           "Project Owner",
		Team:            "Team Name",
		ApNo:            "AP No.",
		LinuxapHost:     "Linuxap Host",
		JavaVersion:     "Java Version",
		TeamRepo:        "Team Repo",
		RepoBranch:      "Repo Branch",
		Environment:     "Env Location",
		ClusterLocation: "Cluster Location",
		CrontabTime:     "Crontab Time",
		DrRun:           "DR RUN",
	}
	return FormList
}

func OutputFirst() int {
	return first
}

func OutputEnd() int {
	return end
}

// 設定映射設置環境變數
func getEnvMapping(configValue int) map[string]string {
	envMappings := []map[string]string{
		{
			"Business":    "VAR1",
			"Description": "VAR2",
			"Owner":       "VAR3",
			"Repository":  "VAR4",
			"AES":         "VAR5",
		},
		{
			"JavaVersion": "VAR1",
			"Team":        "VAR2",
			"Time":        "VAR3",
			"APNum":       "APNUM",
			"Location":    "VAR4",
			"DrRun":       "VAR6",
		},
		{
			"Team":     "team",
			"Name":     "app",
			"Location": "cluster",
		},
		{
			"Env":     "env",
			"Team":    "team",
			"Name":    "name",
			"Volume":  "volume",
			"Bu":      "bu",
			"Server":  "server",
			"Path":    "path",
			"Linuxap": "linuxap",
		},
	}
	return envMappings[configValue]
}

// SetEnvVars 提供映射設置環境變數
func SetEnvVars(vars map[string]string, configValue int) error {
	envMapping := getEnvMapping(configValue)
	for key, envVar := range envMapping {
		value, exists := vars[key]
		if !exists {
			return fmt.Errorf("警告: 沒有找到 '%s' 對應的值", key)
		}
		if err := os.Setenv(envVar, value); err != nil {
			return fmt.Errorf("設置環境變數 %s 失敗: %v", envVar, err)
		}
		fmt.Printf("設置環境變數 %s 成功: %s\n", envVar, value)
	}
	return nil
}

// set Env Value Controller
const (
	ConfigGlobal    = 0
	ConfigCluster   = 1
	ConfigLocatipn  = 2
	ConfigNamespace = 2
	Configpvc       = 3
)

// Description 預設所有 linux ap 的描述都為此 default 值, 需要變更可以請產品單位自己修改 PR
const Description = "linux ap cronjob" // 將 "linux ap cronjob" 使用常量變數取代
