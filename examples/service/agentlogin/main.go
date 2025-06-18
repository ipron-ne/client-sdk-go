package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/service/presence"
	"github.com/ipron-ne/client-sdk-go/types"
)

var (
	API_URL = os.Getenv("IPRON_NE_API_URL")
)

func main() {
	TenantName := flag.String("tenantname", "", "Tenant Name")
	UserID := flag.String("userid", "", "User ID")
	Passwd := flag.String("passwd", "", "User Password")
	DN := flag.String("dn", "", "User Phone DN")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithDebug(true),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 인증 서비스 생성
	auth := auth.NewFromClient(client)

	// 사용자 로그인
	err := auth.Login(*UserID, *Passwd, *TenantName, []code.MediaType{code.Media.Voice},
		code.AgentStatus.NotReady,
		code.AgentStateCauseType("00"), // code.AgentStateCause.NotReady.Idle,
		*DN,
		handlerEvent, handlerError,
	)
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	err = auth.Logout(client.GetTenantID(), auth.GetUserID(), []code.MediaType{code.Media.Voice}, code.AgentStateCauseType("00"))
	if err != nil {
		panic(err)
	}
}

func handlerEvent(e types.Data) {
	eventName := e.Get("data").Get("event").Str()
	switch eventName {
	case presence.EventUserBanishment:
		var data presence.Banishment
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case presence.EventUserStateChanged:
		var data presence.UserStateChanged
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	case presence.EventUserReasonChanged:
		var data presence.UserReasonChanged
		e.Get("data").Unmarshal(&data)
		log.Printf("***** %+v\n", data)
	default:
		log.Printf("***** %+v\n", e)
	}
}

func handlerError(err error) {
	log.Println(err)
}
