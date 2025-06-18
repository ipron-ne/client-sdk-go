// Supported IPRON-NE v1.2
package stream

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	streampb "github.com/ipron-ne/client-sdk-go/service/stream/grpc/proto"
)

type Stream struct {
	types.Client
	grpcConn *grpc.ClientConn
	grpcApi  streampb.StreamClient
}

type StreamSubscription struct {
	CallID       string
	ConnectionID string
	StreamID     string
	StreamAddr   string
}

func NewFromClient(client types.Client) *Stream {
	return &Stream{
		Client: client,
	}
}

func (c *Stream) connect(dns string) error {
	var opts []grpc.DialOption
	var err error

	// 이미 접속된 상태인 경우 이전 세션 사용
	if c.grpcConn != nil {
		return nil
	}

	hostName, _ := os.Hostname()

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithUserAgent(hostName))

	c.grpcConn, err = grpc.NewClient(dns, opts...)
	if err != nil {
		return errors.Wrap(err, "connect")
	}
	c.grpcApi = streampb.NewStreamClient(c.grpcConn)

	return nil
}

func (c *Stream) StreamServiceAlloc(callID, connID string) (*StreamSubscription, error) {
	md := metadata.New(map[string]string{
		"routeName": "stream",
		"apiKey":    c.GetToken(),
	})

	if err := c.connect(c.GetConfig().GetGrpcURI()); err != nil {
		return &StreamSubscription{}, errors.Wrap(err, "StreamServiceAlloc")
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn()

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	resp, err := c.grpcApi.StreamServiceAlloc(ctx, &streampb.StreamServiceAllocReq{
		CallId:       callID,
		ConnectionId: connID,
	})

	subs := StreamSubscription{
		CallID:       callID,
		ConnectionID: resp.GetConnectionId(),
		StreamID:     resp.GetPodName(),
		StreamAddr:   resp.GetPodAddr(),
	}

	return &subs, errors.Wrap(err, "StreamServiceAlloc")
}

// streamType: [rx|tx]
func (c *Stream) StreamAlloc(subs *StreamSubscription, streamType string, eventFn func(data *StreamResponse) bool) error {
	var err error

	md := metadata.New(map[string]string{
		"routeName": subs.StreamID,
		"apiKey":    c.GetToken(),
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	respStream, _ := c.grpcApi.StreamAlloc(ctx, &streampb.StreamAllocReq{
		CallId:       subs.CallID,
		ConnectionId: subs.ConnectionID,
		StreamType:   streamType,
	})

	for {
		var msg *streampb.StreamAllocResp

		msg, err = respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		data := StreamResponse{
			StatusCode:   msg.GetHeader().GetStatusCode(),
			ErrorMessage: msg.GetHeader().GetErrorMessage(),
			CallID:       msg.GetHeader().GetCallId(),
			ConnID:       msg.GetHeader().GetConnectionId(),
			StreamData:   msg.GetChunk().GetStreamData(),
		}

		if !eventFn(&data) {
			break
		}
	}

	return errors.Wrap(err, "StreamAlloc")
}
