@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

:: 设置 BitBrowser 配置
set "BITBROWSER_API_URL=http://127.0.0.1:54345"
set "BITBROWSER_PROFILE_ID=daf4434d555046e4b72be4b19ae6bffc"
set "BITBROWSER_AUTH_TOKEN="
set "PORT=18060"

echo [INFO] 启动 xiaohongshu-mcp (Windows BitBrowser 模式)...
echo [INFO] BitBrowser API: %BITBROWSER_API_URL%
echo [INFO] BitBrowser Profile: %BITBROWSER_PROFILE_ID%
echo [INFO] MCP endpoint: http://127.0.0.1:%PORT%/mcp
echo.

:: 启动 MCP 服务
xiaohongshu-mcp.exe ^
  --browser-mode=bitbrowser ^
  --bitbrowser-api=%BITBROWSER_API_URL% ^
  --bitbrowser-profile-id=%BITBROWSER_PROFILE_ID% ^
  --port=:%PORT%

pause
