package types

import (
	"encoding/json"
	"fmt"
	"net/http"

	"reflect"

	"github.com/ipron-ne/client-sdk-go/config"
)

type Function = func(Data)
type FunctionMap = map[string]func(Data)

type Logger interface {
	Success(fmt string, args ...interface{})
	Error(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Debug(fmt string, args ...interface{})
}

type Client interface {
	GetClientID() string
	GetToken() string
	GetBaseURL() string
	IsDebug() bool

	SetLocalToken(accessToken, resfreshToken string)
	DeleteLocalToken()

	GetTenantID() string
	GetUserID() string

	GetLogger() Logger
	GetRequest() Request
	GetConfig() config.Config
}

type Request interface {
	Post(uri string, bodyJson any) (*Response, error)
	Put(uri string, bodyJson any) (*Response, error)
	Get(uri string, bodyJson any) (*Response, error)
	Request(method, uri string, bodyJson any) (*Response, error)

	DelHeader(name string)
	SetHeader(name string, value string)
}

type GatewayResponse struct {
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
}

type ServiceResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"msg"`
	Code    int    `json:"code"`
}

type Response struct {
	*http.Response `json:"-"`
	Body           map[string]any
}

func (s *Response) SetResult(data map[string]any) {
	s.Body = data
}

func (s *Response) DataUnmarshal(out any) error {
	jdoc, _ := json.Marshal(s.Body["data"])
	return json.Unmarshal(jdoc, out)
}

func (s *Response) ServiceUnmarshal(out any) error {
	jdoc, _ := json.Marshal(s.Body)
	return json.Unmarshal(jdoc, out)
}

func (s *Response) GetBody() map[string]any {
	return s.Body
}

// Platform 으로 부터 수신받은 데이터를 처리하기 위한 구조체
type Data struct {
	any
}

func NewData(v any) Data {
	return Data{any: v}
}

func (d Data) Unmarshal(out any) error {
	b, _ := json.Marshal(d.any)
	return json.Unmarshal(b, out)
}

func (d Data) Array() []Data {
	arr, ok := d.any.([]any)
	if !ok {
		return []Data{}
	}

	ret := []Data{}
	for _, v := range arr {
		ret = append(ret, Data{any: v})
	}

	return ret
}

func (d Data) Object() map[string]Data {
	obj, ok := d.any.(map[string]any)
	if !ok {
		fmt.Printf("convert fail [%s]\n", reflect.TypeOf(d.any))
		return map[string]Data{}
	}

	ret := map[string]Data{}
	for k, v := range obj {
		ret[k] = Data{any: v}
	}

	return ret
}

func (d Data) Get(key string) Data {
	if d.any == nil {
		return Data{}
	}

	obj, ok := d.any.(map[string]any)
	if !ok {
		return Data{}
	}

	ret, ok := obj[key]
	if !ok {
		return Data{}
	}

	return Data{any: ret}
}

func (d Data) Str() string {
	if d.any == nil {
		return ""
	}
	return d.any.(string)
}

func (d Data) Int() int {
	if d.any == nil {
		return 0
	}
	return d.any.(int)
}

func (d Data) Bool() bool {
	if d.any == nil {
		return false
	}
	return d.any.(bool)
}
