// Supported IPRON-NE v1.2
package flowgw

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/flowgw"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type FlowGW struct {
	types.Client
}

func NewFromClient(client types.Client) *FlowGW {
	return &FlowGW{
		Client: client,
	}
}

func (c *FlowGW) FlowStart(providerID, tenantID, flowID, siteID, ani, dnis, userData string) (FlowStartResponse, error) {
	var respData FlowStartResponse

	url := fmt.Sprintf("%s/flow-start", API_NAME)
	body := FlowStartRequest{
		ProviderID: providerID,
		TenantID:   tenantID,
		FlowID:     flowID,
		SiteID:     siteID,
		ANI:        ani,
		DNIS:       dnis,
		Userdata:   userData,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "FlowStart")
}
