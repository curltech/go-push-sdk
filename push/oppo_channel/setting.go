package oppo_channel

type PushMessageRequest struct {
	Title             string `json:"title"`
	SubTitle          string `json:"sub_title"`
	Content           string `json:"content"`
	ClickActionType   int    `json:"click_action_type"`
	ClickActionUrl    string `json:"click_action_url"`
	CallBackUrl       string `json:"call_back_url"`
	CallBackParameter string `json:"call_back_parameter"`
}

type SinglePush struct {
	Message   *SingleMessage `json:"message"`
	AuthToken string         `json:"auth_token"`
}

type SingleMessage struct {
	TargetType   int               `json:"target_type"`
	TargetValue  string            `json:"target_value"`
	Notification map[string]string `json:"notification"`
}

type BroadcastPush struct {
	MessageId   string `json:"message_id"`
	TargetType  int    `json:"target_type"`
	TargetValue string `json:"target_value"`
	AuthToken   string `json:"auth_token"`
}

type PushMessageResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *PushMessageData `json:"data"`
}

type PushMessageData struct {
	BroadcastMessageId string `json:"message_id"`
	SingleMessageId    string `json:"messageId"`
	Status             string `json:"status"`
	TaskId             string `json:"task_id"`
}

type SaveMessageToCloudResponse struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    *SaveMessageToCloudResult `json:"data"`
}

type SaveMessageToCloudResult struct {
	MessageId string `json:"message_id"`
}

type AuthTokenReq struct {
	AppKey       string `json:"app_key"`
	Timestamp    string `json:"timestamp"`
	MasterSecret string `json:"master_secret"`
}

type AuthTokenResp struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    *AuthTokenData `json:"data"`
}

type AuthTokenData struct {
	AuthToken  string `json:"auth_token"`
	CreateTime int64  `json:"create_time"`
}

