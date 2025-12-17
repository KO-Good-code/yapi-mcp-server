# 快速开始指南

## 1. 本地开发和测试

### 安装Go依赖

```bash
go mod download
go mod tidy
```

### 运行程序

```bash
# 直接运行
go run .

# 或者构建后运行
go build -o bin/yapi-mcp-server .
./bin/yapi-mcp-server
```

### 测试MCP服务器

MCP服务器通过stdio通信，可以通过以下方式测试：

```bash
# 发送测试请求
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | go run .
```

## 2. 发布到私有npm仓库

### 步骤1: 配置私有仓库

创建 `.npmrc` 文件：

```bash
@neigri:registry=https://your-private-npm-registry.com
//your-private-npm-registry.com/:_authToken=your-token-here
```

### 步骤2: 更新package.json

修改以下字段：

```json
{
  "name": "@neigri/yapi-mcp-server",
  "version": "1.0.0",
  "repository": {
    "type": "git",
    "url": "git+https://your-private-repo-url.git"
  },
  "publishConfig": {
    "registry": "https://your-private-npm-registry.com"
  }
}
```

### 步骤3: 构建和发布

```bash
# 方式1: 使用npm脚本（自动构建）
npm install
npm publish

# 方式2: 使用Makefile
make build
npm publish

# 方式3: 手动构建
go build -o bin/yapi-mcp-server .
npm publish
```

## 3. 在Claude Desktop中使用

### 配置MCP服务器

编辑Claude Desktop配置文件（macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`）：

```json
{
  "mcpServers": {
    "yapi-mcp-server": {
      "command": "npx",
      "args": [
        "-y",
        "@neigri/yapi-mcp-server@latest"
      ],
      "env": {
        "YAPI_BASE_URL": "http://your-yapi-instance.com",
        "YAPI_TOKEN": "your-yapi-token"
      }
    }
  }
}
```

### 重启Claude Desktop

保存配置后，重启Claude Desktop即可使用。

## 4. 使用示例

### 获取单个接口信息

```
工具: get_yapi_interface
参数: {
  "url": "http://yapi.example.com/project/123/interface/api/456"
}
```

### 获取项目所有接口

```
工具: get_yapi_project_interfaces
参数: {
  "project_url": "http://yapi.example.com/project/123"
}
```

## 5. 环境变量

- `YAPI_BASE_URL`: YApi实例的基础URL（可选）
- `YAPI_TOKEN`: YApi访问令牌（可选，用于API访问）

## 6. 故障排除

### 问题1: npx无法找到包

**解决方案：**
- 确保npm配置正确指向私有仓库
- 检查包是否已成功发布
- 验证scope和包名是否正确

### 问题2: 构建失败

**解决方案：**
- 确保已安装Go 1.21或更高版本
- 运行 `go mod download` 安装依赖
- 检查网络连接

### 问题3: MCP服务器无法启动

**解决方案：**
- 检查Claude Desktop配置文件格式
- 验证命令路径是否正确
- 查看Claude Desktop日志

## 7. 开发建议

- 使用 `make build` 快速构建
- 使用 `make build-all` 构建多平台版本
- 使用 `go test ./...` 运行测试
- 使用 `go fmt ./...` 格式化代码

