package vivo_channel

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitee.com/cristiane/go-push-sdk/push/common/crypt"
	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
)

const (
	actionAuth = "message/auth"
)

type AuthToken struct {
	httpClient *http.Client
}

func NewAuthToken() *AuthToken {
	return &AuthToken{
		httpClient: http.NewClient(timeout),
	}
}

func (a *AuthToken) buildRequest(request *AuthTokenReq) map[string]string {

	request.Timestamp = strconv.FormatInt(time.Now().UTC().UnixNano()/(1e6), 10)

	return map[string]string{
		"appId":     request.AppId,
		"appKey":    request.AppKey,
		"timestamp": request.Timestamp,
		"sign":      a.generateSign(request),
	}
}

func (a *AuthToken) generateSign(request *AuthTokenReq) string {

	signStr := request.AppId + request.AppKey + request.Timestamp + request.AppSecret
	signStr = strings.Trim(signStr, "")

	return strings.ToLower(crypt.MD5([]byte(signStr)))
}

func (a *AuthToken) getUri() string {

	return fmt.Sprintf("%s/%s", urlBase, actionAuth)
}

func (a *AuthToken) Get(ctx context.Context, request *AuthTokenReq) (*AuthTokenResp, error) {
	errCheck := a.checkRequest(request)
	if errCheck != nil {
		return nil, errCheck
	}
	authUri := a.getUri()
	param := json.MarshalToStringNoError(a.buildRequest(request))
	body, err := a.httpClient.PostJson(ctx, authUri, param)
	if err != nil {
		return nil, err
	}

	return a.parseBody(body)
}

func (a *AuthToken) parseBody(body []byte) (*AuthTokenResp, error) {
	resp := &AuthTokenResp{}
	err := json.UnmarshalByte(body, resp)
	if err != nil {
		return nil, errcode.ErrVivoParseBody
	}

	return resp, nil
}

func (a *AuthToken) checkRequest(request *AuthTokenReq) error {
	if request.AppId == "" {
		return errcode.ErrVivoAppIdEmpty
	}
	if request.AppKey == "" {
		return errcode.ErrVivoAppKeyEmpty
	}
	if request.AppSecret == "" {
		return errcode.ErrVivoAppSecretEmpty
	}
	return nil
}
