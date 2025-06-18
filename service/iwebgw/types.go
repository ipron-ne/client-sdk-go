package iwebgw

import "github.com/ipron-ne/client-sdk-go/types"

// 세션 발급
type CreateSessionRequest struct {
	ProviderID string `json:"provider_id"` // 발급된 Provider ID
	Token      string `json:"token"`       // 시나리오에서 발급받은 토큰
	InvokeID   string `json:"invoke_id"`   // API Transaction ID
	ConnType   string `json:"conn_type"`   // 세션 발급 구분 tserver인경우 토큰은 Ani번호 [callgate / tserver]
	DeviceType string `json:"device_type"` // 단말 OS 유형 [android|ios|pc|…]
}

type CreateSessionResponse struct {
	types.ServiceResponse
	Token     string `json:"token"`      // Req로 받은 토큰
	InvokeID  string `json:"invoke_id"`  // Req로 받은 API Transaction ID
	SessionID string `json:"session_id"` // 발급되는 세션 ID
}

// 세션 갱신
type AliveSessionRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	SessionID  string `json:"session_id"`  // 세션 발급시 발급받은 세션 ID
	InvokeID   string `json:"invoke_id"`   // API Transaction ID
}

type AliveSessionResponse struct {
	types.ServiceResponse
	SessionID string `json:"session_id"` // Req로 받은 API Transaction ID
	InvokeID  string `json:"invoke_id"`  // Req로 받은 세션 ID
}

// Web 페이지 접속
type ConnectSessionRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	SessionID  string `json:"session_id"`  // 세션 발급시 발급받은 세션 ID
	InvokeID   string `json:"invoke_id"`   // API Transaction ID
	ConnType   string `json:"conn_type"`   // 세션 발급 구분 tserver인경우 토큰은 Ani번호
	DeviceType string `json:"device_type"` // 단말 OS 유형
}

type ConnectSessionResponse struct {
	types.ServiceResponse
	SessionID    string `json:"session_id"`    // Req로 받은 세션 ID
	InvokeID     string `json:"invoke_id"`     // Req로 받은 API Transaction ID
	PageFilename string `json:"page_filename"` // 연결된 Flow의 최초 페이지 파일명
	PageData     string `json:"page_data"`     // 연결된 Flow의 최초 페이지 데이터
	Bargin       bool   `json:"bargin"`        // Flow에서 전달하는 Flag값
	EndFlag      bool   `json:"end_flag"`      // Flow 종료여부
}

// 페이지 요청
type SubmitSessionRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	SessionID  string `json:"session_id"`  // 세션 발급시 발급받은 세션 ID
	InvokeID   string `json:"invoke_id"`   // API Transaction ID
	SubmitData string `json:"submit_data"` // 웹화면 입력 데이터
	Bargin     bool   `json:"bargin"`      // Flow로 전달하는 Flag값
}

type SubmitSessionResponse struct {
	types.ServiceResponse
	SessionID    string `json:"session_id"`    // Req로 받은 세션 ID
	InvokeID     string `json:"invoke_id"`     // Req로 받은 API Transaction ID
	PageFilename string `json:"page_filename"` // 연결된 Flow의 최초 페이지 파일명
	PageData     string `json:"page_data"`     // 연결된 Flow의 최초 페이지 데이터
	Bargin       bool   `json:"bargin"`        // Flow에서 전달하는 Flag값
	EndFlag      bool   `json:"end_flag"`      // Flow 종료여부
}

// 세션종료
type DisconnectSessionRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	SessionID  string `json:"session_id"`  // 세션 발급시 발급받은 세션 ID
	InvokeID   string `json:"invoke_id"`   // API Transaction ID
}

type DisconnectSessionResponse struct {
	types.ServiceResponse
	SessionID string `json:"session_id"` // Req로 받은 세션 ID
	InvokeID  string `json:"invoke_id"`  // Req로 받은 API Transaction ID
}
