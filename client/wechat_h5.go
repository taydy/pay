package client

import (
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	payError "github.com/taydy/pay/error"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
)

/** 微信 H5 支付客户端。 */
type WeChatH5Client struct {
	*WeChatClient
}

var _ PayClient = &WeChatH5Client{}

/**
 * 微信 H5 支付，返回结果实际有用字段为 mweb_url，将此 HTTP code 设置为 302, 并返回此链接。
 */
func (c *WeChatH5Client) Pay(charge *_struct.Charge) (map[string]interface{}, error) {

	if !util.IsIP(charge.SpbillCreateIp) {
		return nil, payError.NewBadRequestError(payError.PAY_FAILED, "invalid ip!")
	}

	m := c.ChargeToWeChatUnifiedOrder(charge)

	m.TradeType = constant.WX_TRADE_TYPE_H5
	// 加签
	m.Sign = util.WeChatSign(structs.Map(m), c.SecretKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	return structs.Map(payResult), nil
}

func (c *WeChatH5Client) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}
