package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/taydy/pay/struct"
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

func ValidAliSign(data string, sign string, publicKey *rsa.PublicKey) bool {
	return verifySign(data, sign, publicKey)
}

func TransferSuccess(result *_struct.AliTransferResult) bool {
	return result.AlipayFundTransToaccountTransferResponse.Code == "10000" && result.AlipayFundTransToaccountTransferResponse.PayDate != ""
}

