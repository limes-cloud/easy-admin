package types

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

type ResponseList struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id"`
	Total   int64       `json:"total"`
}

func (r *Response) Error() string {
	return r.Msg
}
