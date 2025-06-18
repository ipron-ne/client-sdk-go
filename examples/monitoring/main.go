package main

import (
	"flag"
	"log"
	"os"

	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/info"

	"github.com/ipron-ne/client-sdk-go/monitoring/sse"
	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/ipron-ne/client-sdk-go/utils"
)

var (
	API_URL  = os.Getenv("IPRON_NE_API_URL")
	AppToken = os.Getenv("IPRON_NE_APPKEY")
)

func main() {
	TenantID := flag.String("tenantid", "", "Tenant ID")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithAppToken(AppToken),
		config.WithDebug(true),
		config.WithTenantID(*TenantID),
	)

	client := service.NewFromConfig(cfg)

	listDatasets(client)
	listDataset(client, "flow")
	listDatasource(client)

	// moniEventFlow(client, getFlowIDList(client))
	moniEventUser(client, getUserIDList(client))
}

func getFlowIDList(client types.Client) []string {
	info := info.NewFromClient(client)
	resp, err := info.GetFlowList(client.GetTenantID())
	if err != nil {
		log.Println(err)
	}

	flowList := []string{}
	for _, v := range resp {
		flowList = append(flowList, v.ID)
	}

	return flowList
}

func getUserIDList(client types.Client) []string {
	info := info.NewFromClient(client)
	resp, err := info.GetAllAgentList(client.GetTenantID())
	if err != nil {
		log.Println(err)
	}

	flowList := []string{}
	for _, v := range resp {
		flowList = append(flowList, v.ID)
	}

	return flowList
}

func listDatasets(client types.Client) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Datasets List]\n")
	resp, err := monitor.GetDatasets(utils.NewParam("tntId", client.GetTenantID()))
	if err != nil {
		log.Panic(err)
	}
	for _, item := range resp.Dataset {
		log.Printf("%+v\n", item)
	}
}

func listDataset(client types.Client, datasetName string) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Dataset List:%s]\n", datasetName)
	resp, err := monitor.GetDataset(datasetName, utils.NewParam("tntId", client.GetTenantID()))
	if err != nil {
		log.Panic(err)
	}
	dataset := resp
	for _, item := range dataset {
		log.Printf("[%s] %s\n", item.Name, item.Desc)
	}
}

func listDatasource(client types.Client) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Datasource]\n")
	resp, err := monitor.GetDatasource(utils.NewParam("tntId", client.GetTenantID()))
	if err != nil {
		log.Panic(err)
	}
	for _, dataset := range resp {
		setName := dataset.DatasetName
		fields := monitor.GetDatasourceFields(dataset)

		log.Printf("[%s] \n", setName)
		for _, item := range fields {
			log.Printf("  - %s [%s] %s\n", item.Name, item.Type, item.Desc)
		}
	}
}

func moniEventFlow(client types.Client, resource []string) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[EventListen]\n")

	params := map[string]any{
		"tntId": client.GetTenantID(),
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

	eventSubs, err := monitor.GetEventSource("flow", params)
	if err != nil {
		log.Panic(err)
	}
	eventSubs.OnMessage(func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		for _, v := range data.Array() {
			log.Printf("%+v\n", v.Object())
		}
	})

	eventSubs.EventLoop()
}

func moniEventUser(client types.Client, resource []string) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[EventListen]\n")

	params := map[string]any{
		"tntId": client.GetTenantID(),
		"colFilter": []string{
			"userId", "userName", "mediaType", "usersts1000", "usersts1020",
		},
		"rowFilter": map[string]any{
			"userId": resource,
			"mediaType": []string{
				"voice",
			},
		},
	}

	eventSubs, err := monitor.GetEventSource("user", params)
	if err != nil {
		log.Panic(err)
	}
	eventSubs.OnMessage(func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		for _, v := range data.Array() {
			log.Printf("%+v\n", v.Object())
		}
	})

	eventSubs.EventLoop()
}
