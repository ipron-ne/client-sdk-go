package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	stream "github.com/ipron-ne/client-sdk-go/service/stream/grpc"
)

var (
	API_URL     = os.Getenv("IPRON_NE_API_URL")
	GRPCAPI_URL = os.Getenv("IPRON_NE_GRPCAPI_URL")
	AppKey      = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	TenantID := flag.String("tenantid", "", "Tenant ID")
	callID := flag.String("callid", "", "Call ID")
	connID := flag.String("connid", "", "Conn ID")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	if callID == nil || *callID == "" {
		flag.Usage()
		return
	}
	if connID == nil || *connID == "" {
		flag.Usage()
		return
	}

	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithDebug(true),
		config.WithBaseURL(API_URL),
		config.WithGRPCURL(GRPCAPI_URL),
		config.WithAppToken(AppKey),
		config.WithTenantID(*TenantID),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 정보조회 서비스 생성
	mediaStream := stream.NewFromClient(client)

	streamSubs, err := mediaStream.StreamServiceAlloc(*callID, *connID)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%+v\n", streamSubs)

	err = mediaStream.StreamAlloc(streamSubs, "rx", func(data *stream.StreamResponse) bool {
		if data.GetStatusCode() != stream.OK {
			return false
		}
		log.Printf("[%03d] %+v\n", data.GetStreamDataLen(), data.GetStreamData())
		return true
	})
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(10 * time.Second)
}
