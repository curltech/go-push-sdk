package ios_channel

const (
	PayloadTemplate = `{"aps" : {"content-available": 1,"alert" : {"title" : "%s","subtitle" : "%s","body" : "%s"},"badge" : 1,"sound" : "default"},"body" : "%s"}`
)

type PushMessageResponse struct {
	StatusCode int    `json:"status_code"`
	APNsId     string `json:"ap_ns_id"`
	Reason     string `json:"reason"`
}

