package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const FunctionNum = 4

func RunPWDCmd() error {
	cmd := exec.Command("pwd")
	// 捕獲命令的輸出
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("執行命令時出錯: %v", err)
	}
	// 輸出命令的結果
	fmt.Print(string(output))
	return nil
}

func runGobalCommand() error {
	cmd := "cat .github/hack/Template/app/values.yaml | envsubst '$VAR1 $VAR2 $VAR3 $VAR4 $VAR5' > output.yaml"
	execCmd := exec.Command("bash", "-c", cmd)
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}
	return nil
}

func runClusterCommand() error {

	envVar := "$JAVA_HOME"
	value := "$JAVA_HOME"
	// 執行命令 envsubst 生成複製 cluster.values.yaml 需要文件
	if err := os.Setenv(envVar, value); err != nil {
		return fmt.Errorf("設置環境變數 %s 失敗: %v", envVar, err)
	} else {
		fmt.Printf("設置環境變數 %s 成功: %s\n", envVar, value)
	}

	cmd := "pwd"
	err := runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "ls -al .github/hack/Template/start.sh"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "cat .github/hack/Template/start.sh"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "echo $APNUM"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "echo $JAVA_HOME"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "cat .github/hack/Template/start.sh|envsubst '$APNUM $JAVA_HOME' > content.sh"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	envVar = "VAR5"
	value = "${SCRIPT_CONTENT}"
	if err := os.Setenv(envVar, value); err != nil {
		return fmt.Errorf("設置環境變數 %s 失敗: %v", envVar, err)
	}
	cmd = "cat .github/hack/Template/app/cluster.values.yaml | envsubst '$VAR1 $VAR2 $VAR3 $VAR4 $VAR5 $VAR6' > output.yaml"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}

	return nil
}

func runNamespaceCommand() error {
	// 執行命令 envsubst 生成複製 namespace.values.yaml 需要文件
	// runNamespaceApply()
	err := runNamespaceGenerate()

	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}

	return nil
}

func runNamespaceGenerate() error {
	cmd := "cat .github/hack/github/script/goamNamespace.sh"
	err := runCmd(cmd)

	cmd = "bash .github/hack/github/script/goamNamespace.sh"
	err = runCmd(cmd)

	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	return nil
}

func runNamespaceApply() error {
	var cmd string
	var err error

	cmd = "cat .github/hack/github/templates/namespace.md"
	err = runCmd(cmd)

	cmd = "cat .github/hack/github/templates/namespace.md | envsubst '$team $app $cluster' > template.md"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = ".github/hack/github/script/createNamespace.sh"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}
	return nil
}

func runStorageCommand() error {
	//err := runStorageApply()

	err := runStorageGenerate()

	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}
	return nil
}

func runStorageGenerate() error {
	cmd := "ls -al"
	err := runCmd(cmd)

	cmd = ".github/hack/github/script/goamStorage.sh"
	err = runCmd(cmd)

	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}

	return nil
}

func runStorageApply() error {
	var cmd string
	var err error

	cmd = "cat .github/hack/github/templates/storage.md | envsubst '$env $team $name $volume $bu $server $path' > temp.md"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "cat temp.md | envsubst '$linuxap' > $volume.md"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	cmd = "echo $volume"
	err = runCmd(cmd)

	cmd = ".github/hack/github/script/createStorage.sh $volume"
	err = runCmd(cmd)
	if err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	} else {
		fmt.Printf("執行命令成功: %s\n", cmd)
	}

	return nil
}

func runCmd(msg string) error {
	cmd := msg
	execCmd := exec.Command("bash", "-c", cmd)
	execCmd.Stderr = os.Stderr

	var out bytes.Buffer
	execCmd.Stdout = &out

	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("執行命令失敗: %v", err)
	}

	println(out.String())
	return nil
}

type FuncType func() error

var runCommandList = []FuncType{
	runGobalCommand,
	runClusterCommand,
	runNamespaceCommand,
	runStorageCommand,
}

func RunCommandToFile(controller int) FuncType {
	if controller >= 0 && controller < FunctionNum {
		return runCommandList[controller]
	}
	// 如果索引超出範圍，返回一個錯誤函式
	return func() error {
		return fmt.Errorf("invalid controller index")
	}
}
