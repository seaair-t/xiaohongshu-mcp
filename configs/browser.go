package configs

var (
	useHeadless = true

	binPath = ""
	
	// 浏览器模式: headless, chrome, bitbrowser
	browserMode = "headless"
	
	// BitBrowser 配置
	bitbrowserAPI     = ""
	bitbrowserProfile = ""
	bitbrowserToken   = ""
)

func InitHeadless(h bool) {
	useHeadless = h
}

// IsHeadless 是否无头模式。
func IsHeadless() bool {
	return useHeadless
}

func SetBinPath(b string) {
	binPath = b
}

func GetBinPath() string {
	return binPath
}

// SetBrowserMode 设置浏览器模式
func SetBrowserMode(mode string) {
	browserMode = mode
}

// GetBrowserMode 获取浏览器模式
func GetBrowserMode() string {
	return browserMode
}

// IsBitBrowserMode 是否使用 BitBrowser 模式
func IsBitBrowserMode() bool {
	return browserMode == "bitbrowser"
}

// SetBitBrowserConfig 设置 BitBrowser 配置
func SetBitBrowserConfig(api, profile, token string) {
	bitbrowserAPI = api
	bitbrowserProfile = profile
	bitbrowserToken = token
}

// GetBitBrowserAPI 获取 BitBrowser API 地址
func GetBitBrowserAPI() string {
	return bitbrowserAPI
}

// GetBitBrowserProfile 获取 BitBrowser Profile ID
func GetBitBrowserProfile() string {
	return bitbrowserProfile
}

// GetBitBrowserToken 获取 BitBrowser 认证 Token
func GetBitBrowserToken() string {
	return bitbrowserToken
}
