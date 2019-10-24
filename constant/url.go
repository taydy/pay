package constant

const (
	/** ALI PAY */
	ALI_PAY_GATEWAY = "https://openapi.alipay.com/gateway.do"

	/** WECHAT PAY */
	WECHAT_PAY_ROOT          = "https://api.mch.weixin.qq.com"
	WECHAT_PAY_UNIFIED_ORDER = WECHAT_PAY_ROOT + "/pay/unifiedorder"
	WECHAT_PAY_ORDER_QUERY   = WECHAT_PAY_ROOT + "/pay/orderquery"
	WECHAT_PAY_CLOSE_ORDER   = WECHAT_PAY_ROOT + "/pay/closeorder"
	WECHAT_PAY_REFUND        = WECHAT_PAY_ROOT + "/secapi/pay/refund"
	WECHAT_PAY_REFUND_QUERY  = WECHAT_PAY_ROOT + "/pay/refundquery"
	WECHAT_PAY_TRANSFER      = WECHAT_PAY_ROOT + "/mmpaymkttransfers/promotion/transfers"
	WECHAT_SNS_OAUTH2        = "https://api.weixin.qq.com/sns/oauth2/access_token"
)
