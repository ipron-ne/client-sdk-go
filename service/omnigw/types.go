package omnigw

import "github.com/ipron-ne/client-sdk-go/types"

type RouteRequest struct {
	ProviderID string `json:"provider_id"`
	ANI        string `json:"ani"`
	DNIS       string `json:"dnis"`
	UEI        string `json:"uei"`
}

type RouteResponse struct {
	types.ServiceResponse
	CallID string `json:"call_id"`
}

type ConnectRequest struct {
	ProviderID string `json:"provider_id"`
	CallID     string `json:"call_id"`
}

type ConnectResponse struct {
	types.ServiceResponse
}

type ReleaseRequest struct {
	ProviderID string `json:"provider_id"`
	CallID     string `json:"call_id"`
}

type ReleaseResponse struct {
	types.ServiceResponse
}
