package types

import "github.com/ipron-ne/client-sdk-go/code"

type AgentStateAndMediaD struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	State      string `json:"state"`
	Cause      string `json:"cause"`
	CauseName  string `json:"causeName"`
	DN         string `json:"dn"`
	ChangeTime string `json:"chgtime"`
}

type QueueByUserD struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Extension string   `json:"extension"`
	Skills    []string `json:"skills"`
}

type PresenceResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"msg"`
	Code    int    `json:"code"`
}

type UserLoginRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
	DN       string                   `json:"dn"`
}

type UserLoginResponse struct {
	PresenceResponse
}

type UserLogoutRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type UserLogoutResponse struct {
	PresenceResponse
}

type SetUserStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserStateResponse struct {
	PresenceResponse
}

type GetUserStateRequest struct {
}

type GetUserStateResponse struct {
	PresenceResponse
	StateSet []string `json:"stateset"`
	CauseSet []string `json:"causeset"`
}

type GetUserStateMultiMediaRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
}

type GetUserStateMultiMediaResponse struct {
	PresenceResponse
	StateSet []string `json:"stateset"`
	CauseSet []string `json:"causeset"`
}

type RouteableRequest struct {
}

type RouteableResponse struct {
	PresenceResponse
	Enable bool `json:"enable"`
}

type SetUserAfterStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserAfterStateResponse struct {
	PresenceResponse
}

type SetUserRecallStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserRecallStateResponse struct {
	PresenceResponse
}

type GetQueueStatusRequest struct {
}

type GetQueueStatusResponse struct {
	PresenceResponse
	AllUserCnt   int `json:"all_user_cnt"`
	LoginUserCnt int `json:"login_user_cnt"`
	ReadyUserCnt int `json:"ready_user_cnt"`
	BusyUserCnt  int `json:"busy_user_cnt"`
	AfterUserCnt int `json:"after_user_cnt"`
	CurrCallCnt  int `json:"curr_call_cnt"`
}

type GetSkillStatusRequest struct {
}

type GetSkillStatusResponse struct {
	PresenceResponse
	AllUserCnt   int `json:"all_user_cnt"`
	LoginUserCnt int `json:"login_user_cnt"`
	ReadyUserCnt int `json:"ready_user_cnt"`
	BusyUserCnt  int `json:"busy_user_cnt"`
	AfterUserCnt int `json:"after_user_cnt"`
	CurrCallCnt  int `json:"curr_call_cnt"`
}

type GetUsersByStateAndMediaRequest struct {
}

type GetUsersByStateAndMediaResponse struct {
	PresenceResponse
	Data []AgentStateAndMediaD `json:"data"`
}

type GetQueuesByUserIdRequest struct {
}

type GetQueuesByUserIdResponse struct {
	PresenceResponse
	Data []QueueByUserD `json:"data"`
}
