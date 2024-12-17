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

