package common

//Response 相应请求
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ResponseSuccess 响应成功
func ResponseSuccess(data interface{}) *Response {
	return &Response{
		Code:    1,
		Message: "操作成功",
		Data:    data,
	}
}

//ResponseFailed 响应失败
func ResponseFailed(data interface{}) *Response {
	return &Response{
		Code:    1,
		Message: "操作失败",
		Data:    nil,
	}
}
