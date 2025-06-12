package auth

import (
	"net/http"

	"github.com/ipron-ne/client-sdk-go/types"
)

const (
	API_NAME = "/auth/v1"
)

type Auth struct {
	types.Client
}

func NewFromClient(client types.Client) *Auth {
	return &Auth{
		Client: client,
	}
}

func (c *Auth) FetchToken(email, plainPassword string) {
	resp, err := c.GetRequest().Post(API_NAME+"/token", map[string]any{
		"email":         email,
		"plainPassword": plainPassword,
	})
	if err != nil {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	var tokenResp types.CreateTokenResponse
	resp.DataUnmarshal(&tokenResp)
	if !tokenResp.LoginResult {
		return // fmt.Errorf("failed to fetch token: %v", resp.Msg)
	}

	c.SetLocalToken(tokenResp.AccessToken, tokenResp.RefreshToken)
}

func (c *Auth) FetchTokenByToken(accessToken string) {
	resp, err := c.GetRequest().Post(API_NAME+"/token/"+accessToken, nil)
	if err != nil {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	var tokenResp types.CreateTokenResponse
	resp.DataUnmarshal(&tokenResp)
	if !tokenResp.LoginResult {
		return // fmt.Errorf("failed to fetch token: %v", resp.Msg)
	}

	c.SetLocalToken(tokenResp.AccessToken, tokenResp.RefreshToken)
}

func (c *Auth) DeleteToken(refreshToken string) {
	resp, err := c.GetRequest().Post(API_NAME+"/token/delete", nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	c.DeleteLocalToken()
}

func (c *Auth) VerifyToken(accessToken, refreshToken string) {
	_, err := c.GetRequest().Post(API_NAME+"/token/verify", map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
	if err != nil {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

}

func (c *Auth) RefreshToken(refreshToken string) {
	c.SetLocalToken(refreshToken, refreshToken)

	resp, err := c.GetRequest().Post(API_NAME+"/token/refresh", nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	var tokenResp types.CreateTokenResponse
	resp.DataUnmarshal(&tokenResp)
	if !tokenResp.LoginResult {
		return // fmt.Errorf("failed to fetch token: %v", resp.Msg)
	}

	c.SetLocalToken(tokenResp.AccessToken, tokenResp.RefreshToken)
}
