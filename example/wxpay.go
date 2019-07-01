package example

import (
	"github.com/taydy/pay/client"
	"github.com/taydy/pay/constant"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"time"
)

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
