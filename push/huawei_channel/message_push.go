package huawei_channel

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/curltech/go-push-sdk/push/common/http"
	"github.com/curltech/go-push-sdk/push/common/intent"
	"github.com/curltech/go-push-sdk/push/common/json"
	"github.com/curltech/go-push-sdk/push/common/message"
	"github.com/curltech/go-push-sdk/push/errcode"
	"github.com/curltech/go-push-sdk/push/setting"
)

const (
	urlPush                = "https://api.push.hicloud.com/pushsend.do"
	nspSvcDefault          = "openpush.message.api.send"
	verDefault             = "1" // 用来解决大版本升级的兼容问题
	typeMsg                = 1   // 透传异步消息
	typeMsgNotificationBar = 3   // 系统通知栏异步消息
	typeActionCustom       = 1   // 自定义行为：行为由参数intent定义
	typeActionOpenUrl      = 2   // 打开URL：URL地址由参数url定义
	typeActionOpenApp      = 3   // 打开APP：默认值，打开App的首页
	deviceTokenMax         = 100 // 单次最大设备推送量
	deviceTokenMin         = 1
)

type PushClient struct {
	conf       setting.ConfigHuawei
	httpClient *http.Client
	authClient *AccessToken
}

func NewPushClient(conf setting.ConfigHuawei) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf:       conf,
		httpClient: http.NewClient(timeout),
		authClient: NewAccessToken(),
	}, nil
}

func checkConf(conf setting.ConfigHuawei) error {
	if conf.AppPkgName == "" {
		return errcode.ErrHuaweiAppPkgNameEmpty
	}
	if conf.ClientId == "" {
		return errcode.ErrHuaweiClientIdEmpty
	}
	if conf.ClientSecret == "" {
		return errcode.ErrHuaweiClientSecretEmpty
	}

	return nil
}

func (p *PushClient) checkParam(pushRequest *setting.PushMessageRequest) error {
	// 公共消息参数格式检查
	err := message.CheckMessageParam(pushRequest, deviceTokenMin, deviceTokenMax, true)
	if err != nil {
		return err
	}
	// 其余参数检查

	return nil
}

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	accessTokenReq := &AccessTokenReq{
		ClientId:     p.conf.ClientId,
		ClientSecret: p.conf.ClientSecret,
	}
	return p.authClient.Get(ctx, accessTokenReq)
}

func (p *PushClient) buildUrl() string {
	var buf strings.Builder
	buf.WriteString(urlPush)
	buf.WriteString("?")
	buf.WriteString("nsp_ctx=")
	nspCtx := &NspCtx{
		Ver:   verDefault,
		AppId: p.conf.ClientId,
	}
	nspCtxJson := json.MarshalToStringNoError(nspCtx)
	nspCtxUrlEncoding := url.QueryEscape(nspCtxJson)
	buf.WriteString(nspCtxUrlEncoding)
	urlStr := buf.String()
	return urlStr
}

func (p *PushClient) buildMessage(pushRequest *setting.PushMessageRequest) map[string]string {
	msgPayload := &PayLoad{
		Hps: &Hps{
			Msg: &Msg{
				Type: typeMsgNotificationBar,
				Body: &Body{
					Content: pushRequest.Message.Content,
					Title:   pushRequest.Message.Title,
				},
				Action: &Action{
					Type: typeActionOpenApp,
					Param: &Param{
						AppPkgName: p.conf.AppPkgName,
						Intent:     intent.GenerateIntent(p.conf.AppPkgName, pushRequest.Message.Extra),
					},
				},
			},
		},
	}

	msgMap := map[string]string{
		"access_token":      pushRequest.AccessToken,
		"nsp_svc":           nspSvcDefault,
		"nsp_ts":            strconv.FormatInt(time.Now().Local().Unix(), 10),
		"device_token_list": json.MarshalToStringNoError(pushRequest.DeviceTokens),
		"payload":           json.MarshalToStringNoError(msgPayload),
	}

	return msgMap
}

func (p *PushClient) PushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (interface{}, error) {
	errCheck := p.checkParam(pushRequest)
	if errCheck != nil {
		return nil, errCheck
	}
	return p.pushNotice(ctx, pushRequest)
}

func (p *PushClient) parseBody(body []byte) (*PushMessageResponse, error) {
	resp := &PushMessageResponse{}
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		return nil, errcode.ErrHuaweiParseBody
	}
	return resp, nil
}

func (p *PushClient) doRequest(ctx context.Context, uri string, data map[string]string) ([]byte, error) {

	return p.httpClient.PostForm(ctx, uri, data)
}

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*PushMessageResponse, error) {
	msg := p.buildMessage(pushRequest)
	pushUrl := p.buildUrl()
	body, err := p.doRequest(ctx, pushUrl, msg)
	if err != nil {
		return nil, err
	}

	return p.parseBody(body)
}
