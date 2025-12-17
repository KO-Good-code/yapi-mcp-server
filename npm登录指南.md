# npm 登录指南

## 方案1: 登录公共 npm（推荐用于测试）

### 步骤1: 登录 npm

```bash
npm login
```

会提示输入：
- Username: 你的 npm 用户名
- Password: 你的 npm 密码
- Email: 你的邮箱地址

### 步骤2: 验证登录

```bash
npm whoami
```

如果显示你的用户名，说明登录成功。

### 步骤3: 发布

```bash
npm publish
```

## 方案2: 登录私有 npm 仓库

### 步骤1: 配置私有仓库

创建 `.npmrc` 文件：

```bash
# 如果使用 scope（推荐）
@neigri:registry=https://your-private-npm-registry.com
//your-private-npm-registry.com/:_authToken=your-token-here

# 或者全局配置
registry=https://your-private-npm-registry.com
//your-private-npm-registry.com/:_authToken=your-token-here
```

### 步骤2: 登录私有仓库

```bash
npm login --registry=https://your-private-npm-registry.com
```

或者如果使用 scope：

```bash
npm login --scope=@neigri --registry=https://your-private-npm-registry.com
```

### 步骤3: 发布

```bash
npm publish
```

## 方案3: 使用 npm token（适合 CI/CD）

### 步骤1: 生成 token

在 npm 网站或私有仓库管理界面生成 access token。

### 步骤2: 配置 token

```bash
# 公共 npm
npm config set //registry.npmjs.org/:_authToken=your-token

# 私有 npm
npm config set //your-private-npm-registry.com/:_authToken=your-token
```

或者创建 `.npmrc` 文件：

```
//registry.npmjs.org/:_authToken=your-token
```

### 步骤3: 发布

```bash
npm publish
```

## 常见问题

### Q1: 忘记密码怎么办？

访问 https://www.npmjs.com/forgot 重置密码。

### Q2: 需要两步验证吗？

如果启用了 2FA，需要：
1. 生成 access token
2. 使用 token 登录而不是密码

### Q3: 如何查看当前登录状态？

```bash
npm whoami
```

### Q4: 如何退出登录？

```bash
npm logout
```

### Q5: 发布到私有仓库失败？

检查：
1. `.npmrc` 配置是否正确
2. token 是否有效
3. 是否有发布权限
4. package.json 中的 `name` 是否包含正确的 scope

## 当前项目配置

根据 `package.json`：
- 包名: `@neigri/yapi-mcp-server`（使用 scope）

## 快速命令参考

```bash
# 登录公共 npm
npm login

# 登录私有 npm
npm login --registry=https://your-registry.com

# 查看当前用户
npm whoami

# 查看配置
npm config list

# 发布
npm publish

# 退出登录
npm logout
```

