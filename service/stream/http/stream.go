// Supported IPRON-NE v1.2
package stream

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/stream"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Stream struct {
	types.Client
}

func NewFromClient(client types.Client) *Stream {
	return &Stream{
		Client: client,
	}
}

func (c *Stream) StreamServiceAlloc(callID, connID string) (StreamServiceAllocResponse, error) {
	var respData StreamServiceAllocResponse

	url := fmt.Sprintf("%s/stream/service/alloc", API_NAME)
	body := StreamServiceAllocRequest{
		CallID: callID,
		ConnID: connID,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "StreamServiceAlloc")
}

func (c *Stream) StreamAlloc(streamURI, callID, connID, streamType string) (StreamAllocResponse, error) {
	var respData StreamAllocResponse

	url := fmt.Sprintf("%s/stream/alloc", API_NAME)
	body := StreamAllocRequest{
		CallID:     callID,
		ConnID:     connID,
		StreamType: streamType,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "StreamAlloc")
}

func (c *Stream) Stream(streamURI, streamChannel string, eventFn func(data *StreamResponse) bool) error {
	var err error

	origin := ""
	url := fmt.Sprintf("%s/stream/ws/%s", API_NAME, streamChannel)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		return errors.Wrap(err, "Stream")
	}
	defer ws.Close()

	for {
		var data StreamResponse

		/*
			if n, err = ws.Read(msg); err != nil {
				break
			}

			if err := json.Unmarshal(msg[:n], &data); err != nil {
				continue
			}
		*/

		websocket.JSON.Receive(ws, &data)

		if !eventFn(&data) {
			break
		}
	}

	return nil
}
