### 微信、支付宝支付 GoLand 版

| 支付方式    | 是否支持 |
| ----------- | -------- |
| 阿里 web    | ✅       |
| 阿里 wap    | ✅        |
| 阿里 app    | ✅      |
| 微信 web    | ✅         |
| 微信 h5     | ✅         |
| 微信 app    | ✅         |
| 微信 小程序 | ✅         |

### 示例
#### 微信 App 支付
```
// 微信支付均可参考以下格式来调用
var WXPayClient = &client.WeChatAppClient{
	WeChatClient: &client.WeChatClient{
		AppID:     "",
		MchID:     "",
		SecretKey: "",
	},
}

// 注意 H5 支付需要传客户端的 ip 地址，微信会做校验
func WXAppPay(orderId string, amount int, description string, notifyUrl string, returnUrl string, ip string) (map[string]interface{}, error) {
	now := time.Now()

	charge := &_struct.Charge{}
	charge.OutTradeNo = orderId
	charge.TotalFee = amount
	charge.Body = description
	charge.NotifyUrl = notifyUrl
	charge.ReturnUrl = returnUrl
	charge.SpbillCreateIp = ip

	charge.TimeStart = util.FormatTime(now)
	charge.TimeExpire = util.FormatTime(now.Add(constant.TIMEOUT_PAY))

	result, err := WXPayClient.Pay(charge)
	if err != nil {
		return nil, err
	}
	return result, nil
}
```

#### 支付宝 App 支付
```
var ALIPayClient = &client.AliAppPayClient{
	AliPayClient: &client.AliPayClient{
		AppID:      "",
		PrivateKey: nil,
		PublicKey:  nil,
	},
}

func AliAppPay(orderId string, amount int, description string, notifyUrl string, returnUrl string) (map[string]interface{}, error) {
	charge := &_struct.Charge{}
	charge.Body = description
	charge.OutTradeNo = orderId
	charge.TotalFee = amount
	charge.NotifyUrl = notifyUrl
	charge.ReturnUrl = returnUrl

	result, err := ALIPayClient.Pay(charge)
	if err != nil {
		return nil, err
	}
	return result, nil
}
```