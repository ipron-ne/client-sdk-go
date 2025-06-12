package types

type AgentInfo struct {
	ID                string `json:"_id"`
	TntId             string `json:"tntId"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	AuthLevel         string `json:"authLevel"`
	AccessAuthID      string `json:"accessAuthId"`
	GroupID           string `json:"groupId"`
	Extension         string `json:"extension"`
	DidNum            string `json:"didNum"`
	DidPortID         string `json:"didPortId"`
	PhoneID           string `json:"phoneId"`
	DefaultSkillID    string `json:"defaultSkillId"`
	Enable            string `json:"enable"`
	ScheduleEnable    string `json:"scheduleEnable"`
	ServiceTemplateID string `json:"serviceTemplateId"`
	OutgoingBlock     string `json:"outgoingBlock"`
	IncomingBlock     string `json:"incomingBlock"`
	ForwardUse        string `json:"forwardUse"`
	ReleaseToneUse    string `json:"releaseToneUse"`
	TransferedToneUse string `json:"transferedToneUse"`
	EnableCallWait    string `json:"enableCallWait"`
	ScheduleType      string `json:"scheduleType"`
	ForwardType       string `json:"forwardType"`
	NoAnswerSec       string `json:"noAnswerSec"`
	TransferNum       string `json:"transferNum"`
	Monitor           string `json:"monitor"`
	Coaching          string `json:"coaching"`
	AvoidMonitor      string `json:"avoidMonitor"`
	RecEnable         string `json:"recEnable"`
	RecAutoEnable     string `json:"recAutoEnable"`
	LastLoginDate     string `json:"lastLoginDate"`
	CreateDate        string `json:"createDate"`
	GroupName         string `json:"groupName"`
	UserStatus        string `json:"userStatus"`
}

type QueueInfo struct {
	ID                 string `json:"_id"`
	TntID              string `json:"tntId"`
	Name               string `json:"name"`
	Extension          string `json:"extension"`
	Desc               string `json:"desc"`
	RouteKind          string `json:"routeKind"`
	AgentChoiceMethod  string `json:"agentChoiceMethod"`
	ScheduleID         string `json:"scheduleId"`
	Enable             string `json:"enable"`
	NoAnswerSec        string `json:"noAnswerSec"`
	MaxWaitCalls       string `json:"maxWaitCalls"`
	MaxWaitSec         string `json:"maxWaitSec"`
	MinAbandonSec      string `json:"minAbandonSec"`
	DefaultQueueFlowID string `json:"defaultQueueFlowId"`
	SvcStandardSec     string `json:"svcStandardSec"`
	SvcEnable          string `json:"svcEnable"`
	SvcGoalRate        string `json:"svcGoalRate"`
	InitPromptID       string `json:"initPromptId"`
	WaitPromptID       string `json:"waitPromptId"`
	BlockPromptID      string `json:"blockPromptId"`
	CreateDate         string `json:"createDate"`
}

type GetGroupListResponse struct {
	Name     string   `json:"name"`
	DidNum   string   `json:"didNum"`
	ParentID string   `json:"parentId"`
	Enable   string   `json:"enable"`
	ID       string   `json:"_id"`
	TntID    string   `json:"tntId"`
	Medias   []string `json:"medias"`
	Skills   []string `json:"skills"`
}

type GetGroupInfoResponse struct {
	Name     string `json:"name"`
	DidNum   string `json:"didNum"`
	ParentID string `json:"parentId"`
	Enable   string `json:"enable"`
	ID       string `json:"_id"`
	TntID    string `json:"tntId"`
}

type GetAllAgentListResponse struct {
	AgentInfo
}

type GetAgentListResponse struct {
	AgentInfo
}

type GetAgentInfoResponse struct {
	AgentInfo
}

type GetQueueListResponse struct {
	QueueInfo
}

type GetQueueInfoResponse struct {
	QueueInfo
}

type GetFlowListResponse struct {
	ID               string `json:"_id"`
	ConcurrentCh     int    `json:"concurrentCh"`
	EditLock         bool   `json:"editLock"`
	EditType         string `json:"editType"`
	EditUserID       string `json:"editUserId"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	RsvEnable        bool   `json:"rsvEnable"`
	RsvStatus        string `json:"rsvStatus"`
	RsvTargetTime    string `json:"rsvTargetTime"`
	RsvVersion       string `json:"rsvVersion"`
	SID              string `json:"sId"`
	ServiceOptionFdt string `json:"serviceOptionFdt"`
	ServiceOptionIdt string `json:"serviceOptionIdt"`
	Status           string `json:"status"`
	TntID            string `json:"tntId"`
	Type             string `json:"type"`
	Version          string `json:"version"`
	Versions         string `json:"versions"`
}
