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
	AppKey     = os.Getenv("IPRON_NE_APPKEY")
	TenantID   = os.Getenv("IPRON_NE_TENANTID")
	DN         = "4400"
)

func main() {
	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithDebug(true),
		// config.WithAppToken(AppKey),
		// config.WithTenantID(TenantID),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 인증 서비스 생성
	auth := auth.NewFromClient(client)

	// 사용자 로그인
	err := auth.Login(UserID, Passwd, TenantName, []code.MediaType{code.Media.Voice},
		code.AgentStatus.NotReady,
		code.AgentStateCauseType("00"), // code.AgentStateCause.NotReady.Idle,
		DN,
		handlerEvent, handlerError,
	)
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	err = auth.Logout(client.GetTenantID(), auth.GetUserID(), []code.MediaType{code.Media.Voice}, code.AgentStateCauseType("00"))
	if err != nil {
		panic(err)
	}
}

func main2() {
	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithDebug(true),
		config.WithAppToken(AppKey),
		config.WithTenantID(TenantID),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 정보조회 서비스 생성
	info := info.NewFromClient(client)

	// Group List 조회
	var groupId string
	log.Println("[GROUP LIST]")
	{
		resp, err := info.GetGroupList(client.GetTenantID())
		if err != nil {
			log.Println(err)
		}
		for i, v := range resp {
			log.Printf("[%02d] %+v\n", i, v)
			groupId = v.ID
		}
	}
	log.Println("")

	// Group List 조회
	log.Println("[GROUP INFO]")
	{
		resp, err := info.GetGroupInfo(client.GetTenantID(), groupId)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[%02d] %+v\n", 1, resp)
	}
	log.Println("")

	// Group List 조회
	log.Println("[AGENT ALL LIST]")
	{
		resp, err := info.GetAllAgentList(client.GetTenantID())
		if err != nil {
			log.Println(err)
		}
		for i, v := range resp {
			log.Printf("[%02d] %+v\n", i, v)
		}
	}
	log.Println("")

	// Group List 조회
	var agentId string
	log.Println("[GROUP AGENT LIST]")
	{
		resp, err := info.GetAgentList(client.GetTenantID(), groupId)
		if err != nil {
			log.Println(err)
		}
		for i, v := range resp {
			log.Printf("[%02d] %+v\n", i, v)
			agentId = v.ID
		}
	}
	log.Println("")

	// Group List 조회
	log.Println("[AGENT INFO]")
	{
		resp, err := info.GetAgentInfo(client.GetTenantID(), agentId)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[%02d] %+v\n", 1, resp)
	}
	log.Println("")

	// Group List 조회
	var queueId string
	log.Println("[QUEUE LIST]")
	{
		resp, err := info.GetQueueList(client.GetTenantID())
		if err != nil {
			log.Println(err)
		}
		for i, v := range resp {
			log.Printf("[%02d] %+v\n", i, v)
			queueId = v.ID
		}
	}
	log.Println("")

	// Group List 조회
	log.Println("[QUEUE INFO]")
	{
		resp, err := info.GetQueueInfo(client.GetTenantID(), queueId)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[%02d] %+v\n", 1, resp)
	}
	log.Println("")

	// Flow List 조회
	log.Println("[FLOW LIST]")
	{
		resp, err := info.GetFlowList(client.GetTenantID())
		if err != nil {
			log.Println(err)
		}
		for i, v := range resp {
			log.Printf("[%02d] %+v\n", i, v)
		}
	}
	log.Println("")

	time.Sleep(10 * time.Second)
}

func handlerEvent(e types.Data) {
	log.Printf("***** %+v\n", e)
}

func handlerError(err error) {
	log.Println(err)
}
