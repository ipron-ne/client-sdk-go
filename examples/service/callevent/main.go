package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/call"
	"github.com/ipron-ne/client-sdk-go/service/notify"
	"github.com/ipron-ne/client-sdk-go/types"
)

var (
	API_URL = os.Getenv("IPRON_NE_API_URL")
	AppKey  = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	phoneID := flag.String("phoneid", "", "Phone ID")
	userID := flag.String("userid", "", "User ID")
	TenantID := flag.String("tenantid", "", "Tenant ID")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithDebug(false),
		config.WithBaseURL(API_URL),
		config.WithAppToken(AppKey),
		config.WithTenantID(*TenantID),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 이벤트 수신모듈 생성
	eventNotify := notify.NewFromClient(client)

	// PhoneID 기반 이벤트 수신 게시
	if phoneID != nil && *phoneID != "" {
		err := eventNotify.AddSubscriptions(*TenantID, "phone/"+*phoneID, nil, handlerEvent, handlerError, "provider")
		if err != nil {
			log.Panic(err)
		}
	}

	// UserID 기반 이벤트 수신 게시
	if userID != nil && *userID != "" {
		err := eventNotify.AddSubscriptions(*TenantID, "user/"+*userID, nil, handlerEvent, handlerError, "provider")
		if err != nil {
			log.Panic(err)
		}
	}

	time.Sleep(100 * time.Second)
}

func handlerEvent(e types.Data) {
	eventName := e.Get("data").Get("event").Str()

	switch eventName {
	case call.EventOriginated:
		var data call.Originated
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case call.EventAlerting:
		var data call.Alerting
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case call.EventConnected:
		var data call.Connected
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case call.EventDisconnected:
		var data call.Disconnected
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case call.EventTerminated:
		var data call.Terminated
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	default:
		// log.Printf("***** %+v\n", e)
	}
}

func handlerError(err error) {
	log.Println(err)
}
