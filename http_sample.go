package main

import (
	"log"
	"net/http"
	"strings"
	
	"github.com/gin-gonic/gin"
	"gitee.com/cristiane/go-push-sdk/push"
	"gitee.com/cristiane/go-push-sdk/push/common/json"
	"gitee.com/cristiane/go-push-sdk/push/setting"
)

var (
	register *push.RegisterClient
	err      error
)

func init() {
	register, err = push.NewRegisterClient()
	if err != nil {
		log.Fatalf("NewRegisterClient err: %v", err)
	}
}

func main() {
	
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.NoRoute(Handle404)
	pushRouter := router.Group("/go-push")
	huaweiRouter := pushRouter.Group("/huawei")
	{
		huaweiRouter.GET("/push/access_token", huaweiAccessToken)
		
		huaweiRouter.POST("/push", huaweiPush)
	}
	iphoneRouter := pushRouter.Group("/iphone")
	{
		iphoneRouter.POST("/push/cert", iphoneCert)
		
		iphoneRouter.POST("/push/token", iphoneToken)
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
		oppoRouter.GET("/push/auth_token", oppoAuthToken)
		
		oppoRouter.POST("/push", oppoPush)
	}
	vivoRouter := pushRouter.Group("/vivo")
	{
		vivoRouter.GET("/push/auth_token", vivoAuthToken)
		vivoRouter.POST("/push", vivoPush)
	}
	
	errGinHttp := router.Run(":9090")
	if errGinHttp != nil {
		log.Fatalln(errGinHttp)
	}
}

func huaweiAccessToken(c *gin.Context) {
	
	huaweiClient, err := register.GetHuaweiClient()
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
		"access_token":      accessTokenResp.AccessToken,
		"expires":           accessTokenResp.ExpiresIn,
		"error":             accessTokenResp.Error,
		"scope":             accessTokenResp.Scope,
		"error_description": accessTokenResp.ErrorDescription,
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
	huaweiClient, err := register.GetHuaweiClient()
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
		deviceAccessToken = accessTokenResp.AccessToken
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
		"code":      respPush.Code,
		"msg":       respPush.Msg,
		"requestId": respPush.RequestId,
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
		"status_code": respPush.StatusCode,
		"ap_ns_id":    respPush.APNsId,
		"reason":      respPush.Reason,
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
		"status_code": respPush.StatusCode,
		"ap_ns_id":    respPush.APNsId,
		"reason":      respPush.Reason,
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
	
	meizuClient, err := register.GetMeizuClient()
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
		"code":     respPush.Code,
		"message":  respPush.Message,
		"value":    respPush.Value,
		"redirect": respPush.Redirect,
		"msgId":    respPush.MsgId,
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
	xiaomiClient, err := register.GetXiaomiClient()
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
		"result":      respPush.Result,
		"description": respPush.Description,
		"data":        respPush.Data,
		"code":        respPush.Code,
		"info":        respPush.Info,
		"reason":      respPush.Reason,
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
		"code":    authTokenResp.Code,
		"data":    json.MarshalToStringNoError(authTokenResp.Data),
		"message": authTokenResp.Message,
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
		deviceAccessToken = authTokenResp.Data.AuthToken
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
		"message": respPush.Message,
		"code":    respPush.Code,
		"data":    json.MarshalToStringNoError(respPush.Data),
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
		"result":     authTokenResp.Result,
		"auth_token": authTokenResp.AuthToken,
		"desc":       authTokenResp.Desc,
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
		deviceAccessToken = authTokenResp.AuthToken
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
		"desc":          respPush.Desc,
		"result":        respPush.Result,
		"request_id":    respPush.RequestId,
		"task_id":       respPush.TaskId,
		"invalid_users": respPush.InvalidUsers,
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
