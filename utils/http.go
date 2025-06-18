package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ipron-ne/client-sdk-go/types"
)

type HttpClient struct {
	*http.Client
	BaseURL string
	Headers map[string]string
}

func NewHttpClient(baseURL string, timeout time.Duration, headers map[string]string) *HttpClient {
	return &HttpClient{
		BaseURL: baseURL,
		Headers: headers,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HttpClient) SetHeader(name string, value string) {
	c.Headers[name] = value
}

func (c *HttpClient) DelHeader(name string) {
	delete(c.Headers, name)
}

func (c *HttpClient) Post(uri string, bodyJson any) (*types.Response, error) {
	return c.Request("POST", uri, bodyJson)
}

func (c *HttpClient) Put(uri string, bodyJson any) (*types.Response, error) {
	return c.Request("PUT", uri, bodyJson)
}

func (c *HttpClient) Get(uri string, bodyJson any) (*types.Response, error) {
	return c.Request("GET", uri, bodyJson)
}

func (c *HttpClient) Request(method, uri string, bodyJson any) (*types.Response, error) {
	uri = c.BaseURL + "/" + strings.TrimLeft(uri, "/")

	resp := &types.Response{}

	var bodyReader io.Reader

	if bodyJson != nil {
		body, err := json.Marshal(bodyJson)
		if err != nil {
			return resp, err
		}
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, uri, bodyReader)
	if err != nil {
		return resp, err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	resp.Response, err = c.Client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Response.Body.Close()

	tempResp := make(map[string]any)

	if resp.Response.Header.Get("Content-Type") != "application/json" {
		data, _ := ioutil.ReadAll(resp.Response.Body)
		tempResp["msg"] = string(data)
	} else {
		if err := json.NewDecoder(resp.Response.Body).Decode(&tempResp); err != nil {
			return resp, errors.Wrap(err, "failed to decode token response")
		}
	}

	resp.SetResult(tempResp)

	return resp, errors.Wrap(err, "request")
}
