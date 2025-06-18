package info

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/pkg/errors"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/api"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

type Info struct {
	types.Client
}

func NewFromClient(client types.Client) *Info {
	return &Info{
		Client: client,
	}
}

// GetGroupList retrieves a list of groups for a specific tenant.
func (c *Info) GetGroupList(tenantID string) ([]GetGroupListResponse, error) {
	var respData []GetGroupListResponse

	url := fmt.Sprintf("%s/groups/%s", API_NAME, tenantID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetGroupList")
}

// GetGroupInfo retrieves information about a specific group.
func (c *Info) GetGroupInfo(tenantID, groupID string) (GetGroupListResponse, error) {
	var respData GetGroupListResponse

	url := fmt.Sprintf("%s/group/%s/%s", API_NAME, tenantID, groupID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetGroupInfo")
}

// GetAllAgentList retrieves a list of all agents for a specific tenant.
func (c *Info) GetAllAgentList(tenantID string) ([]GetAllAgentListResponse, error) {
	var respData []GetAllAgentListResponse

	url := fmt.Sprintf("%s/users/%s", API_NAME, tenantID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetAllAgentList")
}

// GetAgentList retrieves a list of agents for a specific tenant and group.
func (c *Info) GetAgentList(tenantID, groupID string) ([]GetAgentListResponse, error) {
	var respData []GetAgentListResponse

	url := fmt.Sprintf("%s/users/%s?groupId=%s", API_NAME, tenantID, groupID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetAgentList")
}

// GetAgentInfo retrieves information about a specific agent.
func (c *Info) GetAgentInfo(tenantID, userID string) (GetAgentInfoResponse, error) {
	var respData GetAgentInfoResponse

	url := fmt.Sprintf("%s/user/%s/%s", API_NAME, tenantID, userID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetAgentInfo")
}

// GetQueueList retrieves a list of queues for a specific tenant.
func (c *Info) GetQueueList(tenantID string) ([]GetQueueListResponse, error) {
	var respData []GetQueueListResponse

	url := fmt.Sprintf("%s/queues/%s", API_NAME, tenantID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetQueueList")
}

// GetQueueInfo retrieves information about a specific queue.
func (c *Info) GetQueueInfo(tenantID, queueID string) (GetQueueInfoResponse, error) {
	var respData GetQueueInfoResponse

	url := fmt.Sprintf("%s/queue/%s/%s", API_NAME, tenantID, queueID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetQueueInfo")
}

func (c *Info) GetFlowList(tenantID string) ([]GetFlowListResponse, error) {
	var respData []GetFlowListResponse

	url := fmt.Sprintf("%s/flows/%s", API_NAME, tenantID)
	resp, err := c.GetRequest().Get(url, nil)
	err = types.GetBackendError(resp, err)
	if err == nil {
		resp.DataUnmarshal(&respData)
	}

	return respData, errors.Wrap(err, "GetFlowList")
}

// Placeholder for future functions

// func (c *Info) GetStateSubcode(queueID, type string) (*types.Response, error) {
// 	// TODO: Implement this function.
// }

// func (c *Info) GetAgentQueueList(userID string) (*types.Response, error) {
// 	// TODO: Implement this function.
// }
