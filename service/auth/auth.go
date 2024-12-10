package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/donovanhide/eventsource"
	"github.com/pkg/errors"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/notify"
	"github.com/ipron-ne/client-sdk-go/service/presence"
	"github.com/ipron-ne/client-sdk-go/utils"
)

const (
	API_PREFIX    = "/webapi/sdk-auth"
	API_VERSION   = "/v1/management"
	API_NAME      = API_PREFIX + API_VERSION
)

func Login(email, plainPassword, tntName string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string, eventCallback any, eventErrorCallback func(error)) error {
	client := service.GetApiClient()

	// client.Lock()
	// defer client.Unlock()

	// 기존 로그인 요청이 진행중인지 확인
	if client.IsLoginInProgress {
		return errors.New("login request is already in progress")
	}

	client.IsLoginInProgress = true
	defer func() { client.IsLoginInProgress = false }()

	// 기존 클라이언트 유저 데이터가 존재하는 경우, sse 이벤트를 사용중인지 확인
	currentUserId := utils.GetStr(client.UserData, "_id")
	if currentUserId != "" && client.EventMap[fmt.Sprintf("user/%s", currentUserId)] != nil {
		return fmt.Errorf("user %s is already logged in", email)
	}

	// 주기가 짧은 토큰 발급 요청
	body := map[string]any{
		"email":         email,
		"plainPassword": plainPassword,
		"tntName":       tntName,
	}
	resp, err := client.Post(fmt.Sprintf("%s/token", API_NAME), body) // Simplified
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch token: %v", err)
	}

	if !resp.GetData().Get("loginResult").Bool() {
		return fmt.Errorf("failed to fetch token: %v", resp.Msg)
	}

	// 짧은 토큰으로 헤더 설정
	client.SetLocalToken(resp.GetData())

	// sse event subscribe을 위한 파라미터 설정
	tntId := utils.GetStr(client.UserData, "tntId")
	userId := utils.GetStr(client.UserData, "_id")
	topic := fmt.Sprintf("user/%s", userId)
	// sse event subscribe
	err = notify.AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, "cti-client")
	if err != nil {
		fmt.Println(err)
	}

	/**
	 * 
	 **/

	// 이벤트 핸들러
	eventRegisteredCallback := func(e eventsource.Event) {
		// CTI login 시도
		resp, err := presence.UserLogin(tntId, userId, mediaSet, state, cause, dn)
		if err != nil {
			client.Log.Error("Failed to login user: %s", err)
			// 실패시 eventSource close
			notify.DelSubscriptions(topic);
			client.DeleteLocalToken();
			return
		}
		// 성공시 토큰 재설정
		client.SetLocalToken(resp.GetData())
	}

	client.EventMap[topic].AddEventListener(string(code.Event.Handler.Registered), eventRegisteredCallback)
	client.EventMap[topic].OnError(func(err error) {
		log.Printf("Catch error eventsource [%s]: %v\n", topic, err)
		Logout(tntId, userId, mediaSet, cause)
		if eventErrorCallback != nil {
			eventErrorCallback(err)
		}
	})

	return nil
}

func Logout(tntId, userId string, mediaSet []code.MediaType, cause code.AgentStateCauseType) error {
	client := service.GetApiClient()
	topic := fmt.Sprintf("user/%s", userId)

	notify.DelSubscriptions(topic);

	_, err := presence.UserLogout(tntId, userId, mediaSet, cause)
	if err != nil {
		log.Printf("Failed to logout user: %s", err)
		return errors.Wrap(err, "failed to logout user")
	}

	delete(client.EventMap, topic)
	client.DeleteLocalToken()

	log.Printf("Logged out user: %s", userId)
	return nil
}
