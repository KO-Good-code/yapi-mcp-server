# 详细 Token 配置步骤（解决 2FA 问题）

## 问题确认

错误信息：
```
Two-factor authentication or granular access token with bypass 2fa enabled is required to publish packages.
```

这说明当前 token **没有 bypass 2FA 权限**。

## 必须使用 Granular Token

**重要**：必须使用 **Granular** 类型的 token，并且**必须启用 bypass 2FA**。

## 详细步骤

### 步骤1: 访问 Token 管理页面

1. 打开浏览器
2. 访问：https://www.npmjs.com/settings/neigri/tokens
3. 确保已登录 npm 账号

### 步骤2: 生成 Granular Token

1. 点击 **"Generate New Token"** 按钮
2. **选择 Token 类型**：
   - ❌ **不要选** "Automation"
   - ❌ **不要选** "Publish"  
   - ✅ **必须选** "**Granular**"

### 步骤3: 配置 Token 权限（关键步骤）

在 Granular token 配置页面：

#### 3.1 设置 Token 名称
```
yapi-mcp-server-publish
```

#### 3.2 设置权限（Permissions）

**必须勾选以下权限**：

1. ✅ **"Bypass 2FA"** - **这是必须的！**
   - 位置：在权限列表的最上方
   - 必须勾选，否则无法发布

2. ✅ **"Publish"** 权限
   - 允许发布包

#### 3.3 设置包范围（Packages）

选择以下之一：
- ✅ **"All packages"** - 可以发布所有包（推荐）
- ✅ **"Selected packages"** - 然后选择 `@neigri/yapi-mcp-server`

#### 3.4 设置过期时间

- 选择 "Never expire"（永久有效）
- 或设置一个较长的过期时间

#### 3.5 生成 Token

1. 点击 **"Generate Token"**
2. **立即复制 token**（格式类似：`npm_xxxxxxxxxxxxxxxxxxxxx`）
3. **保存到安全的地方**（只显示一次！）

### 步骤4: 配置 Token

#### 方法1: 使用 npm config（推荐）

```bash
# 删除可能存在的旧 token
npm config delete //registry.npmjs.org/:_authToken

# 配置新 token（替换 YOUR_TOKEN 为实际 token）
npm config set //registry.npmjs.org/:_authToken=YOUR_TOKEN
```

#### 方法2: 使用 .npmrc 文件

```bash
# 在项目根目录创建或编辑 .npmrc
echo "//registry.npmjs.org/:_authToken=YOUR_TOKEN" > .npmrc

# 确保 .npmrc 在 .gitignore 中（已配置）
```

### 步骤5: 验证配置

```bash
# 检查当前用户
npm whoami

# 应该显示: neigri

# 检查 token 是否配置（不会显示内容）
npm config get //registry.npmjs.org/:_authToken
# 应该显示: (protected) 或 token 值
```

### 步骤6: 测试发布

```bash
cd /Users/makeblock/Desktop/plugin/yapi-mcp-server

# 先测试（不会真正发布）
npm publish --dry-run

# 如果测试通过，正式发布
npm publish --access public
```

## 常见错误

### 错误1: Token 没有 bypass 2FA 权限

**症状**：仍然报错需要 2FA

**解决**：
- 确保选择了 **Granular** token
- 确保勾选了 **"Bypass 2FA"** 选项
- 重新生成 token

### 错误2: Token 过期

**症状**：`Access token expired or revoked`

**解决**：
- 重新生成新 token
- 更新配置

### 错误3: 权限不足

**症状**：`403 Forbidden`

**解决**：
- 确保 token 有 **Publish** 权限
- 确保包范围包含 `@neigri/*`

## 验证 Token 类型

如果已经配置了 token，可以通过以下方式验证：

```bash
# 尝试发布一个测试包
npm publish --dry-run

# 如果仍然报 2FA 错误，说明 token 类型不对
```

## 完整命令序列

```bash
# 1. 删除旧配置
npm config delete //registry.npmjs.org/:_authToken

# 2. 配置新 token（替换 YOUR_TOKEN）
npm config set //registry.npmjs.org/:_authToken=YOUR_TOKEN

# 3. 验证
npm whoami

# 4. 发布
cd /Users/makeblock/Desktop/plugin/yapi-mcp-server
npm publish --access public
```

## 重要提示

1. ✅ **必须使用 Granular token**
2. ✅ **必须启用 Bypass 2FA**
3. ✅ **必须设置 Publish 权限**
4. ✅ **使用 scope 包名需要 `--access public`**

## 如果还是不行

1. 检查 npm 账号是否真的启用了 2FA：
   - 访问：https://www.npmjs.com/settings/neigri/two-factor
   - 如果启用了，必须使用 Granular token with bypass 2FA

2. 尝试临时禁用 2FA（不推荐，安全性降低）

3. 联系 npm 支持：support@npmjs.com

