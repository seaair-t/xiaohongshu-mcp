# BitBrowser 支持修改说明

## 修改概述

为 Windows 平台添加 BitBrowser 指纹浏览器支持，实现方式参考 macOS 版本。

## 修改文件列表

### 1. main.go
**修改内容**：
- 添加 `--browser-mode` 参数（支持 headless/chrome/bitbrowser）
- 添加 `--bitbrowser-api` 参数
- 添加 `--bitbrowser-profile-id` 参数
- 添加 `--bitbrowser-auth-token` 参数
- 支持从环境变量读取配置

### 2. configs/browser.go
**修改内容**：
- 添加 `browserMode` 变量
- 添加 `bitbrowserAPI`, `bitbrowserProfile`, `bitbrowserToken` 变量
- 添加 `SetBrowserMode()` / `GetBrowserMode()` 函数
- 添加 `IsBitBrowserMode()` 函数
- 添加 `SetBitBrowserConfig()` 函数
- 添加 `GetBitBrowserAPI()` / `GetBitBrowserProfile()` / `GetBitBrowserToken()` 函数

### 3. browser/browser.go
**修改内容**：
- 添加 `BitBrowserResponse` 结构体
- 添加 `BrowserWrapper` 包装器结构体
- 重写 `NewBrowser()` 函数返回包装器
- 添加 `NewBitBrowser()` 函数
- 添加 `startBitBrowserWindow()` 函数调用 BitBrowser API
- 添加 `NewPage()` / `Close()` / `GetCookies()` 方法

### 4. service.go
**修改内容**：
- 修改 `newBrowser()` 函数，支持根据配置选择浏览器模式
- BitBrowser 模式下调用 `browser.NewBitBrowser()`
- 默认模式下调用 `browser.NewBrowser()`

## 新增文件

### start-bitbrowser.bat
Windows 批处理启动脚本，方便用户使用。

### BITBROWSER_WINDOWS.md
详细的使用文档和故障排除指南。

## 使用方法

### 命令行参数
```batch
xiaohongshu-mcp.exe ^
  --browser-mode=bitbrowser ^
  --bitbrowser-api=http://127.0.0.1:54345 ^
  --bitbrowser-profile-id=your_profile_id ^
  --bitbrowser-auth-token=your_token ^
  --port=:18060
```

### 环境变量
```batch
set BITBROWSER_PROFILE_ID=your_profile_id
set BITBROWSER_API_URL=http://127.0.0.1:54345
set BITBROWSER_AUTH_TOKEN=your_token

xiaohongshu-mcp.exe --browser-mode=bitbrowser --port=:18060
```

### 启动脚本
双击运行 `start-bitbrowser.bat`

## 与原版对比

| 功能 | 原版 Windows | 修改版 |
|------|-------------|--------|
| `--browser-mode` | ❌ | ✅ |
| `--bitbrowser-api` | ❌ | ✅ |
| `--bitbrowser-profile-id` | ❌ | ✅ |
| `--bitbrowser-auth-token` | ❌ | ✅ |
| 系统 Chrome | ✅ | ✅ |
| BitBrowser | ❌ | ✅ |

## 编译方法

```bash
# 下载依赖
go mod tidy

# 编译 Windows 版本
set GOOS=windows
set GOARCH=amd64
go build -o xiaohongshu-mcp-bitbrowser.exe
```

## 注意事项

1. BitBrowser 客户端必须先启动
2. Profile ID 必须存在
3. 需要 Go 1.21+ 环境编译

## 参考

- 原版项目：https://github.com/xpzouying/xiaohongshu-mcp
- BitBrowser API：http://127.0.0.1:54345/swagger-ui/
