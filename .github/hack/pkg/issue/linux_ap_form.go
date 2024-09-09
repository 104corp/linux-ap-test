package issue

import (
	"fmt"
	"hack/pkg/config"
	"strings"
)

func ParseLinuxApForm(apRequest *[]map[int]string, isu *Issue) error {
	lines := strings.Split(isu.Body, "\n")
	foundKeywords := make(map[int]bool) // 用於追踪哪些關鍵字已被找到

	for i, line := range lines {
		for keyword, value := range config.GetFormListValue() {
			if strings.Contains(line, value) {
				// 確保這不是最後兩行，以防越界
				if i+2 < len(lines) {
					// 有抓取
					foundKeywords[keyword] = true
					fmt.Printf("the keyword: %x, the value: %s \n", keyword, value)
					// 返回下下一行的字串，並移除首尾的空白字符
					*apRequest = append(
						*apRequest,
						map[int]string{keyword: strings.TrimSpace(lines[i+2])},
					)
				}
			}
		}
	}
	// 檢查是否所有關鍵字都已被找到

	for k := config.OutputFirst(); k < config.OutputEnd(); k++ {
		if !foundKeywords[k] {
			err := fmt.Errorf("issue 表單缺少必要的配置信息: %s", config.GetFormListValue()[k])
			return err
		}
	}

	return nil
}
