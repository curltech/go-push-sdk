package setting

type PlatformAll struct {
	Huawei   *PlatformHuawei     `json:"huawei"`
	Xiaomi   *PlatformXiaomi     `json:"xiaomi"`
	Meizu    *PlatformMeizu      `json:"meizu"`
	OPPO     *PlatformOPPO       `json:"oppo"`
	VIVO     *PlatformVIVO       `json:"vivo"`
	Ios      *PlatformAppleCert  `json:"ios"`
	IosToken *PlatformAppleToken `json:"ios-token"`
}

type PlatformHuawei struct {
	AppPkgName   string `json:"appPkgName"`   // 应用包名
	ClientId     string `json:"clientId"`     // 用户在联盟申请的APPID
	ClientSecret string `json:"clientSecret"` // 应用秘钥
}

type PlatformXiaomi struct {
	AppPkgName string `json:"appPkgName"`
	AppSecret  string `json:"appSecret"`
}

type PlatformMeizu struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppSecret  string `json:"appSecret"`
}

type PlatformOPPO struct {
	AppPkgName   string `json:"appPkgName"`
	AppKey       string `json:"appKey"`
	MasterSecret string `json:"masterSecret"`
}

type PlatformVIVO struct {
	AppPkgName string `json:"appPkgName"`
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	AppSecret  string `json:"appSecret"`
}

type PlatformAppleCert struct {
	IsSandBox bool   `json:"isSandBox"`
	CertPath  string `json:"certPath"`
	Password  string `json:"password"`
}

type PlatformAppleToken struct {
	IsSandBox  bool   `json:"isSandBox"`
	TeamId     string `json:"teamId"`
	KeyId      string `json:"keyId"`
	SecretFile string `json:"secretFile"`
	BundleId   string `json:"bundleId"`
}

