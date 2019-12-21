package setting

type PushMessageRequest struct {
	DeviceTokens []string `json:"token"`                  // 设备列表
	AccessToken  string   `json:"access_token,omitempty"` // 认证token
	Message      *Message `json:"message"`                // 消息
	ExpireTime   string   `json:"expire_time,omitempty"`  // 消息超时时间，必填
}

type Message struct {
	BusinessId    string            `json:"businessId"`    // 业务ID
	Title         string            `json:"title"`         // 标题，建议不超过10个汉字
	SubTitle      string            `json:"subTitle"`      // 副标题，建议不超过10个汉字
	Content       string            `json:"content"`       // 内容，建议不超过20个汉字
	Extra         map[string]string `json:"extra"`         // 自定义消息。只支持一维
	CallBack      string            `json:"callback"`      // 送达回执地址，供推送厂商调用，最大128字节
	CallbackParam string            `json:"callbackParam"` // 自定义回执参数
}

type PushMessageResponse struct {
}

