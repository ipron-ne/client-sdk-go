package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/info"
)

var (
	API_URL = os.Getenv("IPRON_NE_API_URL")
	AppKey  = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	TenantID := flag.String("tenantid", "", "Tenant ID")

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

	time.Sleep(3 * time.Second)
}
