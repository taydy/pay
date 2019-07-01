package http

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func XmlPost(path string, body interface{}, timeout time.Duration) (*http.Response, error) {
	xmlBody, err := xml.Marshal(body)
	fmt.Printf("xml : %s \n", xmlBody)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(xmlBody)

	req, err := http.NewRequest("POST", path, reqBody)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	if timeout > 0 {
		client.Timeout = timeout * time.Second
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func XmlSecurePost(path string, body interface{}, timeout time.Duration, TLSConfig *tls.Config) (*http.Response, error) {
	xmlBody, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(xmlBody)

	req, err := http.NewRequest("POST", path, reqBody)
	if err != nil {
		return nil, err
	}

	tlsConfig := TLSConfig
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	if timeout > 0 {
		client.Timeout = timeout * time.Second
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
