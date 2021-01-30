package setting

type PlatformType string

const (
	HuaweiPlatform   PlatformType = "huawei"
	XiaomiPlatform   PlatformType = "xiaomi"
	MeizuPlatform    PlatformType = "meizu"
	VivoPlatform     PlatformType = "vivo"
	OppoPlatform     PlatformType = "oppo"
	IosCertPlatform  PlatformType = "ios"
	IosTokenPlatform PlatformType = "ios-token"
)

type PushConfig struct {
	ConfigHuawei   `json:"huawei"`
	ConfigXiaomi   `json:"xiaomi"`
	ConfigMeizu    `json:"meizu"`
	ConfigOppo     `json:"oppo"`
	ConfigVivo     `json:"vivo"`
	ConfigIosCert  `json:"ios"`
	ConfigIosToken `json:"ios-token"`
}

type ConfigHuawei struct {
	AppPkgName   string `json:"appPkgName"`   // 应用包名
	ClientId     string `json:"clientId"`     // 用户在联盟申请的APPID
	ClientSecret string `json:"clientSecret"` // 应用秘钥
}

type ConfigXiaomi struct {
	AppPkgName string `json:"appPkgName"`
	AppSecret  string `json:"appSecret"`
}

type ConfigMeizu struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
}

type ConfigOppo struct {
	AppPkgName   string `json:"appPkgName"`
	AppKey       string `json:"appKey"`
	MasterSecret string `json:"masterSecret"`
}

type ConfigVivo struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
}

type ConfigIosCert struct {
	IsSandBox bool   `json:"isSandBox"`
	CertPath  string `json:"certPath"`
	Password  string `json:"password"`
}

type ConfigIosToken struct {
	IsSandBox  bool   `json:"isSandBox"`
	TeamId     string `json:"teamId"`
	KeyId      string `json:"keyId"`
	SecretFile string `json:"secretFile"`
	BundleId   string `json:"bundleId"`
}
