package meizu_channel

type PushMessageRequest struct {
	AppId       string `json:"appId"`
	PushIds     string `json:"pushIds"`
	MessageJson string `json:"messageJson"`
	Sign        string `json:"sign"`
}

type Payload struct {
	NoticeBarInfo *NoticeBarInfo `json:"noticeBarInfo"`
	ClickTypeInfo *ClickTypeInfo `json:"clickTypeInfo"`
	Extra         *Extra         `json:"extra"`
}

type Extra struct {
	Callback      string `json:"callback"`
	CallbackParam string `json:"callback.param"`
}

type ClickTypeInfo struct {
	ClickType int    `json:"clickType"`
	Url       string `json:"url"`
}

type NoticeBarInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PushMessageResponse struct {
	Code     string                 `json:"code"`     // 必选,返回码
	Message  string                 `json:"message"`  //可选，返回消息，网页端接口出现错误时使用此消息展示给用户，手机端可忽略此消息，甚至服务端不传输此消息
	Value    map[string]interface{} `json:"value"`    // 必选，返回结果
	Redirect string                 `json:"redirect"` //可选, returnCode=300 重定向时，使用此URL重新请求
	MsgId    string                 `json:"msgId"`    //可选，消息推送msgId
}

