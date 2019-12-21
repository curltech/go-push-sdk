package huawei_channel

type PushMessageRequest struct {
	AccessToken     string   `json:"access_token"`          // 认证token，必填
	NspSvc          string   `json:"nsp_svc"`               // 默认 nspSvcDefault，必填
	NspTs           string   `json:"nsp_ts"`                // 服务请求时间戳，必填
	DeviceTokenList []string `json:"device_token_list"`     // JSON数值字符串，单次最大100
	ExpireTime      string   `json:"expire_time,omitempty"` // 消息超时时间，必填
	PayLoad         *PayLoad `json:"payload,omitempty"`     // 描述投递消息JSON结构体,必填
}

type PushMessageResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type NspCtx struct {
	Ver   string `json:"ver"`
	AppId string `json:"appId"`
}

type PayLoad struct {
	Hps *Hps `json:"hps"`
}

type Hps struct {
	Msg *Msg `json:"msg"`
	Ext *Ext `json:"ext,omitempty"`
}

type Msg struct {
	Type   int     `json:"type"` // 1 透传异步消息，2系统通知栏异步消息
	Body   *Body   `json:"body"` // 消息内容,对于透传类消息可以是字符串或JSON
	Action *Action `json:"action,omitempty"`
}

type Body struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

type Action struct {
	Type  int    `json:"type,omitempty"`
	Param *Param `json:"param,omitempty"`
}

type Param struct {
	Intent     string `json:"intent,omitempty"`
	Url        string `json:"url,omitempty"`
	AppPkgName string `json:"appPkgName,omitempty"`
}

type Customize struct {
	Season  string `json:"season"`
	Weather string `json:"weather"`
}

type Ext struct {
	BiTag       string     `json:"biTag,omitempty"`
	Customize   *Customize `json:"customize,omitempty"`
	BadgeAddNum string     `json:"badgeAddNum,omitempty"`
	BadgeClass  string     `json:"badgeClass,omitempty"`
}

type AccessTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope,omitempty"`
}

type AccessTokenResp struct {
	AccessToken      string `json:"access_token"`      //返回access_token
	ExpiresIn        int    `json:"expires_in"`        //单位是秒
	Scope            string `json:"scope"`             // access_token访问范围
	Error            int    `json:"error"`             // 返回错误码
	ErrorDescription string `json:"error_description"` // 返回错误描述信息
}

