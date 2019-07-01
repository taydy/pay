package client

import (
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
)

/** 微信 Web 扫码支付客户端。 */
type WeChatWebClient struct {
	*WeChatClient
}

var _ PayClient = &WeChatWebClient{}

/**
 * 微信 Web 扫码支付，调用微信统一下单接口生成订单，根据返回的 code_url 生成二维码.
 */
func (c *WeChatWebClient) Pay(charge *_struct.Charge) (map[string]interface{}, error) {

	m := c.ChargeToWeChatUnifiedOrder(charge)

	m.TradeType = constant.WX_TRADE_TYPE_NATIVE
	// 加签
	m.Sign = util.WeChatSign(structs.Map(m), c.SecretKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	pngBase64 := util.PngBase64(payResult.CodeUrl)

	return map[string]interface{}{"png": pngBase64}, nil
}

func (c *WeChatWebClient) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}
