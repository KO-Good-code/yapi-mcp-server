#!/bin/bash

# npm Access Token 快速配置脚本

echo "=== npm Access Token 配置 ==="
echo ""
echo "由于你的账号启用了两步验证（2FA），需要使用 Access Token 来发布包。"
echo ""
echo "步骤1: 生成 Access Token"
echo "1. 访问: https://www.npmjs.com/settings/$(npm whoami 2>/dev/null || echo 'YOUR_USERNAME')/tokens"
echo "2. 点击 'Generate New Token'"
echo "3. 选择 'Automation' 或 'Publish' 类型"
echo "4. 复制生成的 token（只显示一次！）"
echo ""
read -p "请输入你的 Access Token: " NPM_TOKEN

if [ -z "$NPM_TOKEN" ]; then
    echo "错误: Token 不能为空"
    exit 1
fi

# 配置 token
echo ""
echo "配置 token..."
npm config set //registry.npmjs.org/:_authToken="$NPM_TOKEN"

# 验证
echo ""
echo "验证配置..."
if npm whoami > /dev/null 2>&1; then
    echo "✓ 配置成功！当前用户: $(npm whoami)"
    echo ""
    echo "现在可以发布包了："
    echo "  npm publish"
else
    echo "✗ 配置失败，请检查 token 是否正确"
    exit 1
fi

