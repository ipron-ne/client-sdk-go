package omnigw

import "time"

const (
	EventAlerting     = "event.alerting"     // 상담사 연결 이벤트
	EventConnected    = "event.connected"    // 상담원 연결 응답 이벤트
	EventEnd          = "event.end"          // 콜 종료 이벤트 (호 회수와 콜 종료의 Disconnected 구분이 어려움에 따라서 별도 이벤트로 분리 )
	EventPause        = "event.pause"        // 지정콜 보류 요청 이벤트
	EventContinue     = "event.continue"     // 지정콜 보류해제 요청 이벤트
	EventTransfer     = "event.transfer"     // 지정콜 협의호전환 이벤트
	EventDisconnected = "event.disconnected" // 콜 상담원 연결 종료 이벤트
)

type OmniGwMsg struct {
	Event           string    `json:"event,omitempty"`         // 이벤트
	UCID            string    `json:"ucid,omitempty"`          // UCID
	TenantID        string    `json:"tenantId,omitempty"`      // 테넌트 아이디
	CallID          string    `json:"callId,omitempty"`        // 콜 아이디
	TransToCallId   string    `json:"transToCallId,omitempty"` // 협의전환 대상 콜 아이디
	MediaType       string    `json:"mediaType,omitempty"`     // 미디어 타입
	ProviderID      string    `json:"providerId,omitempty"`    // 제공자 ID
	UserID          string    `json:"userId,omitempty"`        // 상담원 ID
	UserName        string    `json:"userName,omitempty"`      // 상담원 이름
	ConnectionState string    `json:"connState,omitempty"`     // 참가자 상태
	Reason          string    `json:"reason,omitempty"`        // 사유
	CreateTime      time.Time `json:"createTime,omitempty"`    // 콜 생성 시간
	NowTime         time.Time `json:"nowTime,omitempty"`       // 현재시간
}
