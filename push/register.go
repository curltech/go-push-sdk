package push

import (
	"errors"
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
	"log"
)

const (
	defaultGlobalFileName = "/usr/local/etc/go-push-sdk/setting.json"
)

type RegisterClient struct {
	configPlatform *setting.PlatformAll
}

func NewRegisterClient() (*RegisterClient, error) {
	fileRead := file.NewFileRead()
	jsonByte, err := fileRead.Read(defaultGlobalFileName)
	if err != nil {
		log.Printf("read setting.json file err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}
	config := &setting.PlatformAll{}
	err = json.UnmarshalByte(jsonByte, config)
	if err != nil {
		log.Printf("parse setting.json file err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}
	
	return &RegisterClient{
		configPlatform: config,
	}, nil
}

func NewRegisterClientWithJsonData(jsonData string) (*RegisterClient, error) {
	if jsonData == "" {
		return nil, errors.New("jsonData is nil")
	}
	config := &setting.PlatformAll{}
	err := json.Unmarshal(jsonData, config)
	if err != nil {
		log.Printf("parse setting.json file err: %v", err)
		return nil, errcode.ErrParseSettingFile
	}
	
	return &RegisterClient{
		configPlatform: config,
	}, nil
}

func (r *RegisterClient) GetHuaweiClient() (*huawei_channel.PushClient, error) {
	
	return huawei_channel.NewPushClient(r.configPlatform.Huawei)
}

func (r *RegisterClient) GetMeizuClient() (*meizu_channel.PushClient, error) {
	
	return meizu_channel.NewPushClient(r.configPlatform.Meizu)
}

func (r *RegisterClient) GetXiaomiClient() (*xiaomi_channel.PushClient, error) {
	
	return xiaomi_channel.NewPushClient(r.configPlatform.Xiaomi)
}

func (r *RegisterClient) GetOPPOClient() (*oppo_channel.PushClient, error) {
	
	return oppo_channel.NewPushClient(r.configPlatform.OPPO)
}

func (r *RegisterClient) GetVIVOClient() (*vivo_channel.PushClient, error) {
	
	return vivo_channel.NewPushClient(r.configPlatform.VIVO)
}

func (r *RegisterClient) GetIosCertClient() (*cert_channel.PushClient, error) {
	
	return cert_channel.NewPushClient(r.configPlatform.Ios)
}

func (r *RegisterClient) GetIosTokenClient() (*token_channel.PushClient, error) {
	
	return token_channel.NewPushClient(r.configPlatform.IosToken)
}
