// Supported IPRON-NE v1.2
package record

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/recgw"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Record struct {
	types.Client
}

func NewFromClient(client types.Client) *Record {
	return &Record{
		Client: client,
	}
}

func (c *Record) RecordStart(callID, connID, dn string) (RecordStartResponse, error) {
	var respData RecordStartResponse

	url := fmt.Sprintf("%s/recordstart/%s/%s/%s", API_NAME, callID, connID, dn)
	body := RecordStartRequest{}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "FlowStart")
}

func (c *Record) RecordStop(callID, connID, reason string) (RecordStartResponse, error) {
	var respData RecordStartResponse

	url := fmt.Sprintf("%s/recordstop/%s/%s/%s", API_NAME, callID, connID, reason)
	body := RecordStartRequest{}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "FlowStart")
}
