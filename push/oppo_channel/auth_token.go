package oppo_channel

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gitee.com/cristiane/go-push-sdk/push/common/crypt"
	"gitee.com/cristiane/go-push-sdk/push/common/http"
	"gitee.com/cristiane/go-push-sdk/push/errcode"
)

const (
	timeout    = 5
	actionAuth = "auth"
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

	request.Timestamp = strconv.FormatInt(time.Now().UnixNano()/(1e6), 10)

	return map[string]string{
		"app_key":   request.AppKey,
		"timestamp": request.Timestamp,
		"sign":      a.generateSign(request),
	}
}

func (a *AuthToken) parseBody(body []byte) (*AuthTokenResp, error) {
	resp := &AuthTokenResp{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		return nil, errcode.ErrOppoParseBody
	}

	return resp, nil
}

func (a *AuthToken) checkRequest(request *AuthTokenReq) error {
	if request.AppKey == "" {
		return errcode.ErrOppoAppKeyEmpty
	}
	if request.MasterSecret == "" {
		return errcode.ErrOppoMasterSecretEmpty
	}
	return nil
}

func (a *AuthToken) getUri() string {

	return fmt.Sprintf("%s/%s", urlBase, actionAuth)
}

func (a *AuthToken) generateSign(request *AuthTokenReq) string {

	signStr := request.AppKey + request.Timestamp + request.MasterSecret

	return crypt.SHA256([]byte(signStr))
}

func (a *AuthToken) Get(ctx context.Context, request *AuthTokenReq) (*AuthTokenResp, error) {
	errCheck := a.checkRequest(request)
	if errCheck != nil {
		return nil, errCheck
	}
	authUri := a.getUri()

	body, err := a.httpClient.PostForm(ctx, authUri, a.buildRequest(request))
	if err != nil {
		return nil, err
	}

	return a.parseBody(body)
}
