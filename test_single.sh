#!/bin/bash

# 单个接口测试脚本 - 带详细输出

if [ -z "$1" ]; then
    echo "用法: ./test_single.sh <yapi-url>"
    echo "示例: ./test_single.sh 'https://yapi.makeblock.com/project/382/interface/api/17888'"
    exit 1
fi

URL="$1"

echo "=== 测试单个接口 ==="
echo "URL: $URL"
echo ""

# 检查环境变量
echo "环境变量:"
echo "  YAPI_BASE_URL: ${YAPI_BASE_URL:-未设置}"
echo "  YAPI_TOKEN: ${YAPI_TOKEN:+已设置}"
echo ""

# 构建测试请求
REQUEST="{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"tools/call\",\"params\":{\"name\":\"get_yapi_interface\",\"arguments\":{\"url\":\"$URL\"}}}"

echo "发送请求:"
echo "$REQUEST" | jq . 2>/dev/null || echo "$REQUEST"
echo ""

echo "响应:"
echo "$REQUEST" | go run . 2>&1 | tee /tmp/yapi_response.json

echo ""
echo "格式化响应:"
cat /tmp/yapi_response.json | jq . 2>/dev/null || cat /tmp/yapi_response.json

echo ""
echo "提取接口信息:"
cat /tmp/yapi_response.json | jq '.content[0].text' 2>/dev/null | sed 's/^"//;s/"$//' | sed 's/\\n/\n/g' | jq . 2>/dev/null || echo "无法解析"

