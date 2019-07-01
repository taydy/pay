package _struct

/*
 *	微信交易所需要的请求响应实体类型，具体可参考页面 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
 */

// 微信统一下单请求参数，不需要证书
type WeChatUnifiedOrder struct {
	Appid          string `xml:"appid" json:"appid,omitempty" structs:"appid,omitempty"`                                  // 应用ID
	MchId          string `xml:"mch_id" json:"mch_id,omitempty" structs:"mch_id,omitempty"`                               // 商户号
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty" structs:"device_info,omitempty"`      // 设备号
	NonceStr       string `xml:"nonce_str" json:"nonce_str,omitempty" structs:"nonce_str,omitempty"`                      // 随机字符串
	Sign           string `xml:"sign" json:"sign,omitempty" structs:"sign,omitempty"`                                     // 签名
	SignType       string `xml:"sign_type,omitempty" json:"sign_type,omitempty" structs:"sign_type,omitempty"`            // 签名类型
	Body           string `xml:"body" json:"body,omitempty" structs:"body,omitempty"`                                     // 商品描述
	Detail         string `xml:"detail,omitempty" json:"detail,omitempty" structs:"detail,omitempty"`                     // 商品详情
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty" structs:"attach,omitempty"`                     // 附加数据
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no,omitempty" structs:"out_trade_no,omitempty"`             // 商户订单号
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty" structs:"fee_type,omitempty"`               // 货币类型
	TotalFee       int    `xml:"total_fee,omitempty" json:"total_fee,omitempty" structs:"total_fee,omitempty"`            // 总金额
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip,omitempty" structs:"spbill_create_ip,omitempty"` // 终端IP
	TimeStart      string `xml:"time_start" json:"time_start,omitempty" structs:"time_start,omitempty"`                   // 交易起始时间
	TimeExpire     string `xml:"time_expire" json:"time_expire,omitempty" structs:"time_expire,omitempty"`                // 交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty" json:"goods_tag,omitempty" structs:"goods_tag,omitempty"`            // 订单优惠标记
	NotifyUrl      string `xml:"notify_url" json:"notify_url,omitempty" structs:"notify_url,omitempty"`                   // 通知地址
	TradeType      string `xml:"trade_type" json:"trade_type,omitempty" structs:"trade_type,omitempty"`                   // 交易类型
	LimitPar       string `xml:"limit_pay,omitempty" json:"limit_pay,omitempty" structs:"limit_pay,omitempty"`            // 指定支付方式
	Receipt        string `xml:"receipt,omitempty" json:"receipt,omitempty" structs:"receipt,omitempty"`                  // 开发票入口开放标识
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty" structs:"openid,omitempty"`                     // 用户标识
}

type WeChatPayResult struct {
	/** Wechat 返回结果  用于判断请求是否成功 */
	ReturnCode string `xml:"return_code" json:"return_code,omitempty" structs:"return_code,omitempty"` // 返回状态码
	ReturnMsg  string `xml:"return_msg" json:"return_msg,omitempty" structs:"return_msg,omitempty"`    // 返回信息

	/** 当 return_code 为 SUCCESS 的时候返回 */
	AppID      string `xml:"appid,omitempty" json:"appid,omitempty" structs:"appid,omitempty"`                      // 应用APPID
	MchID      string `xml:"mch_id,omitempty" json:"mch_id,omitempty" structs:"mch_id,omitempty"`                   // 商户号
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty" structs:"device_info,omitempty"`    // 设备号
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty" structs:"nonce_str,omitempty"`          // 随机字符串
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty" structs:"sign,omitempty"`                         // 签名
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty" structs:"result_code,omitempty"`    // 业务结果
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty" structs:"err_code,omitempty"`             // 错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty" structs:"err_code_des,omitempty"` // 错误代码描述

	/** 在return_code 和 result_code都为SUCCESS的时候返回 */
	TradeType          string `xml:"trade_type,omitempty" json:"trade_type,omitempty" structs:"trade_type,omitempty"`                     // 交易类型
	PrepayId           string `xml:"prepay_id,omitempty" json:"prepay_id,omitempty" structs:"prepay_id,omitempty"`                        // 预支付交易会话标识
	CodeUrl            string `xml:"code_url,omitempty" json:"code_url,omitempty" structs:"code_url,omitempty"`                           // 二维码链接
	MWebUrl            string `xml:"mweb_url,omitempty" json:"mweb_url,omitempty" structs:"mweb_url,omitempty"`                           // 此链接为拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端
	TradeState         string `xml:"trade_state,omitempty" json:"trade_state,omitempty" structs:"trade_state,omitempty"`                  // 交易状态
	TradeStateDesc     string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty" structs:"trade_state_desc,omitempty"`   // 交易状态描述
	OutTradeNo         string `xml:"out_trade_no" json:"out_trade_no,omitempty" structs:"out_trade_no,omitempty"`                         // 商户订单号
	TransactionId      string `xml:"transaction_id" json:"transaction_id,omitempty" structs:"transaction_id,omitempty"`                   // 微信支付订单号
	TimeEnd            string `xml:"time_end" json:"time_end,omitempty" structs:"time_end,omitempty"`                                     // 支付完成时间
	CashFee            int    `xml:"cash_fee" json:"cash_fee,omitempty" structs:"cash_fee,omitempty"`                                     // 订单现金支付金额
	TotalFee           int    `xml:"total_fee" json:"total_fee,omitempty" structs:"total_fee,omitempty"`                                  // 订单总金额，单位为分
	BankType           string `xml:"bank_type" json:"bank_type,omitempty" structs:"bank_type,omitempty"`                                  // 银行类型，采用字符串类型的银行标识
	OpenId             string `xml:"openid" json:"openid,omitempty" structs:"openid,omitempty"`                                           // 用户在商户appid下的唯一标识
	IsSubscribe        string `xml:"is_subscribe" json:"is_subscribe,omitempty" structs:"is_subscribe,omitempty"`                         // 用户是否关注公众账号，Y-关注，N-未关注
	SettlementTotalFee int    `xml:"settlement_total_fee" json:"settlement_total_fee,omitempty" structs:"settlement_total_fee,omitempty"` // 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额
	FeeType            string `xml:"fee_type" json:"fee_type,omitempty" structs:"fee_type,omitempty"`                                     // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CashFeeType        string `xml:"cash_fee_type" json:"cash_fee_type,omitempty" structs:"cash_fee_type,omitempty"`                      // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	CouponFee          int    `xml:"coupon_fee" json:"coupon_fee,omitempty" structs:"coupon_fee,omitempty"`                               // “代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额
	CouponCount        int    `xml:"coupon_count" json:"coupon_count,omitempty" structs:"coupon_count,omitempty"`                         // 代金券使用数量
	Attach             string `xml:"attach" json:"attach,omitempty" structs:"attach"`                                                     // 附加数据，原样返回

	/** 微信转账成功后返回。 */
	PartnerTradeNo string `xml:"partner_trade_no" json:"partner_trade_no,omitempty" structs:"partner_trade_no,omitempty"` // 商户订单号，需保持历史全局唯一性(只能是字母或者数字，不能包含有其他字符)
	PaymentNo      string `xml:"payment_no" json:"payment_no,omitempty" structs:"payment_no,omitempty"`                   // 企业付款成功，返回的微信付款单号
	PaymentTime    string `xml:"payment_time" json:"payment_time,omitempty" structs:"payment_time,omitempty"`             // 企业付款成功时间
}

type WeChatTransfer struct {
	MchAppId       string `xml:"mch_appid" json:"mch_appid,omitempty" structs:"mch_appid,omitempty"`                      // 应用ID
	MchId          string `xml:"mchid" json:"mchid,omitempty" structs:"mchid,omitempty"`                                  // 商户号
	DeviceInfo     string `xml:"device_info,omitempty" json:"device_info,omitempty" structs:"device_info,omitempty"`      // 设备号
	NonceStr       string `xml:"nonce_str" json:"nonce_str,omitempty" structs:"nonce_str,omitempty"`                      // 随机字符串
	Sign           string `xml:"sign" json:"sign,omitempty" structs:"sign,omitempty"`                                     // 签名
	PartnerTradeNo string `xml:"partner_trade_no" json:"partner_trade_no,omitempty" structs:"partner_trade_no,omitempty"` // 商户订单号，需保持唯一性
	OpenId         string `xml:"openid" json:"openid,omitempty" structs:"openid,omitempty"`                               // 商户appid下，某用户的openid
	CheckName      string `xml:"check_name" json:"check_name,omitempty" structs:"check_name,omitempty"`                   // NO_CHECK：不校验真实姓名; FORCE_CHECK：强校验真实姓名
	ReUserName     string `xml:"re_user_name" json:"re_user_name,omitempty" structs:"re_user_name,omitempty"`             // 收款用户真实姓名。 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	Amount         int    `xml:"amount" json:"amount,omitempty" structs:"amount,omitempty"`                               // 企业付款金额，单位为分
	Desc           string `xml:"desc" json:"desc,omitempty" structs:"desc,omitempty"`                                     // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip,omitempty" structs:"spbill_create_ip,omitempty"` // 该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP。
}

type WeChatSnsOauthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
