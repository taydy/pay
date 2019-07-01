package client

import (
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	payError "github.com/taydy/pay/error"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"strconv"
	"time"
)

/** 微信 内置浏览器 支付客户端。 */
type WeChatJSAPIClient struct {
	*WeChatClient
}

func (c *WeChatJSAPIClient) Pay(charge *_struct.Charge) (map[string]interface{}, error) {
	m := c.ChargeToWeChatUnifiedOrder(charge)

	if m.Openid == "" {
		return nil, payError.NewBadRequestError(payError.PAY_FAILED, "invalid openid!")
	}

	m.TradeType = constant.WX_TRADE_TYPE_JSAPI
	// 加签
	m.Sign = util.WeChatSign(structs.Map(m), c.SecretKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	/**
	 * 统一下单接口返回正常的prepay_id，再按签名规范重新生成签名后，将数据传输给APP。
	 * 参与签名的字段名为appid，noncestr，timestamp，package, signType。注意：package的值格式为 prepay_id=***
	 */
	prepay := map[string]interface{}{
		"appId":     c.AppID,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
		"nonceStr":  util.RandomStr(),
		"package":   "prepay_id=" + payResult.PrepayId,
		"signType":  constant.SIGN_TYPE_MD5,
	}
	prepay["paySign"] = util.WeChatSign(prepay, c.SecretKey)

	return prepay, nil
}

func (c *WeChatJSAPIClient) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}

var _ PayClient = &WeChatJSAPIClient{}
