package setting

type PlatformType string

const (
	HUAWEI_PLATFORM    PlatformType = "huawei"
	XIAOMI_PLATFORM    PlatformType = "xiaomi"
	MEIZU_PLATFORM     PlatformType = "meizu"
	VIVO_PLATFORM      PlatformType = "vivo"
	OPPO_PLATFORM      PlatformType = "oppo"
	IOS_CERT_PLATFORM  PlatformType = "ios"
	IOS_TOKEN_PLATFORM PlatformType = "ios-token"
)

type PushConfig struct {
	HUAWEI    `json:"huawei"`
	XIAOMI    `json:"xiaomi"`
	MEIZU     `json:"meizu"`
	OPPO      `json:"oppo"`
	VIVO      `json:"vivo"`
	IOS_CERT  `json:"ios"`
	IOS_TOKEN `json:"ios-token"`
}

type HUAWEI struct {
	AppPkgName   string `json:"appPkgName"`   // 应用包名
	ClientId     string `json:"clientId"`     // 用户在联盟申请的APPID
	ClientSecret string `json:"clientSecret"` // 应用秘钥
}

type XIAOMI struct {
	AppPkgName string `json:"appPkgName"`
	AppSecret  string `json:"appSecret"`
}

type MEIZU struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
}

type OPPO struct {
	AppPkgName   string `json:"appPkgName"`
	AppKey       string `json:"appKey"`
	MasterSecret string `json:"masterSecret"`
}

type VIVO struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
}

type IOS_CERT struct {
	IsSandBox bool   `json:"isSandBox"`
	CertPath  string `json:"certPath"`
	Password  string `json:"password"`
}

type IOS_TOKEN struct {
	IsSandBox  bool   `json:"isSandBox"`
	TeamId     string `json:"teamId"`
	KeyId      string `json:"keyId"`
	SecretFile string `json:"secretFile"`
	BundleId   string `json:"bundleId"`
}
