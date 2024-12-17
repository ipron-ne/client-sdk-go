package presence

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/types"
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

func (c *Presence) UserLogin(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType, dn string) (*types.Response, error) {
	url := fmt.Sprintf("%s/login/%s/%s", API_NAME, tntID, userID)
	body := map[string]any{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
		"dn":       dn,
	}
	return c.GetRequest().Post(url, body)
}

func (c *Presence) UserLogout(tntID, userID string, mediaSet []code.MediaType, cause code.AgentStateCauseType) (*types.Response, error) {
	url := fmt.Sprintf("%s/logout/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"cause":    cause,
	}
	return c.GetRequest().Post(url, body)
}

func (c *Presence) SetUserState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*types.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return c.GetRequest().Put(url, body)
}

func (c *Presence) GetUserState(tntID, userID, mediaType code.MediaType) (*types.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	return c.GetRequest().Get(url, nil)
}

func (c *Presence) GetUserStateMultiMedia(tntID, userID string, mediaSet []code.MediaType) (*types.Response, error) {
	url := fmt.Sprintf("%s/state/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
	}
	return c.GetRequest().Post(url, body)
}

func (c *Presence) Routeable(tntID, userID, mediaType string) (*types.Response, error) {
	url := fmt.Sprintf("%s/routeable/%s/%s/%s", API_NAME, tntID, userID, mediaType)
	return c.GetRequest().Get(url, nil)
}

func (c *Presence) SetUserAfterState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*types.Response, error) {
	url := fmt.Sprintf("%s/afterstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return c.GetRequest().Put(url, body)
}

func (c *Presence) SetUserRecallState(tntID, userID string, mediaSet []code.MediaType, state code.AgentStateType, cause code.AgentStateCauseType) (*types.Response, error) {
	url := fmt.Sprintf("%s/recallstate/%s/%s", API_NAME, tntID, userID)
	body := map[string]interface{}{
		"mediaset": mediaSet,
		"state":    state,
		"cause":    cause,
	}
	return c.GetRequest().Put(url, body)
}

func (c *Presence) GetQueueStatus(tntID, queueDnSet, mediaType code.MediaType) (*types.Response, error) {
	url := fmt.Sprintf("%s/queuestatus/%s/%s/%s", API_NAME, tntID, queueDnSet, mediaType)
	return c.GetRequest().Get(url, nil)
}

func (c *Presence) GetSkillStatus(tntID, skillIDSet, mediaType code.MediaType) (*types.Response, error) {
	url := fmt.Sprintf("%s/skillstatus/%s/%s/%s", API_NAME, tntID, skillIDSet, mediaType)
	return c.GetRequest().Get(url, nil)
}

func (c *Presence) GetUsersByStateAndMedia(tntID, state, media code.MediaType) (*types.Response, error) {
	url := fmt.Sprintf("%s/userlist/%s/%s/%s", API_NAME, tntID, state, media)
	return c.GetRequest().Get(url, nil)
}

func (c *Presence) GetQueuesByUserId(tntID, userID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/userqueuelist/%s/%s", API_NAME, tntID, userID)
	return c.GetRequest().Get(url, nil)
}
