# IPRON-NE 를 이용하기 위한 golang용 client SDK

본 SDK에는 IPRON-NE 의 콜제어 상태 모니터링을 위한 것입니다.

SDK 를 이용하기 위해서는 가입계정 또는 발급된 APP KEY 가 필요합니다.

## 콜제어 & 정보조회

### 초기화 & 로그인

```golang
import (
	"log"
	"os"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/service/info"
	"github.com/ipron-ne/client-sdk-go/utils"
)

func main() {
	var (
		API_URL    = os.Getenv("IPRON_NE_URL")
		UserID     = os.Getenv("USER_ID")
		Passwd     = os.Getenv("USER_PWD")
		TenantName = os.Getenv("TENANT_NAME")
		DN         = "4400"
	)

	service.Init(API_URL, 0, true)

	err := auth.Login(UserID, Passwd, TenantName, []code.MediaType{code.Media.Voice}, 
		code.AgentStatus.NotReady, 
		code.AgentStateCause.NotReady.Idle, 
		DN, 
		handlerEvent, handlerError,
	)
	if err != nil {
		panic(err)
	}
}

func handlerEvent(e utils.Event) {
	log.Println(e)
}

func handlerError(err error) {
	log.Println(err)
}

```

### 정보조회

- Group 목록 조회

```golang
	resp, err := info.GetGroupList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}
```

- Flow 목록 조회

```golang
	resp, err = info.GetFlowList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}
```

## 모니터링

### 초기화

```golang
import (
	"log"
	"os"

	"github.com/ipron-ne/client-sdk-go/monitoring"
	"github.com/ipron-ne/client-sdk-go/monitoring/sse"
	"github.com/ipron-ne/client-sdk-go/utils"
)

func main() {
	var (
		API_URL    = os.Getenv("IPRON_NE_URL")
		TenantID   = os.Getenv("IPRON_NE_TENANTID")
		AppToken   = os.Getenv("IPRON_NE_APPKEY")
	)

	params := map[string]any{"token": AppToken}
	monitoring.Init(API_URL, params, 0, true)
}
```

### 모니터링용 데이터셋 목록 조회

```golang
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
```

### 모니터링용 지정 데이터셋 항목 조회

```golang
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
```

### 모리터링용 데이터셋 전체 조회

```golang
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
```

### 모니터링 변경 데이터 수신

```golang
func moniEvent(resource []string) {
	log.Printf("\n\n[EventListen]\n")

	params := map[string]any{
		"tntId": TenantID,
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
```

### 모니터링 APP KEY로 데이터 조회

```golang
import (
	"log"

	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/info"	
)

func main() {
	service.Init(API_URL, 0, true)
	service.GetApiClient().SetToken(AppToken)
	service.GetApiClient().SetTenant(TenantID)

	resp, err := info.GetFlowList(service.GetApiClient().GetTenantID())
	if err != nil {
		log.Println(err)
	}

	flowList := []string{}
	for _, v := range resp.GetData().Array() {
		flowList = append(flowList, v.Get("_id").Str())
	}

	log.Printf("%+v\n", flowList)
}
```
