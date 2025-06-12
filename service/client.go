package service

import (
	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/ipron-ne/client-sdk-go/utils"
)

const AUTH_HEADER = "Authorization"

type Response = types.Response

type Client struct {
	types.Request

	BaseURL      string
	ClientID     string
	RefreshToken string
	AccessToken  string
	isDebug      bool

	UserData map[string]any // AuthData중 AccessToken 정보에 포함된 사용자 데이터
	Log      utils.Log
}

func NewFromConfig(cfg config.Config) *Client {
	clientID := utils.CreateUUID()

	instance := &Client{
		ClientID: clientID,
		BaseURL:  cfg.BaseURL,
		UserData: make(map[string]any),
		Request: utils.NewHttpClient(cfg.BaseURL, cfg.Timeout, map[string]string{
			"X-CLIENT-ID": clientID,
		}),
		isDebug: cfg.IsDebug,
	}

	// AppToken 으로 인증 토큰 사용
	if cfg.AppToken != "" {
		instance.SetToken(cfg.AppToken)
	}

	if cfg.TenantID != "" {
		instance.SetTenant(cfg.TenantID)
	}

	return instance
}

func (c *Client) GetClientID() string {
	return c.ClientID
}

func (c *Client) GetToken() string {
	return c.AccessToken
}

func (c *Client) GetBaseURL() string {
	return c.BaseURL
}

func (c *Client) IsDebug() bool {
	return c.isDebug
}

func (c *Client) GetLogger() types.Logger {
	return &c.Log
}

func (c *Client) GetRequest() types.Request {
	return c.Request
}

func (c *Client) SetTenant(tenantID string) {
	c.UserData["tntId"] = tenantID
}

func (c *Client) SetToken(token string) {
	c.AccessToken = token
	c.SetHeader(AUTH_HEADER, "Bearer "+token)
}

func (c *Client) SetLocalToken(accessToken, refreshToken string) {
	var err error

	c.UserData, err = utils.DecodeJWT(accessToken)
	if err != nil {
		c.Log.Error("Failed to decode JWT: %s", err)
	}
	c.AccessToken = accessToken
	c.RefreshToken = refreshToken
	c.SetHeader(AUTH_HEADER, "Bearer "+accessToken)
}

func (c *Client) DeleteLocalToken() {
	c.AccessToken = ""
	c.RefreshToken = ""
	c.UserData = make(map[string]any)
	c.DelHeader(AUTH_HEADER)
}

func (c *Client) GetTenantID() string {
	return utils.GetStr(c.UserData, "tntId")
}

func (c *Client) GetUserID() string {
	return utils.GetStr(c.UserData, "_id")
}
