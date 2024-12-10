package auth

import (
	"net/http"

	"github.com/ipron-ne/client-sdk-go/monitoring"
)

const (
	API_NAME = "/auth/v1"
)

func FetchToken(email, plainPassword string) {
	client := monitoring.GetApiClient()

	resp, err := client.Post(API_NAME + "/token", map[string]any{
		"email": email,
		"plainPassword": plainPassword,
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	client.SetLocalToken(resp.Data)
}

func FetchTokenByToken(accessToken string) {
	client := monitoring.GetApiClient()

	resp, err := client.Post(API_NAME + "/token/" + accessToken, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	client.SetLocalToken(resp.Data)	
}

func DeleteToken(refreshToken string) {
	client := monitoring.GetApiClient()

	resp, err := client.Post(API_NAME + "/token/delete", nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	client.DeleteLocalToken()
}

func VerifyToken(accessToken, refreshToken string) {
	client := monitoring.GetApiClient()

	resp, err := client.Post(API_NAME + "/token/verify", map[string]any{
		"accessToken": accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshToken(refreshToken string) {
	client := monitoring.GetApiClient()

	client.SetToken(refreshToken)

	resp, err := client.Post(API_NAME + "/token/refresh", nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		return // fmt.Errorf("failed to fetch token: %v", err)
	}

	client.SetLocalToken(resp.Data)
}

