package flowgw

import "github.com/ipron-ne/client-sdk-go/types"

type FlowStartRequest struct {
	ProviderID string `json:"provider_id"`
	TenantID   string `json:"tenant_id"`
	FlowID     string `json:"flow_id"`
	SiteID     string `json:"site_id"`
	ANI        string `json:"ani"`
	DNIS       string `json:"dnis"`
	Userdata   string `json:"userdata"`
}

type FlowStartResponse struct {
	types.ServiceResponse
	ProviderID string `json:"provider_id"`
	CallID     string `json:"call_id"`
}
