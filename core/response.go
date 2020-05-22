/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package core
//Data Data
type Data struct {
	Data interface{} `json:"data,omitempty"`
}
//Response 响应
type Response struct {
	Code int         `json:"code"`
	Message       string      `json:"Message"`
	Data      interface{} `json:"data,omitempty"`
}

func Fail(Code int) Response {
	return Resp(Code, GetMessage(Code), nil)
}

func SetFail(Code int, Message string) Response {
	return Resp(Code, Message, nil)
}
func ParmsError(Message string) Response {
	return Resp(InvaldParams, Message, nil)
}
func Succeed() Response {
	return Resp(OK, GetMessage(OK), nil)
}
func SetData(data interface{}) Response {
	return Resp(OK, GetMessage(OK), data)
}
func Resp(Code int, Message string, data interface{}) Response {
	res := Response{
		Code: Code,
		Message:       Message,
		Data:      data,
	}
	return res
}
