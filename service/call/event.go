package call

import (
	"encoding/json"
	"time"
)

const (
	EventCreate         = "event.create"
	EventOriginated     = "event.originated"
	EventAlerting       = "event.alerting"
	EventConnected      = "event.connected"
	EventHold           = "event.hold"
	EventUnHold         = "event.unhold"
	EventDisconnected   = "event.disconnected"
	EventTerminated     = "event.terminated"
	EventPartyChanged   = "event.partychanged"
	EventUpdateUserData = "event.updateuserdata"
)

// basic infomation
type PartyInfo struct {
	EndPointType string `json:"epType"`
	EndPointID   string `json:"epId"`
	ConnID       string `json:"connId"`
}

// call event
type Create struct {
	Event            string      `json:"event,omitempty"`       // 이벤트
	UCID             string      `json:"ucid,omitempty"`        // UCID
	CallSeq          json.Number `json:"callSeq,omitempty"`     // 콜 시퀀스
	TenantID         string      `json:"tenantId,omitempty"`    // 테넌트 아이디
	CallID           string      `json:"callId,omitempty"`      // 콜 아이디
	MediaType        string      `json:"mediaType,omitempty"`   // 미디어 타입
	Category         string      `json:"category,omitempty"`    // 콜 카테고리
	CallType         string      `json:"callType,omitempty"`    // 콜 타입
	CallSubType      string      `json:"callSubType,omitempty"` // 콜 서브 타입
	ANI              string      `json:"ani,omitempty"`         // 발신번호
	DNIS             string      `json:"dnis,omitempty"`        // 착신번호
	OriginCallNumber string      `json:"ocn,omitempty"`         // 원 착신번호
	RedirectNumber   string      `json:"rdn,omitempty"`         // 전환 발신번호
	UUI              string      `json:"uui,omitempty"`         // 유저 에이전트 정보
	UEI              string      `json:"uei,omitempty"`         // 사용자 확장정보
	RefCallID        string      `json:"refCallId,omitempty"`   // 연관 콜 ID
	CreateTime       time.Time   `json:"createTime,omitempty"`  // 콜 생성 시간
	NowTime          time.Time   `json:"nowTime,omitempty"`     // 현재시간
}

type Originated struct {
	Event              string      `json:"event,omitempty"`        // 이벤트
	EventEndPointID    string      `json:"eventEpId,omitempty"`    // 이벤트 발생 참가자 아이디
	EventEndPointType  string      `json:"eventEpType,omitempty"`  // 이벤트 발생 참가자 유형
	EventUserID        string      `json:"eventUserId,omitempty"`  // 이벤트 발생 유저 아이디
	EndPointID         string      `json:"epId,omitempty"`         // 이벤트 수신 참가자 아이디
	EndPointType       string      `json:"epType,omitempty"`       // 이벤트 수신 참가자 유형
	EndPointName       string      `json:"epName,omitempty"`       // 이벤트 수신 참가자 명
	UCID               string      `json:"ucid,omitempty"`         // UCID
	CallSeq            json.Number `json:"callSeq,omitempty"`      // 콜 시퀀스
	TenantID           string      `json:"tenantId,omitempty"`     // 테넌트 아이디
	UserID             string      `json:"userId,omitempty"`       // 유저 아이디
	PhoneLineID        string      `json:"phoneLineId,omitempty"`  // 폰 라인 아이디
	CallID             string      `json:"callId,omitempty"`       // 콜 아이디
	ConnectionID       string      `json:"connId,omitempty"`       // 참가자 아이디
	ConnectionOldState string      `json:"connOldState,omitempty"` // 참가자 전 상태
	ConnectionNewState string      `json:"connNewState,omitempty"` // 참가자 현재 상태
	MediaType          string      `json:"mediaType,omitempty"`    // 미디어 타입
	Category           string      `json:"category,omitempty"`     // 콜 카테고리
	CallType           string      `json:"callType,omitempty"`     // 콜 타입
	CallSubType        string      `json:"callSubType,omitempty"`  // 콜 서브 타입
	CallerID           string      `json:"callerId,omitempty"`     // 발신 참가자 아이디
	CallerType         string      `json:"callerType,omitempty"`   // 발신 참가자 유형
	CalleeID           string      `json:"calleeId,omitempty"`     // 착신 참가자 아이디
	CalleeType         string      `json:"calleeType,omitempty"`   // 착신 참가자 유형
	ANI                string      `json:"ani,omitempty"`          // 발신번호
	DNIS               string      `json:"dnis,omitempty"`         // 착신번호
	OriginalNum        string      `json:"on,omitempty"`           // 원착신번호
	RealNum            string      `json:"rn,omitempty"`           // 실착신번호
	AccessCode         string      `json:"ac,omitempty"`           // 접근 코드
	Pattern            string      `json:"pattern,omitempty"`      // 접근 패턴
	UUI                string      `json:"uui,omitempty"`          // 유저 에이전트 정보
	UEI                string      `json:"uei,omitempty"`          // 사용자 확장정보
	Reason             string      `json:"reason,omitempty"`       // 상세이유
	CreateTime         time.Time   `json:"createTime,omitempty"`   // 콜 생성시간
	NowTime            time.Time   `json:"nowTime,omitempty"`      // 현재시간
}

type Alerting struct {
	Event              string      `json:"event,omitempty"`        // 이벤트
	EventEndPointID    string      `json:"eventEpId,omitempty"`    // 이벤트 발생 참가자 아이디
	EventEndPointType  string      `json:"eventEpType,omitempty"`  // 이벤트 발생 참가자 유형
	EventUserID        string      `json:"eventUserId,omitempty"`  // 이벤트 발생 유저 아이디
	EndPointID         string      `json:"epId,omitempty"`         // 이벤트 수신 참가자 아이디
	EndPointType       string      `json:"epType,omitempty"`       // 이벤트 수신 참가자 유형
	EndPointName       string      `json:"epName,omitempty"`       // 이벤트 수신 참가자 명
	UCID               string      `json:"ucid,omitempty"`         // UCID
	CallSeq            json.Number `json:"callSeq,omitempty"`      // 콜 시퀀스
	TenantID           string      `json:"tenantId,omitempty"`     // 테넌트 아이디
	UserID             string      `json:"userId,omitempty"`       // 유저 아이디
	PhoneLineID        string      `json:"phoneLineId,omitempty"`  // 폰 라인 아이디
	CallID             string      `json:"callId,omitempty"`       // 콜 아이디
	ConnectionID       string      `json:"connId,omitempty"`       // 참가자 아이디
	ConnectionOldState string      `json:"connOldState,omitempty"` // 참가자 전 상태
	ConnectionNewState string      `json:"connNewState,omitempty"` // 참가자 현재 상태
	MediaType          string      `json:"mediaType,omitempty"`    // 미디어 타입
	Category           string      `json:"category,omitempty"`     // 콜 카테고리
	CallType           string      `json:"callType,omitempty"`     // 콜 타입
	CallSubType        string      `json:"callSubType,omitempty"`  // 콜 서브 타입
	CallerID           string      `json:"callerId,omitempty"`     // 발신 참가자 아이디
	CallerType         string      `json:"callerType,omitempty"`   // 발신 참가자 유형
	CalleeID           string      `json:"calleeId,omitempty"`     // 착신 참가자 아이디
	CalleeType         string      `json:"calleeType,omitempty"`   // 착신 참가자 유형
	ANI                string      `json:"ani,omitempty"`          // 발신번호
	DNIS               string      `json:"dnis,omitempty"`         // 착신번호
	OriginalNum        string      `json:"on,omitempty"`           // 원착신번호
	RealNum            string      `json:"rn,omitempty"`           // 실착신번호
	AccessCode         string      `json:"ac,omitempty"`           // 접근 코드
	Pattern            string      `json:"pattern,omitempty"`      // 접근 패턴
	UUI                string      `json:"uui,omitempty"`          // 유저 에이전트 정보
	UEI                string      `json:"uei,omitempty"`          // 사용자 확장정보
	Reason             string      `json:"reason,omitempty"`       // 상세이유
	FirstAcdID         string      `json:"firstAcdId,omitempty"`   // 최초 진입 ACD 아이디
	CurrAcdID          string      `json:"currAcdId,omitempty"`    // 현재 진입 ACD 아이디
	AcdTransCount      json.Number `json:"acdTransCnt,omitempty"`  // ACD 전환 횟수
	FirstFlowID        string      `json:"firstFlowId,omitempty"`  // 최초 FLOW 아이디
	CurrFlowID         string      `json:"currFlowId,omitempty"`   // 현재 FLOW 아이디
	UserInCount        json.Number `json:"userInCnt,omitempty"`    // 콜 종료시까지 유저 연결 수
	FirstSkillID       string      `json:"firstSkillId,omitempty"` // 최초 연결 스킬 아이디
	CurrSkillID        string      `json:"currSkillId,omitempty"`  // 현재 연결 스킬 아이디
	CreateTime         time.Time   `json:"createTime,omitempty"`   // 콜 생성 시간
	NowTime            time.Time   `json:"nowTime,omitempty"`      // 현재시간
}

type Connected struct {
	Event              string       `json:"event,omitempty"`          // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`      // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`    // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`    // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`           // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`         // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`         // 이벤트 수신 참가자 명
	UCID               string       `json:"ucid,omitempty"`           // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`        // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`       // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`         // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`    // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`         // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`         // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"`   // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"`   // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`      // 미디어 타입
	Category           string       `json:"category,omitempty"`       // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`       // 콜 타입
	CallSubType        string       `json:"callSubType,omitempty"`    // 콜 서브 타입
	CallerID           string       `json:"callerId,omitempty"`       // 발신 참가자 아이디
	CallerType         string       `json:"callerType,omitempty"`     // 발신 참가자 유형
	CalleeID           string       `json:"calleeId,omitempty"`       // 착신 참가자 아이디
	CalleeType         string       `json:"calleeType,omitempty"`     // 착신 참가자 유형
	PartyCount         json.Number  `json:"partyCnt,omitempty"`       // 참가자 수
	PartyInfoSet       *[]PartyInfo `json:"partyInfoSet,omitempty"`   // 참가자 정보
	ANI                string       `json:"ani,omitempty"`            // 발신번호
	DNIS               string       `json:"dnis,omitempty"`           // 착신번호
	OriginalNum        string       `json:"on,omitempty"`             // 원착신번호
	RealNum            string       `json:"rn,omitempty"`             // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`             // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`        // 접근 패턴
	UUI                string       `json:"uui,omitempty"`            // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`            // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`         // 상세이유
	LastRingTime       time.Time    `json:"lastRingTime,omitempty"`   // 마지막 벨 울림 시간
	FirstAcdID         string       `json:"firstAcdId,omitempty"`     // 최초 ACD 아이디
	CurrAcdID          string       `json:"currAcdId,omitempty"`      // 현재 ACD 아이디
	AcdTransCount      json.Number  `json:"acdTransCnt,omitempty"`    // ACD 전환 횟수
	FirstAcdInTime     time.Time    `json:"firstAcdInTime,omitempty"` // 최초 ACD 인입 시간
	CurrAcdInTime      time.Time    `json:"currAcdInTime,omitempty"`  // 현재 ACD 인입 시간
	FirstFlowID        string       `json:"firstFlowId,omitempty"`    // 최초 FLOW 아이디
	CurrFlowId         string       `json:"currFlowId,omitempty"`     // 현재 FLOW 아이디
	UserInCount        json.Number  `json:"userInCnt,omitempty"`      // 콜 종료시까지 유저 연결 수
	FirstSkillID       string       `json:"firstSkillId,omitempty"`   // 최초 연결 스킬 아이디
	CurrSkillID        string       `json:"currSkillId,omitempty"`    // 현재 연결 스킬 아이디
	CreateTime         time.Time    `json:"createTime,omitempty"`     // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`        // 현재시간
}

type Disconnected struct {
	Event              string       `json:"event,omitempty"`           // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`       // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`     // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`     // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`            // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`          // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`          // 이벤트 수신 참가자 유형
	UCID               string       `json:"ucid,omitempty"`            // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`         // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`        // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`          // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`     // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`          // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`          // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"`    // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"`    // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`       // 미디어 타입
	Category           string       `json:"category,omitempty"`        // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`        // 콜 타입
	CallSubType        string       `json:"callSubType,omitempty"`     // 콜 서브 타입
	CallerID           string       `json:"callerId,omitempty"`        // 발신 참가자 아이디
	CallerType         string       `json:"callerType,omitempty"`      // 발신 참가자 유형
	CalleeID           string       `json:"calleeId,omitempty"`        // 착신 참가자 아이디
	CalleeType         string       `json:"calleeType,omitempty"`      // 착신 참가자 유형
	OldPartyCount      json.Number  `json:"oldPartyCnt,omitempty"`     // 전 참가자 수
	OldPartyInfoSet    *[]PartyInfo `json:"oldPartyInfoSet,omitempty"` // 전 참가자 정보
	NewPartyCount      json.Number  `json:"newPartyCnt,omitempty"`     // 현재 참가자 수
	NewPartyInfoSet    *[]PartyInfo `json:"newPartyInfoSet,omitempty"` // 현재 참가자 정보
	ANI                string       `json:"ani,omitempty"`             // 발신번호
	DNIS               string       `json:"dnis,omitempty"`            // 착신번호
	OriginalNum        string       `json:"on,omitempty"`              // 원착신번호
	RealNum            string       `json:"rn,omitempty"`              // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`              // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`         // 접근 패턴
	UUI                string       `json:"uui,omitempty"`             // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`             // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`          // 상세이유
	TransToType        string       `json:"transToType,omitempty"`     // 호 전환 참가자 타입
	TransToID          string       `json:"transToId,omitempty"`       // 호전환 참가자 아이디
	FirstFlowInTime    time.Time    `json:"firstFlowInTime,omitempty"` // 최초 FLOW 인입 시간
	LastRingTime       time.Time    `json:"lastRingTime,omitempty"`    // 마지막 벨 울림 시간
	LastConnTime       time.Time    `json:"lastConnTime,omitempty"`    // 마지막 연결 시간
	LastHoldTime       time.Time    `json:"lastHoldTime,omitempty"`    // 마지막 보류 시간
	LastConfTime       time.Time    `json:"lastConfTime,omitempty"`    // 마지막 협의 시간
	FirstAcdID         string       `json:"firstAcdId,omitempty"`      // 최초 ACD 아이디
	CurrAcdID          string       `json:"currAcdId,omitempty"`       // 현재 ACD 아이디
	FlowTransCount     json.Number  `json:"flowTransCnt,omitempty"`    // FLOW 전환 수
	AcdTransCount      json.Number  `json:"acdTransCnt,omitempty"`     // ACD 전환 수
	FirstAcdInTime     time.Time    `json:"firstAcdInTime,omitempty"`  // 최초 ACD 인입 시간
	CurrAcdInTime      time.Time    `json:"currAcdInTime,omitempty"`   // 현재 큐 진입 시간
	FirstFlowID        string       `json:"firstFlowId,omitempty"`     // 최초 FLOW 아이디
	CurrFlowID         string       `json:"currFlowId,omitempty"`      // 현재 FLOW 아이디
	FirstSkillID       string       `json:"firstSkillId,omitempty"`    // 최초 연결 스킬 아이디
	CurrSkillID        string       `json:"currSkillId,omitempty"`     // 현재 연결 스킬 아이디
	EndState           bool         `json:"endState"`                  // 통화종료여부
	EndPart            string       `json:"endPart,omitempty"`         // Hop 종료주체
	UserInCount        json.Number  `json:"userInCnt,omitempty"`       // 콜 종료시까지 유저 연결 수
	CreateTime         time.Time    `json:"createTime,omitempty"`      // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`         // 현재시간
}

type Terminated struct {
	Event            string      `json:"event,omitempty"`       // 이벤트
	UCID             string      `json:"ucid,omitempty"`        // UCID
	CallSeq          json.Number `json:"callSeq,omitempty"`     // 콜 시퀀스
	TenantID         string      `json:"tenantId,omitempty"`    // 테넌트 아이디
	CallID           string      `json:"callId,omitempty"`      // 콜 아이디
	MediaType        string      `json:"mediaType,omitempty"`   // 미디어 타입
	Category         string      `json:"category,omitempty"`    // 콜 카테고리
	CallType         string      `json:"callType,omitempty"`    // 콜 타입
	CallSubType      string      `json:"callSubType,omitempty"` // 콜 서브 타입
	ANI              string      `json:"ani,omitempty"`         // 발신번호
	DNIS             string      `json:"dnis,omitempty"`        // 착신번호
	OriginCallNumber string      `json:"ocn,omitempty"`         // 원 착신번호
	RedirectNumber   string      `json:"rdn,omitempty"`         // 전환 발신번호
	UUI              string      `json:"uui,omitempty"`         // 유저 에이전트 정보
	UEI              string      `json:"uei,omitempty"`         // 사용자 확장정보
	RefCallID        string      `json:"refCallId,omitempty"`   // 연관 콜 ID
	EndType          string      `json:"endType,omitempty"`     // 콜 종료타입
	CreateTime       time.Time   `json:"createTime,omitempty"`  // 콜 생성 시간
	NowTime          time.Time   `json:"nowTime,omitempty"`     // 현재시간
}

type PartyChanged struct {
	Event              string       `json:"event,omitempty"`           // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`       // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`     // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`     // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`            // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`          // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`          // 이벤트 수신 참가자 유형
	UCID               string       `json:"ucid,omitempty"`            // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`         // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`        // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`          // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`     // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`          // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`          // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"`    // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"`    // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`       // 미디어 타입
	Category           string       `json:"category,omitempty"`        // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`        // 콜 유형
	CallSubType        string       `json:"callSubType,omitempty"`     // 콜 서브 타입
	ANI                string       `json:"ani,omitempty"`             // 발신번호
	DNIS               string       `json:"dnis,omitempty"`            // 착신번호
	OriginalNum        string       `json:"on,omitempty"`              // 원착신번호
	RealNum            string       `json:"rn,omitempty"`              // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`              // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`         // 접근 패턴
	CallerID           string       `json:"callerId,omitempty"`        // 발신 참가자 아이디
	CallerType         string       `json:"callerType,omitempty"`      // 발신 참가자 유형
	CalleeID           string       `json:"calleeId,omitempty"`        // 착신 참가자 아이디
	CalleeType         string       `json:"calleeType,omitempty"`      // 착신 참가자 유형
	UUI                string       `json:"uui,omitempty"`             // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`             // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`          // 상세이유
	ChangeFromID       string       `json:"changeFromId,omitempty"`    // 호 변경 참가자 아이디
	ChangeFromType     string       `json:"changeFromType,omitempty"`  // 호 변경 참가자 타입
	ChangeToID         string       `json:"changeToId,omitempty"`      // 호 참여 참가자 아이디
	ChangeToType       string       `json:"changeToType,omitempty"`    // 호 참여 참가자 유형
	OldPartyCount      json.Number  `json:"oldPartyCnt,omitempty"`     // 전 참가자 수
	OldPartyInfoSet    *[]PartyInfo `json:"oldPartyInfoSet,omitempty"` // 전 참가자 정보
	NewPartyCount      json.Number  `json:"newPartyCnt,omitempty"`     // 현재 참가자 수
	NewPartyInfoSet    *[]PartyInfo `json:"newPartyInfoSet,omitempty"` // 현재 참가자 정보
	LastConnTime       time.Time    `json:"lastConnTime,omitempty"`    // 마지막 연결 시간
	NewCallSeq         json.Number  `json:"newCallSeq,omitempty"`      // 새로운 콜 시퀀스
	NewCallID          string       `json:"newCallId,omitempty"`       // 새로운 콜 아이디
	NewConnectionID    string       `json:"newConnId,omitempty"`       // 새로운 참가자 아이디
	NewCategory        string       `json:"newCategory,omitempty"`     // 새로운 콜 카테고리
	NewCallType        string       `json:"newCallType,omitempty"`     // 새로운 콜 타입
	NewCallSubType     string       `json:"newCallSubType,omitempty"`  // 새로운 콜 서브 타입
	NewANI             string       `json:"newAni,omitempty"`          // 새로운 콜 발신번호
	NewDNIS            string       `json:"newDnis,omitempty"`         // 새로운 콜 착신번호
	NewCreateTime      time.Time    `json:"newCreateTime,omitempty"`   // 콜 생성 시간
	EndState           bool         `json:"endState"`                  // 통화종료여부
	EndPart            string       `json:"endPart,omitempty"`         // Hop 종료주체
	UserInCount        json.Number  `json:"userInCnt,omitempty"`       // 콜 종료시까지 유저 연결 수
	CreateTime         time.Time    `json:"createTime,omitempty"`      // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`         // 현재시간
}

type Hold struct {
	Event              string       `json:"event,omitempty"`        // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`    // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`  // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`  // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`         // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`       // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`       // 이벤트 수신 참가자 유형
	UCID               string       `json:"ucid,omitempty"`         // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`      // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`     // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`       // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`  // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`       // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`       // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"` // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"` // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`    // 미디어 타입
	Category           string       `json:"category,omitempty"`     // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`     // 콜 유형
	CallSubType        string       `json:"callSubType,omitempty"`  // 콜 서브 타입
	PartyCount         json.Number  `json:"partyCnt,omitempty"`     // 참가자 수
	PartyInfoSet       *[]PartyInfo `json:"partyInfoSet,omitempty"` // 참가자 정보
	ANI                string       `json:"ani,omitempty"`          // 발신번호
	DNIS               string       `json:"dnis,omitempty"`         // 착신번호
	OriginalNum        string       `json:"on,omitempty"`           // 원착신번호
	RealNum            string       `json:"rn,omitempty"`           // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`           // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`      // 접근 패턴
	UUI                string       `json:"uui,omitempty"`          // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`          // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`       // 상세이유
	FirstAcdID         string       `json:"firstAcdId,omitempty"`   // 최초 ACD 아이디
	FirstSkillID       string       `json:"firstSkillId,omitempty"` // 최초 연결 스킬 아이디
	CurrSkillID        string       `json:"currSkillId,omitempty"`  // 현재 연결 스킬 아이디
	CreateTime         time.Time    `json:"createTime,omitempty"`   // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`      // 현재시간
}

type UnHold struct {
	Event              string       `json:"event,omitempty"`        // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`    // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`  // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`  // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`         // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`       // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`       // 이벤트 수신 참가자 유형
	UCID               string       `json:"ucid,omitempty"`         // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`      // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`     // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`       // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`  // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`       // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`       // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"` // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"` // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`    // 미디어 타입
	Category           string       `json:"category,omitempty"`     // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`     // 콜 유형
	CallSubType        string       `json:"callSubType,omitempty"`  // 콜 서브 타입
	PartyCount         json.Number  `json:"partyCnt,omitempty"`     // 참가자 수
	PartyInfoSet       *[]PartyInfo `json:"partyInfoSet,omitempty"` // 참가자 정보
	ANI                string       `json:"ani,omitempty"`          // 발신번호
	DNIS               string       `json:"dnis,omitempty"`         // 착신번호
	OriginalNum        string       `json:"on,omitempty"`           // 원착신번호
	RealNum            string       `json:"rn,omitempty"`           // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`           // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`      // 접근 패턴
	UUI                string       `json:"uui,omitempty"`          // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`          // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`       // 상세이유
	LastHoldTime       time.Time    `json:"lastHoldTime,omitempty"` // 마지막 호 보류 시간
	FirstAcdID         string       `json:"firstAcdId,omitempty"`   // 최초 ACD 아이디
	FirstSkillID       string       `json:"firstSkillId,omitempty"` // 최초 연결 스킬 아이디
	CurrSkillID        string       `json:"currSkillId,omitempty"`  // 현재 연결 스킬 아이디
	CreateTime         time.Time    `json:"createTime,omitempty"`   // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`      // 현재시간
}

type UpdateUserData struct {
	Event              string       `json:"event,omitempty"`        // 이벤트
	EventEndPointID    string       `json:"eventEpId,omitempty"`    // 이벤트 발생 참가자 아이디
	EventEndPointType  string       `json:"eventEpType,omitempty"`  // 이벤트 발생 참가자 유형
	EventUserID        string       `json:"eventUserId,omitempty"`  // 이벤트 발생 유저 아이디
	EndPointID         string       `json:"epId,omitempty"`         // 이벤트 수신 참가자 아이디
	EndPointType       string       `json:"epType,omitempty"`       // 이벤트 수신 참가자 유형
	EndPointName       string       `json:"epName,omitempty"`       // 이벤트 수신 참가자 유형
	UCID               string       `json:"ucid,omitempty"`         // UCID
	CallSeq            json.Number  `json:"callSeq,omitempty"`      // 콜 시퀀스
	TenantID           string       `json:"tenantId,omitempty"`     // 테넌트 아이디
	UserID             string       `json:"userId,omitempty"`       // 유저 아이디
	PhoneLineID        string       `json:"phoneLineId,omitempty"`  // 폰 라인 아이디
	CallID             string       `json:"callId,omitempty"`       // 콜 아이디
	ConnectionID       string       `json:"connId,omitempty"`       // 참가자 아이디
	ConnectionOldState string       `json:"connOldState,omitempty"` // 참가자 전 상태
	ConnectionNewState string       `json:"connNewState,omitempty"` // 참가자 현재 상태
	MediaType          string       `json:"mediaType,omitempty"`    // 미디어 타입
	Category           string       `json:"category,omitempty"`     // 콜 카테고리
	CallType           string       `json:"callType,omitempty"`     // 콜 유형
	CallSubType        string       `json:"callSubType,omitempty"`  // 콜 서브 타입
	PartyCount         json.Number  `json:"partyCnt,omitempty"`     // 참가자 수
	PartyInfoSet       *[]PartyInfo `json:"partyInfoSet,omitempty"` // 참가자 정보
	ANI                string       `json:"ani,omitempty"`          // 발신번호
	DNIS               string       `json:"dnis,omitempty"`         // 착신번호
	OriginalNum        string       `json:"on,omitempty"`           // 원착신번호
	RealNum            string       `json:"rn,omitempty"`           // 실착신번호
	AccessCode         string       `json:"ac,omitempty"`           // 접근 코드
	Pattern            string       `json:"pattern,omitempty"`      // 접근 패턴
	UUI                string       `json:"uui,omitempty"`          // 유저 에이전트 정보
	UEI                string       `json:"uei,omitempty"`          // 사용자 확장정보
	Reason             string       `json:"reason,omitempty"`       // 상세이유
	LastHoldTime       time.Time    `json:"lastHoldTime,omitempty"` // 마지막 호 보류 시간
	FirstAcdID         string       `json:"firstAcdId,omitempty"`   // 최초 ACD 아이디
	FirstSkillID       string       `json:"firstSkillId,omitempty"` // 최초 연결 스킬 아이디
	CurrSkillID        string       `json:"currSkillId,omitempty"`  // 현재 연결 스킬 아이디
	CreateTime         time.Time    `json:"createTime,omitempty"`   // 콜 생성 시간
	NowTime            time.Time    `json:"nowTime,omitempty"`      // 현재시간
}
