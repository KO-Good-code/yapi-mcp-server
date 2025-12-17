#!/bin/bash

# 快速设置YApi Token的辅助脚本

echo "=== YApi Token 快速设置 ==="
echo ""

# 检查是否已有.env文件
if [ -f .env ]; then
    echo "检测到已有 .env 文件"
    read -p "是否要覆盖? (y/N): " overwrite
    if [ "$overwrite" != "y" ] && [ "$overwrite" != "Y" ]; then
        echo "取消操作"
        exit 0
    fi
fi

# 获取YApi基础URL
echo ""
read -p "请输入YApi基础URL (例如: http://yapi.example.com): " base_url
if [ -z "$base_url" ]; then
    base_url="http://localhost"
    echo "使用默认值: $base_url"
fi

# 获取Token
echo ""
read -p "请输入YApi Token (可选，直接回车跳过): " token

# 创建.env文件
cat > .env << EOF
# YApi配置
# 生成时间: $(date)

YAPI_BASE_URL=$base_url
EOF

if [ -n "$token" ]; then
    echo "YAPI_TOKEN=$token" >> .env
    echo ""
    echo "✓ Token已设置"
else
    echo "# YAPI_TOKEN=your-token-here" >> .env
    echo ""
    echo "⚠ Token未设置（可选）"
fi

echo ""
echo "✓ 配置已保存到 .env 文件"
echo ""
echo "使用方法:"
echo "  1. 加载环境变量: source .env"
echo "  2. 运行测试: ./test.sh"
echo ""
echo "注意: .env 文件已添加到 .gitignore，不会被提交到Git"

