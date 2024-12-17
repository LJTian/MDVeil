package markdown

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// ProcessMarkdown 处理 Markdown 文件，返回处理后的内容
func ProcessMarkdown(inputFile string, outputFile string, showHidden bool) (string, error) {
	// 读取输入文件内容
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %v", err)
	}

	var result string
	if showHidden {
		result = string(content) // 如果选择展示隐藏内容，直接返回原内容
	} else {
		// 使用正则表达式匹配对内内容
		re := regexp.MustCompile(`\[\[内\]\].*?\[\[/内\]\]`)
		// 将对内内容替换为空字符串
		result = re.ReplaceAllString(string(content), "")
	}

	// 将处理后的内容写入输出文件
	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return "", fmt.Errorf("写入文件失败: %v", err)
	}

	return result, nil
}

