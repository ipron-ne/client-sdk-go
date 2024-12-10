package monitoring

import (
	"encoding/json"
	"bytes"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/ipron-ne/client-sdk-go/utils"
	"github.com/ipron-ne/client-sdk-go/types"
)

const AUTH_HEADER   = "Authorization"

type Response = types.Response

type ApiClient struct {
	BaseURL  string
	ClientID string
	Token    string
	IsDebug  bool
	Timeout  time.Duration

	AuthData map[string]any
	UserData map[string]any
	Axios    *http.Client
	Headers  map[string]string
	mu       sync.Mutex
	Log      utils.Log
}

var instance *ApiClient
var once sync.Once

func GetApiClient() *ApiClient {
	return instance
}

func Init(baseURL string, param map[string]any, timeout time.Duration, isDebug bool) {
	if baseURL == "" {
		baseURL = os.Getenv("IPRON_NE_BASE_URL")
	}

	baseURL = strings.TrimRight(baseURL, "/")

	once.Do(func() {
		// global.EventSource = nodeEventSource;
		// paramsSerializer: params => {
        //   return qs.stringify(params, { arrayFormat: "comma" });
        // },
        // headers: {"X-CLIENT-ID": apiClientUUID },

        if timeout == 0 {
        	timeout = 30000 * time.Millisecond
        }

        clientID := utils.CreateUUID()

		instance = &ApiClient{
			BaseURL:           baseURL,
			ClientID:          clientID,
			Timeout:           timeout,
			Axios:             &http.Client{
				Timeout: timeout,
			},
			Headers:           make(map[string]string),
			IsDebug:           isDebug,
		}

		instance.SetToken(utils.GetStr(param, "token"))
	})
}

func (c *ApiClient) SetToken(token string) {
	c.Token = token
	c.Headers[AUTH_HEADER] = "Bearer " + token
}

func (c *ApiClient) DeleteLocalToken() {
	c.AuthData = make(map[string]any)
	c.UserData = make(map[string]any)
	delete(c.Headers, AUTH_HEADER)
}

func (c *ApiClient) Post(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("POST", uri, bodyJson)
}

func (c *ApiClient) Put(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("PUT", uri, bodyJson)
}

func (c *ApiClient) Get(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("GET", uri, bodyJson)
}

func (c *ApiClient) Request(method, uri string, bodyJson map[string]any) (*Response, error) {
	uri = c.BaseURL + "/" + strings.TrimLeft(uri, "/")

	resp := &Response{
		Data: make(map[string]any),
	}

	body, err := json.Marshal(bodyJson)
	if err != nil {
		return resp, err
	}

	b := bytes.NewReader(body)
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return resp, err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	resp.Response, err = c.Axios.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Response.Body.Close()

	tempResp := make(map[string]any)
	if err := json.NewDecoder(resp.Body).Decode(&tempResp); err != nil {
		return resp, errors.Wrap(err, "failed to decode token response")
	}

	resp.SetResult(tempResp)

	return resp, err
}
