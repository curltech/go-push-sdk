package huawei_channel

import (
	"context"
	"encoding/json"
	"log"
	
	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
)

const (
	urlAccessToken   = "https://login.cloud.huawei.com/oauth2/v2/token"
	grantTypeDefault = "client_credentials"
	scopeDefault     = "nsp.auth nsp.user nsp.vfs nsp.ping openpush.message"
	timeout          = 5
)

type AccessToken struct {
	httpClient *http.Client
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		httpClient: http.NewClient(timeout),
	}
}

func (a *AccessToken) buildRequest(request *AccessTokenReq) map[string]string {
	params := map[string]string{
		"grant_type":    request.GrantType,
		"client_id":     request.ClientId,
		"client_secret": request.ClientSecret,
		"scope":         request.Scope,
	}
	
	return params
}

func (a *AccessToken) parseBody(body []byte) (*AccessTokenResp, error) {
	resp := &AccessTokenResp{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		log.Printf("parseBody err: %v", err)
		return nil, errcode.ErrParseBody
	}
	return resp, nil
}

func (a *AccessToken) checkRequest(request *AccessTokenReq) error {
	if request.ClientId == "" {
		return errcode.ErrClientIdEmpty
	}
	if request.ClientSecret == "" {
		return errcode.ErrClientSecretEmpty
	}
	request.GrantType = grantTypeDefault
	
	return nil
}

func (a *AccessToken) Get(ctx context.Context,request *AccessTokenReq) (*AccessTokenResp, error) {
	errCheck := a.checkRequest(request)
	if errCheck != nil {
		return nil, errCheck
	}
	body, err := a.httpClient.PostForm(ctx,urlAccessToken, a.buildRequest(request))
	if err != nil {
		return nil, err
	}
	return a.parseBody(body)
}

