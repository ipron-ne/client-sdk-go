// go run . -tenantid 67ca8c9c308a37387be9da78 -provid 684fb9d939052e60f063838f
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/msggw"
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
	SiteID := flag.String("siteid", "", "Site ID")

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
	msgGw := msggw.NewFromClient(client)
	eventNotify := notify.NewFromClient(client)

	err := eventNotify.AddSubscriptions(*TenantID, "provider/"+*ProviderID, handlerEvent, handlerError, "provider")
	if err != nil {
		log.Panic(err)
	}

	resp, err := msgGw.Connect(*ProviderID, *TenantID, *SiteID, "01012345678", "7000", time.Now(), 60)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("CALLID=%s\n", resp.CallID)

	callID := resp.CallID

	time.Sleep(2 * time.Second)

	msgGw.Message(*ProviderID, callID, "테스트 메시지입니다.")
	time.Sleep(2 * time.Second)

	msgGw.Bye(*ProviderID, callID)
	time.Sleep(1 * time.Second)
}

func handlerEvent(e types.Data) {
	eventName := e.Get("data").Get("event").Str()
	switch eventName {
	case msggw.EventMessage:
		var data msggw.Message
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case msggw.EventTransfer:
		var data msggw.TransferChat
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case msggw.EventTerminiated:
		var data msggw.TerminatedChat
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	default:
		log.Printf("unknown event: %s\n", eventName)
	}
}

func handlerError(err error) {
	log.Println(err)
}
