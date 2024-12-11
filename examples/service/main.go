package main

import (
	"log"
	"os"
	"time"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/service/info"
	"github.com/ipron-ne/client-sdk-go/utils"
)

var (
	API_URL    = os.Getenv("IPRON_NE_URL")
	UserID     = os.Getenv("USER_ID")
	Passwd     = os.Getenv("USER_PWD")
	TenantName = os.Getenv("TENANT_NAME")
	DN         = "4400"
)

func main() {
	serviceInit()

	resp, err := info.GetGroupList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}

	resp, err = info.GetFlowList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}

	time.Sleep(10 * time.Second)
}

func handlerEvent(e utils.Event) {
	log.Println(e)
}

func handlerError(err error) {
	log.Println(err)
}

func serviceInit() {
	service.Init(API_URL, 0, true)

	err := auth.Login(UserID, Passwd, TenantName, []code.MediaType{code.Media.Voice}, 
		code.AgentStatus.NotReady, 
		code.AgentStateCause.NotReady.Idle, 
		DN, 
		handlerEvent, handlerError,
	)
	if err != nil {
		panic(err)
	}
}
