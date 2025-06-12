package presence

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/action/presence"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Presence struct {
	types.Client
}

func NewFromClient(client types.Client) *Presence {
	return &Presence{
		Client: client,
	}
}

func (c *Presence) UserLogin(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string) (types.UserLoginResponse, error) {
	var respData types.UserLoginResponse

	url := fmt.Sprintf("%s/login/%s/%s", API_NAME, tntID, userID)
	body := map[string]any{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
		"dn":       dn,
	}

	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "UserLogin")
}

func (c *Presence) UserLogout(tntID, userID string, mediaSet []code.MediaType, cause code.AgentStateCauseType) (types.UserLogoutResponse, error) {
	var respData types.UserLogoutResponse

	url := fmt.Sprintf("%s/logout/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "UserLogout")
}

func (c *Presence) SetUserState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (types.SetUserStateResponse, error) {
	var respData types.SetUserStateResponse

	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserState")
}

func (c *Presence) GetUserState(tntID, userID, mediaType code.MediaType) (types.GetUserStateResponse, error) {
	var respData types.GetUserStateResponse

	url := fmt.Sprintf("%s/state/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUserState")
}

func (c *Presence) GetUserStateMultiMedia(tntID, userID string, mediaSet []code.MediaType) (types.GetUserStateMultiMediaResponse, error) {
	var respData types.GetUserStateMultiMediaResponse

	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
	}
	resp, err := c.GetRequest().Post(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUserStateMultiMedia")
}

func (c *Presence) Routeable(tntID, userID, mediaType string) (types.RouteableResponse, error) {
	var respData types.RouteableResponse

	url := fmt.Sprintf("%s/routeable/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Routeable")
}

func (c *Presence) SetUserAfterState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (types.SetUserAfterStateResponse, error) {
	var respData types.SetUserAfterStateResponse

	url := fmt.Sprintf("%s/afterstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Put(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserAfterState")
}

func (c *Presence) SetUserRecallState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (types.SetUserRecallStateResponse, error) {
	var respData types.SetUserRecallStateResponse

	url := fmt.Sprintf("%s/recallstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Put(url, body)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserRecallState")
}

func (c *Presence) GetQueueStatus(tntID, queueDnSet, mediaType code.MediaType) (types.GetQueueStatusResponse, error) {
	var respData types.GetQueueStatusResponse

	url := fmt.Sprintf("%s/queuestatus/%s/%s/%s", API_NAME, tntID, queueDnSet, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetQueueStatus")
}

func (c *Presence) GetSkillStatus(tntID, skillIDSet, mediaType code.MediaType) (types.GetSkillStatusResponse, error) {
	var respData types.GetSkillStatusResponse

	url := fmt.Sprintf("%s/skillstatus/%s/%s/%s", API_NAME, tntID, skillIDSet, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetSkillStatus")
}

func (c *Presence) GetUsersByStateAndMedia(tntID, state, media code.MediaType) (types.GetUsersByStateAndMediaResponse, error) {
	var respData types.GetUsersByStateAndMediaResponse

	url := fmt.Sprintf("%s/userlist/%s/%s/%s", API_NAME, tntID, state, media)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUsersByStateAndMedia")
}

func (c *Presence) GetQueuesByUserId(tntID, userID string) (types.GetQueuesByUserIdResponse, error) {
	var respData types.GetQueuesByUserIdResponse

	url := fmt.Sprintf("%s/userqueuelist/%s/%s", API_NAME, tntID, userID)
	resp, err := c.GetRequest().Get(url, nil)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUsersByStateAndMedia")
}
