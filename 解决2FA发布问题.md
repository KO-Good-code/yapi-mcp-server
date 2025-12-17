# 解决 npm 2FA 发布问题

## 当前状态

- ✅ 已登录：`neigri`
- ✅ 已配置 token
- ❌ Token 权限不足，无法发布

## 问题原因

错误信息：
```
Two-factor authentication or granular access token with bypass 2fa enabled is required to publish packages.
```

这说明当前使用的 token **没有 bypass 2FA 的权限**。

## 解决方案

### 步骤1: 生成新的 Access Token

1. **访问 token 管理页面**：
   ```
   https://www.npmjs.com/settings/neigri/tokens
   ```

2. **生成新 Token**：
   - 点击 **"Generate New Token"**
   - **重要**：选择 **"Granular"** 类型（不是 Automation 或 Publish）
   - 在权限设置中：
     - ✅ 勾选 **"Bypass 2FA"** 选项
     - ✅ 选择 **"Publish"** 权限
     - ✅ 选择包名：`yapi-mcp-server`（或选择 "All packages"）
   - 设置过期时间（或选择 "Never expire"）
   - 点击 **"Generate Token"**
   - **立即复制 token**（只显示一次！）

### 步骤2: 更新 Token 配置

```bash
# 方法1: 使用配置脚本（推荐）
./快速配置Token.sh

# 方法2: 手动配置
npm config set //registry.npmjs.org/:_authToken=你的新token

# 方法3: 使用 .npmrc 文件
echo "//registry.npmjs.org/:_authToken=你的新token" > .npmrc
```

### 步骤3: 验证配置

```bash
npm whoami
```

### 步骤4: 发布

```bash
npm publish
```

## Token 类型说明

| Token 类型 | 是否支持 bypass 2FA | 用途 |
|-----------|-------------------|------|
| **Granular** | ✅ 是 | 最灵活，可以设置详细权限 |
| Automation | ❌ 否 | 用于 CI/CD |
| Publish | ❌ 否 | 基础发布权限 |

**必须使用 Granular token 并启用 bypass 2FA！**

## 快速操作

```bash
# 1. 在 npm 网站生成 Granular token（启用 bypass 2FA）
# 2. 配置新 token
npm config set //registry.npmjs.org/:_authToken=你的新token

# 3. 验证
npm whoami

# 4. 发布
npm publish
```

## 如果还是失败

### 检查1: Token 权限

确保 token 有：
- ✅ Bypass 2FA 权限
- ✅ Publish 权限
- ✅ 对 `yapi-mcp-server` 包的权限

### 检查2: 包名是否可用

```bash
# 检查包名是否已被占用
npm view yapi-mcp-server
```

如果包名已被占用，需要：
1. 修改 `package.json` 中的 `name` 字段
2. 或使用 scope：`@neigri/yapi-mcp-server`

### 检查3: 清除旧配置

```bash
# 删除旧的 token 配置
npm config delete //registry.npmjs.org/:_authToken

# 重新配置新 token
npm config set //registry.npmjs.org/:_authToken=新token
```

## 推荐配置

使用 Granular token，设置：
- **Type**: Granular
- **Permissions**: 
  - ✅ Bypass 2FA
  - ✅ Publish
- **Packages**: `yapi-mcp-server` 或 "All packages"
- **Expiration**: 根据需要设置

## 验证命令

```bash
# 检查当前用户
npm whoami

# 检查 token 是否配置（不会显示内容，但会显示是否配置）
npm config get //registry.npmjs.org/:_authToken

# 测试发布（使用 --dry-run 不会真正发布）
npm publish --dry-run
```

