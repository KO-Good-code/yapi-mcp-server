#!/bin/bash

# Claude Desktop 自动配置脚本

set -e

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Claude Desktop MCP 配置 ===${NC}\n"

# Claude Desktop 配置文件路径
CONFIG_DIR=~/Library/Application\ Support/Claude
CONFIG_FILE="$CONFIG_DIR/claude_desktop_config.json"

# 当前项目路径
PROJECT_DIR="/Users/makeblock/Desktop/plugin/yapi-mcp-server"
BINARY_PATH="$PROJECT_DIR/bin/yapi-mcp-server"

# 检查二进制文件是否存在
if [ ! -f "$BINARY_PATH" ]; then
    echo -e "${RED}错误: 二进制文件不存在: $BINARY_PATH${NC}"
    echo "请先运行: go build -o bin/yapi-mcp-server ."
    exit 1
fi

# 确保二进制文件可执行
chmod +x "$BINARY_PATH"
echo -e "${GREEN}✓ 二进制文件已设置为可执行${NC}"

# 创建 Claude 配置目录
mkdir -p "$CONFIG_DIR"

# 检查是否已有配置文件
if [ -f "$CONFIG_FILE" ]; then
    echo -e "${YELLOW}检测到已有配置文件${NC}"
    echo "位置: $CONFIG_FILE"
    echo ""
    read -p "是否备份现有配置? (y/N): " backup
    if [ "$backup" = "y" ] || [ "$backup" = "Y" ]; then
        BACKUP_FILE="$CONFIG_FILE.backup.$(date +%Y%m%d_%H%M%S)"
        cp "$CONFIG_FILE" "$BACKUP_FILE"
        echo -e "${GREEN}✓ 已备份到: $BACKUP_FILE${NC}"
    fi
fi

# 获取 Token
echo ""
read -p "请输入 YAPI_TOKEN (直接回车使用默认): " YAPI_TOKEN
if [ -z "$YAPI_TOKEN" ]; then
    echo -e "${YELLOW}YAPI_TOKEN 未设置，请在配置文件中手动设置${NC}"
    YAPI_TOKEN=""
fi

# 获取 Base URL
echo ""
read -p "请输入 YAPI_BASE_URL (直接回车使用 https://yapi.makeblock.com): " YAPI_BASE_URL
if [ -z "$YAPI_BASE_URL" ]; then
    YAPI_BASE_URL="https://yapi.makeblock.com"
fi

# 创建配置文件
echo ""
echo "创建配置文件..."

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

echo -e "${GREEN}✓ Claude Desktop 配置已更新${NC}"
echo ""
echo "配置文件位置: $CONFIG_FILE"
echo ""
echo "配置内容:"
cat "$CONFIG_FILE"
echo ""
echo -e "${YELLOW}=== 下一步 ===${NC}"
echo "1. 完全退出 Claude Desktop (Cmd+Q)"
echo "2. 重新打开 Claude Desktop"
echo "3. 在对话中测试 MCP 工具"
echo ""
echo "测试示例:"
echo '  "请帮我获取这个 YApi 接口信息: https://yapi.makeblock.com/project/382/interface/api/17888"'
echo ""
echo -e "${GREEN}配置完成！${NC}"

