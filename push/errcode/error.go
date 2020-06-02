package errcode

import "errors"

// 错误定义
var (
	ErrCfgFileEmpty        = errors.New("[go-push-sdk] conf file empty")             // 推送超时
	ErrClientIdEmpty       = errors.New("[go-push-sdk] ClientId Empty")              // 客户端ID为空
	ErrClientSecretEmpty   = errors.New("[go-push-sdk] ClientSecret Empty")          // 客户端秘钥为空
	ErrAppSecretEmpty      = errors.New("[go-push-sdk] AppSecret Empty")             // 秘钥为空
	ErrAppPkgNameEmpty     = errors.New("[go-push-sdk] AppPkgName Empty")            // 应用包名为空
	ErrAccessTokenEmpty    = errors.New("[go-push-sdk] accessToken Empty")           // access_token为空
	ErrMessageTitleEmpty   = errors.New("[go-push-sdk] message title empty")         // 消息标题为空
	ErrMessageContentEmpty = errors.New("[go-push-sdk] message content empty")       // 消息内容为空
	ErrDeviceTokenMax      = errors.New("[go-push-sdk] device token max limited")    // 推送设备超过最大限制
	ErrDeviceTokenMin      = errors.New("[go-push-sdk] device token min limited")    // 推送设备超过最小限制
	ErrAppIdEmpty          = errors.New("[go-push-sdk] AppId Empty")                 // 应用ID为空
	ErrAppKeyEmpty         = errors.New("[go-push-sdk] AppKey Empty")                // 应用KEY为空
	ErrMasterSecretEmpty   = errors.New("[go-push-sdk] MasterSecret Empty")          // 应用MasterSecret为空
	ErrSaveMessageToCloud  = errors.New("[go-push-sdk] Save Message to Cloud error") // 推送消息错误
	ErrBusinessIdEmpty     = errors.New("[go-push-sdk] BusinessId Empty")            // 业务ID为空
	ErrParseBody           = errors.New("[go-push-sdk] parse response body error")   // 解析响应体错误
	ErrCertPathEmpty       = errors.New("[go-push-sdk] CertPath Empty")              // CertPath 为空
	ErrPasswordEmpty       = errors.New("[go-push-sdk] PassWord Empty")              // Password 为空
	ErrTeamIdEmpty         = errors.New("[go-push-sdk] TeamId Empty")
	ErrKeyIdEmpty          = errors.New("[go-push-sdk] KeyId Empty")
	ErrSecretFileEmpty     = errors.New("[go-push-sdk] SecretFile Empty")
	ErrBundleIdEmpty       = errors.New("[go-push-sdk] BundleId Empty")
	ErrParseCfgFile        = errors.New("[go-push-sdk] Parse setting.json file error")
	ErrUnknownPlatform     = errors.New("[go-push-sdk] unknown platform ")
)
