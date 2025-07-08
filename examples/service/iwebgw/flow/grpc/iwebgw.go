package grpc

import (
	"context"
	iwebpb "main/grpc/proto"
	"os"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Config struct {
	token string
	uri   string
}

func (c *Config) SetToken(token string) {
	c.token = token
}

func (c *Config) SetGrpcURI(uri string) {
	c.uri = uri
}

func (c Config) GetToken() string {
	return c.token
}

func (c Config) GetGrpcURI() string {
	return c.uri
}

type IWebGW struct {
	Config
	grpcConn *grpc.ClientConn
	grpcApi  iwebpb.IwebgwClient
}

func NewFromClient(config Config) *IWebGW {
	return &IWebGW{
		Config: config,
	}
}

func (c *IWebGW) connect(dns string) error {
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
	c.grpcApi = iwebpb.NewIwebgwClient(c.grpcConn)

	return nil
}

func (c *IWebGW) Regist(tenantID, callID, userANI string) (*iwebpb.IWebGWCallRegistRes, error) {
	md := metadata.New(map[string]string{
		"routeName": "grpc-iwebgw",
		"apiKey":    c.GetToken(),
	})

	if err := c.connect(c.GetGrpcURI()); err != nil {
		return nil, errors.Wrap(err, "Regist")
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn()

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	resp, err := c.grpcApi.IWebGWCallRegist(ctx, &iwebpb.IWebGWCallRegistReq{
		TenantId: tenantID,
		CallId:   callID,
		UserAni:  userANI,
	})

	return resp, errors.Wrap(err, "Regist")
}

func (c *IWebGW) UnRegist(tenantID, callID string) error {
	return nil
}

func (c *IWebGW) RequestPage(tenantID, callID, transID string, pageFileName, pageData string, bargin bool) (*iwebpb.IWebGWRequestPageRes, error) {
	md := metadata.New(map[string]string{
		"routeName": "grpc-iwebgw",
		"apiKey":    c.GetToken(),
	})

	if err := c.connect(c.GetGrpcURI()); err != nil {
		return nil, errors.Wrap(err, "RequestPage")
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn()

	ctx = metadata.NewOutgoingContext(context.Background(), md)
	resp, err := c.grpcApi.IWebGWRequestPage(ctx, &iwebpb.IWebGWRequestPageReq{
		TenantId:      tenantID,
		CallId:        callID,
		TransactionId: transID,
		PageFilename:  pageFileName,
		PageData:      pageData,
		Bargin:        bargin,
	})

	return resp, errors.Wrap(err, "RequestPage")
}

func (c *IWebGW) DisconnectWeb(tenantID, callID, transID string) error {
	return nil
}
