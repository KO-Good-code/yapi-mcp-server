# YApi MCP Server (Go版本)

一个用Go语言实现的Model Context Protocol (MCP) 服务器，用于从YApi文档链接读取接口信息并提供给大模型使用。

## 特性

- 🚀 使用Go实现，性能优异
- 📦 支持通过npm/npx发布和使用
- 📖 从YApi文档链接读取单个接口信息
- 📚 获取YApi项目中所有接口列表
- 🔍 在YApi项目中搜索接口
- 🔌 支持YApi API和HTML页面解析

## 快速开始

### 方式1: 通过npx使用（推荐）

```bash
npx -y @neigri/yapi-mcp-server@latest
```

### 方式2: 本地构建

```bash
# 安装Go依赖
go mod download

# 构建
make build
# 或
go build -o bin/yapi-mcp-server .

# 运行
./bin/yapi-mcp-server
```

## 在Claude Desktop中配置

在Claude Desktop的配置文件中添加：

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
        "YAPI_TOKEN": "your-token-optional"
      }
    }
  }
}
```

如果发布到私有npm仓库：

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
        "YAPI_TOKEN": "your-token-optional"
      }
    }
  }
}
```

## 可用工具

MCP服务器提供以下工具：

### 1. get_yapi_interface

从YApi文档链接读取单个接口信息。

**参数：**
- `url` (string, 必需): YApi文档链接

### 2. get_yapi_project_interfaces

获取YApi项目中所有接口的列表。

**参数：**
- `project_url` (string, 必需): YApi项目URL或项目ID

## 环境变量

- `YAPI_BASE_URL`: YApi实例的基础URL（可选）
- `YAPI_TOKEN`: YApi访问令牌（可选，用于API访问）

## 项目结构

```
yapi-mcp-server/
├── main.go              # 主程序入口
├── yapi_parser.go       # YApi解析器实现
├── go.mod               # Go模块定义
├── package.json         # npm包配置
├── Makefile            # 构建脚本
├── scripts/
│   └── build.js        # npm构建脚本
├── README.md           # 项目说明
├── QUICKSTART.md       # 快速开始指南
└── 发布指南.md         # 发布指南
```

## 开发

### 本地开发

```bash
# 安装Go依赖
go mod download
go mod tidy

# 运行
go run .

# 构建
make build

# 构建多平台
make build-all
```

### 测试

```bash
# 运行测试（需要先编写测试）
go test ./...
```

## 发布到私有npm仓库

详细步骤请参考 [发布指南.md](./发布指南.md)

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

