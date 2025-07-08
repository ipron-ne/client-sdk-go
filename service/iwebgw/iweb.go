// Supported IPRON-NE v1.2
package iwebgw

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/iwebgw"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type IWEB struct {
	types.Client
}

func NewFromClient(client types.Client) *IWEB {
	return &IWEB{
		Client: client,
	}
}

func (c *IWEB) CreateSession(providerID, token, invokeID, connType, deviceType string) (CreateSessionResponse, error) {
	var respData CreateSessionResponse

	url := fmt.Sprintf("%s/session", API_NAME)
	body := CreateSessionRequest{
		ProviderID: providerID,
		Token:      token,
		InvokeID:   invokeID,
		ConnType:   connType,
		DeviceType: deviceType,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "CreateSession")
}

func (c *IWEB) AliveSession(providerID, sessionID, invokeID string) (AliveSessionResponse, error) {
	var respData AliveSessionResponse

	url := fmt.Sprintf("%s/session/alive", API_NAME)
	body := AliveSessionRequest{
		ProviderID: providerID,
		SessionID:  sessionID,
		InvokeID:   invokeID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "AliveSession")
}

func (c *IWEB) ConnectSession(providerID, sessionID, invokeID, connType, deviceType string) (ConnectSessionResponse, error) {
	var respData ConnectSessionResponse

	url := fmt.Sprintf("%s/session/connect", API_NAME)
	body := ConnectSessionRequest{
		ProviderID: providerID,
		SessionID:  sessionID,
		InvokeID:   invokeID,
		ConnType:   connType,
		DeviceType: deviceType,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "ConnectSession")
}

func (c *IWEB) SubmitSession(providerID, sessionID, invokeID, submitData string, bargin bool) (SubmitSessionResponse, error) {
	var respData SubmitSessionResponse

	url := fmt.Sprintf("%s/session/submit", API_NAME)
	body := SubmitSessionRequest{
		ProviderID: providerID,
		SessionID:  sessionID,
		InvokeID:   invokeID,
		SubmitData: submitData,
		Bargin:     bargin,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SubmitSession")
}

func (c *IWEB) DisconnectSession(providerID, sessioinID, invokeID string) (DisconnectSessionResponse, error) {
	var respData DisconnectSessionResponse

	url := fmt.Sprintf("%s/session/disconnect", API_NAME)
	body := DisconnectSessionRequest{
		ProviderID: providerID,
		SessionID:  sessioinID,
		InvokeID:   invokeID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "DisconnectSession")
}
