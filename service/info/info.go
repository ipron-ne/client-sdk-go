package info

import (
	"fmt"

	"github.com/ipron-ne/client-sdk-go/types"
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
func (c *Info) GetGroupList(tenantID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/groups/%s", API_NAME, tenantID)
	return c.GetRequest().Get(url, nil)
}

// GetGroupInfo retrieves information about a specific group.
func (c *Info) GetGroupInfo(tenantID, groupID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/group/%s/%s", API_NAME, tenantID, groupID)
	return c.GetRequest().Get(url, nil)
}

// GetAllAgentList retrieves a list of all agents for a specific tenant.
func (c *Info) GetAllAgentList(tenantID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/users/%s", API_NAME, tenantID)
	return c.GetRequest().Get(url, nil)
}

// GetAgentList retrieves a list of agents for a specific tenant and group.
func (c *Info) GetAgentList(tenantID, groupID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/users/%s?groupId=%s", API_NAME, tenantID, groupID)
	return c.GetRequest().Get(url, nil)
}

// GetAgentInfo retrieves information about a specific agent.
func (c *Info) GetAgentInfo(tenantID, userID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/user/%s/%s", API_NAME, tenantID, userID)
	return c.GetRequest().Get(url, nil)
}

// GetQueueList retrieves a list of queues for a specific tenant.
func (c *Info) GetQueueList(tenantID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/queues/%s", API_NAME, tenantID)
	return c.GetRequest().Get(url, nil)
}

// GetQueueInfo retrieves information about a specific queue.
func (c *Info) GetQueueInfo(tenantID, queueID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/queue/%s/%s", API_NAME, tenantID, queueID)
	return c.GetRequest().Get(url, nil)
}

func (c *Info) GetFlowList(tenantID string) (*types.Response, error) {
	url := fmt.Sprintf("%s/flows/%s", API_NAME, tenantID)
	return c.GetRequest().Get(url, nil)
}


// Placeholder for future functions

// func (c *Info) GetStateSubcode(queueID, type string) (*types.Response, error) {
// 	// TODO: Implement this function.
// }

// func (c *Info) GetAgentQueueList(userID string) (*types.Response, error) {
// 	// TODO: Implement this function.
// }
