# YApi MCP Server 使用说明

## 项目简介

YApi MCP Server 是一个 Model Context Protocol 服务器，用于从 YApi 文档链接读取接口信息并提供给大模型使用。

## 快速开始

### 1. 构建项目

```bash
cd /Users/makeblock/Desktop/plugin/yapi-mcp-server
go build -o bin/yapi-mcp-server .
```

### 2. 在编辑器中使用

#### 在 Claude Desktop 中使用

```bash
./setup_claude.sh
```

详细说明：[Claude_Desktop配置指南.md](./Claude_Desktop配置指南.md)

#### 在 Cursor 编辑器中使用

```bash
./setup_cursor.sh
```

详细说明：[Cursor配置指南.md](./Cursor配置指南.md)

## 配置要求

需要设置两个环境变量：
- `YAPI_BASE_URL`: YApi 实例地址（如 https://yapi.makeblock.com）
- `YAPI_TOKEN`: YApi 访问 Token

## 可用工具

### 1. get_yapi_interface
获取单个 YApi 接口的详细信息

**使用示例：**
```
请帮我获取这个 YApi 接口信息：
https://yapi.makeblock.com/project/382/interface/api/17888
```

**返回信息包括：**
- 接口基本信息（ID、标题、路径、方法）
- 请求参数（req_query）
- 请求头（req_headers）
- 请求体 Schema（req_body_other）
- 响应体 Schema（res_body）
- 其他元数据

### 2. get_yapi_project_interfaces
获取 YApi 项目中所有接口的列表

**使用示例：**
```
请列出项目 382 的所有接口
```

## 文档索引

- 📖 [README.md](./README.md) - 项目总览
- 🚀 [QUICKSTART.md](./QUICKSTART.md) - 快速入门
- 🔧 [Claude_Desktop配置指南.md](./Claude_Desktop配置指南.md) - Claude Desktop 配置
- 💻 [Cursor配置指南.md](./Cursor配置指南.md) - Cursor 编辑器配置
- 🔐 [TOKEN设置说明.md](./TOKEN设置说明.md) - Token 获取和设置
- 🧪 [本地调试指南.md](./本地调试指南.md) - 本地调试方法
- 📦 [发布指南.md](./发布指南.md) - npm 发布指南
- 🧪 [快速测试.md](./快速测试.md) - 测试命令参考

## 测试脚本

- `test.sh` - 完整测试脚本
- `test_single.sh` - 单个接口测试
- `测试Token.sh` - 测试 Token 是否有效
- `setup_claude.sh` - Claude Desktop 自动配置
- `setup_cursor.sh` - Cursor 编辑器自动配置
- `快速设置Token.sh` - 交互式 Token 设置

## 项目结构

```
yapi-mcp-server/
├── bin/
│   └── yapi-mcp-server           # 编译后的二进制文件
├── main.go                        # MCP 服务器主程序
├── yapi_parser.go                 # YApi 解析器
├── main_test.go                   # 单元测试
├── yapi_parser_test.go           # 解析器测试
├── go.mod                         # Go 依赖
├── package.json                   # npm 包配置
├── README.md                      # 项目说明
└── 各种配置和测试脚本...
```

## 常见问题

### Q1: Token 在哪里设置？
A: Token 通过环境变量 `YAPI_TOKEN` 设置，详见 [TOKEN设置说明.md](./TOKEN设置说明.md)

### Q2: 如何获取 YApi Token？
A: 登录 YApi → 项目设置 → Token 配置

### Q3: 支持哪些编辑器？
A: 目前支持：
- Claude Desktop
- Cursor 编辑器
- 任何支持 MCP 协议的工具

### Q4: 如何本地测试？
A: 运行 `./test.sh` 或查看 [本地调试指南.md](./本地调试指南.md)

### Q5: 返回的数据是空的怎么办？
A: 检查：
1. Token 是否正确
2. YAPI_BASE_URL 是否正确
3. 接口 URL 格式是否正确

## 技术栈

- **语言**: Go 1.21+
- **协议**: Model Context Protocol (MCP)
- **依赖**: 
  - github.com/spf13/cobra - CLI 框架
  - golang.org/x/net - HTML 解析

## 版本信息

- **当前版本**: 1.0.0
- **Go 版本要求**: 1.21+
- **Node 版本要求**: 14.0+ (如果使用 npm)

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

