package example

import (
	"github.com/taydy/pay/client"
	"github.com/taydy/pay/struct"
)

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
