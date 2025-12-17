#!/bin/bash

# YApi MCP Server 测试脚本

set -e

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== YApi MCP Server 测试 ===${NC}\n"

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo -e "${RED}错误: 未找到Go，请先安装Go${NC}"
    exit 1
fi

echo -e "${YELLOW}Go版本:${NC}"
go version
echo ""

# 安装依赖
echo -e "${YELLOW}安装Go依赖...${NC}"
go mod download
go mod tidy
echo ""

# 设置环境变量（如果未设置）
if [ -z "$YAPI_BASE_URL" ]; then
    export YAPI_BASE_URL="http://localhost"
    echo -e "${YELLOW}YAPI_BASE_URL未设置，使用默认值: http://localhost${NC}"
fi

if [ -z "$YAPI_TOKEN" ]; then
    echo -e "${YELLOW}YAPI_TOKEN未设置（可选）${NC}"
    echo -e "${YELLOW}提示: 如果需要访问需要认证的YApi实例，请设置: export YAPI_TOKEN='your-token'${NC}"
else
    echo -e "${GREEN}YAPI_TOKEN已设置${NC}"
fi

echo ""

# 测试1: 初始化
echo -e "${GREEN}测试1: 初始化请求${NC}"
if command -v jq &> /dev/null; then
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test-client","version":"1.0.0"}}}' | go run . | jq .
else
    echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test-client","version":"1.0.0"}}}' | go run .
fi
echo ""

# 测试2: 列出工具
echo -e "${GREEN}测试2: 列出可用工具${NC}"
if command -v jq &> /dev/null; then
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | go run . | jq .
else
    echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | go run .
fi
echo ""

# 测试3: 调用工具（需要有效的YApi URL）
if [ -n "$1" ]; then
    TEST_URL="$1"
else
    TEST_URL="http://yapi.example.com/project/123/interface/api/456"
fi

echo -e "${GREEN}测试3: 调用get_yapi_interface工具${NC}"
echo -e "${YELLOW}使用URL: $TEST_URL${NC}"
if command -v jq &> /dev/null; then
    echo "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"tools/call\",\"params\":{\"name\":\"get_yapi_interface\",\"arguments\":{\"url\":\"$TEST_URL\"}}}" | go run . | jq .
else
    echo "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"tools/call\",\"params\":{\"name\":\"get_yapi_interface\",\"arguments\":{\"url\":\"$TEST_URL\"}}}" | go run .
fi
echo ""

# 测试4: 获取项目接口列表
if [ -n "$2" ]; then
    PROJECT_URL="$2"
else
    PROJECT_URL="http://yapi.example.com/project/123"
fi

echo -e "${GREEN}测试4: 获取项目接口列表${NC}"
echo -e "${YELLOW}使用项目URL: $PROJECT_URL${NC}"
if command -v jq &> /dev/null; then
    echo "{\"jsonrpc\":\"2.0\",\"id\":4,\"method\":\"tools/call\",\"params\":{\"name\":\"get_yapi_project_interfaces\",\"arguments\":{\"project_url\":\"$PROJECT_URL\"}}}" | go run . | jq .
else
    echo "{\"jsonrpc\":\"2.0\",\"id\":4,\"method\":\"tools/call\",\"params\":{\"name\":\"get_yapi_project_interfaces\",\"arguments\":{\"project_url\":\"$PROJECT_URL\"}}}" | go run .
fi
echo ""

echo -e "${GREEN}=== 测试完成 ===${NC}"
echo ""
echo "提示:"
echo "  - 使用 jq 可以更好地格式化JSON输出: brew install jq"
echo "  - 传入自定义URL测试: ./test.sh 'http://your-yapi.com/project/123/interface/api/456'"
echo ""
echo "环境变量设置:"
echo "  - 设置YApi基础URL: export YAPI_BASE_URL='http://your-yapi.com'"
echo "  - 设置YApi Token: export YAPI_TOKEN='your-token'"
echo ""
echo "如何获取YApi Token:"
echo "  1. 登录YApi系统"
echo "  2. 进入项目设置页面"
echo "  3. 在'Token配置'或'接口设置'中找到Token"
echo "  4. 复制Token并设置: export YAPI_TOKEN='your-token'"

