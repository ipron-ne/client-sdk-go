// Supported IPRON-NE v1.2
package msggw

import (
	"fmt"
	"time"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/msggw"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type MsgGW struct {
	types.Client
}

func NewFromClient(client types.Client) *MsgGW {
	return &MsgGW{
		Client: client,
	}
}

func (c *MsgGW) Connect(providerID, tenantID, siteID, ani, dnis string, inTime time.Time, timeout int) (ConnectResponse, error) {
	var respData ConnectResponse

	url := fmt.Sprintf("%s/flow-start", API_NAME)
	body := ConnectRequest{
		ProviderID: providerID,
		TenantID:   tenantID,
		SiteID:     siteID,
		ANI:        ani,
		DNIS:       dnis,
		Time:       inTime.Format(time.RFC3339),
		Timeout:    timeout,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Connect")
}

func (c *MsgGW) Message(providerID, callID, text string) (MessageResponse, error) {
	var respData MessageResponse

	url := fmt.Sprintf("%s/flow-start", API_NAME)
	body := MessageRequest{
		ProviderID: providerID,
		CallID:     callID,
		Text:       text,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Message")
}

func (c *MsgGW) Bye(providerID, callID string) (ByeResponse, error) {
	var respData ByeResponse

	url := fmt.Sprintf("%s/flow-start", API_NAME)
	body := ByeRequest{
		ProviderID: providerID,
		CallID:     callID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Bye")
}
