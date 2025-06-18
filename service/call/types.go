package call

import "github.com/ipron-ne/client-sdk-go/types"

type CallRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`
}

type MakeCallRequest struct {
	Tenant    string `json:"tenant"`
	UserID    string `json:"user_id"`
	CallID    string `json:"call_id"`
	ANI       string `json:"ani"`
	DNIS      string `json:"dnis"`
	UserANI   string `json:"user_ani"`
	MediaType string `json:"media_type"`
}

type MakeCallResponse struct {
	types.ServiceResponse
	CallID string `json:"call_id"`
}

type MakeCallExRequest struct {
	Tenant    string      `json:"tenant"`
	UserID    string      `json:"user_id"`
	ANI       string      `json:"ani"`
	DNIS      string      `json:"dnis"`
	UserANI   string      `json:"user_ani"`
	MediaType string      `json:"media_type"`
	UEI       string      `json:"uei"`
	UUI       string      `json:"uui"`
	RouteOpt  RouteOption `json:"route_opt"`
}

type RouteOption struct {
	Type            int    `json:"type"`
	Priority        int    `json:"priority"`
	RelationMethod  int    `json:"relation_method"`
	RelationAgentID string `json:"relation_agent_id"`
	RelationTimeout int    `json:"relation_timeout"`
	Method          int    `json:"method"`
	SkillID         string `json:"skill_id"`
	SkillLevel      int    `json:"skill_level"`
	GroupID         string `json:"group_id"`
}

type MakeCallExResponse struct {
	types.ServiceResponse
	CallID string `json:"call_id"`
}

type AnswerRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`
}

type AnswerResponse struct {
	types.ServiceResponse
}

type ReleaseCallRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`
}

type ReleaseCallResponse struct {
	types.ServiceResponse
}

type HoldRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`
}

type HoldResponse struct {
	types.ServiceResponse
}

type UnholdRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`
}

type UnholdResponse struct {
	types.ServiceResponse
}

type SingleStepTransferRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`

	DNIS     string      `json:"dnis"`
	UserAni  string      `json:"user_ani"`
	UUI      string      `json:"uui"`
	UEI      string      `json:"uei"`
	RouteOpt RouteOption `json:"route_opt"`
}

type SingleStepTransferResponse struct {
	types.ServiceResponse
}

type MuteTransferRequest struct {
	Tenant       string `json:"tenant"`
	HoldCallID   string `json:"hold_call_id"`
	HoldConnID   string `json:"hold_conn_id"`
	ActiveCallID string `json:"active_call_id"`
}

type MuteTransferResponse struct {
	types.ServiceResponse
}

type SingleStepConferenceRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	ConnID string `json:"conn_id"`

	DNIS      string `json:"dnis"`
	UserAni   string `json:"user_ani"`
	UUI       string `json:"uui"`
	UEI       string `json:"uei"`
	PartyType string `json:"party_type"`
}

type SingleStepConferenceResponse struct {
	types.ServiceResponse
}

type MuteConferenceRequest struct {
	Tenant       string `json:"tenant"`
	HoldCallID   string `json:"hold_call_id"`
	HoldConnID   string `json:"hold_conn_id"`
	ActiveCallID string `json:"active_call_id"`
	PartyType    string `json:"party_type"`
}

type MuteConferenceResponse struct {
	types.ServiceResponse
}

type JoinCallRequest struct {
	Tenant     string `json:"tenant"`
	UserID     string `json:"user_id"`
	JoinCallID string `json:"join_call_id"`
	JoinConnID string `json:"join_conn_id"`
	JoinType   string `json:"join_type"`
}

type JoinCallResponse struct {
	types.ServiceResponse
}

type RouteRequest struct {
	Tenant      string `json:"tenant"`
	CallID      string `json:"call_id"`
	PartID      string `json:"part_id"`
	PartType    string `json:"part_type"`
	ANI         string `json:"ani"`
	DNIS        string `json:"dnis"`
	CallbackURI string `json:"callback_uri"`
	Timeout     int    `json:"timeout"`
	RetryCall   bool   `json:"retry_call"`
	AutoAnswer  bool   `json:"auto_answer"`
}

type RouteResponse struct {
	types.ServiceResponse
}

type NumberplanRequest struct {
	Tenant string `json:"tenant"`
	SiteID string `json:"site_id"`
	DNIS   string `json:"dnis"`
}

type NumberplanResponse struct {
	types.ServiceResponse
	NumberplanID   string `json:"numberplan_id"`
	NumberplanName string `json:"numberplan_name"`
	PartType       string `json:"part_type"`
	PartID         string `json:"part_id"`
	PartName       string `json:"part_name"`
	PartStatus     string `json:"part_status"`
	ChangeDNIS     string `json:"change_dnis"`
}

type SetUserdataRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
	UEI    string `json:"uei"`
	UUI    string `json:"uui"`
}

type SetUserdataResponse struct {
	types.ServiceResponse
}

type GetUserdataRequest struct {
	Tenant string `json:"tenant"`
	CallID string `json:"call_id"`
}

type GetUserdataResponse struct {
	types.ServiceResponse
	UEI string `json:"uei"`
	UUI string `json:"uui"`
}

type RemoteTransferRequest struct {
	Tenant  string `json:"tenant"`
	CallID  string `json:"call_id"`
	ConnID  string `json:"conn_id"`
	DNIS    string `json:"dnis"`
	UserANI string `json:"user_ani"`
	UUI     string `json:"uui"`
}

type RemoteTransferResponse struct {
	types.ServiceResponse
}
