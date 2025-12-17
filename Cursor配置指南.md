# Cursor 编辑器 MCP 配置指南

## 前提条件

✅ 已构建好二进制文件：`/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server`

## 配置步骤

### 1. 找到 Cursor 配置文件

Cursor 的 MCP 配置文件位置（macOS）：
```
~/.cursor/mcp_settings.json
```

或者：
```
~/Library/Application Support/Cursor/User/globalStorage/mcp_settings.json
```

### 2. 创建或编辑配置文件

如果文件不存在，创建它：

```bash
# 创建目录
mkdir -p ~/.cursor

# 创建配置文件
touch ~/.cursor/mcp_settings.json
```

### 3. 添加 MCP 服务器配置

编辑配置文件，添加以下内容：

```json
{
  "mcpServers": {
    "yapi-mcp-server": {
      "command": "/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server",
      "env": {
        "YAPI_BASE_URL": "https://yapi.makeblock.com",
        "YAPI_TOKEN": "your-yapi-token-here"
      }
    }
  }
}
```

### 4. 使用 VS Code 打开配置文件

```bash
code ~/.cursor/mcp_settings.json
```

或在 Cursor 中：
1. 按 `Cmd+Shift+P` 打开命令面板
2. 输入 "MCP Settings"
3. 选择 "Edit MCP Settings"

### 5. 重启 Cursor

配置完成后，完全退出并重启 Cursor：
- macOS: `Cmd+Q` 退出
- 重新打开 Cursor

## Cursor 内置方式配置

Cursor 可能也支持通过 UI 界面配置 MCP：

### 方法1: 通过设置界面

1. 打开 Cursor Settings（`Cmd+,`）
2. 搜索 "MCP" 或 "Model Context Protocol"
3. 找到 MCP 服务器配置选项
4. 添加新的 MCP 服务器

### 方法2: 通过命令面板

1. 按 `Cmd+Shift+P` 打开命令面板
2. 输入 "MCP"
3. 选择 "Add MCP Server" 或类似选项
4. 按提示输入配置信息

## 完整配置示例

```json
{
  "mcpServers": {
    "yapi-mcp-server": {
      "command": "/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server",
      "args": [],
      "env": {
        "YAPI_BASE_URL": "https://yapi.makeblock.com",
        "YAPI_TOKEN": "your-yapi-token-here"
      },
      "disabled": false
    }
  }
}
```

## 自动配置脚本

创建 `setup_cursor.sh` 脚本：

```bash
#!/bin/bash

# Cursor MCP 自动配置脚本

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}=== Cursor MCP 配置 ===${NC}\n"

# Cursor 配置文件路径（尝试两个可能的位置）
CONFIG_FILE1=~/.cursor/mcp_settings.json
CONFIG_FILE2=~/Library/Application\ Support/Cursor/User/globalStorage/mcp_settings.json

# 确定使用哪个配置文件
if [ -f "$CONFIG_FILE2" ]; then
    CONFIG_FILE="$CONFIG_FILE2"
else
    CONFIG_FILE="$CONFIG_FILE1"
    mkdir -p ~/.cursor
fi

# 二进制文件路径
BINARY_PATH="/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server"

# 检查二进制文件
if [ ! -f "$BINARY_PATH" ]; then
    echo "错误: 二进制文件不存在"
    exit 1
fi

chmod +x "$BINARY_PATH"

# 获取配置信息
read -p "请输入 YAPI_TOKEN (直接回车使用默认): " YAPI_TOKEN
if [ -z "$YAPI_TOKEN" ]; then
    YAPI_TOKEN="f5bb438ab7b19cc117f127cd5b05c234a20d21178dd117735c1752affe0ba4d6"
fi

read -p "请输入 YAPI_BASE_URL (直接回车使用 https://yapi.makeblock.com): " YAPI_BASE_URL
if [ -z "$YAPI_BASE_URL" ]; then
    YAPI_BASE_URL="https://yapi.makeblock.com"
fi

# 备份现有配置
if [ -f "$CONFIG_FILE" ]; then
    cp "$CONFIG_FILE" "$CONFIG_FILE.backup.$(date +%Y%m%d_%H%M%S)"
    echo -e "${YELLOW}已备份现有配置${NC}"
fi

# 创建配置
cat > "$CONFIG_FILE" << EOF
{
  "mcpServers": {
    "yapi-mcp-server": {
      "command": "$BINARY_PATH",
      "env": {
        "YAPI_BASE_URL": "$YAPI_BASE_URL",
        "YAPI_TOKEN": "$YAPI_TOKEN"
      }
    }
  }
}
EOF

echo -e "${GREEN}✓ Cursor MCP 配置完成${NC}"
echo "配置文件: $CONFIG_FILE"
echo ""
echo "下一步："
echo "1. 重启 Cursor (Cmd+Q 后重新打开)"
echo "2. 在 Cursor AI 聊天中测试"
echo ""
```

保存并运行：
```bash
chmod +x setup_cursor.sh
./setup_cursor.sh
```

## 在 Cursor 中使用

配置完成后，在 Cursor 的 AI 聊天窗口中：

### 方法1: 通过聊天

直接在聊天中询问：
```
请帮我获取这个 YApi 接口的信息：
https://yapi.makeblock.com/project/382/interface/api/17888
```

### 方法2: 通过 @ 符号

在聊天中使用 `@` 符号：
```
@yapi-mcp-server 获取接口信息: https://yapi.makeblock.com/project/382/interface/api/17888
```

### 方法3: 通过命令

按 `Cmd+K` 或 `Cmd+L` 打开 AI 面板，然后输入：
```
使用 YApi MCP 工具获取接口信息
```

## 可用工具

1. **get_yapi_interface**
   - 功能：获取单个 YApi 接口的详细信息
   - 参数：url (YApi 接口链接)

2. **get_yapi_project_interfaces**
   - 功能：获取 YApi 项目的所有接口列表
   - 参数：project_url (YApi 项目链接)

## 使用示例

### 示例1: 获取接口详情

```
👤 你：请分析这个 API 接口：https://yapi.makeblock.com/project/382/interface/api/17888

🤖 Cursor AI 会：
   1. 使用 get_yapi_interface 工具
   2. 解析接口参数、返回值
   3. 提供接口分析和建议
```

### 示例2: 获取项目接口列表

```
👤 你：列出项目 382 的所有接口

🤖 Cursor AI 会：
   1. 使用 get_yapi_project_interfaces 工具
   2. 显示所有接口列表
   3. 可以进一步询问特定接口详情
```

### 示例3: 生成代码

```
👤 你：根据这个 YApi 接口生成 TypeScript 类型定义：
https://yapi.makeblock.com/project/382/interface/api/17888

🤖 Cursor AI 会：
   1. 获取接口信息
   2. 分析请求和响应结构
   3. 生成对应的 TypeScript 类型
```

## 验证配置

### 检查配置文件

```bash
cat ~/.cursor/mcp_settings.json
```

### 检查二进制文件

```bash
ls -la /Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server
```

### 测试二进制文件

```bash
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | /Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server
```

## 故障排查

### 问题1: Cursor 没有识别到 MCP

**解决方案：**
1. 确认配置文件位置正确
2. 检查 JSON 格式是否正确
3. 完全重启 Cursor
4. 查看 Cursor 开发者工具控制台（Help > Toggle Developer Tools）

### 问题2: 工具调用失败

**解决方案：**
1. 检查 Token 和 URL 配置
2. 确认二进制文件有执行权限
3. 查看 Cursor 日志

### 问题3: 找不到配置文件位置

**解决方案：**
```bash
# 搜索 Cursor 配置目录
find ~ -name "Cursor" -type d 2>/dev/null | head -20

# 或者直接在 Cursor 中打开
# Cmd+Shift+P -> "Open Settings (JSON)"
```

## 查看 Cursor 日志

```bash
# macOS
tail -f ~/Library/Logs/Cursor/*.log
```

## 更新配置

如果需要更新 Token 或 URL：
1. 编辑配置文件
2. 重启 Cursor

无需重新构建二进制文件。

## 与 Claude Desktop 的区别

| 特性 | Claude Desktop | Cursor |
|------|---------------|---------|
| 配置文件位置 | `~/Library/Application Support/Claude/` | `~/.cursor/` |
| 配置文件名 | `claude_desktop_config.json` | `mcp_settings.json` |
| 使用方式 | 独立应用聊天 | 编辑器内 AI 助手 |
| 工具调用 | 自动 | 自动或手动 @ |

## 高级用法

### 在代码编辑器中直接使用

1. 选中代码
2. 按 `Cmd+K` 或 `Cmd+L`
3. 输入：`使用 YApi 接口信息优化这段代码`

### 生成测试用例

```
根据 YApi 接口 https://yapi.makeblock.com/project/382/interface/api/17888
生成 Jest 测试用例
```

### 生成 API 文档

```
为这些 YApi 接口生成 Markdown 文档
```

