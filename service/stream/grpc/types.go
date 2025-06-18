package stream

import (
	streampb "github.com/ipron-ne/client-sdk-go/service/stream/grpc/proto"
)

type StreamStatus = streampb.StreamStatus

const (
	OK                    = streampb.StreamStatus_OK
	INVALID_PARAM         = streampb.StreamStatus_INVALID_PARAM
	CREATE_CHANNEL_FAILED = streampb.StreamStatus_CREATE_CHANNEL_FAILED
	UPDATE_URI_ERROR      = streampb.StreamStatus_UPDATE_URI_ERROR
	GET_CHANNEL_FAILED    = streampb.StreamStatus_GET_CHANNEL_FAILED
	MEDIA_URL_NOT_FOUND   = streampb.StreamStatus_MEDIA_URL_NOT_FOUND
	STREAM_CHANNEL_CLOSED = streampb.StreamStatus_STREAM_CHANNEL_CLOSED
)

type StreamResponse struct {
	StatusCode   StreamStatus `json:"status_code"`   // 상태코드
	ErrorMessage string       `json:"error_message"` // 오류명
	CallID       string       `json:"call_id"`       // 전송 Stream의 Call ID
	ConnID       string       `json:"connection_id"` // 전송 Stream의 Connection ID
	StreamData   []byte       `json:"stream_data"`   // Stream 데이터
}

func (s StreamResponse) GetStatusCode() StreamStatus {
	return s.StatusCode
}

func (s StreamResponse) GetMessage() string {
	return s.ErrorMessage
}

func (s StreamResponse) GetCallID() string {
	return s.CallID
}

func (s StreamResponse) GetConnID() string {
	return s.ConnID
}

func (s StreamResponse) GetStreamData() []byte {
	return s.StreamData
}

func (s StreamResponse) GetStreamDataLen() int {
	return len(s.StreamData)
}
