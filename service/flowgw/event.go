package flowgw

import "time"

const (
	EventStart = "event.flowstart"
	EventDial  = "event.makecallresult"
	EventEnd   = "event.flowend"
)

type FlowGWStart struct {
	Event      string    `json:"event"`       // event
	ProviderID string    `json:"provide_id"`  // provide id
	TenantID   string    `json:"tenant_id"`   // 테넌트 ID
	CallID     string    `json:"call_id"`     // call ID
	ConnID     string    `json:"conn_id"`     // conn ID
	Ani        string    `json:"ani"`         // ani
	Dnis       string    `json:"dnis"`        // dnis
	CreateTime time.Time `json:"create_time"` // flow 시작 시간
	NowTime    time.Time `json:"now_time"`    // 이벤트 발생 시간
}

type FlowGWDial struct {
	Event      string        `json:"event"`          // event
	ProviderID string        `json:"provide_id"`     // provide id
	CallID     string        `json:"call_id"`        // call ID
	Ani        string        `json:"ani"`            // ani
	Dnis       string        `json:"dnis"`           // dnis
	Result     int           `json:"result"`         // 1:success, 0:fail, 9:tonedetect
	Reason     string        `json:"reason"`         // 성공 실패
	TonDTInfo  TonDetectInfo `json:"tondetect_info"` // tondetect 결과 정보   result 가 9일 때
	NowTime    time.Time     `json:"now_time"`       // 이벤트 발생 시간
}

type TonDetectInfo struct {
	Code int `json:"code"` // code
	Desc int `json:"desc"` // desc
}

type FlowGWEnd struct {
	Event       string    `json:"event"`         // event
	ProviderID  string    `json:"provide_id"`    // provide id
	TenantID    string    `json:"tenant_id"`     // 테넌트 ID
	CallID      string    `json:"call_id"`       // call ID
	ConnID      string    `json:"conn_id"`       // conn ID
	Ani         string    `json:"ani"`           // ani
	Dnis        string    `json:"dnis"`          // dnis
	Reason      string    `json:"reason"`        // 성공 실패
	EndPart     string    `json:"endPart"`       // 종료 주체
	UserSvcDrop bool      `json:"user_svc_drop"` // 사용자 포기여부
	AgentConn   bool      `json:"agent_conn"`    // 상담사 연결 여부
	Duration    string    `json:"duration"`      // Flow 점유 시간
	CreateTime  time.Time `json:"create_time"`   // flow 시작 시간
	NowTime     time.Time `json:"now_time"`      // 이벤트 발생 시간
}
