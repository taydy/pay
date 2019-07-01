package util

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
)

const (
	DefaultSize     = 512
	PngBase64Prefix = "data:image/png;base64,%s"
)

func New(body string) ([]byte, error) {
	return qrcode.Encode(body, qrcode.Medium, DefaultSize)
}

/*
 * 生成 png 图片 base64 码。
 */
func PngBase64(body string) string {
	var bytes []byte
	bytes, err := New(body)
	if err != nil {
		return ""
	}
	data := fmt.Sprintf(PngBase64Prefix, base64.StdEncoding.EncodeToString(bytes))
	return data
}

/**
 * 默认 PNG 图片。
 */
func WriteToFile(body string, fileName string) error {
	return qrcode.WriteFile(body, qrcode.Medium, DefaultSize, fmt.Sprintf("%s.png", fileName))
}
