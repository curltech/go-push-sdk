package push

import (
	"errors"
	"log"

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
		log.Printf("[go-push-sdk] read conf file err: %v", err)
		return nil, errcode.ErrParseCfgFile
	}

	return NewRegisterClientWithConf(convert.Byte2Str(jsonByte))
}

func newRegisterClient(cfgJson string, obj interface{}) (*RegisterClient, error) {
	if cfgJson == "" {
		return nil, errors.New("jsonData is nil")
	}
	err := json.Unmarshal(cfgJson, obj)
	if err != nil {
		log.Printf("[go-push-sdk] parse json conf err: %v", err)
		return nil, errcode.ErrParseCfgFile
	}

	return &RegisterClient{
		cfg: obj,
	}, nil
}

func NewRegisterClientWithConf(cfgJson string) (*RegisterClient, error) {
	obj := &setting.PushConfig{}
	return newRegisterClient(cfgJson, obj)
}

func (r *RegisterClient) GetPlatformClient(platform string) (setting.PushClientInterface, error) {
	if platform == "huawei" {
		return r.GetHUAWEIClient()
	}
	if platform == "meizu" {
		return r.GetMEIZUClient()
	}
	if platform == "oppo" {
		return r.GetOPPOClient()
	}
	if platform == "vivo" {
		return r.GetVIVOClient()
	}
	if platform == "xiaomi" {
		return r.GetXIAOMIClient()
	}
	if platform == "ios" {
		return r.GetIosCertClient()
	}
	if platform == "ios-token" {
		return r.GetIosTokenClient()
	}
	return nil, errors.New("UNKNOWN REGISTER PLATFORM ？")
}

func (r *RegisterClient) GetHUAWEIClient() (setting.PushClientInterface, error) {

	return huawei_channel.NewPushClient(r.cfg.(*setting.PushConfig).HUAWEI)
}

func (r *RegisterClient) GetMEIZUClient() (setting.PushClientInterface, error) {

	return meizu_channel.NewPushClient(r.cfg.(*setting.PushConfig).MEIZU)
}

func (r *RegisterClient) GetXIAOMIClient() (setting.PushClientInterface, error) {

	return xiaomi_channel.NewPushClient(r.cfg.(*setting.PushConfig).XIAOMI)
}

func (r *RegisterClient) GetOPPOClient() (setting.PushClientInterface, error) {

	return oppo_channel.NewPushClient(r.cfg.(*setting.PushConfig).OPPO)
}

func (r *RegisterClient) GetVIVOClient() (setting.PushClientInterface, error) {

	return vivo_channel.NewPushClient(r.cfg.(*setting.PushConfig).VIVO)
}

func (r *RegisterClient) GetIosCertClient() (setting.PushClientInterface, error) {

	return cert_channel.NewPushClient(r.cfg.(*setting.PushConfig).IOS_CERT)
}

func (r *RegisterClient) GetIosTokenClient() (setting.PushClientInterface, error) {

	return token_channel.NewPushClient(r.cfg.(*setting.PushConfig).IOS_TOKEN)
}
