package callapi

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/service"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/call"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

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

func MakeCall(tntID, userID, callID, ani, dnis, userANI, mediaType string) (*service.Response, error) {
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

	return service.GetApiClient().Post(url, body)
}

func MakeCallEx(tntID, userID, ani, dnis, userANI, mediaType, uei, uui string, routeOption RouteOption) (*service.Response, error) {
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

	return service.GetApiClient().Post(url, body)
}

func Answer(tntID, callID, connID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/answer/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return service.GetApiClient().Post(url, body)
}

func ReleaseCall(tntID, callID, connID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/release/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return service.GetApiClient().Post(url, body)
}

func Hold(tntID, callID, connID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/hold/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return service.GetApiClient().Post(url, body)
}

func Unhold(tntID, callID, connID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/unhold/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"conn_id": connID,
	}

	return service.GetApiClient().Post(url, body)
}

func SingleStepTransfer(tntID, callID, connID, dnis, userANI, uui, uei string, routeOption RouteOption) (*service.Response, error) {
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

	return service.GetApiClient().Post(url, body)
}

func MuteTransfer(tntID, holdCallID, holdConnID, activeCallID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/mute-transfer/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":       tntID,
		"hold_call_id": holdCallID,
		"hold_conn_id": holdConnID,
		"active_call_id": activeCallID,
	}

	return service.GetApiClient().Post(url, body)
}

func SingleStepConference(tntID, callID, connID, dnis, userANI, uui, uei, partyType string) (*service.Response, error) {
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

	return service.GetApiClient().Post(url, body)
}

func MuteConference(tntID, holdCallID, holdConnID, activeCallID, partyType string) (*service.Response, error) {
	url := fmt.Sprintf("%s/mute-conference/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":        tntID,
		"hold_call_id":  holdCallID,
		"hold_conn_id":  holdConnID,
		"active_call_id": activeCallID,
		"party_type":    partyType,
	}

	return service.GetApiClient().Post(url, body)
}

func JoinCall(tntID, userID, joinCallID, joinConnID, joinType string) (*service.Response, error) {
	url := fmt.Sprintf("%s/join/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":       tntID,
		"user_id":      userID,
		"join_call_id": joinCallID,
		"join_conn_id": joinConnID,
		"join_type":    joinType,
	}

	return service.GetApiClient().Post(url, body)
}

func Route(tntID, callID, partID, partType, ani, dnis, callbackURI string, timeout int, retryCall, autoAnswer bool) (*service.Response, error) {
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

	return service.GetApiClient().Post(url, body)
}

func Numberplan(tntID, siteID, dnis string) (*service.Response, error) {
	url := fmt.Sprintf("%s/numberplan/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"site_id": siteID,
		"dnis":    dnis,
	}

	return service.GetApiClient().Post(url, body)
}

func SetUserdata(tntID, callID, uei, uui string) (*service.Response, error) {
	url := fmt.Sprintf("%s/userdata/%s", API_NAME, tntID)
	body := map[string]interface{}{
		"tenant":  tntID,
		"call_id": callID,
		"uei":     uei,
		"uui":     uui,
	}

	return service.GetApiClient().Post(url, body)
}

func GetUserdata(tntID, callID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/userdata/%s/%s", API_NAME, tntID, callID)
	return service.GetApiClient().Get(url, nil)
}
