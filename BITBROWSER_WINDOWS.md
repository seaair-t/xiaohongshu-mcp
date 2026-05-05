# Windows BitBrowser 支持

这个修改版本为 Windows 平台添加了 BitBrowser 指纹浏览器支持。

## 修改内容

### 1. main.go
- 添加了 BitBrowser 相关命令行参数：
  - `--browser-mode`: 浏览器模式 (headless/chrome/bitbrowser)
  - `--bitbrowser-api`: BitBrowser API 地址
  - `--bitbrowser-profile-id`: BitBrowser Profile ID
  - `--bitbrowser-auth-token`: BitBrowser 认证 Token (可选)

### 2. configs/browser.go
- 添加了浏览器模式配置
- 添加了 BitBrowser 配置支持
- 新增函数：
  - `SetBrowserMode()` / `GetBrowserMode()`
  - `IsBitBrowserMode()`
  - `SetBitBrowserConfig()`
  - `GetBitBrowserAPI()` / `GetBitBrowserProfile()` / `GetBitBrowserToken()`

### 3. browser/browser.go
- 重写了浏览器创建逻辑
- 添加了 `BrowserWrapper` 包装器，支持多种浏览器模式
- 添加了 `NewBitBrowser()` 函数，通过 BitBrowser API 启动窗口
- 添加了 `startBitBrowserWindow()` 函数，调用 BitBrowser API

### 4. service.go
- 修改了 `newBrowser()` 函数，支持根据配置选择浏览器模式

## 使用方法

### 方式 1：命令行参数
```batch
xiaohongshu-mcp.exe ^
  --browser-mode=bitbrowser ^
  --bitbrowser-api=http://127.0.0.1:54345 ^
  --bitbrowser-profile-id=daf4434d555046e4b72be4b19ae6bffc ^
  --port=:18060
```

### 方式 2：环境变量
```batch
set BITBROWSER_API_URL=http://127.0.0.1:54345
set BITBROWSER_PROFILE_ID=daf4434d555046e4b72be4b19ae6bffc
set BITBROWSER_AUTH_TOKEN=your_token_here

xiaohongshu-mcp.exe --browser-mode=bitbrowser --port=:18060
```

### 方式 3：使用启动脚本
双击运行 `start-bitbrowser.bat`

## 编译方法

### 前提条件
- 安装 Go 1.21 或更高版本
- 确保 BitBrowser 已安装并运行

### 编译步骤
```bash
cd xiaohongshu-mcp-src

# 下载依赖
go mod tidy

# 编译 Windows 版本
set GOOS=windows
set GOARCH=amd64
go build -o xiaohongshu-mcp-bitbrowser.exe

# 或者编译当前平台版本
go build -o xiaohongshu-mcp.exe
```

## 与原版对比

| 功能 | 原版 Windows | 修改版 |
|------|-------------|--------|
| `--browser-mode` | ❌ 不支持 | ✅ 支持 |
| `--bitbrowser-api` | ❌ 不支持 | ✅ 支持 |
| `--bitbrowser-profile-id` | ❌ 不支持 | ✅ 支持 |
| `--bitbrowser-auth-token` | ❌ 不支持 | ✅ 支持 |
| 系统 Chrome | ✅ 支持 | ✅ 支持 |
| 系统 Edge | ✅ 支持 | ✅ 支持 |
| BitBrowser | ❌ 不支持 | ✅ 支持 |

## 注意事项

1. **BitBrowser 必须先启动**：确保 BitBrowser 客户端已运行，API 服务可用
2. **Profile ID 必须存在**：确保指定的 Profile ID 在 BitBrowser 中存在
3. **端口冲突**：如果 18060 端口被占用，可以修改 `--port` 参数
4. **Cookie 持久化**：登录状态会保存到 cookies.json 文件

## 故障排除

### 问题 1：BitBrowser API 连接失败
```
Failed to call BitBrowser API: Post "http://127.0.0.1:54345/browser/open": dial tcp 127.0.0.1:54345: connectex: No connection could be made
```
**解决**：确保 BitBrowser 客户端已启动，API 服务已启用

### 问题 2：Profile ID 不存在
```
BitBrowser API failed: profile not found
```
**解决**：检查 Profile ID 是否正确，或在 BitBrowser 中创建新的 Profile

### 问题 3：端口被占用
```
listen tcp :18060: bind: Only one usage of each socket address
```
**解决**：修改 `--port` 参数为其他端口，如 `:18061`

## 参考

- BitBrowser API 文档：http://127.0.0.1:54345/swagger-ui/
- 原版项目：https://github.com/xpzouying/xiaohongshu-mcp
