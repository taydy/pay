package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/taydy/pay/constant"
	"github.com/taydy/pay/struct"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// 将分转换为以元为单位，保留两位小数的金额
func CentsToYuan(cents int) string {
	yuan := cents / 100
	cent := cents % 100

	yuanStr := strconv.Itoa(yuan)
	centStr := strconv.Itoa(cent)
	if cent < 10 {
		centStr = "0" + strconv.Itoa(cent)
	}

	amount := yuanStr + "." + centStr
	return amount
}

// 将以元为单位，保留两位小数的金额，转换成以分为单位的数值
func YuanToCents(amount string) (int, bool) {
	parts := strings.Split(amount, ".")
	if len(parts) > 2 {
		return 0, false
	}

	yuan, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, false
	}
	total := yuan * 100

	if len(parts) == 2 {
		cents, err := strconv.Atoi(parts[1])
		if err != nil || cents < 0 || cents >= 100 {
			return 0, false
		}
		total += cents
	}
	return total, true
}

// 阿里支付参数加签
func AliSign(params map[string]interface{}, privateKey *rsa.PrivateKey) string {
	var keys []string
	var origin []string

	for k, v := range params {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		origin = append(origin, fmt.Sprintf("%s=%v", k, params[k]))
	}

	str := strings.Join(origin, "&")
	sign, err := rsaSign(str, privateKey)
	if err != nil {
		return ""
	}
	return sign
}

func rsaSign(originalData string, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(originalData))
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

func verifySign(originalData string, signData string, publicKey *rsa.PublicKey) bool {
	fmt.Println(originalData)
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return false
	}
	hash := sha256.New()
	hash.Write([]byte(originalData))
	hashed := hash.Sum(nil)

	verifyError := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, sign)
	if verifyError != nil {
		return false
	}
	return true
}

// 校验 ali 支付签名
func ValidAliSign(data string, sign string, publicKey *rsa.PublicKey) bool {
	return verifySign(data, sign, publicKey)
}

func TransferSuccess(result *_struct.AliTransferResult) bool {
	return result.AlipayFundTransToaccountTransferResponse.Code == "10000" && result.AlipayFundTransToaccountTransferResponse.PayDate != ""
}

func RefundSuccess(result *_struct.AliTradeRefundResult) bool {
	return result.AlipayTradeRefundResponse.Code == "10000" && strings.ToUpper(result.AlipayTradeRefundResponse.Msg) == constant.SUCCESS
}

// 加载支付宝私钥
func LoadAliPrivateKey(path string) *rsa.PrivateKey {
	PrivateKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error)
	}
	blockPrivate, _ := pem.Decode(PrivateKeyBytes)
	if blockPrivate == nil {
		panic("No RSA Private Key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(blockPrivate.Bytes)
	if err != nil {
		panic(err.Error)
	}
	return privateKey
}

// 加载支付宝公钥
func LoadAliPublicKey(path string) *rsa.PublicKey {
	aliPublicKeyByte, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error)
	}
	aliPublicKey, _ := base64.StdEncoding.DecodeString(string(aliPublicKeyByte))
	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}
	return pubInterface.(*rsa.PublicKey)
}
