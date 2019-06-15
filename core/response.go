package core

type Data struct {
	Data interface{} `json:"data,omitempty"`
}

type Response struct {
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
}

func Fail(errorCode int) Response {
	return Resp(errorCode, GetMsg(errorCode), nil)
}
func FailMsg(errorCode int, msg string) Response {
	return Resp(errorCode, msg, nil)
}
func Succeed() Response {
	return Resp(CODE_OK, GetMsg(CODE_OK), nil)
}
func SetData(data interface{}) Response {
	return Resp(CODE_OK, GetMsg(CODE_OK), data)
}
func Resp(errorCode int, msg string, data interface{}) Response {
	res := Response{
		ErrorCode: errorCode,
		Msg:       msg,
		Data:      data,
	}
	return res
}
