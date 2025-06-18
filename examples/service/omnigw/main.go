// go run . -tenantid 67ca8c9c308a37387be9da78 -provid 684fe08a9a023d7038d09f15
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/notify"
	"github.com/ipron-ne/client-sdk-go/service/omnigw"
	"github.com/ipron-ne/client-sdk-go/types"
)

var (
	API_URL = os.Getenv("IPRON_NE_API_URL")
	AppKey  = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	TenantID := flag.String("tenantid", "", "Tenant ID")
	ProviderID := flag.String("provid", "", "Provider ID")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithDebug(true),
		config.WithBaseURL(API_URL),
		config.WithAppToken(AppKey),
		config.WithTenantID(*TenantID),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)
	omniGw := omnigw.NewFromClient(client)
	eventNotify := notify.NewFromClient(client)

	err := eventNotify.AddSubscriptions(*TenantID, "provider/"+*ProviderID, nil, handlerEvent, handlerError, "provider")
	if err != nil {
		log.Panic(err)
	}

	// 67d3d8ad3242cc5bded53d4c, acs_test0
	resp, err := omniGw.Route(*ProviderID, "901073220804", "8000", "")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("CALLID=%+v\n", resp.CallID)

	time.Sleep(100 * time.Second)
}

func handlerEvent(e types.Data) {
	var data omnigw.OmniGwMsg

	eventName := e.Get("data").Get("event").Str()
	switch eventName {
	case omnigw.EventAlerting:
		fallthrough
	case omnigw.EventConnected:
		fallthrough
	case omnigw.EventEnd:
		fallthrough
	case omnigw.EventPause:
		fallthrough
	case omnigw.EventContinue:
		fallthrough
	case omnigw.EventTransfer:
		fallthrough
	case omnigw.EventDisconnected:
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	default:
		log.Printf("unknown event: %s, %+v\n", eventName, e.Get("data"))
	}
}

func handlerError(err error) {
	log.Println(err)
}
