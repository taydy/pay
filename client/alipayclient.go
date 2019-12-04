package client

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/taydy/pay/constant"
	payErrors "github.com/taydy/pay/error"
	"github.com/taydy/pay/struct"
	"github.com/taydy/pay/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

/** 阿里支付默认客户端。 */
type AliPayClient struct {
	AppID      string          // 支付宝分配给开发者的应用ID ps: 查询订单用
	PrivateKey *rsa.PrivateKey // 私钥
	PublicKey  *rsa.PublicKey  // 公钥
}

/**
 * 支付宝统一下单。
 */
func (c *AliPayClient) UnifiedOrder(unifiedOrder *_struct.AliPayUnifiedOrder) (map[string]interface{}, error) {
	url := ToURL(constant.ALI_PAY_GATEWAY, structs.Map(unifiedOrder))
	fmt.Printf("ali pay unified order url : %s \n", url)
	return map[string]interface{}{"url": url}, nil
}

func (c *AliPayClient) TradePrecreate(unifiedOrder *_struct.AliPayUnifiedOrder) (*_struct.AlipayTradePrecreateResponse, error) {
	url := ToURL(constant.ALI_PAY_GATEWAY, structs.Map(unifiedOrder))
	fmt.Printf("ali pay trade precreate order url : %s \n", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("alipay unified order error: %v \n", err)
		return nil, payErrors.ErrAlipayError
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ali pay order query response: %s \n", string(data))

	err = c.validResult(data, constant.ALI_PAY_TRADE_PRECREATE_RESPONSE)
	if err != nil {
		return nil, err
	}

	payResult := &_struct.AlipayTradePrecreateResponse{}
	decodeErr := json.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("alipay order query error, error: %v \n", decodeErr)
		return nil, payErrors.ErrAlipayError
	}

	return payResult, nil
}

/**
 * 订单查询。
 */
func (c *AliPayClient) OrderQuery(unifiedOrder *_struct.AliPayUnifiedOrder) (*_struct.AliPayResult, error) {
	url := ToURL(constant.ALI_PAY_GATEWAY, structs.Map(unifiedOrder))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("alipay unified order error: %v \n", err)
		return nil, payErrors.ErrAlipayError
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ali pay order query response: %s \n", string(data))

	err = c.validResult(data, constant.ALI_PAY_TRADE_QUERY_RESPONSE)
	if err != nil {
		return nil, err
	}

	payResult := &_struct.AliPayResult{}
	decodeErr := json.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("alipay order query error, error: %v \n", decodeErr)
		return nil, payErrors.ErrAlipayError
	}

	return payResult, nil
}

/**
 * 通过内部订单号查询订单。
 */
func (c *AliPayClient) OrderQueryByOutTradeNo(outTradeNo string) (*_struct.AliPayResult, error) {
	unifiedOrder := _struct.AliPayUnifiedOrder{}
	unifiedOrder.AppId = c.AppID
	unifiedOrder.Method = constant.ALI_PAY_API_QUERY
	unifiedOrder.Charset = constant.CHARSET_UTF8
	unifiedOrder.Timestamp = time.Now().Format(constant.TIME_FORMAT)
	unifiedOrder.Version = "1.0"
	unifiedOrder.SignType = constant.SIGN_TYPE_RSA2

	var bizContent = &_struct.BizContent{}
	bizContent.OutTradeNo = outTradeNo
	bizContentByte, _ := json.Marshal(bizContent)
	unifiedOrder.BizContent = string(bizContentByte)

	unifiedOrder.Sign = util.AliSign(structs.Map(unifiedOrder), c.PrivateKey)

	return c.OrderQuery(&unifiedOrder)
}

/**
 * 关闭订单。
 */
func (c *AliPayClient) CloseOrder(unifiedOrder *_struct.AliPayUnifiedOrder) (*_struct.AliPayResult, error) {
	url := ToURL(constant.ALI_PAY_GATEWAY, structs.Map(unifiedOrder))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("alipay close order error: %v \n", err)
		return nil, payErrors.ErrAlipayCloseOrderError
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ali pay order query response: %s \n", string(data))

	err = c.validResult(data, constant.ALI_PAY_TRADE_CLOSE_RESPONSE)
	if err != nil {
		return nil, err
	}

	payResult := &_struct.AliPayResult{}
	decodeErr := json.Unmarshal(data, payResult)
	if decodeErr != nil {
		fmt.Printf("alipay order query error, error: %v \n", decodeErr)
		return nil, payErrors.ErrAlipayError
	}

	return payResult, nil
}

/**
 * 通过内部订单号关闭订单。
 */
func (c *AliPayClient) CloseOrderByOutTradeNo(outTradeNo string) (*_struct.AliPayResult, error) {
	unifiedOrder := _struct.AliPayUnifiedOrder{}
	unifiedOrder.AppId = c.AppID
	unifiedOrder.Method = constant.ALI_PAY_API_CLOSE
	unifiedOrder.Charset = constant.CHARSET_UTF8
	unifiedOrder.Timestamp = time.Now().Format(constant.TIME_FORMAT)
	unifiedOrder.Version = "1.0"
	unifiedOrder.SignType = constant.SIGN_TYPE_RSA2

	var bizContent = &_struct.BizContent{}
	bizContent.OutTradeNo = outTradeNo
	bizContentByte, _ := json.Marshal(bizContent)
	unifiedOrder.BizContent = string(bizContentByte)

	unifiedOrder.Sign = util.AliSign(structs.Map(unifiedOrder), c.PrivateKey)

	return c.CloseOrder(&unifiedOrder)
}

func (c *AliPayClient) Transfer(transfer *_struct.AliTransfer) (*_struct.AliTransferResult, error) {
	unifiedOrder := _struct.AliPayUnifiedOrder{}
	unifiedOrder.AppId = c.AppID
	unifiedOrder.Method = constant.ALI_PAY_API_TRANSFER
	unifiedOrder.Charset = constant.CHARSET_UTF8
	unifiedOrder.Timestamp = time.Now().Format(constant.TIME_FORMAT)
	unifiedOrder.Version = "1.0"
	unifiedOrder.SignType = constant.SIGN_TYPE_RSA2

	bizContentByte, _ := json.Marshal(transfer)
	unifiedOrder.BizContent = string(bizContentByte)

	unifiedOrder.Sign = util.AliSign(structs.Map(unifiedOrder), c.PrivateKey)

	return c.doTransfer(&unifiedOrder)
}

/**
 * 单笔转账到支付宝账户。
 */
func (c *AliPayClient) doTransfer(unifiedOrder *_struct.AliPayUnifiedOrder) (*_struct.AliTransferResult, error) {
	url := ToURL(constant.ALI_PAY_GATEWAY, structs.Map(unifiedOrder))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("alipay transfer error: %v \n", err)
		return nil, payErrors.ErrAlipayCloseOrderError
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ali pay transfer query response: %s \n", string(data))

	err = c.validResult(data, constant.ALI_PAY_FUND_TRANS_TO_ACCOUNT_TRANSFER_RESPONSE)
	if err != nil {
		return nil, err
	}

	transferResult := &_struct.AliTransferResult{}
	decodeErr := json.Unmarshal(data, transferResult)
	if decodeErr != nil {
		fmt.Printf("alipay order query error, error: %v \n", decodeErr)
		return nil, payErrors.ErrAlipayError
	}

	return transferResult, nil
}

func (c *AliPayClient) validResult(data []byte, responseKey string) error {

	temp := string(data)
	keyIndex := strings.Index(temp, responseKey)
	signIndex := strings.Index(temp, "sign")
	var keyStart, keyEnd, signStart, signEnd int
	if keyIndex < signIndex {
		keyStart = keyIndex + len(responseKey+"\":")
		keyEnd = signIndex - len(",\"")
		signStart = signIndex + len("sign"+"\":\"")
		signEnd = len(temp) - len("\"}")
	} else {
		keyStart = keyIndex + len(responseKey+"\":")
		keyEnd = len(temp) - len("}")
		signStart = signIndex + len("sign"+"\":\"")
		signEnd = keyIndex - len("\",\"")
	}

	if !util.ValidAliSign(temp[keyStart:keyEnd], temp[signStart:signEnd], c.PublicKey) {
		fmt.Printf("alipay order query valid sign error")
		return payErrors.ErrAlipaySignCheckError
	}
	return nil
}

func (c *AliPayClient) ValidAsyncResultByStruct(PayNotifyResult *_struct.ALIPayNotifyResult) bool {
	return c.ValidAsyncResult(structs.Map(PayNotifyResult))
}

func (c *AliPayClient) ValidAsyncResult(params map[string]interface{}) bool {
	log.Println(params)
	var keys []string
	var origin []string

	for k, v := range params {
		if k != "sign" && k != "sign_type" && v != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		origin = append(origin, fmt.Sprintf("%s=%v", k, params[k]))
	}

	str := strings.Join(origin, "&")
	return util.ValidAliSign(str, params["sign"].(string), c.PublicKey)
}

/**
 * 退款。
 */
func (c *AliPayClient) Refund(tradeNo string, refundFee int) (bool, error) {
	bizContent := make(map[string]interface{})
	bizContent["trade_no"] = tradeNo
	bizContent["refund_amount"] = util.CentsToYuan(refundFee)
	bizContentByte, _ := json.Marshal(bizContent)
	refundReq := map[string]interface{}{
		"app_id":         c.AppID,
		"method":         constant.ALI_PAY_API_REFUND,
		"charset":        "UTF-8",
		"sign_type":      "RSA2",
		"timestamp":      time.Now().Format(constant.TIME_FORMAT),
		"version":        "1.0",
		"biz_content":    string(bizContentByte),
		"out_request_no": fmt.Sprintf("%s%s", tradeNo, strconv.FormatInt(time.Now().UnixNano(), 10)[10:16]),
	}
	refundReq["sign"] = util.AliSign(refundReq, c.PrivateKey)

	url := ToURL(constant.ALI_PAY_GATEWAY, refundReq)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("alipay refund %s error: %v \n", tradeNo, err)
		return false, payErrors.ErrAlipayCloseOrderError
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("ali pay refund %s query response: %s \n", tradeNo, string(data))

	err = c.validResult(data, constant.ALI_PAY_TRADE_REFUND_RESPONSE)
	if err != nil {
		return false, err
	}

	refundResult := &_struct.AliTradeRefundResult{}
	decodeErr := json.Unmarshal(data, refundResult)
	if decodeErr != nil {
		fmt.Printf("alipay refund order query error, error: %v \n", decodeErr)
		return false, payErrors.ErrAlipayError
	}

	return util.RefundSuccess(refundResult), nil
}

/**
 * 生成蚂蚁金服支付接口连接。
 */
func ToURL(payUrl string, m map[string]interface{}) string {
	var keys []string
	var sorted []string
	for k, v := range m {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%v", k, url.QueryEscape(m[k].(string))))
	}
	sorted = append(sorted, fmt.Sprintf("%s=%v", "sign", url.QueryEscape(m["sign"].(string))))

	return fmt.Sprintf("%s?%s", payUrl, strings.Join(sorted, "&"))
}
