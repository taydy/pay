package constant

import "time"

const (
	WX_TRADE_TYPE_JSAPI  = "JSAPI"
	WX_TRADE_TYPE_NATIVE = "NATIVE"
	WX_TRADE_TYPE_APP    = "APP"
	WX_TRADE_TYPE_H5     = "MWEB"
)

const (
	SIGN_TYPE_MD5  = "MD5"
	SIGN_TYPE_SHA  = "HMAC-SHA256"
	SIGN_TYPE_RSA2 = "RSA2"
	SIGN_TYPE_RSA  = "RSA"
)

const (
	TRADE_STATE_SUCCESS    = "SUCCESS"    // 支付成功
	TRADE_STATE_REFUND     = "REFUND"     // 转入退款
	TRADE_STATE_NOTPAY     = "NOTPAY"     // 未支付
	TRADE_STATE_CLOSED     = "CLOSED"     // 已关闭
	TRADE_STATE_REVOKED    = "REVOKED"    // 已撤销（付款码支付）
	TRADE_STATE_USERPAYING = "USERPAYING" // 用户支付中（付款码支付）
	TRADE_STATE_PAYERROR   = "PAYERROR"   // 支付失败(其他原因，如银行返回失败)

	TRADE_STATE_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	TRADE_STATE_TRADE_CLOSED   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	TRADE_STATE_TRADE_SUCCESS  = "TRADE_SUCCESS"  // 交易支付成功
	TRADE_STATE_TRADE_FINISHED = "TRADE_FINISHED" // 交易结束，不可退款
)

const (
	ALI_PAY_API_PAGE_PAY     = "alipay.trade.page.pay"                          // 统一收单下单并支付页面接口 https://docs.open.alipay.com/api_1/alipay.trade.page.pay/
	ALI_PAY_API_WAP_PAY      = "alipay.trade.wap.pay"                           // 手机端统一下单接口
	ALI_PAY_API_APP_PAY      = "alipay.trade.app.pay"                           // app 端统一下单
	ALI_PAY_API_PRE_CREATE   = "alipay.trade.precreate"                         // 统一收单线下交易预创建（扫码支付）
	ALI_PAY_API_QUERY        = "alipay.trade.query"                             // 统一收单线下交易查询接口 https://docs.open.alipay.com/api_1/alipay.trade.query
	ALI_PAY_API_REFUND       = "alipay.trade.refund"                            // 统一收单交易退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund
	ALI_PAY_API_REFUND_QUERY = "alipay.trade.fastpay.refund.query"              // 统一收单交易退款查询接口 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
	ALI_PAY_API_CLOSE        = "alipay.trade.close"                             // 统一收单交易关闭接口 https://docs.open.alipay.com/api_1/alipay.trade.close
	ALI_PAY_API_BILL         = "alipay.data.dataservice.bill.downloadurl.query" // 查询对账单下载地址 https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
	ALI_PAY_API_TRANSFER     = "alipay.fund.trans.toaccount.transfer"           // 单笔转账到支付宝账户接口 https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
)

const (
	ALI_PAY_FUND_TRANS_TO_ACCOUNT_TRANSFER_RESPONSE = "alipay_fund_trans_toaccount_transfer_response"
	ALI_PAY_TRADE_QUERY_RESPONSE                    = "alipay_trade_query_response"
	ALI_PAY_TRADE_CLOSE_RESPONSE                    = "alipay_trade_close_response"
	ALI_PAY_TRADE_PRECREATE_RESPONSE                = "alipay_trade_precreate_response"
	ALI_PAY_TRADE_REFUND_RESPONSE                   = "alipay_trade_refund_response"
)

const (
	ALI_PAY_PRODUCT_CODE_WAP             = "QUICK_WAP_WAY"
	ALI_PAY_PRODUCT_CODE_MSECURITY       = "QUICK_MSECURITY_PAY"
	ALI_PAY_PRODUCT_CODE_INSTANT         = "FAST_INSTANT_TRADE_PAY"
	ALI_PAY_PRODUCT_DACE_TO_FACE_PAYMENT = "FACE_TO_FACE_PAYMENT"
)

const (
	QR_PAY_MOD_IFRAME_SIMPLE = "0" // 订单码-简约前置模式，对应 iframe 宽度不能小于600px，高度不能小于300px
	QR_PAY_MOD_IFRAME        = "1" // 订单码-前置模式，对应iframe 宽度不能小于 300px，高度不能小于600px
	QR_PAY_MOD_IFRAME_MINI   = "3" // 订单码-迷你前置模式，对应 iframe 宽度不能小于 75px，高度不能小于75px
	QR_PAY_MOD_IFRAME_AUTO   = "4" // 订单码-可定义宽度的嵌入式二维码，商户可根据需要设定二维码的大小
	QR_PAY_MOD_JUMP          = "2" // 订单码-跳转模式
)

const (
	CHARSET_UTF8   = "UTF-8"
	CHARSET_GBK    = "GBK"
	CHARSET_GB2312 = "GB2312"
)

const (
	TIME_FORMAT        = "2006-01-02 15:04:05"
	EXPIRE_TIME_FORMAT = "2006-01-02 15:04"
)

const (
	TIMEOUT_PAY = 5 * time.Minute
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

const (
	NO_CHECK    = "NO_CHECK"    // 不校验真实姓名
	FORCE_CHECK = "FORCE_CHECK" // 强校验真实姓名
)
