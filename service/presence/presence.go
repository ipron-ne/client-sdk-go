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

func (c *Presence) UserLogin(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string) (UserLoginResponse, error) {
	var respData UserLoginResponse

	url := fmt.Sprintf("%s/login/%s/%s", API_NAME, tntID, userID)
	body := map[string]any{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
		"dn":       dn,
	}

	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "UserLogin")
}

func (c *Presence) UserLogout(tntID, userID string, mediaSet []code.MediaType, cause code.AgentStateCauseType) (UserLogoutResponse, error) {
	var respData UserLogoutResponse

	url := fmt.Sprintf("%s/logout/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "UserLogout")
}

func (c *Presence) SetUserState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (SetUserStateResponse, error) {
	var respData SetUserStateResponse

	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserState")
}

func (c *Presence) GetUserState(tntID, userID, mediaType code.MediaType) (GetUserStateResponse, error) {
	var respData GetUserStateResponse

	url := fmt.Sprintf("%s/state/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUserState")
}

func (c *Presence) GetUserStateMultiMedia(tntID, userID string, mediaSet []code.MediaType) (GetUserStateMultiMediaResponse, error) {
	var respData GetUserStateMultiMediaResponse

	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUserStateMultiMedia")
}

func (c *Presence) Routeable(tntID, userID, mediaType string) (RouteableResponse, error) {
	var respData RouteableResponse

	url := fmt.Sprintf("%s/routeable/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "Routeable")
}

func (c *Presence) SetUserAfterState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (SetUserAfterStateResponse, error) {
	var respData SetUserAfterStateResponse

	url := fmt.Sprintf("%s/afterstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Put(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserAfterState")
}

func (c *Presence) SetUserRecallState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (SetUserRecallStateResponse, error) {
	var respData SetUserRecallStateResponse

	url := fmt.Sprintf("%s/recallstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	resp, err := c.GetRequest().Put(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetUserRecallState")
}

func (c *Presence) GetQueueStatus(tntID, queueDnSet, mediaType code.MediaType) (GetQueueStatusResponse, error) {
	var respData GetQueueStatusResponse

	url := fmt.Sprintf("%s/queuestatus/%s/%s/%s", API_NAME, tntID, queueDnSet, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetQueueStatus")
}

func (c *Presence) GetSkillStatus(tntID, skillIDSet, mediaType code.MediaType) (GetSkillStatusResponse, error) {
	var respData GetSkillStatusResponse

	url := fmt.Sprintf("%s/skillstatus/%s/%s/%s", API_NAME, tntID, skillIDSet, mediaType)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetSkillStatus")
}

func (c *Presence) GetUsersByStateAndMedia(tntID, state, media code.MediaType) (GetUsersByStateAndMediaResponse, error) {
	var respData GetUsersByStateAndMediaResponse

	url := fmt.Sprintf("%s/userlist/%s/%s/%s", API_NAME, tntID, state, media)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUsersByStateAndMedia")
}

func (c *Presence) GetQueuesByUserId(tntID, userID string) (GetQueuesByUserIdResponse, error) {
	var respData GetQueuesByUserIdResponse

	url := fmt.Sprintf("%s/userqueuelist/%s/%s", API_NAME, tntID, userID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetUsersByStateAndMedia")
}

// Supported v1.2
func (c *Presence) ForceUpdateState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (ForceUpdateStateResponse, error) {
	var respData ForceUpdateStateResponse

	url := fmt.Sprintf("%s/forceupdatestate/%s/%s", API_NAME, tntID, userID)
	body := ForceUpdateStateRequest{
		Mediaset: mediaSet,
		State:    state,
		Cause:    cause,
	}
	resp, err := c.GetRequest().Put(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "ForceUpdateState")
}

// Supported v1.2
func (c *Presence) GetCallInfo(tntID, userID string) (GetCallInfoResponse, error) {
	var respData GetCallInfoResponse

	url := fmt.Sprintf("%s/callinfo/%s/%s", API_NAME, tntID, userID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetCallInfo")
}

// Supported v1.2
func (c *Presence) GetDoNotDistube(tntID, userID string, mediaSet []code.MediaType) (GetDoNotDistubeResponse, error) {
	var respData GetDoNotDistubeResponse

	url := fmt.Sprintf("%s/getdnd/%s/%s", API_NAME, tntID, userID)
	body := GetDoNotDistubeRequest{
		Mediaset: mediaSet,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetDoNotDistube")
}

// Supported v1.2
func (c *Presence) SetDoNotDistube(tntID, userID string, mediaSet []code.MediaType, enable bool) (SetDoNotDistubeResponse, error) {
	var respData SetDoNotDistubeResponse

	url := fmt.Sprintf("%s/dnd/%s/%s", API_NAME, tntID, userID)
	body := SetDoNotDistubeRequest{
		Mediaset: mediaSet,
		Enable:   enable,
	}
	resp, err := c.GetRequest().Put(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetDoNotDistube")
}

// Supported v1.2
func (c *Presence) GetMediaReady(tntID, userID string, mediaSet []code.MediaType) (GetMediaReadyResponse, error) {
	var respData GetMediaReadyResponse

	url := fmt.Sprintf("%s/getmediaready/%s/%s", API_NAME, tntID, userID)
	body := GetMediaReadyRequest{
		Mediaset: mediaSet,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetMediaReady")
}

// Supported v1.2
func (c *Presence) SetMediaReady(tntID, userID string, mediaSet []code.MediaType, enable bool) (SetMediaReadyResponse, error) {
	var respData SetMediaReadyResponse

	url := fmt.Sprintf("%s/mediaready/%s/%s", API_NAME, tntID, userID)
	body := SetMediaReadyRequest{
		Mediaset: mediaSet,
		Enable:   enable,
	}
	resp, err := c.GetRequest().Put(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetDoNotDistube")
}

// Supported v1.2
func (c *Presence) SetReserveAniUEI(tntID, userID string, uei string) (SetReserveAniUEIResponse, error) {
	var respData SetReserveAniUEIResponse

	url := fmt.Sprintf("%s/reserve-ani-uei/%s/%s", API_NAME, tntID, userID)
	body := SetReserveAniUEIRequest{
		UEI: uei,
	}
	resp, err := c.GetRequest().Post(url, body)
	err = types.GetServiceError(resp, err)
	if err == nil {
		resp.ServiceUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "SetReserveAniUEI")
}
