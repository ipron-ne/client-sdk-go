package presence

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/v1/service"
	"github.com/ipron-ne/client-sdk-go/v1/code"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/presence"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

func UserLogin(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string) (*service.Response, error) {
	url := fmt.Sprintf("%s/login/%s/%s", API_NAME, tntID, userID)
	body := map[string]any{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
		"dn":       dn,
	}
	return service.GetApiClient().Post(url, body)
}

func UserLogout(tntID, userID string, mediaSet []code.MediaType, cause code.AgentStateCauseType) (*service.Response, error) {
	url := fmt.Sprintf("%s/logout/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"cause":    cause,
	}
	return service.GetApiClient().Post(url, body)
}

func SetUserState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*service.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return service.GetApiClient().Put(url, body)
}

func GetUserState(tntID, userID, mediaType code.MediaType) (*service.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	return service.GetApiClient().Get(url, nil)
}

func GetUserStateMultiMedia(tntID, userID string, mediaSet []code.MediaType) (*service.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
	}
	return service.GetApiClient().Post(url, body)
}

func Routeable(tntID, userID, mediaType string) (*service.Response, error) {
	url := fmt.Sprintf("%s/routeable/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	return service.GetApiClient().Get(url, nil)
}

func SetUserAfterState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*service.Response, error) {
	url := fmt.Sprintf("%s/afterstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return service.GetApiClient().Put(url, body)
}

func SetUserRecallState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*service.Response, error) {
	url := fmt.Sprintf("%s/recallstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return service.GetApiClient().Put(url, body)
}

func GetQueueStatus(tntID, queueDnSet, mediaType code.MediaType) (*service.Response, error) {
	url := fmt.Sprintf("%s/queuestatus/%s/%s/%s", API_NAME, tntID, queueDnSet, mediaType)
	return service.GetApiClient().Get(url, nil)
}

func GetSkillStatus(tntID, skillIDSet, mediaType code.MediaType) (*service.Response, error) {
	url := fmt.Sprintf("%s/skillstatus/%s/%s/%s", API_NAME, tntID, skillIDSet, mediaType)
	return service.GetApiClient().Get(url, nil)
}

func GetUsersByStateAndMedia(tntID, state, media code.MediaType) (*service.Response, error) {
	url := fmt.Sprintf("%s/userlist/%s/%s/%s", API_NAME, tntID, state, media)
	return service.GetApiClient().Get(url, nil)
}

func GetQueuesByUserId(tntID, userID string) (*service.Response, error) {
	url := fmt.Sprintf("%s/userqueuelist/%s/%s", API_NAME, tntID, userID)
	return service.GetApiClient().Get(url, nil)
}
