package errcode

import "errors"

// 错误定义
var (
	ErrCfgFileEmpty        = errors.New("conf file empty")             // 推送超时
	ErrClientIdEmpty       = errors.New("ClientId Empty")              // 客户端ID为空
	ErrClientSecretEmpty   = errors.New("ClientSecret Empty")          // 客户端秘钥为空
	ErrAppSecretEmpty      = errors.New("AppSecret Empty")             // 秘钥为空
	ErrAppPkgNameEmpty     = errors.New("AppPkgName Empty")            // 应用包名为空
	ErrAccessTokenEmpty    = errors.New("accessToken Empty")           // access_token为空
	ErrMessageTitleEmpty   = errors.New("message title empty")         // 消息标题为空
	ErrMessageContentEmpty = errors.New("message content empty")       // 消息内容为空
	ErrDeviceTokenMax      = errors.New("device token max limited")    // 推送设备超过最大限制
	ErrDeviceTokenMin      = errors.New("device token min limited")    // 推送设备超过最小限制
	ErrAppIdEmpty          = errors.New("AppId Empty")                 // 应用ID为空
	ErrAppKeyEmpty         = errors.New("AppKey Empty")                // 应用KEY为空
	ErrMasterSecretEmpty   = errors.New("MasterSecret Empty")          // 应用MasterSecret为空
	ErrSaveMessageToCloud  = errors.New("Save Message to Cloud error") // 推送消息错误
	ErrBusinessIdEmpty     = errors.New("BusinessId Empty")            // 业务ID为空
	ErrParseBody           = errors.New("parse response body error")   // 解析响应体错误
	ErrCertPathEmpty       = errors.New("CertPath Empty")              // CertPath 为空
	ErrPasswordEmpty       = errors.New("PassWord Empty")              // Password 为空
	ErrTeamIdEmpty         = errors.New("TeamId Empty")
	ErrKeyIdEmpty          = errors.New("KeyId Empty")
	ErrSecretFileEmpty     = errors.New("SecretFile Empty")
	ErrBundleIdEmpty       = errors.New("BundleId Empty")
	ErrParseCfgFile        = errors.New("Parse setting.json file error")
)
