package client

import (
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"strconv"
	"time"
)

/** 微信 app 支付客户端。 */
type WeChatAppClient struct {
	*WeChatClient
}

var _ PayClient = &WeChatAppClient{}

func (c *WeChatAppClient) Pay(charge *_struct.Charge) (map[string]interface{}, error) {

	m := c.ChargeToWeChatUnifiedOrder(charge)

	m.TradeType = constant.WX_TRADE_TYPE_APP
	// 加签
	m.Sign = util.WeChatSign(structs.Map(m), c.SecretKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	/**
	 * 统一下单接口返回正常的prepay_id，再按签名规范重新生成签名后，将数据传输给APP。
	 * 参与签名的字段名为appid，partnerid，prepayid，noncestr，timestamp，package。注意：package的值格式为 Sign=WXPay
	 */
	prepay := map[string]interface{}{
		"appid":     c.AppID,
		"partnerid": c.MchID,
		"prepayid":  payResult.PrepayId,
		"package":   "Sign=wxpay",
		"noncestr":  util.RandomStr(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}
	prepay["sign"] = util.WeChatSign(prepay, c.SecretKey)

	return prepay, nil
}

func (c *WeChatAppClient) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}
