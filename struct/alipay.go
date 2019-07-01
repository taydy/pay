package _struct

/*
 *	阿里支付所需要的请求响应实体类型，具体可参考页面 https://docs.open.alipay.com/api_1/alipay.trade.page.pay
 */

// 阿里支付统一下单请求参数，不需要证书
type AliPayUnifiedOrder struct {
	/** 公共请求参数 */
	AppId        string `xml:"app_id" json:"app_id,omitempty" structs:"app_id,omitempty"`                         // 支付宝分配给开发者的应用ID
	Method       string `xml:"method" json:"method,omitempty" structs:"method,omitempty"`                         // 接口名称
	Format       string `xml:"format" json:"format,omitempty" structs:"format,omitempty"`                         // 仅支持JSON
	ReturnUrl    string `xml:"return_url" json:"return_url,omitempty" structs:"return_url,omitempty"`             // HTTP/HTTPS开头字符串
	Charset      string `xml:"charset" json:"charset,omitempty" structs:"charset,omitempty"`                      // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType     string `xml:"sign_type" json:"sign_type,omitempty" structs:"sign_type,omitempty"`                // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Sign         string `xml:"sign" json:"sign,omitempty" structs:"sign,omitempty"`                               // 商户请求参数的签名串
	Timestamp    string `xml:"timestamp" json:"timestamp,omitempty" structs:"timestamp,omitempty"`                // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `xml:"version" json:"version,omitempty" structs:"version,omitempty"`                      // 调用的接口版本，固定为：1.0
	NotifyUrl    string `xml:"notify_url" json:"notify_url,omitempty" structs:"notify_url,omitempty"`             // 支付宝服务器主动通知商户服务器里指定的页面http/https路径
	AppAuthToken string `xml:"app_auth_token" json:"app_auth_token,omitempty" structs:"app_auth_token,omitempty"` // 应用ID
	BizContent   string `xml:"biz_content" json:"biz_content,omitempty" structs:"biz_content,omitempty"`          // 业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

type BizContent struct {
	/** 请求参数 */
	OutTradeNo         string `json:"out_trade_no,omitempty"`         // 商户网站唯一订单号
	ProductCode        string `json:"product_code,omitempty"`         // 销售产品码，商家和支付宝签约的产品码
	TotalAmount        string `json:"total_amount,omitempty"`         // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	Subject            string `json:"subject,omitempty"`              // 商品的标题/交易标题/订单标题/订单关键字等。
	Body               string `json:"body,omitempty"`                 // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	TimeExpire         string `json:"time_expire,omitempty"`          // 绝对超时时间，格式为yyyy-MM-dd HH:mm。 注：1）以支付宝系统时间为准；2）如果和timeout_express参数同时传入，以time_expire为准。
	PassBackParams     string `json:"passback_params,omitempty"`      // 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	GoodTypes          string `json:"goods_type,omitempty"`           // 商品主类型：0—虚拟类商品，1—实物类商品注：虚拟类商品不支持使用花呗渠道
	TimeoutExpress     string `json:"timeout_express,omitempty"`      // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。注：若为空，则默认为15d。
	PromoParams        string `json:"promo_params,omitempty"`         // 优惠参数注：仅与支付宝协商后可用
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`  // 可用渠道，用户只能在指定渠道范围内支付当有多个渠道时用“,”分隔注：与disable_pay_channels互斥
	StoreId            string `json:"store_id,omitempty"`             // 商户门店编号。该参数用于请求参数中以区分各门店，非必传项
	DisablePayChannels string `json:"disable_pay_channels,omitempty"` // 禁用渠道，用户不可用指定渠道支付当有多个渠道时用“,”分隔注：与enable_pay_channels互斥
	QrPayMode          string `json:"qr_pay_mode,omitempty"`          // PC扫码支付的方式，支持前置模式和跳转模式
	QrcodeWidth        int    `json:"qrcode_width,omitempty"`         // 商户自定义二维码宽度 注：qr_pay_mode=4时该参数生效
	IntegrationType    string `json:"integration_type,omitempty"`     // 请求后页面的集成方式,默认值为PCWEB  1.ALIAPP：支付宝钱包内; 2.PCWEB：PC端访问
	RequestFromUrl     string `json:"request_from_url,omitempty"`     // 请求来源地址。如果使用ALIAPP的集成方式，用户中途取消支付会返回该地址。
	BusinessParams     string `json:"business_params,omitempty"`      // 商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式
}

type AliPayResult struct {
	AlipayTradeQueryResponse struct {
		Code    string `json:"code,omitempty" structs:"code,omitempty"`         // 网关返回码
		Msg     string `json:"msg,omitempty" structs:"msg,omitempty"`           // 网关返回码描述
		SubCode string `json:"sub_code,omitempty" structs:"sub_code,omitempty"` // 业务返回码
		SubMsg  string `json:"sub_msg,omitempty" structs:"sub_msg,omitempty"`   // 业务返回码描述

		TradeNo         string `json:"trade_no,omitempty" structs:"trade_no,omitempty"`             // 支付宝交易号
		OutTradeNo      string `json:"out_trade_no,omitempty" structs:"out_trade_no,omitempty"`     // 商户订单号
		BuyerLogonId    string `json:"buyer_logon_id,omitempty" structs:"buyer_logon_id,omitempty"` // 买家支付宝账号
		TradeStatus     string `json:"trade_status,omitempty" structs:"trade_status,omitempty"`     // 交易状态
		TotalAmount     string `json:"total_amount,omitempty" structs:"total_amount,omitempty"`     // 交易金额
		TransCurrency   string `json:"trans_currency,omitempty" structs:"trans_currency"`           // 标价币种
		SettleCurrency  string `json:"settle_currency,omitempty" structs:"settle_currency"`         // 订单结算币种
		SettleAmount    string `json:"settle_amount,omitempty" structs:"settle_amount"`             // 结算币种订单金额
		PayCurrency     string `json:"pay_currency,omitempty" structs:"pay_currency"`               // 订单支付币种
		PayAmount       string `json:"pay_amount,omitempty" structs:"pay_amount"`                   // 支付币种订单金额
		SettleTransRate string `json:"settle_trans_rate,omitempty" structs:"settle_trans_rate"`     // 结算币种兑换标价币种汇率
		TransPayRate    string `json:"trans_pay_rate,omitempty" structs:"trans_pay_rate"`           // 标价币种兑换支付币种汇率
		BuyerPayAmount  string `json:"buyer_pay_amount,omitempty" structs:"buyer_pay_amount"`       // 买家实付金额，单位为元，两位小数
		PointAmount     string `json:"point_amount,omitempty" structs:"point_amount"`               // 积分支付的金额，单位为元，两位小数
		Subject         string `json:"subject,omitempty,omitempty" structs:"subject"`               // 订单标题
		InvoiceAmount   string `json:"invoice_amount,omitempty" structs:"invoice_amount"`           // 交易中用户支付的可开具发票的金额，单位为元，两位小数。该金额代表该笔交易中可以给用户开具发票的金额
		ReceiptAmount   string `json:"receipt_amount,omitempty" structs:"receipt_amount"`           // 实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
		SellerId        string `json:"seller_id,omitempty" structs:"seller_id"`                     // 收款支付宝账号对应的支付宝唯一用户号
	} `json:"alipay_trade_query_response" structs:"alipay_trade_query_response,omitempty"`
	Sign string `json:"sign" structs:"sign,omitempty"`
}

type AliTransfer struct {
	OutBizNo      string `json:"out_biz_no"`      // 商户转账唯一订单号。发起转账来源方定义的转账单据ID，用于将转账回执通知给来源方。
	PayeeType     string `json:"payee_type"`      // 收款方账户类型
	PayeeAccount  string `json:"payee_account"`   // 收款方账户。与payee_type配合使用。付款方和收款方不能是同一个账户。
	Amount        string `json:"amount"`          // 转账金额，单位：元。 只支持2位小数，小数点前最大支持13位，金额必须大于等于0.1元。最大转账金额以实际签约的限额为准。
	PayerShowName string `json:"payer_show_name"` // 付款方姓名（最长支持100个英文/50个汉字）。显示在收款方的账单详情页。如果该字段不传，则默认显示付款方的支付宝认证姓名或单位名称。
	PayeeRealName string `json:"payee_real_name"` // 收款方真实姓名（最长支持100个英文/50个汉字）。 如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。
	Remark        string `json:"remark"`          // 转账备注（支持200个英文/100个汉字）。 当付款方为企业账户，且转账金额达到（大于等于）50000元，remark不能为空。收款方可见，会展示在收款用户的收支详情中。
}

type AliTransferResult struct {
	AlipayFundTransToaccountTransferResponse struct {
		Code    string `json:"code,omitempty"`     // 网关返回码
		Msg     string `json:"msg,omitempty"`      // 网关返回码描述
		SubCode string `json:"sub_code,omitempty"` // 业务返回码
		SubMsg  string `json:"sub_msg,omitempty"`  // 业务返回码描述

		OutBizNo string `json:"out_biz_no"` // 商户转账唯一订单号。发起转账来源方定义的转账单据ID，用于将转账回执通知给来源方。
		OrderId  string `json:"order_id"`   // 支付宝转账单据号，成功一定返回，失败可能不返回也可能返回。
		PayDate  string `json:"pay_date"`   // 支付时间：格式为yyyy-MM-dd HH:mm:ss，仅转账成功返回。
	} `json:"alipay_fund_trans_toaccount_transfer_response"`
	Sign string `json:"sign"`
}
