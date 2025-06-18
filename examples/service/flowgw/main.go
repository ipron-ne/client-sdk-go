// go run . -tenantid 67ca8c9c308a37387be9da78 -provid 684fb9d939052e60f063838f
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/flowgw"
	"github.com/ipron-ne/client-sdk-go/service/notify"
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
	flowGw := flowgw.NewFromClient(client)
	eventNotify := notify.NewFromClient(client)

	err := eventNotify.AddSubscriptions(*TenantID, "provider/"+*ProviderID, handlerEvent, handlerError, "provider")
	if err != nil {
		log.Panic(err)
	}

	// 67d3d8ad3242cc5bded53d4c, acs_test0
	resp, err := flowGw.FlowStart(*ProviderID, *TenantID, "67d3d8ad3242cc5bded53d4c", "Default", "0234304114", "901073220804", "")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("CALLID=%s\n", resp.CallID)

	time.Sleep(100 * time.Second)
}

func handlerEvent(e types.Data) {
	eventName := e.Get("data").Get("event").Str()
	switch eventName {
	case flowgw.EventStart:
		var data flowgw.FlowGWStart
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case flowgw.EventDial:
		var data flowgw.FlowGWDial
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case flowgw.EventEnd:
		var data flowgw.FlowGWEnd
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	default:
		log.Printf("unknown event: %s\n", eventName)
	}
}

func handlerError(err error) {
	log.Println(err)
}
