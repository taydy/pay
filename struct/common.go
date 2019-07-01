package _struct

// 默认支付参数
type Charge struct {
	Body           string `xml:"body" json:"body,omitempty"`                         // 商品描述
	Detail         string `xml:"detail" json:"detail,omitempty"`                     // 商品详情
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no,omitempty"`         // 商户订单号
	FeeType        string `xml:"fee_type" json:"fee_type,omitempty"`                 // 货币类型
	TotalFee       int    `xml:"total_fee" json:"total_fee,omitempty"`               // 总金额
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip,omitempty"` // 终端IP
	NotifyUrl      string `xml:"notify_url" json:"notify_url,omitempty"`             // 通知地址
	TradeType      string `xml:"trade_type" json:"trade_type,omitempty"`             // 交易类型
	OpenID         string `json:"openid,omitempty"`                                  // open id
	TimeStart      string `xml:"time_start" json:"time_start,omitempty"`             // 交易起始时间
	TimeExpire     string `xml:"time_expire" json:"time_expire,omitempty"`           // 交易结束时间
	ReturnUrl      string `xml:"return_url" json:"return_url,omitempty"`             // 回调地址
}
