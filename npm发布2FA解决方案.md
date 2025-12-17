# npm 发布 2FA 解决方案

## 问题原因

错误信息显示：
```
Two-factor authentication or granular access token with bypass 2fa enabled is required to publish packages.
```

这说明你的 npm 账号启用了**两步验证（2FA）**，需要使用 **access token** 而不是密码来发布包。

## 解决方案

### 方法1: 生成 Access Token（推荐）

#### 步骤1: 登录 npm 网站

访问 https://www.npmjs.com/ 并登录你的账号。

#### 步骤2: 生成 Access Token

1. 点击右上角头像 → **Access Tokens**
2. 或者直接访问：https://www.npmjs.com/settings/YOUR_USERNAME/tokens
3. 点击 **Generate New Token**
4. 选择 Token 类型：
   - **Automation** - 用于 CI/CD（推荐）
   - **Publish** - 用于发布包
   - **Granular** - 细粒度权限控制
5. 设置过期时间（或选择 "Never expire"）
6. 点击 **Generate Token**
7. **重要**：复制生成的 token（只显示一次！）

#### 步骤3: 使用 Token 登录

```bash
# 方法1: 使用 npm login（推荐）
npm login

# 当提示输入密码时，输入你的 access token（不是账号密码）

# 方法2: 直接配置 token
npm config set //registry.npmjs.org/:_authToken=your-access-token-here
```

#### 步骤4: 验证登录

```bash
npm whoami
```

#### 步骤5: 发布

```bash
npm publish
```

### 方法2: 使用 .npmrc 文件

创建或编辑项目根目录的 `.npmrc` 文件：

```bash
# .npmrc
//registry.npmjs.org/:_authToken=your-access-token-here
```

**注意**：确保 `.npmrc` 在 `.gitignore` 中，不要提交到 Git！

### 方法3: 使用环境变量

```bash
export NPM_TOKEN=your-access-token-here
npm config set //registry.npmjs.org/:_authToken=$NPM_TOKEN
npm publish
```

## 完整流程示例

```bash
# 1. 生成 token（在 npm 网站上）
# 2. 配置 token
npm config set //registry.npmjs.org/:_authToken=your-token-here

# 3. 验证
npm whoami

# 4. 发布
cd /Users/makeblock/Desktop/plugin/yapi-mcp-server
npm publish
```

## 常见问题

### Q1: Token 在哪里生成？

访问：https://www.npmjs.com/settings/YOUR_USERNAME/tokens

### Q2: Token 类型选择哪个？

- **Automation**: 适合 CI/CD 和自动化脚本
- **Publish**: 只用于发布包
- **Granular**: 最灵活，可以设置详细权限

### Q3: Token 过期了怎么办？

重新生成一个新的 token，然后更新配置。

### Q4: 如何查看当前使用的 token？

```bash
npm config get //registry.npmjs.org/:_authToken
```

### Q5: 如何删除 token？

在 npm 网站的 Access Tokens 页面删除即可。

## 安全建议

1. ✅ 使用 **Automation** 或 **Granular** token
2. ✅ 设置合理的过期时间
3. ✅ 不要将 token 提交到 Git
4. ✅ 定期轮换 token
5. ✅ 使用环境变量存储 token（CI/CD）

## 快速命令

```bash
# 配置 token
npm config set //registry.npmjs.org/:_authToken=YOUR_TOKEN

# 验证
npm whoami

# 发布
npm publish
```

