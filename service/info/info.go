package info

import (
	"fmt"
	"log"
	"net/http"
)

const (
	API_PREFIX  = "/webapi"
	API_MODULE  = "/api"
	API_VERSION = "/v1"
	API_NAME    = API_PREFIX + API_MODULE + API_VERSION
)

// GetGroupList retrieves a list of groups for a specific tenant.
func GetGroupList(tenantID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/groups/%s", API_NAME, tenantID)
	return api.GetApiClient().Get(url, nil)
}

// GetGroupInfo retrieves information about a specific group.
func GetGroupInfo(tenantID, groupID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/group/%s/%s", API_NAME, tenantID, groupID)
	return api.GetApiClient().Get(url, nil)
}

// GetAllAgentList retrieves a list of all agents for a specific tenant.
func GetAllAgentList(tenantID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/users/%s", API_NAME, tenantID)
	return api.GetApiClient().Get(url, nil)
}

// GetAgentList retrieves a list of agents for a specific tenant and group.
func GetAgentList(tenantID, groupID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/users/%s?groupId=%s", API_NAME, tenantID, groupID)
	return api.GetApiClient().Get(url, nil)
}

// GetAgentInfo retrieves information about a specific agent.
func GetAgentInfo(tenantID, userID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/user/%s/%s", API_NAME, tenantID, userID)
	return api.GetApiClient().Get(url, nil)
}

// GetQueueList retrieves a list of queues for a specific tenant.
func GetQueueList(tenantID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/queues/%s", API_NAME, tenantID)
	return api.GetApiClient().Get(url, nil)
}

// GetQueueInfo retrieves information about a specific queue.
func GetQueueInfo(tenantID, queueID string) (*http.Response, map[string]any, error) {
	url := fmt.Sprintf("%s/queue/%s/%s", API_NAME, tenantID, queueID)
	return api.GetApiClient().Get(url, nil)
}

// Placeholder for future functions

// func GetStateSubcode(queueID, type string) (*http.Response, map[string]any, error) {
// 	// TODO: Implement this function.
// }

// func GetAgentQueueList(userID string) (*http.Response, error) {
// 	// TODO: Implement this function.
// }
