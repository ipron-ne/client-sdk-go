package main

import (
	"fmt"
	"time"

	"github.com/donovanhide/eventsource"

	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/code"
)

func main() {
	service.Init("http://100.100.103.160/", 0, true)
	err := auth.Login("chiwoo@bridgetec.co.kr", "1andromeda", "연구1팀", []code.MediaType{code.Media.Voice}, 
		code.AgentStatus.NotReady, 
		code.AgentStateCause.NotReady.Idle, 
		"4400", 
		handlerEvent, handlerError,
	)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(100 * time.Second)
}

func handlerEvent(e eventsource.Event) {
	fmt.Println(e)
}

func handlerError(err error) {
	fmt.Println(err)
}
