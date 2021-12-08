package oppo_channel

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/curltech/go-push-sdk/push/common/http"
	"github.com/curltech/go-push-sdk/push/common/intent"
	"github.com/curltech/go-push-sdk/push/common/json"
	"github.com/curltech/go-push-sdk/push/common/message"
	"github.com/curltech/go-push-sdk/push/errcode"
	"github.com/curltech/go-push-sdk/push/setting"
)

const (
	deviceTokenMax           = 1000
	deviceTokenMin           = 1
	clickActionTypeZero      = 0 // Launch App
	clickActionTypeFive      = 5 // Intent scheme URL
	targetTypeTwo            = 2
	urlBase                  = "https://api.push.oppomobile.com/server/v1"
	actionPush               = "message/notification/unicast"
	actionBroadcast          = "message/notification/broadcast"
	actionSaveMessageContent = "message/notification/save_message_content"
	pushTypeSingle           = 1
	pushTypeBroadcast        = 2
)

type PushClient struct {
	httpClient *http.Client
	conf       setting.ConfigOppo
	authClient *AuthToken
}

func NewPushClient(conf setting.ConfigOppo) (setting.PushClientInterface, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf:       conf,
		httpClient: http.NewClient(timeout),
		authClient: NewAuthToken(),
	}, nil
}

func checkConf(conf setting.ConfigOppo) error {
	if conf.AppPkgName == "" {
		return errcode.ErrOppoAppPkgNameEmpty
	}
	if conf.AppKey == "" {
		return errcode.ErrOppoAppKeyEmpty
	}
	if conf.MasterSecret == "" {
		return errcode.ErrOppoMasterSecretEmpty
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

	err := message.CheckMessageParam(pushRequest, deviceTokenMin, deviceTokenMax, true)
	if err != nil {
		return err
	}
	// 其余参数检查

	return nil
}

func (p *PushClient) pushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*PushMessageResponse, error) {
	msg := p.buildMessage(pushRequest)
	body, err := p.pushGateWay(ctx, pushRequest, msg)
	if err != nil {
		return nil, err
	}

	return p.parseBody(body)
}

func (p *PushClient) pushGateWay(ctx context.Context, pushRequest *setting.PushMessageRequest, message map[string]string) ([]byte, error) {
	if len(pushRequest.DeviceTokens) > deviceTokenMin {
		messageId, err := p.saveMessageToCloud(ctx, pushRequest.AccessToken, message)
		if err != nil {
			return nil, err
		}
		if messageId == "" {
			return nil, errcode.ErrOppoSaveMessageToCloud
		}
		body, err := p.pushBroadcast(ctx, messageId, pushRequest)
		return body, err
	} else {
		body, err := p.pushSingle(ctx, pushRequest, message)
		return body, err
	}
}

func (p *PushClient) buildMessage(pushRequest *setting.PushMessageRequest) map[string]string {

	//pushRequest.DeviceTokens = slice.RemoveDuplicateElement(pushRequest.DeviceTokens)
	messageMap := map[string]string{
		"title":             pushRequest.Message.Title,
		"sub_title":         pushRequest.Message.SubTitle,
		"content":           pushRequest.Message.Content,
		"click_action_type": strconv.Itoa(clickActionTypeZero),
		"click_action_url":  intent.GenerateIntent(p.conf.AppPkgName, pushRequest.Message.Extra),
	}
	if pushRequest.Message.CallBack != "" {
		messageMap["call_back_url"] = pushRequest.Message.CallBack
		if pushRequest.Message.CallbackParam != "" {
			messageMap["call_back_parameter"] = pushRequest.Message.CallbackParam
		}
	}

	return messageMap
}

func (p *PushClient) saveMessageToCloud(ctx context.Context, authToken string, message map[string]string) (string, error) {
	message["auth_token"] = authToken

	uri := p.buildSaveMessageContentPushUrl()

	body, err := p.httpClient.PostForm(ctx, uri, message)
	if err != nil {
		return "", err
	}
	result := &SaveMessageToCloudResponse{}
	errParse := json.UnmarshalByte(body, result)
	if errParse != nil {
		return "", errcode.ErrOppoParseBody
	}
	if result.Code == 11 {
		return "", errcode.ErrOppoParseBody
	}

	return result.Data.MessageId, nil
}

func (p *PushClient) pushBroadcast(ctx context.Context, messageId string, pushRequest *setting.PushMessageRequest) ([]byte, error) {

	msg := map[string]string{
		"message_id":   messageId,
		"target_type":  strconv.Itoa(targetTypeTwo),
		"target_value": strings.Join(pushRequest.DeviceTokens, ","),
		"auth_token":   pushRequest.AccessToken,
	}
	url := p.buildBroadcastPushUrl()

	return p.httpClient.PostForm(ctx, url, msg)
}

func (p *PushClient) pushSingle(ctx context.Context, pushRequest *setting.PushMessageRequest, message map[string]string) ([]byte, error) {

	param := map[string]string{
		"message": json.MarshalToStringNoError(&SingleMessage{
			TargetType:   targetTypeTwo,
			TargetValue:  strings.Join(pushRequest.DeviceTokens, ","),
			Notification: message,
		}),
		"auth_token": pushRequest.AccessToken,
	}
	uri := p.buildSinglePushUrl()

	return p.httpClient.PostForm(ctx, uri, param)
}

func (p *PushClient) buildSinglePushUrl() string {

	return fmt.Sprintf("%s/%s", urlBase, actionPush)
}

func (p *PushClient) buildBroadcastPushUrl() string {

	return fmt.Sprintf("%s/%s", urlBase, actionBroadcast)
}

func (p *PushClient) buildSaveMessageContentPushUrl() string {

	return fmt.Sprintf("%s/%s", urlBase, actionSaveMessageContent)
}

func (p *PushClient) parseBody(body []byte) (*PushMessageResponse, error) {
	resp := &PushMessageResponse{}
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		return nil, errcode.ErrOppoParseBody
	}

	return resp, nil
}

func (p *PushClient) GetAccessToken(ctx context.Context) (interface{}, error) {

	authTokenReq := &AuthTokenReq{
		AppKey:       p.conf.AppKey,
		MasterSecret: p.conf.MasterSecret,
	}
	return p.authClient.Get(ctx, authTokenReq)
}
