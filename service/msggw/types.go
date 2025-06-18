package msggw

import "github.com/ipron-ne/client-sdk-go/types"

// Chat 연결 요청
type ConnectRequest struct {
	ProviderID string `json:"provider_id"` // 수신 Provider ID
	TenantID   string `json:"tenant_id"`   // 테넌트 ID
	SiteID     string `json:"site_id"`     // 사이트 ID
	ANI        string `json:"ani"`         // 발신번호
	DNIS       string `json:"dnis"`        // 착신번호
	Time       string `json:"time"`        // 인입 시간
	Timeout    int    `json:"timeout"`     // 타임아웃
}

type ConnectResponse struct {
	types.ServiceResponse
	ProviderID string `json:"provider_id"` // 요청 Provider ID
	CallID     string `json:"call_id"`     // 생성 된 콜 ID
}

// 고객 입력 내용 전달
type MessageRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	CallID     string `json:"call_id"`     // 콜 ID
	Text       string `json:"text"`        // 고객 입력내용
}

type MessageResponse struct {
	types.ServiceResponse
	ProviderID string `json:"provider_id"` // 요청 Provider ID
	CallID     string `json:"call_id"`     // 생성 된 콜 ID
}

// Chat 종료 요청
type ByeRequest struct {
	ProviderID string `json:"provider_id"` // Provider ID
	CallID     string `json:"call_id"`     // 콜 ID
}

type ByeResponse struct {
	types.ServiceResponse
	ProviderID string `json:"provider_id"` // 요청 Provider ID
	CallID     string `json:"call_id"`     // 생성 된 콜 ID
}
