package stream

type StreamServiceAllocRequest struct {
	CallID string `json:"call_id"`       // Stream을 사용할 Call ID
	ConnID string `json:"connection_id"` // Stream 대상 Connection ID
}

type StreamServiceAllocResponse struct {
	ConnID string `json:"connection_id"` // Stream 대상 Connection ID
	URI    string `json:"pod_name"`      // 현재 처리 중인 Stream POD 이름
}

type StreamAllocRequest struct {
	CallID     string `json:"call_id"`       // Stream을 사용할 Call ID
	ConnID     string `json:"connection_id"` // Stream 대상 Connection ID
	StreamType string `json:"stream_type"`   // 처리될  Stream 타입 지정 [rx|tx]
}

type StreamAllocResponse struct {
	CallID        string `json:"call_id"`        // Stream을 사용할 Call ID
	ConnID        string `json:"connection_id"`  // Stream 대상 Connection ID
	StreamChannel string `json:"stream_channel"` // Stream 웹소켓 채널
}

type StreamRequest struct {
}

type StreamResponse struct {
	CallID     string `json:"call_id"`       // 전송 Stream의 Call ID
	ConnID     string `json:"connection_id"` // 전송 Stream의 Connection ID
	StreamData []byte `json:"stream_data"`   // Stream 데이터
}
