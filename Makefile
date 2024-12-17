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

