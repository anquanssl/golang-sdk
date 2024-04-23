package anquanssl

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/anquanssl/golang-sdk/utils"
	"github.com/google/uuid"
)

const ORIGIN_API = "https://api.orion.pki.plus/api/v1"

func sign(baseString string, accessKeySecret string) string {
	h := hmac.New(sha256.New, []byte(accessKeySecret))
	h.Write([]byte(baseString))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

type Client struct {
	AccessKeyID     string
	AccessKeySecret string
	APIOrigin       string
}

func NewClient(accessKeyID string, accessKeySecret string, apiOrigin string) *Client {
	if apiOrigin == "" {
		apiOrigin = ORIGIN_API
	}
	return &Client{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		APIOrigin:       apiOrigin,
	}
}

func (c *Client) Get(uri string, query map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	return c.call("GET", uri, query, body)
}

func (c *Client) Post(uri string, query map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	return c.call("POST", uri, query, body)
}

func (c *Client) call(method string, uri string, query map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	query["accessKeyId"] = c.AccessKeyID
	query["nonce"] = strings.ReplaceAll(uuid.New().String(), "-", "")
	var timeLocal = time.Local.String()
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	query["timestamp"] = time.Now().In(cstSh).Format("2006-01-02T15:04:05Z")
	time.LoadLocation(timeLocal)

	parameter := make(map[string]interface{})
	for k, v := range query {
		if v != "" {
			parameter[k] = v
		}
	}
	requestBody := make(map[string]interface{})
	for k, v := range body {
		if v == nil || v == "" {
			continue
		}
		requestBody[k] = v
		parameter[k] = v
	}

	u, err := url.Parse(c.APIOrigin + uri)
	if err != nil {
		return nil, fmt.Errorf("parse uri error: %v", err)
	}

	signature := sign(u.Path+"?"+utils.HttpBuildQuery(parameter, ""), c.AccessKeySecret)

	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}
	values.Add("sign", signature)
	u.RawQuery = values.Encode()

	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshal body error: %v", err)
	}
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, fmt.Errorf("new request error: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %v", err)
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %v", err)
	}
	respBody := make(map[string]interface{})
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response body error: %v", err)
	}
	return respBody, nil
}
