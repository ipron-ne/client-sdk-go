package types

import (
	"errors"
	"fmt"
	"net/http"
)

type APIError struct {
	Code   int    `json:"code"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
}

func NewAPIError(resp *Response) error {
	if resp.Code == 0 {
		if resp.StatusCode != http.StatusOK {
			return &APIError{
				Code:   -1,
				Status: resp.StatusCode,
				Title:  "",
				Msg:    "http error response",
			}
		}
		return nil
	}

	return &APIError{
		Code:   resp.Code,
		Status: resp.Status,
		Title:  resp.Title,
		Msg:    resp.Msg,
	}
}

func GetAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
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
