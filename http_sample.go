package main

import (
	"github.com/curltech/go-push-sdk/push/huawei_channel"
	"github.com/curltech/go-push-sdk/push/ios_channel"
	"github.com/curltech/go-push-sdk/push/meizu_channel"
	"github.com/curltech/go-push-sdk/push/oppo_channel"
	"github.com/curltech/go-push-sdk/push/vivo_channel"
	"github.com/curltech/go-push-sdk/push/xiaomi_channel"
	"log"
	"net/http"
	"strings"

	"github.com/curltech/go-push-sdk/push"
	"github.com/curltech/go-push-sdk/push/common/json"
	"github.com/curltech/go-push-sdk/push/setting"
	"github.com/gin-gonic/gin"
)

var (
	register *push.RegisterClient
	err      error
)

func init() {
	register, err = push.NewRegisterClient(push.DefaultConfFile)
	if err != nil {
		log.Fatalf("NewRegisterClient err: %v", err)
	}
}

func main() {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.NoRoute(Handle404)
	pushRouter := router.Group("/go-push-sdk")
	huaweiRouter := pushRouter.Group("/ios")
	{
		huaweiRouter.GET("/access_token", huaweiAccessToken)

		huaweiRouter.POST("/push", huaweiPush)
	}
	iphoneRouter := pushRouter.Group("/ios")
	{
		iphoneRouter.POST("/cert-push", iphoneCert)

		iphoneRouter.POST("/token-push", iphoneToken)
	}
	meizuRouter := pushRouter.Group("/meizu")
	{
		meizuRouter.POST("/push", meizuPush)
	}
	xiaomiRouter := pushRouter.Group("/xiaomi")
	{
		xiaomiRouter.POST("/push", xiaomiPush)
	}
	oppoRouter := pushRouter.Group("/oppo")
	{
		oppoRouter.GET("/auth_token", oppoAuthToken)

		oppoRouter.POST("/push", oppoPush)
	}
	vivoRouter := pushRouter.Group("/vivo")
	{
		vivoRouter.GET("/auth_token", vivoAuthToken)
		vivoRouter.POST("/push", vivoPush)
	}

	errGinHttp := router.Run(":80")
	if errGinHttp != nil {
		log.Fatalln(errGinHttp)
	}
}

func huaweiAccessToken(c *gin.Context) {

	huaweiClient, err := register.GetHUAWEIClient()
	if err != nil {
		log.Println("huawei get access_token err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"access_token": "",
			"expires":      0,
			"reason":       err.Error(),
		})
		return
	}
	accessTokenResp, err := huaweiClient.GetAccessToken(c)
	if err != nil {
		log.Println("huawei get access_token err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"access_token": "",
			"expires":      0,
			"reason":       err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"access_token":      accessTokenResp.(*huawei_channel.AccessTokenResp).AccessToken,
		"expires":           accessTokenResp.(*huawei_channel.AccessTokenResp).ExpiresIn,
		"error":             accessTokenResp.(*huawei_channel.AccessTokenResp).Error,
		"scope":             accessTokenResp.(*huawei_channel.AccessTokenResp).Scope,
		"error_description": accessTokenResp.(*huawei_channel.AccessTokenResp).ErrorDescription,
	})
}

func huaweiPush(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")
	deviceAccessToken := c.PostForm("accessToken")
	huaweiClient, err := register.GetHUAWEIClient()
	if err != nil {
		log.Println("huawei push get access_token err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"access_token": "",
			"expires":      0,
			"reason":       err.Error(),
		})
		return
	}
	if deviceAccessToken == "" {
		accessTokenResp, err := huaweiClient.GetAccessToken(c)
		if err != nil {
			log.Println("huawei push get access_token err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"access_token": "",
				"expires":      0,
				"reason":       err.Error(),
			})
			return
		}
		deviceAccessToken = accessTokenResp.(*huawei_channel.AccessTokenResp).AccessToken
	}

	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  deviceAccessToken,
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			log.Println("messageExtraMap Unmarshal err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	respPush, err := huaweiClient.PushNotice(c, msg)
	if err != nil {
		log.Println("huawei push err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"reason":    "",
		"code":      respPush.(*huawei_channel.PushMessageResponse).Code,
		"msg":       respPush.(*huawei_channel.PushMessageResponse).Msg,
		"requestId": respPush.(*huawei_channel.PushMessageResponse).RequestId,
	})
}

func iphoneCert(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")
	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  "",
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	iosClient, err := register.GetIosCertClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	respPush, err := iosClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": respPush.(*ios_channel.PushMessageResponse).StatusCode,
		"ap_ns_id":    respPush.(*ios_channel.PushMessageResponse).APNsId,
		"reason":      respPush.(*ios_channel.PushMessageResponse).Reason,
	})
}

func iphoneToken(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")

	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  "",
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	iosTokenClient, err := register.GetIosTokenClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	respPush, err := iosTokenClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"status_code": respPush.(*ios_channel.PushMessageResponse).StatusCode,
		"ap_ns_id":    respPush.(*ios_channel.PushMessageResponse).APNsId,
		"reason":      respPush.(*ios_channel.PushMessageResponse).Reason,
	})
}

func meizuPush(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")

	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  "",
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}

	meizuClient, err := register.GetMEIZUClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	respPush, err := meizuClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"code":     respPush.(*meizu_channel.PushMessageResponse).Code,
		"message":  respPush.(*meizu_channel.PushMessageResponse).Message,
		"value":    respPush.(*meizu_channel.PushMessageResponse).Value,
		"redirect": respPush.(*meizu_channel.PushMessageResponse).Redirect,
		"msgId":    respPush.(*meizu_channel.PushMessageResponse).MsgId,
	})
}

func xiaomiPush(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")

	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  "",
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	xiaomiClient, err := register.GetXIAOMIClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	respPush, err := xiaomiClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"result":      respPush.(*xiaomi_channel.PushMessageResponse).Result,
		"description": respPush.(*xiaomi_channel.PushMessageResponse).Description,
		"data":        respPush.(*xiaomi_channel.PushMessageResponse).Data,
		"code":        respPush.(*xiaomi_channel.PushMessageResponse).Code,
		"info":        respPush.(*xiaomi_channel.PushMessageResponse).Info,
		"reason":      respPush.(*xiaomi_channel.PushMessageResponse).Reason,
	})
}

func oppoAuthToken(c *gin.Context) {
	oppoClient, err := register.GetOPPOClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	authTokenResp, err := oppoClient.GetAccessToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    authTokenResp.(*oppo_channel.AuthTokenResp).Code,
		"data":    json.MarshalToStringNoError(authTokenResp.(*oppo_channel.AuthTokenResp).Data),
		"message": authTokenResp.(*oppo_channel.AuthTokenResp).Message,
	})
}

func oppoPush(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")
	deviceAccessToken := c.PostForm("accessToken")
	oppoClient, err := register.GetOPPOClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	if deviceAccessToken == "" {
		authTokenResp, err := oppoClient.GetAccessToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		deviceAccessToken = authTokenResp.(*oppo_channel.AuthTokenResp).Data.AuthToken
	}
	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  deviceAccessToken,
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	respPush, err := oppoClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": respPush.(*oppo_channel.PushMessageResponse).Message,
		"code":    respPush.(*oppo_channel.PushMessageResponse).Code,
		"data":    json.MarshalToStringNoError(respPush.(*oppo_channel.PushMessageResponse).Data),
	})
}

func vivoAuthToken(c *gin.Context) {
	vivoClient, err := register.GetVIVOClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	authTokenResp, err := vivoClient.GetAccessToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"result":     authTokenResp.(*vivo_channel.AuthTokenResp).Result,
		"auth_token": authTokenResp.(*vivo_channel.AuthTokenResp).AuthToken,
		"desc":       authTokenResp.(*vivo_channel.AuthTokenResp).Desc,
	})
}

func vivoPush(c *gin.Context) {

	deviceTokens := c.PostForm("deviceTokens")
	messageBusinessId := c.PostForm("messageBusinessId")
	messageTitle := c.PostForm("messageTitle")
	messageSubTitle := c.PostForm("messageSubTitle")
	messageContent := c.PostForm("messageContent")
	messageExtra := c.PostForm("messageExtraJson")
	messageCallBack := c.PostForm("messageCallBack")
	messageCallBackParam := c.PostForm("messageCallBackParam")
	deviceAccessToken := c.PostForm("accessToken")

	vivoClient, err := register.GetVIVOClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	if deviceAccessToken == "" {
		authTokenResp, err := vivoClient.GetAccessToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		deviceAccessToken = authTokenResp.(*vivo_channel.AuthTokenResp).AuthToken
	}

	msg := &setting.PushMessageRequest{
		DeviceTokens: strings.Split(deviceTokens, ","),
		AccessToken:  deviceAccessToken,
		Message: &setting.Message{
			BusinessId:    messageBusinessId,
			Title:         messageTitle,
			SubTitle:      messageSubTitle,
			Content:       messageContent,
			CallBack:      messageCallBack,
			CallbackParam: messageCallBackParam,
		},
	}
	if messageExtra != "" {
		messageExtraMap := map[string]string{}
		err := json.Unmarshal(messageExtra, &messageExtraMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  err.Error(),
			})
			return
		}
		msg.Message.Extra = messageExtraMap
	}
	respPush, err := vivoClient.PushNotice(c, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"desc":          respPush.(*vivo_channel.PushMessageResponse).Desc,
		"result":        respPush.(*vivo_channel.PushMessageResponse).Result,
		"request_id":    respPush.(*vivo_channel.PushMessageResponse).RequestId,
		"task_id":       respPush.(*vivo_channel.PushMessageResponse).TaskId,
		"invalid_users": respPush.(*vivo_channel.PushMessageResponse).InvalidUsers,
	})
}

func Handle404(c *gin.Context) {
	HandleMessage(c, "糟糕，页面找不到了")
}

func HandleMessage(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": message,
	})
}
