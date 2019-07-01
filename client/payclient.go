package client

import (
	"github.com/taydy/pay/struct"
)

/** 支付接口，包含用户支付和支付给用户。 */
type PayClient interface {
	Pay(charge *_struct.Charge) (map[string]interface{}, error)
	PayToClient(charge *_struct.Charge) (map[string]interface{}, error)
}
