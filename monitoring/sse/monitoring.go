package sse

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/ipron-ne/client-sdk-go/utils"
)

// Constants
const (
	apiPrefix  = "/webapi/sse"
	apiModule  = "/monitoring"
	apiVersion = "/v1"
)

const (
	apiName = apiPrefix + apiVersion + apiModule
)

type SSE struct {
	types.Client
}

func NewFromClient(client types.Client) *SSE {
	return &SSE{
		Client: client,
	}
}

/**
 * [요청]
 * GET http://100.100.103.160/webapi/sse/v1/monitoring/{tenant-id}/datasource
 * Authorization: Bearer xxxxxx
 * Content-Type: application/json
 * 
 * [응답]
 * {
 *     "result": true,
 *     "code": "0",
 *     "status": 200,
 *     "title": "success",
 *     "msg": "success",
 *     "data": [
 *         {
 *             "datasetName": "dnis",
 *             "dataSource": "kafka",
 *             "redisKey": "REALTIME:DNIS:tenantId:dnis:mediaType",
 *             "jsonData": "..."
 *         },
 * 
 **/
func (c *SSE) GetDatasource(params utils.Param) (*types.Response, error) {
	url := fmt.Sprintf("%s/%s/datasource", apiName, params.Get("tntId"))
	resp, err := c.GetRequest().Get(url, nil)
	return resp, errors.Wrap(err, "GetDatasource")
}

/**
 * [요청]
 * GET http://100.100.103.160/webapi/sse/v1/monitoring/{tenant-id}/dataset
 * Authorization: Bearer xxxxxx
 * Content-Type: application/json
 *
 * [응답]
 * {
 *     "result": true,
 *     "code": "0",
 *     "status": 200,
 *     "title": "success",
 *     "msg": "success",
 *     "data": {
 *         "name": "dataset",
 *         "dataset": [
 *             "tenant", 
 *             "dnis", 
 *             "flow", 
 *             "scenario", 
 *             "menu", 
 *             "servicecode", 
 *             "queue", 
 *             "group", 
 *             "user", 
 *             "interaction"
 *         ]
 *     }
 * }
 * 
 * params  : tntId
 * 
 **/
func (c *SSE) GetDatasets(params utils.Param) (*types.Response, error) {
	url := fmt.Sprintf("%s/%s/dataset", apiName, params.Get("tntId"))
	resp, err := c.GetRequest().Get(url, nil)
	return resp, errors.Wrap(err, "GetDatasets")
}

/**
 * [요청]
 * GET http://100.100.103.160/webapi/sse/v1/monitoring/{tenant-id}/dataset/{dataset}
 * Authorization: Bearer xxxxxx
 * Content-Type: application/json
 *
 * [응답]
 * {
 *     "result": true,
 *     "code": "0",
 *     "status": 200,
 *     "title": "success",
 *     "msg": "success",
 *     "data": {
 *         "flow": [
 *             {"desc": "", "name": "", "pk": false, "type": "string"},
 *             {"desc": "", "name": "", "pk": false, "type": "string"},
 *         ]
 *     }
 * }
 * 
 * dataset : dnis / group / interaction / ivr / menu / queue / scenario / serviceCode / tenant / user
 * params  : tntId
 * 
 **/
func (c *SSE) GetDataset(dataset string, params utils.Param) (*types.Response, error) {
	url := fmt.Sprintf("%s/%s/dataset/%s", apiName, params.Get("tntId"), dataset)
	resp, err := c.GetRequest().Get(url, nil)
	return resp, errors.Wrap(err, "GetDataset")
}

/**
 * [요청]
 * POST http://100.100.103.160/webapi/sse/v1/monitoring/{tenant-id}/filterValue
 * Authorization: Bearer xxxxxx
 * Content-Type: application/json
 *
 * [응답]
 * {
 *     "result": true,
 *     "code": "0",
 *     "status": 200,
 *     "title": "success",
 *     "msg": "success",
 *     "data": "MONITORING:FILTER:{filter-key}"
 * }
 **/
func (c *SSE) FetchFilterKey(token string, params utils.Param) (*types.Response, error) {
	url := fmt.Sprintf("%s/%s/filterValue", apiName, params.Get("tntId"))
	resp, err := c.GetRequest().Post(url, params)
	return resp, errors.Wrap(err, "FetchFilterKey")
}

/**
 * [요청]
 * GET http://100.100.103.160/webapi/sse/v1/monitoring/{tenant-id}/dnis?filterKey=MONITORING:FILTER:{filter-key}&bcloudToken={token}
 * Accept: text/event-stream
 * 
 * [응답]
 * <EventStream>
 * 
 * 
 **/
// dataset : dnis / group / interaction / ivr / menu / queue / scenario / serviceCode / tenant / user
func (c *SSE) GetEventSource(dataset string, params utils.Param) (*utils.EventSubscription, error) {
	resp, err := c.FetchFilterKey(dataset, params)
	if err != nil {
		return nil, errors.Wrap(err, "GetEventSource")
	}
	if resp.Code != 0 {
		return nil, errors.Errorf("code=%s", resp.Code )
	}

	filter := resp.Data.(string)
	if c.GetToken() != "" {
		filter = filter + "&bcloudToken=" + c.GetToken() 
	}

	sseURL := fmt.Sprintf("%s%s/%s/%s?filterKey=%s", c.GetBaseURL(), apiName, params.Get("tntId"), dataset, filter)
	eventSubs, err := utils.NewEventSubscription(sseURL, "")

	return eventSubs, nil
}
