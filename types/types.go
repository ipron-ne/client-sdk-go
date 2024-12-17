package types

import (
	"fmt"
	"net/http"
	"strconv"

	"reflect"
)

type Function  = func(Data)
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

	SetLocalToken(data Data)
	DeleteLocalToken()

	GetTenantID() string
	GetUserID() string

	GetLogger() Logger
	GetRequest() Request
}

type Request interface {
	Post(uri string, bodyJson map[string]any) (*Response, error)
	Put(uri string, bodyJson map[string]any) (*Response, error)
	Get(uri string, bodyJson map[string]any) (*Response, error)
	Request(method, uri string, bodyJson map[string]any) (*Response, error)

	DelHeader(name string)
	SetHeader(name string, value string)
}

type Response struct {
	*http.Response `json:"-,omitempty"`
	Code   int     `json:"code"`
	Status int     `json:"status"`
	Title  string  `json:"title"`
	Msg    string  `json:"msg"`
	Data   any     `json:"data"`
}

func (s *Response) SetResult(data map[string]any) {
	switch v := data["code"].(type) {
	case int:
		s.Code = v
	case string:
		s.Code, _ = strconv.Atoi(v)
	}

	if v, ok := data["status"]; ok {
		s.Status = int(v.(float64))
	}
	if v, ok := data["title"]; ok {
		s.Title = v.(string)
	}
	if v, ok := data["msg"]; ok {
		s.Msg = v.(string)
	}
	if v, ok := data["data"]; ok {
		s.Data = v
	}
}

func (s *Response) GetData() Data {
	return Data{any: s.Data}
}

// Platform 으로 부터 수신받은 데이터를 처리하기 위한 구조체
type Data struct {
	any
}

func NewData(v any) Data {
	return Data{any: v}
}

func (d Data) Array() []Data {
	arr, ok := d.any.([]any)
	if !ok {
		return []Data{}
	}

	ret := []Data{}
	for _, v := range arr {
		ret = append(ret, Data{any:v})
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
		ret[k] = Data{any:v}
	}

	return ret
}

func (d Data) Get(key string) Data {
	obj, ok := d.any.(map[string]any)
	if !ok {
		return Data{}
	}

	ret, ok := obj[key]
	if !ok {
		return Data{}
	}

	return Data{any:ret}
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
