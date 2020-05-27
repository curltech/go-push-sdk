package push

import (
	"errors"
	"log"

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
	conf *setting.PushConfig
}

func NewRegisterClient(configFilePath string) (*RegisterClient, error) {
	if configFilePath == "" {
		return nil, errcode.ErrConfFileEmpty
	}
	fileRead := file.NewFileRead()
	jsonByte, err := fileRead.Read(configFilePath)
	if err != nil {
		log.Printf("[go-push-sdk] read conf file err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}

	conf := &setting.PushConfig{}
	err = json.UnmarshalByte(jsonByte, conf)
	if err != nil {
		log.Printf("[go-push-sdk] parse conf file err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}

	return &RegisterClient{
		conf: conf,
	}, nil
}

func NewRegisterClientWithConf(jsonData string) (*RegisterClient, error) {
	if jsonData == "" {
		return nil, errors.New("jsonData is nil")
	}
	conf := &setting.PushConfig{}
	err := json.Unmarshal(jsonData, conf)
	if err != nil {
		log.Printf("[go-push-sdk] parse json conf err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}

	return &RegisterClient{
		conf: conf,
	}, nil
}

func (r *RegisterClient) GetHuaweiClient() (*huawei_channel.PushClient, error) {

	return huawei_channel.NewPushClient(r.conf.HUAWEI)
}

func (r *RegisterClient) GetMeizuClient() (*meizu_channel.PushClient, error) {

	return meizu_channel.NewPushClient(r.conf.MEIZU)
}

func (r *RegisterClient) GetXiaomiClient() (*xiaomi_channel.PushClient, error) {

	return xiaomi_channel.NewPushClient(r.conf.XIAOMI)
}

func (r *RegisterClient) GetOPPOClient() (*oppo_channel.PushClient, error) {

	return oppo_channel.NewPushClient(r.conf.OPPO)
}

func (r *RegisterClient) GetVIVOClient() (*vivo_channel.PushClient, error) {

	return vivo_channel.NewPushClient(r.conf.VIVO)
}

func (r *RegisterClient) GetIosCertClient() (*cert_channel.PushClient, error) {

	return cert_channel.NewPushClient(r.conf.IOS_CERT)
}

func (r *RegisterClient) GetIosTokenClient() (*token_channel.PushClient, error) {

	return token_channel.NewPushClient(r.conf.IOS_TOKEN)
}
