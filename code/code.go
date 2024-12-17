// IPRON SDK for go language
// Constant Code Definition
// 
// http://git.bridgetec.co.kr/IPRON-CLOUD-WEB/SDK/cti_sdk/blob/master/src/common/constant.js
package code

import (
  "time"
)

// TIMEOUT 관련 상수
const (
  RECONNECT_TIMEOUT = 5 * time.Second
)

// defined for account auth level type
type UserAuthLevelType string
const (
  UserAuthSuperAdmin  = UserAuthLevelType("superadmin")
  UserAuthAdmin       = UserAuthLevelType("admin")
  UserAuthSvcManager  = UserAuthLevelType("svcmanager")
  UserAuthUserManager = UserAuthLevelType("usermanager")
  UserAuthUser        = UserAuthLevelType("user")
)

type UserScheduleType string
const (
  UserScheduleGroup = UserScheduleType("group")
  UserScheduleUser  = UserScheduleType("user")
)

type SvcForwardType string
const (
  SvcForwardUncondition = SvcForwardType("uncondition")
  SvcForwardNoAnswer    = SvcForwardType("noanswer")
  SvcForwardBusy        = SvcForwardType("busy")
)

type QueueRouteKind string
const (
  QueueRouteStandard = QueueRouteKind("standard")
  QueueRouteExpend   = QueueRouteKind("expend")
)

type AgentChoiceMethod string
const (
  AgentChoiceLastLongWait = AgentChoiceMethod("lastLongWait")
  AgentChoiceAccLongWait  = AgentChoiceMethod("accLongWait")
  AgentChoiceTotMinCall   = AgentChoiceMethod("totMinCall")
  AgentChoiceQueMinCall   = AgentChoiceMethod("queMinCall")
  AgentChoiceTotMinTime   = AgentChoiceMethod("totMinTime")
  AgentChoiceQueMinTime   = AgentChoiceMethod("queMinTime")
  AgentChoiceRoundrobin   = AgentChoiceMethod("roundrobin")
)


// defined for protocol type
type ProtocolType string
const (
  ProtocolTCP ProtocolType = "tcp"
  ProtocolUDP ProtocolType = "udp"
  ProtocolTLS ProtocolType = "tls"
  ProtocolWSS ProtocolType = "wss"
)

// defined for internal use protocol type
type protocol struct {
  TCP ProtocolType
  UDP ProtocolType
  TLS ProtocolType
  WSS ProtocolType
}

// defined for external use protocol type
var Protocol = protocol {
  TCP: ProtocolTCP,
  UDP: ProtocolUDP,
  TLS: ProtocolTLS,
  WSS: ProtocolWSS,
}


// defined for media type
type MediaType string
const (
  MediaVoice MediaType = "voice"
  MediaVideo MediaType = "video"
  MediaChat  MediaType = "chat"
  MediaEmail MediaType = "email"
)

// defined for internal use media type
type media struct {
  Voice MediaType
  Video MediaType
  Chat  MediaType
  Email MediaType
}

// defined for external use media type
var Media = media {
  Voice: MediaVoice,
  Video: MediaVideo,
  Chat:  MediaChat,
  Email: MediaEmail,
}


// defined for agent state type
type AgentStateType string
const (
  AgentStateLogout    AgentStateType = "logout"
  AgentStateLogin     AgentStateType = "login"
  AgentStateReady     AgentStateType = "ready"
  AgentStateInReady   AgentStateType = "inready"
  AgentStateOutReady  AgentStateType = "outready"
  AgentStateNotReady  AgentStateType = "notready"
  AgentStateAfterWork AgentStateType = "afterwork"
  AgentStateDialing   AgentStateType = "dialing"
  AgentStateRinging   AgentStateType = "ringing"
  AgentStateOnline    AgentStateType = "online"
)

// defined for internal use agent state type
type agent_status struct {
  Logout    AgentStateType
  Login     AgentStateType
  Ready     AgentStateType
  InReady   AgentStateType
  OutReady  AgentStateType
  NotReady  AgentStateType
  AfterWork AgentStateType
  Dialing   AgentStateType
  Ringing   AgentStateType
  Online    AgentStateType
}

// defined for external use agent state type
var AgentStatus = agent_status {
  Logout:    AgentStateLogout,
  Login:     AgentStateLogin,
  Ready:     AgentStateReady,
  InReady:   AgentStateInReady,
  OutReady:  AgentStateOutReady,
  NotReady:  AgentStateNotReady,
  AfterWork: AgentStateAfterWork,
  Dialing:   AgentStateDialing,
  Ringing:   AgentStateRinging,
  Online:    AgentStateOnline,
}

// defined for agent state select type
type AgentStateSelectType string
const (
  AgentStateSelectReady    AgentStateSelectType = "ready"
  AgentStateSelectInReady  AgentStateSelectType = "inready"
  AgentStateSelectOutReady AgentStateSelectType = "outready"
  AgentStateSelectNotReady AgentStateSelectType = "notready"
  AgentStateSelectAcw      AgentStateSelectType = "acw"
  AgentStateSelectBusy     AgentStateSelectType = "busy"
  AgentStateSelectDialing  AgentStateSelectType = "dialing"
  AgentStateSelectRinging  AgentStateSelectType = "ringing"
)

// defined for internal use agent state select type
type agent_state_select struct {
  Ready    AgentStateSelectType
  InReady  AgentStateSelectType
  OutReady AgentStateSelectType
  NotReady AgentStateSelectType
  Acw      AgentStateSelectType
  Busy     AgentStateSelectType
  Dialing  AgentStateSelectType
  Ringing  AgentStateSelectType
}

// defined for external use agent state select type
var AgentStateSelect = agent_state_select {
  Ready:    AgentStateSelectReady,
  InReady:  AgentStateSelectInReady,
  OutReady: AgentStateSelectOutReady,
  NotReady: AgentStateSelectNotReady,
  Acw:      AgentStateSelectAcw,
  Busy:     AgentStateSelectBusy,
  Dialing:  AgentStateSelectDialing,
  Ringing:  AgentStateSelectRinging,
}

// defined for agent state cause type
type AgentStateCauseType string
const (
  AgentStateCauseInbound    AgentStateCauseType = "inbound"
  AgentStateCauseOutbound   AgentStateCauseType = "outbound"
  AgentStateCauseBusy       AgentStateCauseType = "busy"
  AgentStateCauseAway       AgentStateCauseType = "away"
  AgentStateCauseBreak      AgentStateCauseType = "break"
  AgentStateCauseIdle       AgentStateCauseType = "idle"
  AgentStateCauseUserDefine AgentStateCauseType = "userdefine"
  AgentStateCauseAcw        AgentStateCauseType = "acw"
)

// defined for internal use busy in agent state cause type
type agent_state_cause_busy struct {
  Inbound  AgentStateCauseType
  Outbound AgentStateCauseType
}

// defined for internal use notready in agent state cause type
type agent_state_cause_notready struct {
  Busy       AgentStateCauseType
  Away       AgentStateCauseType
  Break      AgentStateCauseType
  Idle       AgentStateCauseType
  UserDefine AgentStateCauseType
}

// defined for internal use agent state cause type
type agent_state_cause struct {
  Busy     agent_state_cause_busy
  NotReady agent_state_cause_notready
  Acw      AgentStateCauseType
}

// defined for external use agent state cause type
var AgentStateCause = agent_state_cause {
  Busy: agent_state_cause_busy {
    Inbound:  AgentStateCauseInbound,
    Outbound: AgentStateCauseOutbound,
  },
  NotReady: agent_state_cause_notready {
    Busy:       AgentStateCauseBusy,
    Away:       AgentStateCauseAway,
    Break:      AgentStateCauseBreak,
    Idle:       AgentStateCauseIdle,
    UserDefine: AgentStateCauseUserDefine,
  },
  Acw: AgentStateCauseAcw,
}


// defined for subject type
type SubjectType string
const (
  SubjectAccount SubjectType = "account"
  SubjectCall    SubjectType = "call"
  SubjectRouting SubjectType = "routing"
  SubjectUser    SubjectType = "user"
  SubjectPhone   SubjectType = "phone"
  SubjectFlow    SubjectType = "flow"
  SubjectQueue   SubjectType = "queue"
  SubjectTrunk   SubjectType = "trunk"
)

// defined for internal use subject type
type subject struct {
  Account SubjectType
  Call    SubjectType
  Routing SubjectType
  User    SubjectType
  Phone   SubjectType
  Flow    SubjectType
  Queue   SubjectType
  Trunk   SubjectType
}

// defined for external use subject type
var Subject = subject {
  Account: SubjectAccount,
  Call:    SubjectCall,
  Routing: SubjectRouting,
  User:    SubjectUser,
  Phone:   SubjectPhone,
  Flow:    SubjectFlow,
  Queue:   SubjectQueue,
  Trunk:   SubjectTrunk,
}

// defined for party type
type PartyType string
const (
  PartyTrunk PartyType = "trunk"
  PartyUser  PartyType = "user"
  PartyAcd   PartyType = "acd"
  PartyFlow  PartyType = "flow"
  PartyPhone PartyType = "phone"
)

// defined for internal use party type
type party struct {
  Trunk PartyType
  User  PartyType
  Acd   PartyType
  Flow  PartyType
  Phone PartyType
}

// defined for external use party type
var Party = party {
  Trunk: PartyTrunk,
  User:  PartyUser,
  Acd:   PartyAcd,
  Flow:  PartyFlow,
  Phone: PartyPhone,
}

// defined for call category type
type CallCategoryType string
const (
  CallCategoryIn       CallCategoryType = "in"
  CallCategoryOut      CallCategoryType = "out"
  CallCategoryInternal CallCategoryType = "internal"
)

// defined for internal use call category type
type call_category struct {
  In       CallCategoryType
  Out      CallCategoryType
  Internal CallCategoryType
}

// defined for external use call category type
var CallCategory = call_category {
  In:       CallCategoryIn,
  Out:      CallCategoryOut,
  Internal: CallCategoryInternal,
}

// defined for call type
type CallType string
const (
  CallNormal       CallType = "nornal"
  CallUserTransfer CallType = "user transfer"
  CallConsult      CallType = "consult"
  CallFeature      CallType = "feature"
)

// defined for internal use call type
type call_type struct {
  Normal       CallType
  UserTransfer CallType
  Consult      CallType
  Feature      CallType
}

// defined for external use call type
var Call = call_type {
  Normal:       CallNormal,
  UserTransfer: CallUserTransfer,
  Consult:      CallConsult,
  Feature:      CallFeature,
}

// defined for call state type
type RouteMethodType string
const (
  RouteMethodSkill      RouteMethodType = "skill"
  RouteMethodGroup      RouteMethodType = "group"
  RouteMethodSkillGroup RouteMethodType = "skillgroup"
  RouteMethodFeature    RouteMethodType = "feature"
)

// defined for internal use route method type
type route_method struct {
  Skill      RouteMethodType
  Group      RouteMethodType
  SkillGroup RouteMethodType
  Feature    RouteMethodType
}

// defined for external use route method type
var RouteMethod = route_method {
  Skill:      RouteMethodSkill,
  Group:      RouteMethodGroup,
  SkillGroup: RouteMethodSkillGroup,
  Feature:    RouteMethodFeature,
}

// defined for call state type
type ConnStateType string
const (
  ConnStateNull       ConnStateType = "null"
  ConnStateOriginated ConnStateType = "originated"
  ConnStateAlerting   ConnStateType = "alerting"
  ConnStateDialing    ConnStateType = "dialing"
  ConnStateConnected  ConnStateType = "connected"
  ConnStateHold       ConnStateType = "hold"
  ConnStateRouting    ConnStateType = "routing"
  ConnStateContacting ConnStateType = "contacting"
)

// defined for internal use that connection state type
type conn_state struct {
  Null       ConnStateType
  Originated ConnStateType
  Alerting   ConnStateType
  Dialing    ConnStateType
  Connected  ConnStateType
  Hold       ConnStateType
  Routing    ConnStateType
  Contacting ConnStateType
}

// defined for external use that connection state type
var ConnState = conn_state {
  Null:       ConnStateNull,
  Originated: ConnStateOriginated,
  Alerting:   ConnStateAlerting,
  Dialing:    ConnStateDialing,
  Connected:  ConnStateConnected,
  Hold:       ConnStateHold,
  Routing:    ConnStateRouting,
  Contacting: ConnStateContacting,
}

// defined for connection type
type ConnType string
const (
  ConnToCall  ConnType = "tocall"
  ConnToParty ConnType = "toparty"
)

// defined for internal use that connection type
type conn_type struct {
  ToCall  ConnType
  ToParty ConnType
}

// defined for external use that connection type
var Conn = conn_type {
  ToCall:  ConnToCall,
  ToParty: ConnToParty,
}

// defined for event handler type
type EventHandler string
const (
  REGISTER     EventHandler = "sse.notify.register"
  REGISTERED   EventHandler = "sse.notify.registered"
  PUSH         EventHandler = "sse.notify.push"
  PROBEREQ     EventHandler = "sse.notify.probereq"
  BANISHMENT   EventHandler = "sse.notify.banishment"
)

// defined event type for call
type EventType string
const (
  EVENT_ORIGINATED     EventType = "event.originated"
  EVENT_ALERTING       EventType = "event.alerting"
  EVENT_CONNECTED      EventType = "event.connected"
  EVENT_HOLD           EventType = "event.hold"
  EVENT_UNHOLD         EventType = "event.unhold"
  EVENT_DISCONNECTED   EventType = "event.disconnected"
  EVENT_TERMINATED     EventType = "event.terminated"
  EVENT_PARTYCHANGED   EventType = "event.partychanged"
  EVENT_UPDATEUSERDATA EventType = "event.updateuserdata"
  EVENT_RECORDSTART    EventType = "event.recordstart"
  EVENT_RECORDSTOP     EventType = "event.recordstop"
)

// defined event type for routing
type EventTypeRouting string
const (
  EVENT_ROUTING       EventTypeRouting = "event.routing"
  EVENT_ROUTECANCELED EventTypeRouting = "event.routecanceled"
  EVENT_SELECTED      EventTypeRouting = "event.selected"
  EVENT_ROUTED        EventTypeRouting = "event.routed"
)

// defined event type for user state
type EventTypeUser string
const (
  EVENT_USERSTATECHANGED  EventTypeUser = "event.userstatechanged"
  EVENT_USERREASONCHANGED EventTypeUser = "event.userreasonchanged"
)

// defined event type for phone
type EventTypePhone string
const (
  EVENT_REGISTERED   EventType = "event.registered"
  EVENT_UNREGISTERED EventType = "event.unregistered"
  EVENT_BUSY         EventType = "event.busy"
  EVENT_IDLE         EventType = "event.idle"
)

// defined event type for flow
type EventTypeFlow string
const (
  EVENT_FLOWSTART  EventTypeFlow = "event.flowstart"
  EVENT_FLOWEND    EventTypeFlow = "event.flowend"
  EVENT_FLOWCHANGE EventTypeFlow = "event.flowchange"
)

// defined event type for tracking
type EventTypeTracking string
const (
  EVENT_SEGMENTSTART EventTypeTracking = "event.segmentstart"
  EVENT_SEGMENTEND   EventTypeTracking = "event.segmentend"
  EVENT_ACTION       EventTypeTracking = "event.action"
)


// defined main parent event struct
type event struct {
  Handler handler
  Type    _type
}

// defined sse notify event handler name
type handler struct {
  Register   EventHandler
  Registered EventHandler
  Push       EventHandler
  ProbeReq   EventHandler
  Banishment EventHandler
}

type _type struct {
  Call     call
  Routing  routing
  User     user
  Phone    phone
  Flow     flow
  Tracking tracking
}


type call struct {
  Originated     EventType
  Alerting       EventType
  Connected      EventType
  Hold           EventType
  Unhold         EventType
  Disconnected   EventType
  Terminated     EventType
  PartyChanged   EventType
  UpdateUserData EventType
  RecordStart    EventType
  RecordStop     EventType
}

type routing struct {
  Routing       EventTypeRouting
  RouteCanceled EventTypeRouting
  Selected      EventTypeRouting
  Routed        EventTypeRouting
}

type user struct {
  UserStateChanged  EventTypeUser
  UserReasonChanged EventTypeUser
}

type phone struct {
  Registered   EventType
  Unregistered EventType
  Busy         EventType
  Idle         EventType
}

type flow struct {
  FlowStart  EventTypeFlow
  FlowEnd    EventTypeFlow
  FlowChange EventTypeFlow
}

type tracking struct {
  SegmentStart EventTypeTracking
  SegmentEnd   EventTypeTracking
  Action       EventTypeTracking
}


// defined code value for event name
var Event = event {
  Handler: handler {
    Register:   REGISTER,
    Registered: REGISTERED,
    Push:       PUSH,
    ProbeReq:   PROBEREQ,
    Banishment: BANISHMENT,
  },
  Type: _type {
    Call: call {
      Originated:     EVENT_ORIGINATED,
      Alerting:       EVENT_ALERTING,
      Connected:      EVENT_CONNECTED,
      Hold:           EVENT_HOLD,
      Unhold:         EVENT_UNHOLD,
      Disconnected:   EVENT_DISCONNECTED,
      Terminated:     EVENT_TERMINATED,
      PartyChanged:   EVENT_PARTYCHANGED,
      UpdateUserData: EVENT_UPDATEUSERDATA,
      RecordStart:    EVENT_RECORDSTART,
      RecordStop:     EVENT_RECORDSTOP,
    },
    Routing: routing {
      Routing:       EVENT_ROUTING,
      RouteCanceled: EVENT_ROUTECANCELED,
      Selected:      EVENT_SELECTED,
      Routed:        EVENT_ROUTED,
    },
    User: user {
      UserStateChanged:  EVENT_USERSTATECHANGED,
      UserReasonChanged: EVENT_USERREASONCHANGED,
    },
    Phone: phone {
      Registered:   EVENT_REGISTERED,
      Unregistered: EVENT_UNREGISTERED,
      Busy:         EVENT_BUSY,
      Idle:         EVENT_IDLE,
    },
    Flow: flow {
      FlowStart:  EVENT_FLOWSTART,
      FlowEnd:    EVENT_FLOWEND,
      FlowChange: EVENT_FLOWCHANGE,
    },
    Tracking: tracking {
      SegmentStart: EVENT_SEGMENTSTART,
      SegmentEnd:   EVENT_SEGMENTEND,
      Action:       EVENT_ACTION,
    },
  },
}


