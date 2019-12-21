# go-push-sdk

## 版本说明
当前版本：v3   

| 版本号 | 修改人 | 修改日期 | 描述 |
| :-----| ----: | :----: | :----: |
| v2 | 杨强 | 2019/7/15 | 支持华为，小米，魅族，OPPO，vivo，iOS（cert模式和token模式）厂商消息推送 |
| v3 | 杨强 | 2019/8/5 | 获取授权token，推送消息增加上下文参数（为接入打点提供支持） |
| v3.1 | 杨强 | 2019/11/28 | 结构优化，更新引用三方包apns2 |

  
  
## 环境要求
+ go > 1.7

## 使用方式一
1. 添加推送配置文件   
将根目录下文件setting.json复制到/usr/local/etc/go-push-sdk/目录下，并配置相应推送平台参数（iOS平台还需要配置证书文件路径，建议配置证书的绝对路径）   
2. 参考根目录sample.go示例代码进行消息推送   
   2.1  初始化一个推送平台注册器   
```
	register, err := push.NewRegisterClient()
	if err != nil {
		fmt.Printf("NewRegisterClient err: %v", err)
		return
	}
```
   2.1  从步骤1获取到的注册器获取相应推送平台客户端   
```
	iosClient, err := register.GetIosCertClient()
	if err != nil {
		fmt.Printf("GetIosClient err: %v", err)
		return
	}
```   
   2.3  构造消息推送参数（注意平台差异，华为平台，OPPO平台，vivo平台需要在Message结构里附加AccessToken 值）   
```
	var deviceTokens = []string{
		"xxxxxxxxxxxxxxxxxx",
	}
	msg := &setting.PushMessageRequest{
		AccessToken:"",
		DeviceTokens: deviceTokens,
		Message: &setting.Message{
			BusinessId: "a4cc3c24-b3f1-4368-af2a-4c0898ba15e8",
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
```   
2.4  开始消息推送   
```
	ctx := context.Background()
	respPush, err := iosClient.PushNotice(ctx, msg)
	if err != nil {
		fmt.Printf("ios push err: %v", err)
		return
	}
	fmt.Printf("ios push resp: %+v", respPush)
```
## 使用方式二   
使用详情参照v1版本-使用方式二 


## HTTP测试接口
支持平台：华为，小米，OPPO，vivo，魅族，iPhone   
本机需要配置host   
10.5.24.47  go-push.xxx.com   
+ 华为手机
   + 需要先获取access_token   
URL：http://go-push.xxx.com:9090/go-push/huawei/push/access_token   
请求方式：get   
body: 无   
响应结果：   
```
{
  "access_token": "CF0dt1Gt5FDhfF/pUudKCUMphIGxVFxuYaJk3nQnvN2MHJ9MXx+hD0ELpQmNJ6ltPAluBsoVldn34hb2bR8T9w==",
  "error": 0,
  "error_description": "",
  "expires": 3600,
  "scope": "",
  "success": true
}
```
   + 消息推送   
URL：http://go-push.xxx.com:9090/go-push/huawei/push   
请求方式：post表单   
消息参数如下：   

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| accessToken | string | 接入token | 选填（不传则自动获取） |
| deviceTokens | string | 设备token（多个用英文逗号隔开，单次最大100个） | 必填 |
| messageBusinessId | string | 消息业务ID | 非必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   

响应结果：   
```
{
  "code": "80000000",
  "msg": "Success",
  "reason": "",
  "requestId": "156222523472361878003501",
  "success": true
}
```
+ 魅族手机   
+ 不需要授权token   
+ 消息推送   
URL：http://go-push.xxx.com:9090/go-push/meizu/push   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最大1000个） | 必填 |
| messageBusinessId | string | 消息业务ID | 非必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填（填了會校驗格式是否正確） |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   

响应结果：   
```
{
  "code": "200",
  "message": "",
  "msgId": "NS20190704153256153_0_11579908_1_3",
  "redirect": "",
  "success": true,
  "value": {
    "110002": [
      "ULY6c596e6a7d5b714a475a60527c6b5f7f655a6d6370"
    ]
  }
}
```

+ 小米手机   
+ 不需要授权token   
+ 消息推送   
URL：http://go-push.xxx.com:9090/go-push/xiaomi/push   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最多100个） | 必填 |
| messageBusinessId | string | 消息业务ID | 非必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   

响应结果：   
```
{
  "code": 0,
  "data": {
    "id": "sdm00940562225761669wO"
  },
  "description": "成功",
  "info": "Received push messages for 1 REGID",
  "reason": "",
  "result": "ok",
  "success": true
}
```

+ oppo手机   
+ 获取auth_token   
URL：http://go-push.xxx.com:9090/go-push/oppo/push/auth_token   
请求方式：get   
body: 无   
返回：   
```
{
  "code": 0,
  "data": "{\"auth_token\":\"4e85a703-1636-444d-a3a2-c26abb2b1ffb\",\"create_time\":1562207540879}",
  "message": "Success",
  "success": true
}
```
+ 消息推送   
URL：http://go-push.xxx.com:9090/go-push/oppo/push   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| accessToken | string | 接入token | 选填（不传则自动获取） |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最大1000个） | 必填 |
| messageBusinessId | string | 消息业务ID | 非必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   

响应结果：   
```
{
  "code": 0,
  "data": "{\"message_id\":\"\",\"messageId\":\"5d1dad59cac4826aedd7ee4a\",\"status\":\"call_success\",\"task_id\":\"\"}",
  "message": "Success",
  "success": true
}
```

+ vivo手机   
+ 获取auth_token   
URL：http://go-push.xxx.com:9090/go-push/vivo/push/auth_token   
请求方式：get   
body: 无   
返回：   
```
{
  "auth_token": "aa05871a-dd0b-4ab1-a303-b4e8177fb2e1",
  "desc": "请求成功",
  "result": 0,
  "success": true
}
```
+ 消息推送
URL：http://go-push.xxx.com:9090/go-push/vivo/push   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| accessToken | string | 接入token | 选填（不传则自动获取） |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最大1000个） | 必填 |
| messageBusinessId | string | 消息业务ID | 必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   
响应结果：   
```
{
  "desc": "请求成功",
  "invalid_users": null,
  "request_id": "ab-7h-98-io8",
  "result": 0,
  "success": true,
  "task_id": ""
}
```   
+ iphone（基于证书）
URL：http://go-push.xxx.com:9090/go-push/iphone/push/cert   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| isSandBox | bool | 是否开启沙盒推送（默认false） | 选填 |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最多100个） | 必填 |
| messageBusinessId | string | 消息业务ID（注意格式） | 必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   

响应结果：   
```
{
  "ap_ns_id": "62E3A03F-8CE9-56F1-0D1A-2FED55CAF9C6",
  "reason": "",
  "status_code": 200,
  "success": true
}
```   
+ iPhone（基于token）   
URL：http://go-push.xxx.com:9090/go-push/iphone/push/token   
请求方式：post表单   
消息参数如下：

| 参数 | 类型 | 含义 | 是否必填 |
| :-----| ----: | :----: | :----: |
| isSandBox | bool | 是否开启沙盒推送（默认false） | 选填 |
| deviceTokens | string | 设备token（多个用英文逗号隔开，最多100个） | 必填 |
| messageBusinessId | string | 消息业务ID（注意格式） | 必填 |
| messageTitle | string | 消息标题 | 必填 |
| messageSubTitle | string | 消息副标题 | 非必填 |
| messageContent | string | 消息正文 | 必填 |
| messageExtraJson | string | 消息附带参数 键值为string的map序列化结果，如 "{"param1":"value1","param2":"value2"}" | 非必填 |
| messageCallBack | string | 消息回调地址 | 非必填 |
| messageCallBackParam | string | 消息回调附带参数 | 非必填 |   
响应结果：   
```

  "ap_ns_id": "62E3A03F-8CE9-56F1-0D1A-2FED55CAF9C6",
  "reason": "",
  "status_code": 200,
  "success": true
}
```   