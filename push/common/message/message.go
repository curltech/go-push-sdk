package message

import (
	"gitee.com/cristiane/go-push-sdk/push/errcode"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

// 公共消息格式检查
func CheckMessageParam(pushRequest *setting.PushMessageRequest, deviceTokenMin, deviceTokenMax int, isCheckAccessToken bool) error {
	if pushRequest.Message.Title == "" {
		return errcode.ErrMessageTitleEmpty
	}
	if pushRequest.Message.Content == "" {
		return errcode.ErrMessageContentEmpty
	}
	if len(pushRequest.DeviceTokens) < deviceTokenMin {
		return errcode.ErrDeviceTokenMin
	}
	if len(pushRequest.DeviceTokens) > deviceTokenMax {
		return errcode.ErrDeviceTokenMax
	}
	if isCheckAccessToken && pushRequest.AccessToken == "" {
		return errcode.ErrAccessTokenEmpty
	}
	// 其余参数检查
	
	return nil
}

