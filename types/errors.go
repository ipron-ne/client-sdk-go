package types

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type APIError struct {
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s: code=%d status=%d msg=%s", e.Title, e.Code, e.Status, e.Msg)
}

func (e APIError) GetCode() int {
	return e.Code
}

func (e APIError) GetStatus() int {
	return e.Status
}

func (e APIError) GetTitle() string {
	return e.Title
}

func (e APIError) GetMessage() string {
	return e.Msg
}

func GetAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// Service API 의 Common 응답정보 분석
func GetServiceError(resp *Response, perr error) error {
	var err APIError

	if perr != nil || resp == nil {
		return perr
	}

	data := resp.GetBody()

	switch v := data["code"].(type) {
	case int:
		err.Code = v
	case string:
		err.Code, _ = strconv.Atoi(v)
	}

	if v, ok := data["msg"]; ok {
		err.Msg = v.(string)
	}
	if v, ok := data["result"]; ok {
		if v.(bool) {
			err.Status = 1
		} else {
			err.Status = 0
		}
	}

	if err.Code == 0 {
		if resp.StatusCode != http.StatusOK {
			if err.Msg == "" {
				err.Msg = "http error response"
			}
			return &APIError{
				Code:   -1,
				Status: resp.StatusCode,
				Title:  "",
				Msg:    err.Msg,
			}
		}
		return nil
	}

	return &err
}

// Backend API 의 Common 응답정보 분석
func GetBackendError(resp *Response, perr error) error {
	var err APIError

	if perr != nil || resp == nil {
		return perr
	}

	data := resp.GetBody()

	switch v := data["code"].(type) {
	case int:
		err.Code = v
	case string:
		err.Code, _ = strconv.Atoi(v)
	}

	if v, ok := data["status"]; ok {
		err.Status = int(v.(float64))
	}
	if v, ok := data["title"]; ok {
		err.Title = v.(string)
	}
	if v, ok := data["msg"]; ok {
		err.Msg = v.(string)
	}

	if err.Code == 0 {
		if resp.StatusCode != http.StatusOK {
			if err.Msg == "" {
				err.Msg = "http error response"
			}
			return &APIError{
				Code:   -1,
				Status: resp.StatusCode,
				Title:  "",
				Msg:    err.Msg,
			}
		}
		return nil
	}

	return &err
}
