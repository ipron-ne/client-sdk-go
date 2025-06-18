package presence

import (
	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/types"
)

type AgentStateAndMediaD struct {
	ID         string `json:"id"`        // 상담원 ID
	Name       string `json:"name"`      // 상담원 이름
	Email      string `json:"email"`     // 상담원 이메일
	State      string `json:"state"`     // 상담원 현재 상태
	Cause      string `json:"cause"`     // 상담원 현재 상태 부가코드
	CauseName  string `json:"causeName"` // 부가코드 명
	DN         string `json:"dn"`        // 내선 번호
	ChangeTime string `json:"chgtime"`   // 상태 변경 시간
}

type QueueByUserD struct {
	ID        string   `json:"id"`        // 큐 ID
	Name      string   `json:"name"`      // 큐 이름
	Extension string   `json:"extension"` // 큐 번호
	Skills    []string `json:"skills"`    // 큐 스킬 목록
}

type CallInfoD struct {
	CallID    string `json:"callId"`    // 콜 ID
	ConnID    string `json:"connId"`    // 상담사 기준 Connection ID
	ANI       string `json:"ani"`       // ANI
	DNIS      string `json:"dnis"`      // DNIS
	State     string `json:"state"`     // 연결 상태
	MediaType string `json:"mediaType"` // 미디어 타입
}

type UserLoginRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
	DN       string                   `json:"dn"`
}

type UserLoginResponse struct {
	types.ServiceResponse
	Data struct {
		LoginResult  bool   `json:"loginResult"`
		LoginStatus  string `json:"loginStatus"`
		TenantID     string `json:"tenantId"`
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
}

type UserLogoutRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type UserLogoutResponse struct {
	types.ServiceResponse
}

type SetUserStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserStateResponse struct {
	types.ServiceResponse
}

type GetUserStateRequest struct {
}

type GetUserStateResponse struct {
	types.ServiceResponse
	StateSet []string `json:"stateset"`
	CauseSet []string `json:"causeset"`
}

type GetUserStateMultiMediaRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
}

type GetUserStateMultiMediaResponse struct {
	types.ServiceResponse
	StateSet []string `json:"stateset"`
	CauseSet []string `json:"causeset"`
}

type RouteableRequest struct {
}

type RouteableResponse struct {
	types.ServiceResponse
	Enable bool `json:"enable"`
}

type SetUserAfterStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserAfterStateResponse struct {
	types.ServiceResponse
}

type SetUserRecallStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type SetUserRecallStateResponse struct {
	types.ServiceResponse
}

type GetQueueStatusRequest struct {
}

type GetQueueStatusResponse struct {
	types.ServiceResponse
	AllUserCnt   int `json:"all_user_cnt"`   // 전체 상담원 수
	LoginUserCnt int `json:"login_user_cnt"` // 로그인 상담원 수
	ReadyUserCnt int `json:"ready_user_cnt"` // 수신대기 상담원 수
	BusyUserCnt  int `json:"busy_user_cnt"`  // 통화 중 상담원 수
	AfterUserCnt int `json:"after_user_cnt"` // 후처리 상담원 수
	CurrCallCnt  int `json:"curr_call_cnt"`  // 현재 대기 콜 수
}

type GetSkillStatusRequest struct {
}

type GetSkillStatusResponse struct {
	types.ServiceResponse
	AllUserCnt   int `json:"all_user_cnt"`   // 전체 상담원 수
	LoginUserCnt int `json:"login_user_cnt"` // 로그인 상담원 수
	ReadyUserCnt int `json:"ready_user_cnt"` // 수신대기 상담원 수
	BusyUserCnt  int `json:"busy_user_cnt"`  // 통화 중 상담원 수
	AfterUserCnt int `json:"after_user_cnt"` // 후처리 상담원 수
	CurrCallCnt  int `json:"curr_call_cnt"`  // 현재 대기 콜 수
}

type GetUsersByStateAndMediaRequest struct {
}

type GetUsersByStateAndMediaResponse struct {
	types.ServiceResponse
	Data []AgentStateAndMediaD `json:"data"`
}

type GetQueuesByUserIdRequest struct {
}

type GetQueuesByUserIdResponse struct {
	types.ServiceResponse
	Data []QueueByUserD `json:"data"`
}

type ForceUpdateStateRequest struct {
	Mediaset []code.MediaType         `json:"mediaset"`
	State    code.AgentStateType      `json:"state"`
	Cause    code.AgentStateCauseType `json:"cause"`
}

type ForceUpdateStateResponse struct {
	types.ServiceResponse
}

type GetCallInfoRequest struct {
}

type GetCallInfoResponse struct {
	types.ServiceResponse
	Data []CallInfoD `json:"data"` // 콜 정보 리스트
}

type GetDoNotDistubeRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
}

type GetDoNotDistubeResponse struct {
	types.ServiceResponse
	Data map[string]bool `json:"data"`
}

type SetDoNotDistubeRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
	Enable   bool             `json:"enable"`
}

type SetDoNotDistubeResponse struct {
	types.ServiceResponse
}

type GetMediaReadyRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
}

type GetMediaReadyResponse struct {
	types.ServiceResponse
	Data map[string]bool `json:"data"`
}

type SetMediaReadyRequest struct {
	Mediaset []code.MediaType `json:"mediaset"`
	Enable   bool             `json:"enable"`
}

type SetMediaReadyResponse struct {
	types.ServiceResponse
}

type SetReserveAniUEIRequest struct {
	UEI string `json:"uei"`
}

type SetReserveAniUEIResponse struct {
	types.ServiceResponse
}
