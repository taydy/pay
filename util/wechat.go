package util

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

// 生成 sign
func WeChatSign(params map[string]interface{}, key string) string {
	var keys []string
	var sorted []string

	for k, v := range params {
		if k != "sign" && v != "" && v != 0 {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%v", k, params[k]))
	}

	str := strings.Join(sorted, "&")
	str += "&key=" + key
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

// 校验 sign
func WeChatValidSign(params map[string]interface{}, key string) bool {
	if _, ok := params["sign"]; !ok {
		return false
	}
	return WeChatSign(params, key) == params["sign"]
}

func FormatTime(time time.Time) string {
	return time.Format("20060102150405")
}

// 加载微信证书
func LoadWxTlsConfig(certKeyPath string, KeyPath string, caPath string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(certKeyPath, KeyPath)
	if err != nil {
		panic(err)
	}

	caData, err := ioutil.ReadFile(caPath)
	if err != nil {
		panic(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)
	return &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            pool,
	}
}
