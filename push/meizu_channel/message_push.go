package meizu_channel

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"gitee.com/cristiane/go-push-sdk/push/common/crypt"
	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/common/intent"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/common/message"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

const (
	deviceTokenMax = 1000
	deviceTokenMin = 1
	timeout        = 5
	clickTypeTwo   = 2
	urlPush        = "http://server-api-push.meizu.com/garcia/api/server/push/varnished/pushByPushId"
)

type PushClient struct {
	httpClient *http.Client
	conf       setting.MEIZU
}

func NewPushClient(conf setting.MEIZU) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf:       conf,
		httpClient: http.NewClient(timeout),
	}, nil
}

func checkConf(conf setting.MEIZU) error {
	if conf.AppPkgName == "" {
		return errcode.ErrAppPkgNameEmpty
	}
	if conf.AppId == "" {
		return errcode.ErrAppIdEmpty
	}
	if conf.AppSecret == "" {
		return errcode.ErrAppSecretEmpty
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

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*PushMessageResponse, error) {
	msg := p.buildMessage(pushRequest)
	pushUrl := p.buildUrl()
	body, err := p.buildRequest(ctx, pushUrl, msg)
	resp, err := p.parseBody(body)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (p *PushClient) buildRequest(ctx context.Context, uri string, data map[string]string) ([]byte, error) {

	return p.httpClient.PostForm(ctx, uri, data)
}

func (p *PushClient) buildMessage(pushRequest *setting.PushMessageRequest) map[string]string {
	payload := &Payload{
		NoticeBarInfo: &NoticeBarInfo{
			Title:   pushRequest.Message.Title,
			Content: pushRequest.Message.Content,
		},
		ClickTypeInfo: &ClickTypeInfo{
			ClickType: clickTypeTwo,
			Url:       intent.GenerateIntent(p.conf.AppPkgName, pushRequest.Message.Extra),
		},
	}
	if pushRequest.Message.CallBack != "" {
		payload.Extra = &Extra{}
		payload.Extra.Callback = pushRequest.Message.CallBack
		if pushRequest.Message.CallbackParam != "" {
			payload.Extra.CallbackParam = pushRequest.Message.CallbackParam
		}
	}
	msg := map[string]string{
		"appId":       p.conf.AppId,
		"pushIds":     strings.Join(pushRequest.DeviceTokens, ","),
		"messageJson": json.MarshalToStringNoError(payload),
	}
	msg["sign"] = p.generateSign(msg)

	p.generateSign(msg)

	return msg
}

func (p *PushClient) generateSign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key, _ := range params {
		keys = append(keys, key)
	}

	str := ""
	sort.Strings(keys)
	for _, k := range keys {
		str += fmt.Sprintf("%v=%v", k, params[k])
	}
	str += p.conf.AppSecret

	return crypt.MD5([]byte(str))
}

func (p *PushClient) buildUrl() string {

	return urlPush
}

func (p *PushClient) parseBody(body []byte) (*PushMessageResponse, error) {
	resp := &PushMessageResponse{}
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		log.Printf("[go-push-sdk] meizu message push parseBody err: %v", err)
		return nil, err
	}
	return resp, nil
}

func (p *PushClient) checkParam(pushRequest *setting.PushMessageRequest) error {

	err := message.CheckMessageParam(pushRequest, deviceTokenMin, deviceTokenMax, false)
	if err != nil {
		return err
	}
	// 其余参数检查

	return nil
}

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	return nil, nil
}
