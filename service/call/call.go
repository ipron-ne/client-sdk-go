package call

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/call"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Call struct {
	types.Client
}

func NewFromClient(client types.Client) *Call {
	return &Call{
		Client: client,
	}
}


type RouteOption struct {
	Type               int    `json:"type"`
	Priority           int    `json:"priority"`
	RelationMethod     int    `json:"relation_method"`
	RelationAgentID    string `json:"relation_agent_id"`
	RelationTimeout    int    `json:"relation_timeout"`
	Method             int    `json:"method"`
	SkillID            string `json:"skill_id"`
	SkillLevel         int    `json:"skill_level"`
	GroupID            string `json:"group_id"`
}

func (c *Call) MakeCall(tntID, userID, callID, ani, dnis, userANI, mediaType string) (*types.Response, error) {
	url := fmt.Sprintf("%s/makecall/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":    tntID,
		"user_id":   userID,
		"call_id":   callID,
		"ani":       ani,
		"dnis":      dnis,
		"user_ani":  userANI,
		"media_type": mediaType,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) MakeCallEx(tntID, userID, ani, dnis, userANI, mediaType, uei, uui string, routeOption RouteOption) (*types.Response, error) {
	url := fmt.Sprintf("%s/makecallex/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":    tntID,
		"user_id":   userID,
		"ani":       ani,
		"dnis":      dnis,
		"user_ani":  userANI,
		"media_type": mediaType,
		"uei":       uei,
		"uui":       uui,
		"route_opt": routeOption,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) Answer(tntID, callID, connID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/answer/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) ReleaseCall(tntID, callID, connID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/release/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) Hold(tntID, callID, connID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/hold/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) Unhold(tntID, callID, connID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/unhold/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) SingleStepTransfer(tntID, callID, connID, dnis, userANI, uui, uei string, routeOption RouteOption) (*types.Response, error) {
	url := fmt.Sprintf("%s/singlestep-transfer/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":    tntID,
		"call_id":   callID,
		"conn_id":   connID,
		"dnis":      dnis,
		"user_ani":  userANI,
		"uui":       uui,
		"uei":       uei,
		"route_opt": routeOption,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) MuteTransfer(tntID, holdCallID, holdConnID, activeCallID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/mute-transfer/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":       tntID,
		"hold_call_id": holdCallID,
		"hold_conn_id": holdConnID,
		"active_call_id": activeCallID,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) SingleStepConference(tntID, callID, connID, dnis, userANI, uui, uei, partyType string) (*types.Response, error) {
	url := fmt.Sprintf("%s/singlestep-conference/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":     tntID,
		"call_id":    callID,
		"conn_id":    connID,
		"dnis":       dnis,
		"user_ani":   userANI,
		"uui":        uui,
		"uei":        uei,
		"party_type": partyType,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) MuteConference(tntID, holdCallID, holdConnID, activeCallID, partyType string) (*types.Response, error) {
	url := fmt.Sprintf("%s/mute-conference/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":        tntID,
		"hold_call_id":  holdCallID,
		"hold_conn_id":  holdConnID,
		"active_call_id": activeCallID,
		"party_type":    partyType,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) JoinCall(tntID, userID, joinCallID, joinConnID, joinType string) (*types.Response, error) {
	url := fmt.Sprintf("%s/join/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":       tntID,
		"user_id":      userID,
		"join_call_id": joinCallID,
		"join_conn_id": joinConnID,
		"join_type":    joinType,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) Route(tntID, callID, partID, partType, ani, dnis, callbackURI string, timeout int, retryCall, autoAnswer bool) (*types.Response, error) {
	url := fmt.Sprintf("%s/route/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":        tntID,
		"call_id":       callID,
		"part_id":       partID,
		"part_type":     partType,
		"ani":           ani,
		"dnis":          dnis,
		"callback_uri":  callbackURI,
		"timeout":       timeout,
		"retry_call":    retryCall,
		"auto_answer":   autoAnswer,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) Numberplan(tntID, siteID, dnis string) (*types.Response, error) {
	url := fmt.Sprintf("%s/numberplan/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"site_id": siteID,
		"dnis":    dnis,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) SetUserdata(tntID, callID, uei, uui string) (*types.Response, error) {
	url := fmt.Sprintf("%s/userdata/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"uei":     uei,
		"uui":     uui,
	}

	return c.GetRequest().Post(url, body)
}

func (c *Call) GetUserdata(tntID, callID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/userdata/%s/%s", API_NAME, tntID, callID)
	return c.GetRequest().Get(url, nil)
}
