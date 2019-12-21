package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitee.com/cristiane/go-push-sdk/push"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

func main() {
	register, err := push.NewRegisterClient()
	if err != nil {
		fmt.Printf("NewRegisterClient err: %v", err)
		return
	}
	iosClient, err := register.GetIosTokenClient()
	if err != nil {
		fmt.Printf("GetIosClient err: %v", err)
		return
	}
	var deviceTokens = []string{
		"c79784384b65f2464249ce2bc22c6896975b9becb124976d4bce5b3760cf96ac",
	}
	msg := &setting.PushMessageRequest{
		AccessToken:  "",
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
			CallBack:      "https://vengineer.myysq.com.cn/app-push/callback-oppo",
			CallbackParam: "",
		},
	}
	ctx := context.Background()
	respPush, err := iosClient.PushNotice(ctx, msg)
	if err != nil {
		fmt.Printf("ios push err: %v", err)
		return
	}
	fmt.Printf("ios push resp: %+v", respPush)
}

