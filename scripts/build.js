#!/usr/bin/env node

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const platform = os.platform();
const arch = os.arch();

// 确保bin目录存在
const binDir = path.join(__dirname, '..', 'bin');
if (!fs.existsSync(binDir)) {
  fs.mkdirSync(binDir, { recursive: true });
}

// 确定Go二进制文件名
let binaryName = 'yapi-mcp-server';
if (platform === 'windows') {
  binaryName += '.exe';
}

// 设置Go环境变量
const goos = platform === 'win32' ? 'windows' : platform === 'darwin' ? 'darwin' : 'linux';
const goarch = arch === 'x64' ? 'amd64' : arch === 'arm64' ? 'arm64' : '386';

process.env.GOOS = goos;
process.env.GOARCH = goarch;

console.log(`构建 ${goos}-${goarch} 版本的二进制文件...`);

try {
  // 检查Go是否安装
  try {
    execSync('go version', { stdio: 'ignore' });
  } catch (e) {
    console.error('错误: 未找到Go编译器，请先安装Go: https://golang.org/dl/');
    process.exit(1);
  }

  // 安装Go依赖
  console.log('安装Go依赖...');
  execSync('go mod download', {
    stdio: 'inherit',
    cwd: path.join(__dirname, '..')
  });

  // 构建Go二进制文件
  console.log('编译Go程序...');
  execSync(`go build -o bin/${binaryName} .`, {
    stdio: 'inherit',
    cwd: path.join(__dirname, '..')
  });
  
  // 设置可执行权限（Unix系统）
  if (platform !== 'win32') {
    fs.chmodSync(path.join(binDir, binaryName), '755');
  }
  
  console.log('构建成功！');
} catch (error) {
  console.error('构建失败:', error.message);
  process.exit(1);
}

