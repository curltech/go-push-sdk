package xiaomi_channel

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/common/intent"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/common/message"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

const (
	urlPush             = "https://api.xmpush.xiaomi.com/v3/message/regid"
	urlPushTest         = "https://sandbox.xmpush.xiaomi.com/v2/message/regid"
	timeout             = 5
	passThroughTypeZero = 0   // 0 表示通知栏消息
	passThroughTypeOne  = 1   // 1 表示透传消息
	notifyEffectTwo     = "2" // 通知栏点击后打开app的任一Activity
	deviceTokenMax      = 100 // 单次最大设备推送量
	deviceTokenMin      = 1
)

type PushClient struct {
	httpClient *http.Client
	conf       setting.XIAOMI
}

func NewPushClient(conf setting.XIAOMI) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf:       conf,
		httpClient: http.NewClient(timeout),
	}, nil
}

func checkConf(conf setting.XIAOMI) error {
	if conf.AppPkgName == "" {
		return errcode.ErrAppPkgNameEmpty
	}
	if conf.AppSecret == "" {
		return errcode.ErrAppSecretEmpty
	}
	return nil
}

func (p *PushClient) checkParam(pushRequest *setting.PushMessageRequest) error {

	if err := message.CheckMessageParam(pushRequest, deviceTokenMin, deviceTokenMax, false); err != nil {
		return err
	}
	// 其余参数检查

	return nil
}

func (p *PushClient) PushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (interface{}, error) {
	errCheck := p.checkParam(pushRequest)
	if errCheck != nil {
		return nil, errCheck
	}

	return p.pushNotice(ctx, pushRequest)
}

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*PushMessageResponse, error) {
	msg := p.buildMessage(pushRequest)
	pushUrl := p.buildUrl()
	body, err := p.buildRequest(ctx, pushUrl, msg)
	if err != nil {
		return nil, err
	}

	return p.parseBody(body)
}

func (p *PushClient) buildMessage(pushRequest *setting.PushMessageRequest) map[string]string {

	messageMap := map[string]string{
		"payload":                 url.QueryEscape(pushRequest.Message.Content),
		"restricted_package_name": p.conf.AppPkgName,
		"pass_through":            strconv.Itoa(passThroughTypeZero),
		"title":                   pushRequest.Message.Title,
		"description":             pushRequest.Message.Content,
		"registration_id":         strings.Join(pushRequest.DeviceTokens, ","),
		"extra.notify_effect":     notifyEffectTwo,
		"extra.intent_uri":        intent.GenerateIntent(p.conf.AppPkgName, pushRequest.Message.Extra),
	}
	if pushRequest.Message.CallBack != "" {
		messageMap["extra.callback"] = pushRequest.Message.CallBack
		if pushRequest.Message.CallbackParam != "" {
			messageMap["extra.callback.param"] = pushRequest.Message.CallbackParam
		}
	}

	return messageMap
}

func (p *PushClient) parseBody(body []byte) (*PushMessageResponse, error) {
	resp := &PushMessageResponse{}
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		log.Printf("[go-push-sdk] xiaomi message push parseBody err: %v", err)
		return nil, errcode.ErrParseBody
	}
	return resp, nil
}

func (p *PushClient) buildRequest(ctx context.Context, uri string, data map[string]string) ([]byte, error) {
	request, err := p.httpClient.BuildRequest(ctx, "POST", uri, data)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("key=%s", p.conf.AppSecret))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return p.httpClient.Do(ctx, request)
}

func (p *PushClient) buildUrl() string {

	return urlPush
}

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	return nil, nil
}
