package client

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	payErrors "github.com/taydy/pay/error"
	"github.com/taydy/pay/http"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"io/ioutil"
	"strconv"
	"time"
)

/** 微信默认客户端。 */
type WeChatClient struct {
	AppID          string // 公众账号ID
	AppIDSecretKey string // 公众账号secret key
	MchID          string // 商户号ID
	SecretKey      string // 密钥
	PrivateKey     string // 私钥文件内容
	PublicKey      string // 公钥文件内容
	TLSConfig      *tls.Config
}

/**
 *	微信统一下单，已检验签名信息。
 */
func (c *WeChatClient) UnifiedOrder(unifiedOrder *_struct.WeChatUnifiedOrder) (*_struct.WeChatPayResult, error) {
	xmlResponse, err := http.XmlPost(constant.WECHAT_PAY_UNIFIED_ORDER, unifiedOrder, time.Second*10)
	if err != nil {
		fmt.Printf("wechat unified order error: %v \n", err)
		return nil, payErrors.ErrWXPayError
	}

	defer xmlResponse.Body.Close()

	data, _ := ioutil.ReadAll(xmlResponse.Body)
	fmt.Printf("xml response: %s \n", data)

	return c.validResult(data)
}

/**
 * 订单查询，微信订单号(transaction_id)和商户订单号(out_trade_no)二选一。
 */
func (c *WeChatClient) OrderQuery(unifiedOrder *_struct.WeChatUnifiedOrder) (*_struct.WeChatPayResult, error) {
	xmlResponse, err := http.XmlPost(constant.WECHAT_PAY_ORDER_QUERY, unifiedOrder, time.Second*10)
	if err != nil {
		fmt.Printf("wechat unified order error: %v \n", err)
		return nil, payErrors.ErrWXPayError
	}

	defer xmlResponse.Body.Close()

	data, _ := ioutil.ReadAll(xmlResponse.Body)
	fmt.Printf("wechat order query response: %s \n", data)

	return c.validResult(data)
}

func (c *WeChatClient) OrderQueryByOutTradeNo(outTradeNo string) (*_struct.WeChatPayResult, error) {
	unifiedOrder := _struct.WeChatUnifiedOrder{}
	unifiedOrder.Appid = c.AppID
	unifiedOrder.MchId = c.MchID
	unifiedOrder.OutTradeNo = outTradeNo
	unifiedOrder.NonceStr = util.RandomStr()
	unifiedOrder.Sign = util.WeChatSign(structs.Map(unifiedOrder), c.SecretKey)

	return c.OrderQuery(&unifiedOrder)
}

/**
 * 获取订单支付结果，简单判断交易状态是否为 SUCCESS.
 */
func (c *WeChatClient) isPaySuccess(outTradeNo string) bool {
	unifiedOrder := _struct.WeChatUnifiedOrder{}
	unifiedOrder.Appid = c.AppID
	unifiedOrder.MchId = c.MchID
	unifiedOrder.OutTradeNo = outTradeNo
	unifiedOrder.NonceStr = util.RandomStr()
	unifiedOrder.Sign = util.WeChatSign(structs.Map(unifiedOrder), c.SecretKey)

	payResult, err := c.OrderQuery(&unifiedOrder)
	if err != nil {
		fmt.Printf("Get order payment results error, out trade no: %s, error: %v \n", outTradeNo, err)
		return false
	}
	return payResult.TradeState == constant.TRADE_STATE_SUCCESS
}

func (c *WeChatClient) Transfer(transfer *_struct.WeChatTransfer, tlsConfig *tls.Config) (*_struct.WeChatPayResult, error) {
	transfer.MchAppId = c.AppID
	transfer.MchId = c.MchID
	transfer.NonceStr = util.RandomStr()
	transfer.SpbillCreateIp = util.IfFunc(transfer.SpbillCreateIp != "", transfer.SpbillCreateIp, func() string { return util.LocalIP() }).(string)
	transfer.CheckName = util.If(transfer.ReUserName == "", constant.NO_CHECK, constant.FORCE_CHECK).(string)
	transfer.Sign = util.WeChatSign(structs.Map(transfer), c.SecretKey)

	return c.doTransfer(transfer, tlsConfig)
}

/**
 * 单笔转账到微信账户。
 */
func (c *WeChatClient) doTransfer(transfer *_struct.WeChatTransfer, tlsConfig *tls.Config) (*_struct.WeChatPayResult, error) {
	xmlResponse, err := http.XmlSecurePost(constant.WECHAT_PAY_TRANSFER, transfer, time.Second*10, tlsConfig)
	if err != nil {
		fmt.Printf("wechat unified order error: %v \n", err)
		return nil, payErrors.ErrWXPayError
	}

	defer xmlResponse.Body.Close()

	data, _ := ioutil.ReadAll(xmlResponse.Body)
	fmt.Printf("wechat order query response: %s \n", data)

	payResult := &_struct.WeChatPayResult{}
	decodeErr := xml.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("wechat transfer error, error: %v \n", decodeErr)
		return nil, payErrors.ErrWXPayError
	}
	if payResult.ReturnCode == constant.SUCCESS || payResult.ResultCode == constant.SUCCESS {
		return payResult, nil
	}
	return nil, payErrors.NewBadRequestError(payErrors.WX_PAY_ERROR, payResult.ErrCodeDes)
}

func (c *WeChatClient) validResult(data []byte) (*_struct.WeChatPayResult, error) {
	payResult := &_struct.WeChatPayResult{}
	decodeErr := xml.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("wechat unified order error, error: %v \n", decodeErr)
		return nil, payErrors.ErrWXPayError
	}
	if payResult.ReturnCode == constant.FAIL {
		fmt.Printf("wechat unified order error, error: %s \n", payResult.ReturnMsg)
		return nil, payErrors.ErrWXPayError
	}
	if payResult.ResultCode == constant.FAIL {
		fmt.Printf("wechat unified order error, error code: %s, error desc: %s \n", payResult.ResultCode, payResult.ErrCodeDes)
		return nil, payErrors.NewBadRequestError(payErrors.PAY_FAILED, payResult.ErrCodeDes)
	}

	if !util.WeChatValidSign(structs.Map(payResult), c.SecretKey) {
		fmt.Printf("wechat response sign error,  result: %v \n", payResult)
		return nil, payErrors.ErrWXPayError
	}
	return payResult, nil
}

func (c *WeChatClient) ChargeToWeChatUnifiedOrder(charge *_struct.Charge) *_struct.WeChatUnifiedOrder {
	var m = &_struct.WeChatUnifiedOrder{}
	m.Appid = c.AppID
	m.MchId = c.MchID
	m.Body = charge.Body
	m.Detail = charge.Detail
	m.OutTradeNo = charge.OutTradeNo
	m.TimeStart = charge.TimeStart
	m.TimeExpire = charge.TimeExpire
	m.FeeType = charge.FeeType
	m.TotalFee = charge.TotalFee
	m.NotifyUrl = charge.NotifyUrl
	m.Openid = charge.OpenID
	m.SignType = constant.SIGN_TYPE_MD5
	m.SpbillCreateIp = util.IfFunc(charge.SpbillCreateIp != "", charge.SpbillCreateIp, func() string { return util.LocalIP() }).(string)
	m.NonceStr = util.RandomStr()

	return m
}

func (c *WeChatClient) SnsOauthAccessToken(code string) (*_struct.WeChatSnsOauthResponse, error) {
	url := fmt.Sprintf(constant.WECHAT_SNS_OAUTH2+"?appid=%s&secret=%s&code=%s&grant_type=authorization_code", c.AppID, c.AppIDSecretKey, code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	oauth := &_struct.WeChatSnsOauthResponse{}
	err = json.Unmarshal(data, oauth)
	if err != nil {
		return nil, err
	}
	if oauth.ErrMsg != "" {
		return nil, payErrors.NewBadRequestError(oauth.ErrCode, oauth.ErrMsg)
	}
	return oauth, nil
}

/**
 * 退款。
 */
func (c *WeChatClient) Refund(tradeNo string, refundFee int) (bool, error) {
	refundReq := &_struct.WXRefundReq{
		AppID:         c.AppID,
		MchID:         c.MchID,
		NonceStr:      util.RandomStr(),
		TotalFee:      refundFee,
		OutRefundNo:   fmt.Sprintf("%s%s", tradeNo, strconv.FormatInt(time.Now().UnixNano(), 10)[10:16]),
		TransactionID: tradeNo,
		RefundFee:     refundFee,
	}
	refundReq.Sign = util.WeChatSign(structs.Map(refundReq), c.SecretKey)
	xmlResponse, err := http.XmlSecurePost(constant.WECHAT_PAY_REFUND, refundReq, time.Second*10, c.TLSConfig)
	if err != nil {
		fmt.Printf("wechat refund order error: %v \n", err)
		return false, payErrors.ErrWXPayError
	}

	defer xmlResponse.Body.Close()

	data, _ := ioutil.ReadAll(xmlResponse.Body)
	fmt.Printf("wechat refund query response: %s \n", data)

	payResult := &_struct.WXRefundResp{}
	decodeErr := xml.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("wechat transfer error, error: %v \n", decodeErr)
		return false, payErrors.ErrWXPayError
	}
	if payResult.ReturnCode == constant.SUCCESS && payResult.ResultCode == constant.SUCCESS {
		return true, nil
	}
	return false, payErrors.NewBadRequestError(payErrors.WX_PAY_ERROR, payResult.ErrCodeDes)
}
