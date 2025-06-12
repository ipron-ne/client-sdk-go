package types

type Datasource struct {
	DatasetName string `json:"datasetName"`
	DataSource  string `json:"dataSource"`
	RedisKey    string `json:"redisKey"`
	JsonData    string `json:"jsonData"`
}

type DatasourceField struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}

type Datasets struct {
	Name    string   `json:"name"`
	Dataset []string `json:"dataset"`
}

type Dataset struct {
	Tenant      []DataField `json:"tenant"`
	Dnis        []DataField `json:"dnis"`
	Flow        []DataField `json:"flow"`
	Scenario    []DataField `json:"scenario"`
	Menu        []DataField `json:"menu"`
	Servicecode []DataField `json:"servicecode"`
	Queue       []DataField `json:"queue"`
	Group       []DataField `json:"group"`
	User        []DataField `json:"user"`
	Interaction []DataField `json:"interaction"`
}

type Dataset2 map[string][]DataField

type DataField struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	PrKey bool   `json:"pk"`
	Desc  string `json:"desc"`
}

type MonitoringEvent map[string]any
