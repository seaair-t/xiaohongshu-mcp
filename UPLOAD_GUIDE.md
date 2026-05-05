# 上传修改到 GitHub 指南

## 前置条件

1. 安装 Git：https://git-scm.com/download/win
2. 有 GitHub 账号
3. 已经 Fork 了原仓库

## 步骤

### 步骤 1：Fork 原仓库

1. 打开 https://github.com/xpzouying/xiaohongshu-mcp
2. 点击右上角的 **Fork** 按钮
3. 选择你的账号，等待 Fork 完成

Fork 完成后，你的仓库地址会是：
```
https://github.com/你的用户名/xiaohongshu-mcp
```

### 步骤 2：运行上传脚本

在 `xiaohongshu-mcp-src` 目录中，双击运行 `git-setup.bat`

或者在命令行中执行：
```batch
cd D:\code\xhsimprove\xiaohongshu-mcp-src
git-setup.bat
```

按提示输入你的 GitHub 仓库地址。

### 步骤 3：手动上传（如果脚本失败）

如果脚本失败，可以手动执行：

```batch
cd D:\code\xhsimprove\xiaohongshu-mcp-src

:: 添加你的仓库
git remote add myrepo https://github.com/你的用户名/xiaohongshu-mcp.git

:: 创建新分支
git checkout -b bitbrowser-support

:: 添加修改的文件
git add main.go configs/browser.go browser/browser.go service.go
git add start-bitbrowser.bat BITBROWSER_WINDOWS.md CHANGES.md

:: 提交
git commit -m "feat: add BitBrowser support for Windows"

:: 推送
git push myrepo bitbrowser-support
```

### 步骤 4：触发 GitHub Actions 编译

1. 访问你的 GitHub 仓库
2. 切换到 `bitbrowser-support` 分支
3. 点击 "Actions" 标签
4. 找到 "Build and Release" 工作流
5. 点击 "Run workflow" 手动触发编译

或者，你可以创建一个 Pull Request 到原仓库，这样也会触发编译。

### 步骤 5：下载编译好的二进制文件

1. 等待 GitHub Actions 完成（约 2-3 分钟）
2. 点击 "Releases" 标签
3. 找到最新的 Release
4. 下载 `xiaohongshu-mcp-windows-amd64.zip`
5. 解压后使用

## 使用方法

下载编译好的文件后：

```batch
:: 方式 1：命令行参数
xiaohongshu-mcp.exe ^
  --browser-mode=bitbrowser ^
  --bitbrowser-profile-id=daf4434d555046e4b72be4b19ae6bffc ^
  --port=:18060

:: 方式 2：环境变量
set BITBROWSER_PROFILE_ID=daf4434d555046e4b72be4b19ae6bffc
xiaohongshu-mcp.exe --browser-mode=bitbrowser --port=:18060

:: 方式 3：启动脚本
start-bitbrowser.bat
```

## 文件说明

### 修改的文件
- `main.go` - 添加命令行参数
- `configs/browser.go` - 添加配置支持
- `browser/browser.go` - 添加 BitBrowser 客户端
- `service.go` - 修改浏览器初始化

### 新增的文件
- `start-bitbrowser.bat` - Windows 启动脚本
- `BITBROWSER_WINDOWS.md` - 使用文档
- `CHANGES.md` - 修改说明
- `mcp_bitbrowser_proxy.py` - Python 代理脚本（备用方案）
- `start-bitbrowser-proxy.bat` - Python 代理启动脚本

## 故障排除

### Git 推送失败
```
fatal: unable to access 'https://github.com/...': The requested URL returned error: 403
```
**解决**：需要配置 GitHub 认证
```batch
git config --global user.name "你的用户名"
git config --global user.email "你的邮箱"
```

### 编译失败
在 GitHub Actions 中查看错误日志，通常是依赖问题。

### 运行时找不到浏览器
确保 BitBrowser 客户端已启动，Profile ID 正确。

## 联系

如有问题，可以：
1. 在原仓库提交 Issue
2. 在你的仓库提交 Issue
3. 查看 `BITBROWSER_WINDOWS.md` 故障排除部分
