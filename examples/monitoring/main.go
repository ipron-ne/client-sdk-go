package main

import (
	"log"
	"os"

	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/info"

	"github.com/ipron-ne/client-sdk-go/monitoring"
	"github.com/ipron-ne/client-sdk-go/monitoring/sse"
	"github.com/ipron-ne/client-sdk-go/utils"
)

var (
	API_URL    = os.Getenv("IPRON_NE_URL")
	TenantID   = os.Getenv("IPRON_NE_TENANTID")
	AppToken   = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	params := map[string]any{"token": AppToken}
	monitoring.Init(API_URL, params, 0, true)

	listDatasets()
	listDataset("flow")
	listDatasource()

	service.Init(API_URL, 0, true)
	service.GetApiClient().SetToken(AppToken)
	service.GetApiClient().SetTenant(TenantID)


	moniEvent(getFlowIDList())
}

func getFlowIDList() []string {
	resp, err := info.GetFlowList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}

	flowList := []string{}
	for _, v := range resp.GetData().Array() {
		flowList = append(flowList, v.Get("_id").Str())
	}

	return flowList
}

func listDatasets() {
	log.Printf("\n\n[Datasets List]\n")
	resp, err := sse.GetDatasets(utils.NewParam("tntId", TenantID))
	if err != nil {
		log.Panic(err)
	}
	for _, item := range resp.GetData().Get("dataset").Array() {
		log.Printf("%+v\n", item.Str())
	}
}

func listDataset(datasetName string) {
	log.Printf("\n\n[Dataset List:%s]\n", datasetName)
	resp, err := sse.GetDataset(datasetName, utils.NewParam("tntId", TenantID))
	if err != nil {
		log.Panic(err)
	}
	dataset := resp.GetData().Get(datasetName).Array()
	for _, item := range dataset {
		obj := item.Object()
		log.Printf("[%s] %s\n", obj["name"].Str(), obj["desc"].Str())
	}
}


func listDatasource() {
	log.Printf("\n\n[Datasource]\n")
	resp, err := sse.GetDatasource(utils.NewParam("tntId", TenantID))
	if err != nil {
		log.Panic(err)
	}
	for _, dataset := range resp.GetData().Array() {
		setName := dataset.Get("datasetName")
		setData := dataset.Get("jsonData")
		jsonData := utils.JSONParse(setData.Str())

		log.Printf("[%s] \n", setName.Str())
		for _, item := range jsonData.Get(setName.Str()).Array() {
			log.Printf("  - %s [%s] %s\n", item.Get("name").Str(), item.Get("type").Str(), item.Get("desc").Str())
		}
	}
}


func moniEvent(resource []string) {
	log.Printf("\n\n[EventListen]\n")

	params := map[string]any{
		"tntId": "656d849405006b6f6092ab3d",
		"colFilter": []string{
			"flowName", "mediaType", "ivr1000", "ivr1010", "ivr1020", "ivr1040", "ivrmon1000",
		},
		"rowFilter": map[string]any{
			"flowId": resource,
			"mediaType": []string{
				"voice",
			},
		},
	}

	eventSubs, err := sse.GetEventSource("flow", params)
	if err != nil {
		log.Panic(err)
	}
	eventSubs.OnMessage(func(e utils.Event){
		data := utils.JSONParse(e.Data())
		for _, v := range data.Array() {
			log.Printf("%+v\n", v.Object())
		}
	})

	eventSubs.EventLoop()
}

