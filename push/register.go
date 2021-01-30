package push

import (
	"gitee.com/cristiane/go-push-sdk/push/common/convert"
	"gitee.com/cristiane/go-push-sdk/push/common/file"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/huawei_channel"
	"gitee.com/cristiane/go-push-sdk/push/ios_channel/cert_channel"
	"gitee.com/cristiane/go-push-sdk/push/ios_channel/token_channel"
	"gitee.com/cristiane/go-push-sdk/push/meizu_channel"
	"gitee.com/cristiane/go-push-sdk/push/oppo_channel"
	"gitee.com/cristiane/go-push-sdk/push/setting"
	"gitee.com/cristiane/go-push-sdk/push/vivo_channel"
	"gitee.com/cristiane/go-push-sdk/push/xiaomi_channel"
)

const (
	// 推荐的配置文件存放路径
	DefaultConfFile = "/usr/local/etc/go-push-sdk/setting.json"
)

type RegisterClient struct {
	cfg interface{}
}

func NewRegisterClient(configFilePath string) (*RegisterClient, error) {
	if configFilePath == "" {
		return nil, errcode.ErrCfgFileEmpty
	}
	fileRead := file.NewFileRead()
	jsonByte, err := fileRead.Read(configFilePath)
	if err != nil {
		return nil, errcode.ErrParseConfigFile
	}

	return NewRegisterClientWithConf(convert.Byte2Str(jsonByte), "")
}

func newRegisterClient(cfgJson string, obj interface{}) (*RegisterClient, error) {
	if cfgJson == "" {
		return nil, errcode.ErrConfigEmpty
	}
	err := json.Unmarshal(cfgJson, obj)
	if err != nil {
		return nil, errcode.ErrParseConfigFile
	}

	return &RegisterClient{
		cfg: obj,
	}, nil
}

func NewRegisterClientWithConf(cfgJson string, platformType setting.PlatformType) (*RegisterClient, error) {
	var obj interface{}
	switch platformType {
	case setting.HuaweiPlatform:
		obj = &setting.ConfigHuawei{}
	case setting.MeizuPlatform:
		obj = &setting.ConfigMeizu{}
	case setting.OppoPlatform:
		obj = &setting.ConfigOppo{}
	case setting.VivoPlatform:
		obj = &setting.ConfigVivo{}
	case setting.XiaomiPlatform:
		obj = &setting.ConfigXiaomi{}
	case setting.IosCertPlatform:
		obj = &setting.ConfigIosCert{}
	case setting.IosTokenPlatform:
		obj = &setting.ConfigIosToken{}
	default:
		obj = &setting.PushConfig{}
	}
	return newRegisterClient(cfgJson, obj)
}

func (r *RegisterClient) GetPlatformClient(platform setting.PlatformType) (setting.PushClientInterface, error) {
	switch platform {
	case setting.HuaweiPlatform:
		return r.GetHUAWEIClient()
	case setting.MeizuPlatform:
		return r.GetMEIZUClient()
	case setting.OppoPlatform:
		return r.GetOPPOClient()
	case setting.VivoPlatform:
		return r.GetVIVOClient()
	case setting.XiaomiPlatform:
		return r.GetXIAOMIClient()
	case setting.IosCertPlatform:
		return r.GetIosCertClient()
	case setting.IosTokenPlatform:
		return r.GetIosTokenClient()
	default:
		return nil, errcode.ErrUnknownPlatform
	}
}

func (r *RegisterClient) GetHUAWEIClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigHuawei); ok {
		return huawei_channel.NewPushClient(*conf)
	}
	return huawei_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigHuawei)
}

func (r *RegisterClient) GetMEIZUClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigMeizu); ok {
		return meizu_channel.NewPushClient(*conf)
	}
	return meizu_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigMeizu)
}

func (r *RegisterClient) GetXIAOMIClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigXiaomi); ok {
		return xiaomi_channel.NewPushClient(*conf)
	}
	return xiaomi_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigXiaomi)
}

func (r *RegisterClient) GetOPPOClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigOppo); ok {
		return oppo_channel.NewPushClient(*conf)
	}
	return oppo_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigOppo)
}

func (r *RegisterClient) GetVIVOClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigVivo); ok {
		return vivo_channel.NewPushClient(*conf)
	}
	return vivo_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigVivo)
}

func (r *RegisterClient) GetIosCertClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigIosCert); ok {
		return cert_channel.NewPushClient(*conf)
	}
	return cert_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigIosCert)
}

func (r *RegisterClient) GetIosTokenClient() (setting.PushClientInterface, error) {
	if conf, ok := r.cfg.(*setting.ConfigIosToken); ok {
		return token_channel.NewPushClient(*conf)
	}
	return token_channel.NewPushClient(r.cfg.(*setting.PushConfig).ConfigIosToken)
}
