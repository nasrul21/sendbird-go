package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/nasrul21/sendbird-go/errors"
)

type Client interface {
	Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error
}

type ClientImpl struct {
	HttpClient *http.Client
	BaseURL    string
	Headers    map[string]string
}

func New(client *http.Client, baseURL string, headers map[string]string) Client {
	return &ClientImpl{
		HttpClient: client,
		BaseURL:    baseURL,
		Headers:    headers,
	}
}

func (c *ClientImpl) Call(ctx context.Context, method string, url string, header http.Header, body interface{}, result interface{}) *errors.Error {
	reqBody := []byte{}
	var req *http.Request
	var err error

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return errors.FromGoErr(err)
		}
	}

	urlWithPath := fmt.Sprintf("%s%s", c.BaseURL, url)

	req, err = http.NewRequestWithContext(
		ctx,
		method,
		urlWithPath,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return errors.FromGoErr(err)
	}

	if header != nil {
		req.Header = header
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	return c.doRequest(req, result)
}

func (c *ClientImpl) doRequest(req *http.Request, result interface{}) *errors.Error {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return errors.FromGoErr(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.FromGoErr(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.FromHTTPErr(resp.StatusCode, respBody)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return errors.FromGoErr(err)
	}

	return nil
}
