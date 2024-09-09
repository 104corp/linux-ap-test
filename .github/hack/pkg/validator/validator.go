package validator

import (
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"hack/pkg/config"
	"hack/schema"
)

func ValidateRequestForm(apRequest []map[int]string) error {
	for key, value := range apRequest {
		documentMap := map[string]string{
			config.OutputItems[key]: value[key],
		}
		if err := validateRequestForm(schema.GetSchemaPath(key), documentMap); err != nil {
			return err
		}
	}
	return nil
}

func validateRequestForm(schemaFile string, documentMap map[string]string) error {
	println(schemaFile)
	schemaJSON, err := schema.EmbedFiles.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("JSON 讀檔編碼失敗: %v", err)
	}

	documentJSON, err := json.Marshal(documentMap)
	print(string(documentJSON))
	if err != nil {
		return fmt.Errorf("JSON 編碼失敗: %v", err)
	}
	schemaLoader := gojsonschema.NewStringLoader(string(schemaJSON))
	documentLoader := gojsonschema.NewStringLoader(string(documentJSON))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("JSON 驗證失敗: %v", err)
	}

	if !result.Valid() {
		var errorMessages string
		for _, desc := range result.Errors() {
			errorMessages += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf("JSON 文件驗證失敗: \n%s", errorMessages)
	}
	return nil
}
