// Supported IPRON-NE v1.2
package omnigw

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/omnigw"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Omni struct {
	types.Client
}

func NewFromClient(client types.Client) *Omni {
	return &Omni{
		Client: client,
	}
}

func (c *Omni) Route(providerID, ani, dnis, uei string) (RouteResponse, error) {
	var respData RouteResponse

	url := fmt.Sprintf("%s/route", API_NAME)
	body := RouteRequest{
		ProviderID: providerID,
		ANI:        ani,
		DNIS:       dnis,
		UEI:        uei,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Route")
}

func (c *Omni) Connect(providerID, callID string) (ConnectResponse, error) {
	var respData ConnectResponse

	url := fmt.Sprintf("%s/connect", API_NAME)
	body := ConnectRequest{
		ProviderID: providerID,
		CallID:     callID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Connect")
}

func (c *Omni) Release(providerID, callID string) (ReleaseResponse, error) {
	var respData ReleaseResponse

	url := fmt.Sprintf("%s/release", API_NAME)
	body := ReleaseRequest{
		ProviderID: providerID,
		CallID:     callID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Release")
}
