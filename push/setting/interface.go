package setting

import "context"

type PushClientInterface interface {
	GetAccessToken(ctx context.Context) (interface{}, error)
	PushNotice(ctx context.Context, pushRequest *PushMessageRequest) (interface{}, error)
}
