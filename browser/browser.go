package browser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/headless_browser"
	"github.com/xpzouying/xiaohongshu-mcp/cookies"
)

type browserConfig struct {
	binPath string
}

type Option func(*browserConfig)

func WithBinPath(binPath string) Option {
	return func(c *browserConfig) {
		c.binPath = binPath
	}
}

// maskProxyCredentials masks username and password in proxy URL for safe logging.
func maskProxyCredentials(proxyURL string) string {
	u, err := url.Parse(proxyURL)
	if err != nil || u.User == nil {
		return proxyURL
	}
	if _, hasPassword := u.User.Password(); hasPassword {
		u.User = url.UserPassword("***", "***")
	} else {
		u.User = url.User("***")
	}
	return u.String()
}

// BitBrowserResponse BitBrowser API 响应结构
type BitBrowserResponse struct {
	Success bool `json:"success"`
	Data    struct {
		HTTP      string `json:"http"`
		WebSocket string `json:"ws"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// BrowserWrapper 浏览器包装器，支持多种模式
type BrowserWrapper struct {
	browser       *headless_browser.Browser
	rodBrowser    *rod.Browser
	isBitBrowser  bool
	bitbrowserWS  string
}

// NewBrowser 创建浏览器实例（支持 headless_browser 和 BitBrowser 模式）
func NewBrowser(headless bool, options ...Option) *BrowserWrapper {
	cfg := &browserConfig{}
	for _, opt := range options {
		opt(cfg)
	}

	opts := []headless_browser.Option{
		headless_browser.WithHeadless(headless),
	}
	if cfg.binPath != "" {
		opts = append(opts, headless_browser.WithChromeBinPath(cfg.binPath))
	}

	// Read proxy from environment variable
	if proxy := os.Getenv("XHS_PROXY"); proxy != "" {
		opts = append(opts, headless_browser.WithProxy(proxy))
		logrus.Infof("Using proxy: %s", maskProxyCredentials(proxy))
	}

	// 加载 cookies
	cookiePath := cookies.GetCookiesFilePath()
	cookieLoader := cookies.NewLoadCookie(cookiePath)

	if data, err := cookieLoader.LoadCookies(); err == nil {
		opts = append(opts, headless_browser.WithCookies(string(data)))
		logrus.Debugf("loaded cookies from filesuccessfully")
	} else {
		logrus.Warnf("failed to load cookies: %v", err)
	}

	return &BrowserWrapper{
		browser: headless_browser.New(opts...),
	}
}

// NewBitBrowser 通过 BitBrowser API 创建浏览器连接
func NewBitBrowser(apiURL, profileID, token string) (*BrowserWrapper, error) {
	// 调用 BitBrowser API 启动窗口
	wsURL, err := startBitBrowserWindow(apiURL, profileID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to start BitBrowser window: %w", err)
	}

	logrus.Infof("Connecting to BitBrowser: %s", wsURL)

	// 使用 go-rod 连接 BitBrowser
	browser := rod.New().ControlURL(wsURL).MustConnect()

	return &BrowserWrapper{
		rodBrowser:   browser,
		isBitBrowser: true,
		bitbrowserWS: wsURL,
	}, nil
}

// startBitBrowserWindow 调用 BitBrowser API 启动窗口
func startBitBrowserWindow(apiURL, profileID, token string) (string, error) {
	if apiURL == "" {
		apiURL = "http://127.0.0.1:54345"
	}

	// 构建请求体
	reqBody := map[string]interface{}{
		"id": profileID,
	}
	if token != "" {
		reqBody["authToken"] = token
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	url := fmt.Sprintf("%s/browser/open", apiURL)
	
	logrus.Infof("Starting BitBrowser window: %s", url)
	
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to call BitBrowser API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("BitBrowser API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result BitBrowserResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !result.Success {
		return "", fmt.Errorf("BitBrowser API failed: %s", result.Msg)
	}

	wsURL := result.Data.WebSocket
	if wsURL == "" {
		return "", fmt.Errorf("BitBrowser API returned empty WebSocket URL")
	}

	logrus.Infof("BitBrowser window started: %s", wsURL)
	return wsURL, nil
}

// NewPage 创建新页面
func (bw *BrowserWrapper) NewPage() *rod.Page {
	if bw.isBitBrowser && bw.rodBrowser != nil {
		return bw.rodBrowser.MustPage()
	}
	return bw.browser.NewPage()
}

// Close 关闭浏览器
func (bw *BrowserWrapper) Close() error {
	if bw.isBitBrowser && bw.rodBrowser != nil {
		return bw.rodBrowser.Close()
	}
	if bw.browser != nil {
		bw.browser.Close()
	}
	return nil
}

// GetCookies 获取 cookies
func (bw *BrowserWrapper) GetCookies() ([]*rod.Cookie, error) {
	if bw.isBitBrowser && bw.rodBrowser != nil {
		return bw.rodBrowser.GetCookies()
	}
	// headless_browser 不直接支持 GetCookies，需要通过页面获取
	page := bw.browser.NewPage()
	defer page.Close()
	return page.Browser().GetCookies()
}
