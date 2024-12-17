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
	"github.com/ipron-ne/client-sdk-go/config"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/service/auth"
	"github.com/ipron-ne/client-sdk-go/types"
)

func main() {
	var (
		API_URL    = os.Getenv("IPRON_NE_URL")
		UserID     = os.Getenv("USER_ID")
		Passwd     = os.Getenv("USER_PWD")
		TenantName = os.Getenv("TENANT_NAME")
		DN         = "4400"
	)

	// 접속환경 설정
	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithDebug(true),
	)

	// Client 생성
	client := service.NewFromConfig(cfg)

	// Client 인스턴스로 인증 서비스 생성
	auth := auth.NewFromClient(client)
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

func handlerEvent(e types.Data) {
	log.Println(e)
}

func handlerError(err error) {
	log.Println(err)
}

```

### 정보조회

- Group 목록 조회

```golang
	// Client 인스턴스로 정보조회 서비스 생성
	info := info.NewFromClient(client)

	// Group List 조회
	resp, err := info.GetGroupList(client.GetTenantID())
	if err != nil {
		log.Println(err)
	}
	for i, v := range resp.GetData().Array() {
		log.Printf("[%02d] %+v\n", i,v.Object())
	}
```

- Flow 목록 조회

```golang
	// Flow List 조회
	resp, err = info.GetFlowList(client.GetTenantID())
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

	"github.com/ipron-ne/client-sdk-go/monitoring/sse"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/utils"
)

func main() {
	var (
		API_URL    = os.Getenv("IPRON_NE_URL")
		TenantID   = os.Getenv("IPRON_NE_TENANTID")
		AppToken   = os.Getenv("IPRON_NE_APPKEY")
	)

	cfg := config.NewConfig(
		config.WithBaseURL(API_URL),
		config.WithAppToken(AppToken),
		config.WithDebug(true),
		config.WithTenantID(TenantID),
	)

	client := service.NewFromConfig(cfg)
}
```

### 모니터링용 데이터셋 목록 조회

```golang
func listDatasets(client types.Client) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Datasets List]\n")
	resp, err := monitor.GetDatasets(utils.NewParam("tntId", TenantID))
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
func listDataset(client types.Client, datasetName string) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Dataset List:%s]\n", datasetName)
	resp, err := monitor.GetDataset(datasetName, utils.NewParam("tntId", TenantID))
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
func listDatasource(client types.Client) {
	monitor := sse.NewFromClient(client)

	log.Printf("\n\n[Datasource]\n")
	resp, err := monitor.GetDatasource(utils.NewParam("tntId", TenantID))
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
func moniEvent(client types.Client, resource []string) {
	monitor := sse.NewFromClient(client)

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

	eventSubs, err := monitor.GetEventSource("flow", params)
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
