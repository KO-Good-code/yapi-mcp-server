.PHONY: build build-all clean install test

# 构建当前平台的二进制文件
build:
	@echo "构建当前平台..."
	@mkdir -p bin
	@go build -o bin/yapi-mcp-server .

# 构建所有平台的二进制文件
build-all:
	@echo "构建所有平台..."
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/yapi-mcp-server-linux-amd64 .
	@GOOS=linux GOARCH=arm64 go build -o bin/yapi-mcp-server-linux-arm64 .
	@GOOS=darwin GOARCH=amd64 go build -o bin/yapi-mcp-server-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o bin/yapi-mcp-server-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o bin/yapi-mcp-server-windows-amd64.exe .
	@echo "构建完成！"

# 清理构建文件
clean:
	@rm -rf bin/
	@echo "清理完成"

# 安装依赖
install:
	@go mod download
	@go mod tidy

# 运行测试
test:
	@go test ./...

# 运行程序
run:
	@go run .

