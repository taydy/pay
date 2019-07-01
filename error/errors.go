package error

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int               `json:"statusCode"`
	ErrCode    int               `json:"code"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data,omitempty"`
}

func (e Error) Error() string {
	bs, _ := json.Marshal(e)
	return string(bs)
}

func NewBadRequestError(errCode int, message string) Error {
	return Error{
		StatusCode: http.StatusBadRequest,
		ErrCode:    errCode,
		Message:    message,
	}
}

func ParseError(bytes []byte) Error {
	err := Error{}
	if err := json.Unmarshal(bytes, &err); err != nil {
		return ErrServiceUnavailable
	}
	return err
}

const (
	WX_PAY_ERROR        = 10205
	ALI_PAY_ERROR       = 10206
	PAY_FAILED          = 10207
	SERVICE_UNAVAILABLE = 10002
)

var (
	ErrWXPayError = Error{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    WX_PAY_ERROR,
		Message:    "wxpay error",
	}
	ErrAlipayError = Error{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ALI_PAY_ERROR,
		Message:    "alipay error",
	}
	ErrAlipaySignCheckError = Error{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ALI_PAY_ERROR,
		Message:    "alipay check sign error",
	}
	ErrAlipayCloseOrderError = Error{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ALI_PAY_ERROR,
		Message:    "alipay close order error",
	}
	ErrPayFailed = Error{
		StatusCode: http.StatusBadRequest,
		ErrCode:    PAY_FAILED,
		Message:    "pay failed",
	}
	ErrServiceUnavailable = Error{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SERVICE_UNAVAILABLE,
		Message:    "service unavailable temporarily",
	}
)
