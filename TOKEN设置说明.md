# YApi Token 设置说明

## Token的作用

YApi Token用于访问需要认证的YApi实例。如果YApi实例是公开的或者不需要认证，可以不设置Token。

## 设置Token的方法

### 方法1: 临时设置（当前终端会话有效）

```bash
export YAPI_TOKEN="your-token-here"
```

设置后，在当前终端运行测试：

```bash
./test.sh
# 或
go run .
```

### 方法2: 在测试命令中直接设置

```bash
YAPI_TOKEN="your-token-here" ./test.sh
```

或测试单个命令：

```bash
YAPI_TOKEN="your-token-here" echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"get_yapi_interface","arguments":{"url":"http://yapi.example.com/project/123/interface/api/456"}}}' | go run .
```

### 方法3: 永久设置（推荐用于开发）

创建 `.env` 文件（注意：不要提交到Git）：

```bash
# .env 文件
YAPI_BASE_URL=http://your-yapi-instance.com
YAPI_TOKEN=your-token-here
```

然后在运行前加载：

```bash
# 加载环境变量
source .env
# 或
export $(cat .env | xargs)

# 运行测试
./test.sh
```

### 方法4: 在shell配置文件中设置（全局）

编辑 `~/.bashrc` 或 `~/.zshrc`：

```bash
export YAPI_BASE_URL="http://your-yapi-instance.com"
export YAPI_TOKEN="your-token-here"
```

然后重新加载配置：

```bash
source ~/.zshrc  # 或 source ~/.bashrc
```

## 如何获取YApi Token

### 步骤1: 登录YApi系统

访问你的YApi实例，使用账号登录。

### 步骤2: 进入项目设置

1. 选择你要访问的项目
2. 点击项目设置（通常在右上角或项目菜单中）
3. 找到"Token配置"或"接口设置"选项

### 步骤3: 获取Token

- 如果已有Token，直接复制
- 如果没有Token，点击"生成Token"或"创建Token"
- 复制生成的Token字符串

### 步骤4: 设置Token

```bash
export YAPI_TOKEN="你复制的token"
```

## 验证Token是否设置成功

### 方法1: 检查环境变量

```bash
echo $YAPI_TOKEN
```

如果输出你的token，说明设置成功。

### 方法2: 运行测试

```bash
# 设置token和base URL
export YAPI_BASE_URL="http://your-yapi-instance.com"
export YAPI_TOKEN="your-token"

# 运行测试
./test.sh "http://your-yapi-instance.com/project/123/interface/api/456"
```

如果能够成功获取接口信息，说明token设置正确。

## 常见问题

### Q1: Token是必需的吗？

**A:** 不是。Token是可选的：
- 如果YApi实例是公开的，不需要Token
- 如果YApi实例需要认证，则需要设置Token
- 如果不设置Token，程序会尝试从HTML页面解析（可能不如API方式准确）

### Q2: Token在哪里使用？

**A:** Token在以下场景使用：
- 调用 `get_yapi_interface` 工具时，如果提供了Token，会优先使用YApi API获取接口信息
- 调用 `get_yapi_project_interfaces` 工具时，如果提供了Token，会使用API获取项目接口列表

### Q3: Token安全吗？

**A:** 注意事项：
- Token具有访问权限，请妥善保管
- 不要将Token提交到Git仓库
- 建议将 `.env` 文件添加到 `.gitignore`
- 如果Token泄露，及时在YApi中重新生成

### Q4: 如何测试Token是否有效？

**A:** 运行以下命令：

```bash
export YAPI_BASE_URL="http://your-yapi-instance.com"
export YAPI_TOKEN="your-token"

# 测试获取接口信息
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"get_yapi_interface","arguments":{"url":"http://your-yapi-instance.com/project/123/interface/api/456"}}}' | go run .
```

如果返回接口信息而不是错误，说明Token有效。

## 示例：完整的测试流程

```bash
# 1. 进入项目目录
cd yapi-mcp-server

# 2. 设置环境变量
export YAPI_BASE_URL="http://yapi.yourcompany.com"
export YAPI_TOKEN="abc123def456ghi789"

# 3. 运行测试脚本
./test.sh "http://yapi.yourcompany.com/project/123/interface/api/456"

# 4. 或者手动测试
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | go run .
```

## 在Claude Desktop中配置Token

在Claude Desktop配置文件中：

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
        "YAPI_TOKEN": "your-token-here"
      }
    }
  }
}
```

保存后重启Claude Desktop即可。

