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

func (r *RegisterClient) GetHUAWEIClient() (*huawei_channel.PushClient, error) {

	return huawei_channel.NewPushClient(r.cfg.(*setting.PushConfig).HUAWEI)
}

func (r *RegisterClient) GetMEIZUClient() (*meizu_channel.PushClient, error) {

	return meizu_channel.NewPushClient(r.cfg.(*setting.PushConfig).MEIZU)
}

func (r *RegisterClient) GetXIAOMIClient() (*xiaomi_channel.PushClient, error) {

	return xiaomi_channel.NewPushClient(r.cfg.(*setting.PushConfig).XIAOMI)
}

func (r *RegisterClient) GetOPPOClient() (*oppo_channel.PushClient, error) {

	return oppo_channel.NewPushClient(r.cfg.(*setting.PushConfig).OPPO)
}

func (r *RegisterClient) GetVIVOClient() (*vivo_channel.PushClient, error) {

	return vivo_channel.NewPushClient(r.cfg.(*setting.PushConfig).VIVO)
}

func (r *RegisterClient) GetIosCertClient() (*cert_channel.PushClient, error) {

	return cert_channel.NewPushClient(r.cfg.(*setting.PushConfig).IOS_CERT)
}

func (r *RegisterClient) GetIosTokenClient() (*token_channel.PushClient, error) {

	return token_channel.NewPushClient(r.cfg.(*setting.PushConfig).IOS_TOKEN)
}
