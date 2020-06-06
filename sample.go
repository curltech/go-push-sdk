package main

import (
	"context"
	"fmt"
	"gitee.com/cristiane/go-push-sdk/push"
	"gitee.com/cristiane/go-push-sdk/push/setting"
	"github.com/google/uuid"
)

func main() {
	register, err := push.NewRegisterClient(push.DefaultConfFile)
	if err != nil {
		fmt.Printf("NewRegisterClient err: %v", err)
		return
	}
	huaweiClient, err := register.GetHUAWEIClient()
	if err != nil {
		fmt.Printf("GetIosClient err: %v", err)
		return
	}
	//resp, err := huaweiClient.GetAccessToken(context.Background())
	//if err != nil {
	//	fmt.Printf("GetAccessToken err: %v", err)
	//	return
	//}
	//token := resp.(*huawei_channel.AccessTokenResp).AccessToken
	//log.Println("token==", token)
	var deviceTokens = []string{
		"AL8xAEeqhZFah9hMpnL9mqNz9Quf1bDfIEQ0yhrwrLgLFwwNehHz3VlKPzTunXJP6V1kuxVUOVu_E8hIFaUWxOPD1THjyLLfDxkBy3InSXoKqHQA689CkQQjFSwBHxJVWw",
	}
	msg := &setting.PushMessageRequest{
		AccessToken:  "CV7WMbw6a/hcGdN0YNjHj4RdOYLYHSuRPbMii8Abzg8/kpqSDsaY9GE0P8uDc6Qz0ffuUjbPfo/DMJmGGDaIgYMLiofo",
		DeviceTokens: deviceTokens,
		Message: &setting.Message{
			BusinessId: uuid.New().String(),
			Title:      "待办任务提醒",
			SubTitle:   "您有待办任务哦",
			Content:    "早上好！新的一天开始了，目前您还有任务需要马上处理~",
			Extra: map[string]string{
				"type":        "TodoRemind",
				"link_type":   "TaskList",
				"link_params": "[]",
			},
			CallBack:      "https://vengineer.xxx.com.cn/app-push/callback-oppo",
			CallbackParam: "",
		},
	}
	ctx := context.Background()
	respPush, err := huaweiClient.PushNotice(ctx, msg)
	if err != nil {
		fmt.Printf("ios push err: %v", err)
		return
	}
	fmt.Printf("ios push resp: %+v", respPush)
}
