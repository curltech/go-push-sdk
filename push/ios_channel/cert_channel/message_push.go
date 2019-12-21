package cert_channel

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"gitee.com/cristiane/go-push-sdk/push/common/message"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/ios_channel"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

const (
	deviceTokenMax = 100
	deviceTokenMin = 1
)

func (p *PushClient) buildRequest(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {
	var (
		client *apns2.Client
	)
	cert, err := certificate.FromP12File(p.conf.CertPath, p.conf.Password)
	if err != nil {
		return nil, err
	}
	payloadStr := fmt.Sprintf(ios_channel.PayloadTemplate, pushRequest.Message.Title, pushRequest.Message.SubTitle, pushRequest.Message.Content, pushRequest.Message.Extra)
	
	notification := &apns2.Notification{
		DeviceToken: strings.Join(pushRequest.DeviceTokens, ","),
		ApnsID:      pushRequest.Message.BusinessId,
		CollapseID:  pushRequest.Message.BusinessId,
		Payload:     []byte(payloadStr),
	}
	
	if p.conf.IsSandBox {
		client = apns2.NewClient(cert).Development()
	} else {
		client = apns2.NewClient(cert).Production()
	}
	res, err := client.PushWithContext(ctx, notification)
	if err != nil {
		return nil, err
	}
	return &ios_channel.PushMessageResponse{
		StatusCode: res.StatusCode,
		APNsId:     res.ApnsID,
		Reason:     res.Reason,
	}, nil
}

type PushClient struct {
	conf *setting.PlatformAppleCert
}

func NewPushClient(conf *setting.PlatformAppleCert) (*PushClient, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf: conf,
	}, nil
}

func checkConf(conf *setting.PlatformAppleCert) error {
	if conf.CertPath == "" {
		return errcode.ErrCertPathEmpty
	}
	if conf.Password == "" {
		return errcode.ErrPasswordEmpty
	}
	return nil
}

func (p *PushClient) PushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {
	errCheck := p.checkParam(pushRequest)
	if errCheck != nil {
		return nil, errCheck
	}
	return p.pushNotice(ctx, pushRequest)
}

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {
	
	return p.buildRequest(ctx, pushRequest)
}

func (p *PushClient) checkParam(pushRequest *setting.PushMessageRequest) error {
	
	if err := message.CheckMessageParam(pushRequest, deviceTokenMin, deviceTokenMax, false); err != nil {
		return err
	}
	if pushRequest.Message.BusinessId == "" {
		return errcode.ErrBusinessIdEmpty
	}
	// 其余参数检查
	
	return nil
}

func (p *PushClient) GetAccessToken() (interface{}, error) {
	
	return nil, nil
}
