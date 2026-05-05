package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xpzouying/xiaohongshu-mcp/configs"
)

func main() {
	var (
		headless bool
		binPath  string // 浏览器二进制文件路径
		port     string
		
		// BitBrowser 模式参数
		browserMode       string // 浏览器模式: headless, chrome, bitbrowser
		bitbrowserAPI     string // BitBrowser API 地址
		bitbrowserProfile string // BitBrowser Profile ID
		bitbrowserToken   string // BitBrowser 认证 Token
	)
	flag.BoolVar(&headless, "headless", true, "是否无头模式")
	flag.StringVar(&binPath, "bin", "", "浏览器二进制文件路径")
	flag.StringVar(&port, "port", ":18060", "端口")
	
	// BitBrowser 参数
	flag.StringVar(&browserMode, "browser-mode", "headless", "浏览器模式: headless, chrome, bitbrowser")
	flag.StringVar(&bitbrowserAPI, "bitbrowser-api", "", "BitBrowser API 地址 (默认: http://127.0.0.1:54345)")
	flag.StringVar(&bitbrowserProfile, "bitbrowser-profile-id", "", "BitBrowser Profile ID")
	flag.StringVar(&bitbrowserToken, "bitbrowser-auth-token", "", "BitBrowser 认证 Token (可选)")
	flag.Parse()

	if len(binPath) == 0 {
		binPath = os.Getenv("ROD_BROWSER_BIN")
	}
	
	// BitBrowser 环境变量支持
	if browserMode == "bitbrowser" {
		if bitbrowserAPI == "" {
			bitbrowserAPI = os.Getenv("BITBROWSER_API_URL")
			if bitbrowserAPI == "" {
				bitbrowserAPI = "http://127.0.0.1:54345"
			}
		}
		if bitbrowserProfile == "" {
			bitbrowserProfile = os.Getenv("BITBROWSER_PROFILE_ID")
		}
		if bitbrowserToken == "" {
			bitbrowserToken = os.Getenv("BITBROWSER_AUTH_TOKEN")
		}
		
		logrus.Infof("BitBrowser 模式: API=%s, Profile=%s", bitbrowserAPI, bitbrowserProfile)
	}

	configs.InitHeadless(headless)
	configs.SetBinPath(binPath)
	configs.SetBrowserMode(browserMode)
	configs.SetBitBrowserConfig(bitbrowserAPI, bitbrowserProfile, bitbrowserToken)

	// 初始化服务
	xiaohongshuService := NewXiaohongshuService()

	// 创建并启动应用服务器
	appServer := NewAppServer(xiaohongshuService)
	if err := appServer.Start(port); err != nil {
		logrus.Fatalf("failed to run server: %v", err)
	}
}
