package main

import (
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/service/info"
	"github.com/ipron-ne/client-sdk-go/types"
)

var (
	API_URL    = os.Getenv("IPRON_NE_URL")
	UserID     = os.Getenv("USER_ID")
	Passwd     = os.Getenv("USER_PWD")
	TenantName = os.Getenv("TENANT_NAME")
	DN         = "4400"
)

func main() {
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
	err := auth.Login(UserID, Passwd, TenantName, []code.MediaType{code.Media.Voice}, 
		code.AgentStatus.NotReady, 
		code.AgentStateCause.NotReady.Idle, 
		DN, 
		handlerEvent, handlerError,
	)
	if err != nil {
		panic(err)
	}

	// Client 인스턴스로 정보조회 서비스 생성
	info := info.NewFromClient(client)

	// Group List 조회
	resp, err := info.GetGroupList(client.GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}

	// Flow List 조회
	resp, err = info.GetFlowList(client.GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}

	time.Sleep(10 * time.Second)
}

func handlerEvent(e types.Data) {
	log.Printf("***** %+v\n", e)
}

func handlerError(err error) {
	log.Println(err)
}
