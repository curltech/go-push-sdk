package vivo_channel

type PushMessageRequest struct {
}

type PushMessageResponse struct {
	Result       int         `json:"result"`    // 0 表示成功，非0失败
	Desc         string      `json:"desc"`      // 文字描述接口调用情况
	RequestId    string      `json:"requestId"` // 请求ID
	InvalidUsers interface{} `json:"invalidUsers"`
	TaskId       string      `json:"taskId"` // 任务ID
}

type PushSingleNotify struct {
	RegId       string             `json:"regId"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	SkipType    int                `json:"skipType"`
	SkipContent string             `json:"skipContent"`
	RequestId   string             `json:"requestId"`
	NotifyType  int                `json:"notifyType"`
	Extra       *SingleNotifyExtra `json:"extra,omitempty"`
}

type SingleNotifyExtra struct {
	CallBack      string `json:"callback,omitempty"`
	CallBackParam string `json:"callback.param,omitempty"`
}

type SaveMessageToCloud struct {
	Title       string                   `json:"title"`
	Content     string                   `json:"content"`
	SkipType    int                      `json:"skipType"`
	SkipContent string                   `json:"skipContent"`
	RequestId   string                   `json:"requestId"`
	NotifyType  int                      `json:"notifyType"`
	Extra       *SaveMessageToCloudExtra `json:"extra,omitempty"`
}

type SaveMessageToCloudExtra struct {
	CallBack      string `json:"callback,omitempty"`
	CallBackParam string `json:"callback.param,omitempty"`
}

type SaveMessageToCloudResponse struct {
	Result int    `json:"result"`
	Desc   string `json:"desc"`
	TaskId string `json:"taskId"`
}

type PushMultiNotify struct {
	RegIds    []string `json:"regIds"`
	TaskId    string   `json:"taskId"`
	RequestId string   `json:"requestId"`
}

type AuthTokenReq struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	Timestamp string `json:"timestamp"`
}

type AuthTokenResp struct {
	Result    int    `json:"result"`    // 0 成功，非0失败
	Desc      string `json:"desc"`      // 文字描述接口调用情况
	AuthToken string `json:"authToken"` // 默认有效一天
}

