package unpla

// 参考文档： https://wiki.hangjia.online/pages/viewpage.action?pageId=5473883

import "encoding/json"

//BaseResponse 另一种写法。这里面用到了一个
//golang的语法糖，间接实现了类的继承与方法合并
type BaseResponse struct {
	Result  int    `json:"result"`
	Message string `json:"msg"`
}

// StandardResponse 标准返回结果
type StandardResponse struct {
	BaseResponse
	Data interface{} `json:"data"`
}

// ErrorResponse 错误返回结果
type ErrorResponse struct {
	BaseResponse
}

// NewStandResponse 快速创建一个标准返回
func NewStandResponse(result int, msg string, data interface{}) *StandardResponse {
	return &StandardResponse{
		BaseResponse: BaseResponse{
			Result:  result,
			Message: msg},
		Data: data,
	}
}

// NewErrorResponse 快速创建一个标准返回
func NewErrorResponse(result int, msg string) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			Result:  result,
			Message: msg,
		},
	}
}

// JSON 返回JSON字符串
func (sr *StandardResponse) JSON() []byte {
	js, _ := json.Marshal(&sr)
	return js
}
