package client

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	payError "github.com/taydy/pay/error"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"strconv"
	"time"
)

/** 微信小程序 支付客户端。 */
type WeChatMpClient struct {
	*WeChatClient
}

var _ PayClient = &WeChatMpClient{}

func (c *WeChatMpClient) Pay(charge *_struct.Charge) (map[string]interface{}, error) {

	if charge.OpenID == "" {
		return nil, payError.NewBadRequestError(payError.PAY_FAILED, "invalid open id!")
	}

	m := c.ChargeToWeChatUnifiedOrder(charge)

	m.TradeType = constant.WX_TRADE_TYPE_JSAPI
	// 加签
	m.Sign = util.WeChatSign(structs.Map(m), c.SecretKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	prepay := map[string]interface{}{
		"appId":     c.AppID,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
		"nonceStr":  util.RandomStr(),
		"package":   fmt.Sprintf("prepay_id=%s", payResult.PrepayId),
		"signType":  constant.SIGN_TYPE_MD5,
	}
	prepay["paySign"] = util.WeChatSign(prepay, c.SecretKey)

	return prepay, nil
}

func (c *WeChatMpClient) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}
