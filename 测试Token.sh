#!/bin/bash

# 测试 Token 是否有效的脚本

if [ -z "$1" ]; then
    echo "用法: ./测试Token.sh <your-token>"
    echo "示例: ./测试Token.sh 'abc123def456'"
    exit 1
fi

TOKEN="$1"
INTERFACE_ID="17888"
BASE_URL="https://yapi.makeblock.com"

echo "=== 测试 YApi Token ==="
echo "Base URL: $BASE_URL"
echo "Interface ID: $INTERFACE_ID"
echo "Token: ${TOKEN:0:10}..."
echo ""

echo "测试 API 调用:"
API_URL="${BASE_URL}/api/interface/get?id=${INTERFACE_ID}&token=${TOKEN}"

echo "URL: $API_URL"
echo ""

echo "响应:"
curl -s "$API_URL" | jq . || curl -s "$API_URL"

echo ""
echo "---"
echo "如果看到 errcode=0，说明 Token 有效"
echo "如果看到 errcode=40011，说明需要登录或 Token 无效"

