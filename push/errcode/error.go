package errcode

import (
	"fmt"
)

const (
	tempErr = "\u001B[34m[go-push-sdk]\u001B[0m: \x1b[31m%v\x1b[0m\n"
)

var (
	ErrCfgFileEmpty            = fmt.Errorf(tempErr, "conf file empty")
	ErrHuaweiClientIdEmpty     = fmt.Errorf(tempErr, "huawei clientId Empty")
	ErrHuaweiClientSecretEmpty = fmt.Errorf(tempErr, "huawei clientSecret Empty")
	ErrXiaomiAppSecretEmpty    = fmt.Errorf(tempErr, "xiaomi appSecret Empty")
	ErrVivoAppSecretEmpty      = fmt.Errorf(tempErr, "vivo appSecret Empty")
	ErrMeizuAppSecretEmpty     = fmt.Errorf(tempErr, "meizu appSecret Empty")
	ErrXiaomiAppPkgNameEmpty   = fmt.Errorf(tempErr, "xiaomi appPkgName Empty")
	ErrVivoAppPkgNameEmpty     = fmt.Errorf(tempErr, "vivo appPkgName Empty")
	ErrOppoAppPkgNameEmpty     = fmt.Errorf(tempErr, "oppo appPkgName Empty")
	ErrMeizuAppPkgNameEmpty    = fmt.Errorf(tempErr, "meizu appPkgName Empty")
	ErrHuaweiAppPkgNameEmpty   = fmt.Errorf(tempErr, "huawei appPkgName Empty")
	ErrAccessTokenEmpty        = fmt.Errorf(tempErr, "accessToken Empty")
	ErrMessageTitleEmpty       = fmt.Errorf(tempErr, "message title empty")
	ErrMessageContentEmpty     = fmt.Errorf(tempErr, "message content empty")
	ErrDeviceTokenMax          = fmt.Errorf(tempErr, "device token max limited")
	ErrDeviceTokenMin          = fmt.Errorf(tempErr, "device token min limited")
	ErrVivoAppIdEmpty          = fmt.Errorf(tempErr, "vivo appId Empty")
	ErrMeizuAppIdEmpty         = fmt.Errorf(tempErr, "meizu appId Empty")
	ErrVivoAppKeyEmpty         = fmt.Errorf(tempErr, "vivo appKey Empty")
	ErrOppoAppKeyEmpty         = fmt.Errorf(tempErr, "oppo appKey Empty")
	ErrOppoMasterSecretEmpty   = fmt.Errorf(tempErr, "oppo masterSecret Empty")
	ErrOppoSaveMessageToCloud  = fmt.Errorf(tempErr, "oppo save Message to Cloud error")
	ErrBusinessIdEmpty         = fmt.Errorf(tempErr, "businessId Empty")
	ErrXiaomiParseBody         = fmt.Errorf(tempErr, "xiaomi parse response body error")
	ErrVivoParseBody           = fmt.Errorf(tempErr, "vivo parse response body error")
	ErrOppoParseBody           = fmt.Errorf(tempErr, "oppo parse response body error")
	ErrMeizuParseBody          = fmt.Errorf(tempErr, "meizu parse response body error")
	ErrHuaweiParseBody         = fmt.Errorf(tempErr, "huawei parse response body error")
	ErrIosCertPathEmpty        = fmt.Errorf(tempErr, "ios certPath Empty")
	ErrIosPasswordEmpty        = fmt.Errorf(tempErr, "ios passWord Empty")
	ErrIosTeamIdEmpty          = fmt.Errorf(tempErr, "ios teamId Empty")
	ErrIosKeyIdEmpty           = fmt.Errorf(tempErr, "ios keyId Empty")
	ErrIosSecretFileEmpty      = fmt.Errorf(tempErr, "ios secretFile Empty")
	ErrIosBundleIdEmpty        = fmt.Errorf(tempErr, "ios bundleId Empty")
	ErrUnknownPlatform         = fmt.Errorf(tempErr, "unknown platform ")
	ErrParseConfigFile         = fmt.Errorf(tempErr, "parse config file err")
	ErrConfigEmpty             = fmt.Errorf(tempErr, "config file null")
)
