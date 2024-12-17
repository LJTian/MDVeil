# 起因：

我在工作中遇到了一个问题, 我最近工作需要输出一个文档，这个文档里面有一些内容是不可对外展示的。这种情况下，一般我需要写两份文档，对于我来说，有点不太能接受。因为两份文档就表明了，我以后的修改内容都需要进行两次修改，工作量翻一倍。所以我需要一个程序将第二个文档自动生成出来。

# 思考：

将问题提纯：

如何将 **markdown** 文件内的内容进行分块标记，并且让程序自动识别进行隐藏？

# 生成代码：

询问 ChatGPT:

问题1：MD 文件一般是什么格式？

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/ec500d2c-c8f8-4fff-b940-f5b099d03b65/432e4949-da59-4e66-af5d-57f8aaddf57a/image.png)

问题2：我想用标记语言方式将Markdown 的内容分为对外和对内，被标记语言标注的内容、对外展示时隐藏，用 go 语言写一下

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/ec500d2c-c8f8-4fff-b940-f5b099d03b65/6a68fd6f-e1c7-4ce5-84bb-8728928cf921/image.png)

问题3：我希望它可以标记多行，而且标记语言的符号简单一些

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/ec500d2c-c8f8-4fff-b940-f5b099d03b65/2dc118cd-dc5b-4e81-9cef-4af8a3d274e8/image.png)

问题4：修改一下，markdownContent 的内容是从参数 -I 获取，此参数为文档名称 参数 -O 是文档输出名称 参数 -S 为是否展示隐藏内容

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/ec500d2c-c8f8-4fff-b940-f5b099d03b65/8d5402f0-0e80-4a1b-985d-df098d6742fb/image.png)

问题5：使用 cobra

问题6：将函数提取一下，我希望以后的命令修改不涉及到main.go 文件

问题7：生成 makefile 文件，支持生成各个架构命令

问题8：整理上述内容，给我一个目录结构以及各个文件的内容

> 内容就不展示了，浪费流量。
> 

## 最终结果：

### 目录结构

```bash
markdown-processor/
├── Makefile
├── main.go
└── markdown/
    └── processor.go
```

### Makefile：用于构建不同架构的 Go 应用程序。

```makefile
# 定义变量
APP_NAME = markdownProcessor
SRC = main.go markdown/processor.go
OUTPUT_DIR = bin
BUILD_DIR = $(OUTPUT_DIR)/$(APP_NAME)

# 默认目标
all: build

# 构建目标
build: linux-amd64 linux-arm64 darwin-amd64 windows-amd64

# 构建 Linux AMD64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)-linux-amd64 $(SRC)

# 构建 Linux ARM64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)-linux-arm64 $(SRC)

# 构建 macOS AMD64
darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)-darwin-amd64 $(SRC)

# 构建 Windows AMD64
windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)-windows-amd64.exe $(SRC)

# 清理目标
clean:
	rm -rf $(OUTPUT_DIR)

# 显示帮助信息
help:
	@echo "Makefile for building markdownProcessor"
	@echo "Usage:"
	@echo "  make all           # Build all targets"
	@echo "  make clean         # Clean up build files"
	@echo "  make help          # Show this help message"

```

### main.go：主程序，负责命令行解析和调用处理函数。

```go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"markdown-processor/markdown"
)

func main() {
	var inputFile string
	var outputFile string
	var showHidden bool

	var rootCmd = &cobra.Command{
		Use:   "markdownProcessor",
		Short: "处理 Markdown 文档",
		Run: func(cmd *cobra.Command, args []string) {
			// 处理 Markdown 内容
			result, err := markdown.ProcessMarkdown(inputFile, outputFile, showHidden)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("处理完成，输出文件:", outputFile)
			fmt.Println(result) // 可选：输出处理结果
		},
	}

	rootCmd.Flags().StringVarP(&inputFile, "input", "I", "", "输入文档名称")
	rootCmd.Flags().StringVarP(&outputFile, "output", "O", "", "输出文档名称")
	rootCmd.Flags().BoolVarP(&showHidden, "show", "S", false, "是否展示隐藏内容")

	// 检查必需的参数
	cobra.OnInitialize(func() {
		if inputFile == "" || outputFile == "" {
			fmt.Println("请提供输入和输出文档名称.")
			os.Exit(1)
		}
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

```

### markdown/processor.go：处理 Markdown 文件的逻辑。

```go
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

```

# 调试代码：

大部分代码已经编写完成了，如果你有一定的 linux shell 基础，就能看出，这个实现的逻辑不过就是正则表达式替换功能，其实就是一条 sed 命令的事，很失望的结果。另外：linux 三剑客（awk、sed、grep ）真是强大。

 

![image.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/ec500d2c-c8f8-4fff-b940-f5b099d03b65/4d7270d7-b550-4d59-930f-95c9d2102015/image.png)

由于 window 存在我觉得还是用吧。但是我懒得调试了,只把最终代码提交上去了，等有时间我再调试一下。

> 调试就是验证一下程序，如果有问题的话，进行一下修改。这部分就得靠自己了
> 

# 提交 GitHub:

地址：https://github.com/LJTian/MDVeil

# 号外：这个 [内] 标记符，我还是不太喜欢，如果应该用 [H] 更简单一些。
