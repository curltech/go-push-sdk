package token_channel

import (
	"context"
	"fmt"
	"github.com/curltech/go-push-sdk/push/common/convert"
	"strings"

	"github.com/curltech/go-push-sdk/push/common/message"
	"github.com/curltech/go-push-sdk/push/errcode"
	"github.com/curltech/go-push-sdk/push/ios_channel"
	"github.com/curltech/go-push-sdk/push/setting"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

const (
	deviceTokenMax = 100
	deviceTokenMin = 1
)

type PushClient struct {
	conf setting.ConfigIosToken
}

func NewPushClient(conf setting.ConfigIosToken) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf: conf,
	}, nil
}

func checkConf(conf setting.ConfigIosToken) error {
	if conf.TeamId == "" {
		return errcode.ErrIosTeamIdEmpty
	}
	if conf.KeyId == "" {
		return errcode.ErrIosKeyIdEmpty
	}
	if conf.SecretFile == "" {
		return errcode.ErrIosSecretFileEmpty
	}
	if conf.BundleId == "" {
		return errcode.ErrIosBundleIdEmpty
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

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {

	return p.buildRequest(ctx, pushRequest)
}

func (p *PushClient) buildRequest(ctx context.Context, pushRequest *setting.PushMessageRequest) (*ios_channel.PushMessageResponse, error) {
	var (
		client *apns2.Client
	)
	authKey, err := token.AuthKeyFromFile(p.conf.SecretFile)
	if err != nil {
		return nil, err
	}
	tokenClient := &token.Token{
		AuthKey: authKey,
		KeyID:   p.conf.KeyId,
		TeamID:  p.conf.TeamId,
	}
	payloadStr := fmt.Sprintf(ios_channel.PayloadTemplate, pushRequest.Message.Title, pushRequest.Message.SubTitle, pushRequest.Message.Content,
		pushRequest.Message.Extra)

	notification := &apns2.Notification{
		CollapseID:  pushRequest.Message.BusinessId,
		ApnsID:      pushRequest.Message.BusinessId,
		DeviceToken: strings.Join(pushRequest.DeviceTokens, ","),
		Topic:       p.conf.BundleId,
		Payload:     convert.Str2Byte(payloadStr),
	}
	if p.conf.IsSandBox {
		client = apns2.NewTokenClient(tokenClient).Development()
	} else {
		client = apns2.NewTokenClient(tokenClient).Production()
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

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	return nil, nil
}
