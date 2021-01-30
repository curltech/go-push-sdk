package cert_channel

import (
	"context"
	"fmt"
	"strings"

	"gitee.com/cristiane/go-push-sdk/push/common/convert"
	"gitee.com/cristiane/go-push-sdk/push/common/message"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/ios_channel"
	"gitee.com/cristiane/go-push-sdk/push/setting"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

const (
	deviceTokenMax = 100
	deviceTokenMin = 1
)

type PushClient struct {
	conf setting.ConfigIosCert
}

func NewPushClient(conf setting.ConfigIosCert) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf: conf,
	}, nil
}

func (p *PushClient) buildRequest(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {
	var (
		client *apns2.Client
	)
	cert, err := certificate.FromP12File(p.conf.CertPath, p.conf.Password)
	if err != nil {
		return nil, err
	}
	payloadStr := fmt.Sprintf(ios_channel.PayloadTemplate, pushRequest.Message.Title, pushRequest.Message.SubTitle, pushRequest.Message.Content,
		pushRequest.Message.Extra)

	notification := &apns2.Notification{
		DeviceToken: strings.Join(pushRequest.DeviceTokens, ","),
		ApnsID:      pushRequest.Message.BusinessId,
		CollapseID:  pushRequest.Message.BusinessId,
		Payload:     convert.Str2Byte(payloadStr),
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

func checkConf(conf setting.ConfigIosCert) error {
	if conf.CertPath == "" {
		return errcode.ErrIosCertPathEmpty
	}
	if conf.Password == "" {
		return errcode.ErrIosPasswordEmpty
	}
	return nil
}

func (p *PushClient) PushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (interface{}, error) {
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

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	return nil, nil
}
