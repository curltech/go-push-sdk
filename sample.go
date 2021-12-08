package main

import (
	"context"
	"fmt"
	"github.com/curltech/go-push-sdk/push"
	"github.com/curltech/go-push-sdk/push/setting"
	"github.com/google/uuid"
)

func main() {
	register, err := push.NewRegisterClient(push.DefaultConfFile)
	if err != nil {
		fmt.Printf("NewRegisterClient err: %v", err)
		return
	}
	iosTokenClient, err := register.GetIosTokenClient()
	if err != nil {
		fmt.Printf("GetIosClient err: %v", err)
		return
	}
	var deviceTokens = []string{
		"58fb0a845812c3516819cd24ba50e237c388d6fbfde21fb997e8eeb725066d35",
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
	respPush, err := iosTokenClient.PushNotice(ctx, msg)
	if err != nil {
		fmt.Printf("ios push err: %v", err)
		return
	}
	fmt.Printf("ios push resp: %+v", respPush)
}
