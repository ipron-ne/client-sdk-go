package auth

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/service/notify"
	"github.com/ipron-ne/client-sdk-go/service/presence"
	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/ipron-ne/client-sdk-go/utils"
)

const (
	API_PREFIX  = "/webapi/sdk-auth"
	API_VERSION = "/v1/management"
	API_NAME    = API_PREFIX + API_VERSION
)

type Auth struct {
	types.Client
	*notify.Notify

	InProgress bool
}

func NewFromClient(client types.Client) *Auth {
	return &Auth{
		Client:     client,
		Notify:     notify.NewFromClient(client),
		InProgress: false,
	}
}

func (s *Auth) Login(email, plainPassword, tntName string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string, eventCallback any, eventErrorCallback func(error)) error {
	// 기존 로그인 요청이 진행중인지 확인
	if s.InProgress {
		return errors.New("login request is already in progress")
	}

	s.InProgress = true
	defer func() { s.InProgress = false }()

	// 기존 클라이언트 유저 데이터가 존재하는 경우, sse 이벤트를 사용중인지 확인
	currentUserId := s.GetUserID()
	if currentUserId != "" && s.GetSubscriptions(fmt.Sprintf("user/%s", currentUserId)) != nil {
		return fmt.Errorf("user %s is already logged in", email)
	}

	// 주기가 짧은 토큰 발급 요청
	resp, err := s.GetRequest().Post(fmt.Sprintf("%s/token", API_NAME),
		types.CreateTokenRequest{
			Email:         email,
			PlainPassword: plainPassword,
			TntName:       tntName,
		})
	if err != nil {
		return errors.Wrap(err, "failed to fetch token")
	}

	var tokenResp types.CreateTokenResponse
	resp.DataUnmarshal(&tokenResp)
	if !tokenResp.LoginResult {
		return fmt.Errorf("failed to fetch token: %v", resp.Msg)
	}

	// 짧은 토큰으로 헤더 설정
	s.SetLocalToken(tokenResp.AccessToken, tokenResp.RefreshToken)

	// sse event subscribe을 위한 파라미터 설정
	tntId := s.GetTenantID()
	userId := s.GetUserID()
	topic := fmt.Sprintf("user/%s", userId)
	// sse event subscribe
	err = s.AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, "cti-client")
	if err != nil {
		fmt.Println(err)
	}

	/**
	 *
	 **/

	// 이벤트 핸들러
	eventRegisteredCallback := func(e utils.Event) {
		// CTI login 시도
		ps := presence.NewFromClient(s)
		_, err := ps.UserLogin(tntId, userId, mediaSet, state, cause, dn)
		if err != nil {
			s.GetLogger().Error("Failed to login user: %s", err)
			// 실패시 eventSource close
			s.DelSubscriptions(topic)
			s.DeleteLocalToken()
			return
		}
		// 성공시 토큰 재설정
		s.SetLocalToken(tokenResp.AccessToken, tokenResp.RefreshToken)
	}

	s.GetSubscriptions(topic).AddEventListener(string(code.Event.Handler.Registered), eventRegisteredCallback)
	s.GetSubscriptions(topic).OnError(func(err error) {
		s.GetLogger().Error("Catch error eventsource [%s]: %v\n", topic, err)
		s.Logout(tntId, userId, mediaSet, cause)
		if eventErrorCallback != nil {
			eventErrorCallback(err)
		}
	})

	return nil
}

func (s *Auth) Logout(tntId, userId string, mediaSet []code.MediaType, cause code.AgentStateCauseType) error {
	s.DelUserSubscriptions(userId)

	ps := presence.NewFromClient(s)
	_, err := ps.UserLogout(tntId, userId, mediaSet, cause)
	if err != nil {
		s.GetLogger().Error("Failed to logout user: %s", err)
		return errors.Wrap(err, "failed to logout user")
	}

	s.DeleteLocalToken()

	s.GetLogger().Debug("Logged out user: %s", userId)
	return nil
}
