package client

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
)

/** 阿里 Web 支付客户端。 */
type AliWebPayClient struct {
	*AliPayClient
}

var _ PayClient = &AliWebPayClient{}

/** web支付，生成连接并使用私钥加签。 */
func (c *AliWebPayClient) Pay(charge *_struct.Charge) (map[string]interface{}, error) {

	var m = &_struct.AliPayUnifiedOrder{}
	m.AppId = c.AppID
	m.Method = constant.ALI_PAY_API_PAGE_PAY
	m.Charset = constant.CHARSET_UTF8
	m.Timestamp = util.DongBaTime().Format(constant.TIME_FORMAT)
	m.Version = "1.0"
	m.NotifyUrl = charge.NotifyUrl
	m.ReturnUrl = charge.ReturnUrl

	var bizContent = &_struct.BizContent{}
	bizContent.OutTradeNo = charge.OutTradeNo
	bizContent.ProductCode = constant.ALI_PAY_PRODUCT_CODE_INSTANT
	bizContent.TotalAmount = util.CentsToYuan(charge.TotalFee)
	bizContent.Subject = charge.Body
	bizContent.TimeExpire = util.DongBaTime().Add(constant.TIMEOUT_PAY).Format(constant.TIME_FORMAT)
	bizContent.TimeoutExpress = fmt.Sprintf("%.fm", constant.TIMEOUT_PAY.Minutes())
	bizContent.QrPayMode = constant.QR_PAY_MOD_IFRAME

	bizContentByte, _ := json.Marshal(bizContent)
	m.BizContent = string(bizContentByte)

	m.SignType = constant.SIGN_TYPE_RSA2
	m.Sign = util.AliSign(structs.Map(m), c.PrivateKey)

	payResult, err := c.UnifiedOrder(m)
	if err != nil {
		return nil, err
	}

	return payResult, nil
}

func (c *AliWebPayClient) PayToClient(charge *_struct.Charge) (map[string]interface{}, error) {
	panic("implement me")
}
