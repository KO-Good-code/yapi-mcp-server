# Claude Desktop 配置指南

## 前提条件

✅ 已构建好二进制文件：`/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server`

## 配置步骤

### 1. 找到 Claude Desktop 配置文件

配置文件位置（macOS）：
```
~/Library/Application Support/Claude/claude_desktop_config.json
```

### 2. 编辑配置文件

使用任意文本编辑器打开配置文件：

```bash
# 使用 VS Code 打开
code ~/Library/Application\ Support/Claude/claude_desktop_config.json

# 或使用 nano 打开
nano ~/Library/Application\ Support/Claude/claude_desktop_config.json

# 或使用 vim 打开
vim ~/Library/Application\ Support/Claude/claude_desktop_config.json
```

### 3. 添加 MCP 服务器配置

在配置文件中添加以下内容：

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

**注意事项：**
- 使用**绝对路径**指向二进制文件
- 替换 `YAPI_TOKEN` 为你的真实 token
- 如果文件已有其他 MCP 服务器配置，在 `mcpServers` 对象中添加即可

### 4. 如果已有其他 MCP 配置

如果配置文件已经存在并有其他 MCP 服务器，添加到现有配置中：

```json
{
  "mcpServers": {
    "existing-server": {
      "command": "...",
      "args": ["..."]
    },
    "yapi-mcp-server": {
      "command": "/Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server",
      "env": {
        "YAPI_BASE_URL": "https://yapi.makeblock.com",
        "YAPI_TOKEN": "your-token"
      }
    }
  }
}
```

### 5. 重启 Claude Desktop

配置完成后，**完全退出并重启 Claude Desktop**：
1. 点击菜单栏的 Claude 图标
2. 选择 "Quit Claude"（或按 Cmd+Q）
3. 重新打开 Claude Desktop

### 6. 验证配置

重启后，在 Claude Desktop 中：
1. 开始新对话
2. 检查是否可以看到 MCP 工具
3. 尝试使用工具

示例对话：
```
你：请帮我获取 YApi 接口信息：https://yapi.makeblock.com/project/382/interface/api/17888

Claude 会使用 get_yapi_interface 工具获取接口信息
```

## 完整配置示例

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

## 快速配置脚本

也可以使用以下脚本自动配置：

```bash
#!/bin/bash

# Claude Desktop 配置文件路径
CONFIG_FILE=~/Library/Application\ Support/Claude/claude_desktop_config.json

# 创建目录（如果不存在）
mkdir -p ~/Library/Application\ Support/Claude

# MCP 服务器配置
cat > "$CONFIG_FILE" << 'EOF'
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
EOF

echo "✓ Claude Desktop 配置已更新"
echo "请重启 Claude Desktop 以应用更改"
```

保存为 `setup_claude.sh` 并运行：
```bash
chmod +x setup_claude.sh
./setup_claude.sh
```

## 可用工具

配置成功后，Claude 可以使用以下工具：

### 1. get_yapi_interface
获取单个 YApi 接口信息

**示例：**
```
请帮我查看这个 YApi 接口的详细信息：
https://yapi.makeblock.com/project/382/interface/api/17888
```

### 2. get_yapi_project_interfaces
获取 YApi 项目中所有接口列表

**示例：**
```
请列出项目 382 的所有接口
```

## 故障排查

### 问题1: Claude Desktop 看不到 MCP 工具

**解决方案：**
1. 确认配置文件路径正确
2. 确认 JSON 格式正确（使用 JSON 验证器）
3. 确认二进制文件路径正确
4. 完全退出并重启 Claude Desktop

### 问题2: MCP 工具调用失败

**解决方案：**
1. 检查 Token 是否正确
2. 检查 YAPI_BASE_URL 是否正确
3. 测试二进制文件是否可以单独运行：
   ```bash
   /Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server
   ```

### 问题3: 权限问题

**解决方案：**
```bash
chmod +x /Users/makeblock/Desktop/plugin/yapi-mcp-server/bin/yapi-mcp-server
```

## 查看 Claude Desktop 日志

如果遇到问题，可以查看 Claude Desktop 日志：

```bash
# macOS
tail -f ~/Library/Logs/Claude/mcp*.log
```

## 更新 MCP 服务器

如果更新了代码，需要：
1. 重新编译：`go build -o bin/yapi-mcp-server .`
2. 重启 Claude Desktop

无需修改配置文件。

