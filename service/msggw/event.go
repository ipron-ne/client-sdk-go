package msggw

const (
	EventMessage     = "event.message"      // message 표시 요청 이벤트
	EventTransfer    = "event.transferchat" // Chat 상담사 Transfer 요청 이벤트
	EventTerminiated = "event.terminated"   // Chat 종료 (시나리오 종료)
)

type Message struct {
	Event      string      `json:"event"`     // event
	ProviderID string      `json:"provideId"` // provide id
	CallID     string      `json:"callId"`    // 콜 id
	ConnID     string      `json:"connId"`    // conn id
	TenantID   string      `json:"tntId"`     // 테넌트 ID
	Text       string      `json:"text"`      // 고객에게 보낼 text
	Chatcode   string      `json:"chatCode"`  // 메시지 표현코드
	ChatFmt    interface{} `json:"chatFmt"`   // 메시지 표현상제코드
}

type TransferChat struct {
	Event      string      `json:"event"`     // event
	CallID     string      `json:"callId"`    // 콜 id
	ProviderID string      `json:"provideId"` // provide id
	ConnID     string      `json:"connId"`    // conn id
	TenantID   string      `json:"tntId"`     // 테넌트 ID
	Dnis       string      `json:"dnis"`      // dnis
	UserAni    string      `json:"userAni"`   // userani
	Uui        string      `json:"uui"`       // uui
	Uei        string      `json:"uei"`       // uei
	RouteOpt   RouteOption `json:"routeOpt"`
}

type RouteOption struct {
	Type            int32  `json:"type"`     // 테넌트 ID
	Priority        int32  `json:"priority"` //priority
	RelationMethod  int32  `json:"relation_method"`
	RelationAgentId string `json:"relation_agent_id"`
	RelationTimeout int32  `json:"relation_timeout"`
	Method          int32  `json:"method"`
	SkillId         string `json:"skill_id"`
	SkillLevel      int32  `json:"skill_level"`
	GroupId         string `json:"group_id"`
}

type TerminatedChat struct {
	Event    string `json:"event"`
	Provider string `json:"provideId"` // callid
	CallId   string `json:"callId"`    // callid
	ConnId   string `json:"connId"`    // conn id
}
