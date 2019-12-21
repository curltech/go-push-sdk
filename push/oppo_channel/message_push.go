package oppo_channel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	
	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/common/intent"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/common/message"
	"gitee.com/cristiane/go-push-sdk/push/common/slice"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

const (
	deviceTokenMax           = 1000
	deviceTokenMin           = 1
	clickActionTypeFive      = 5
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
	conf       *setting.PlatformOPPO
}

func NewPushClient(conf *setting.PlatformOPPO) (*PushClient, error) {
	errCheck := checkConf(conf)
	if errCheck != nil {
		return nil, errCheck
	}
	return &PushClient{
		conf:       conf,
		httpClient: http.NewClient(timeout),
	}, nil
}

func checkConf(conf *setting.PlatformOPPO) error {
	if conf.AppPkgName == "" {
		return errcode.ErrAppPkgNameEmpty
	}
	if conf.AppKey == "" {
		return errcode.ErrAppKeyEmpty
	}
	if conf.MasterSecret == "" {
		return errcode.ErrMasterSecretEmpty
	}
	
	return nil
}

func (p *PushClient) PushNotice(ctx context.Context, pushRequest *setting.PushMessageRequest) (*PushMessageResponse, error) {
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
			return nil, errcode.ErrSaveMessageToCloud
		}
		body, err := p.pushBroadcast(ctx, messageId, pushRequest)
		return body, err
	} else {
		body, err := p.pushSingle(ctx, pushRequest, message)
		return body, err
	}
}

func (p *PushClient) buildMessage(pushRequest *setting.PushMessageRequest) map[string]string {
	
	pushRequest.DeviceTokens = slice.RemoveDuplicateElement(pushRequest.DeviceTokens)
	messageAll := &PushMessageRequest{
		Title:             pushRequest.Message.Title,
		SubTitle:          pushRequest.Message.SubTitle,
		Content:           pushRequest.Message.Content,
		ClickActionType:   clickActionTypeFive,
		ClickActionUrl:    intent.GenerateIntent(p.conf.AppPkgName, pushRequest.Message.Extra),
		CallBackUrl:       pushRequest.Message.CallBack,
		CallBackParameter: pushRequest.Message.CallbackParam,
	}
	
	messageMap := map[string]string{
		"title":             messageAll.Title,
		"sub_title":         messageAll.SubTitle,
		"content":           messageAll.Content,
		"click_action_type": strconv.Itoa(messageAll.ClickActionType),
		"click_action_url":  messageAll.ClickActionUrl,
	}
	if messageAll.CallBackUrl != "" {
		messageMap["call_back_url"] = messageAll.CallBackUrl
		if messageAll.CallBackParameter != "" {
			messageMap["call_back_parameter"] = messageAll.CallBackParameter
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
		log.Printf("parse saveMessageToCloud Response err: %v", errParse)
		return "", errcode.ErrParseBody
	}
	if result.Code == 11 {
		return "", errors.New(result.Message)
	}
	
	return result.Data.MessageId, nil
}

func (p *PushClient) pushBroadcast(ctx context.Context, messageId string, pushRequest *setting.PushMessageRequest) ([]byte, error) {
	broadcastPushMessage := &BroadcastPush{
		MessageId:   messageId,
		TargetType:  targetTypeTwo,
		TargetValue: strings.Join(pushRequest.DeviceTokens, ","),
		AuthToken:   pushRequest.AccessToken,
	}
	msg := map[string]string{
		"message_id":   broadcastPushMessage.MessageId,
		"target_type":  strconv.Itoa(broadcastPushMessage.TargetType),
		"target_value": broadcastPushMessage.TargetValue,
		"auth_token":   broadcastPushMessage.AuthToken,
	}
	url := p.buildBroadcastPushUrl()
	
	return p.httpClient.PostForm(ctx, url, msg)
}

func (p *PushClient) pushSingle(ctx context.Context, pushRequest *setting.PushMessageRequest, message map[string]string) ([]byte, error) {
	singlePushMessage := &SinglePush{
		AuthToken: pushRequest.AccessToken,
		Message: &SingleMessage{
			TargetType:   targetTypeTwo,
			TargetValue:  strings.Join(pushRequest.DeviceTokens, ","),
			Notification: message,
		},
	}
	
	param := map[string]string{
		"message":    json.MarshalToStringNoError(singlePushMessage.Message),
		"auth_token": singlePushMessage.AuthToken,
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
	log.Printf("body : %v", string(body))
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		log.Printf("parseBody err: %v", err)
		return nil, errcode.ErrParseBody
	}
	
	return resp, nil
}

func (p *PushClient) GetAccessToken(ctx context.Context) (*AuthTokenResp, error) {
	authToken := NewAuthToken()
	authTokenReq := &AuthTokenReq{
		AppKey:       p.conf.AppKey,
		MasterSecret: p.conf.MasterSecret,
	}
	return authToken.Get(ctx, authTokenReq)
}

