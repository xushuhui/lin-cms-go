/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package request

import (
	"lin-cms-gin/utils"
)

type CreateBook struct {
	Title   string `json:"title" valid:"Required"`
	Author  string `json:"author" valid:"Required"`
	Summary string `json:"summary" valid:"Required"`
	Image   string `json:"image" valid:"Required"`
}

type UpdateBook struct {
	Id      int    `json:"id" valid:"Required"`
	Title   string `json:"title" valid:"Required"`
	Author  string `json:"author" valid:"Required"`
	Summary string `json:"summary" valid:"Required"`
	Image   string `json:"image" valid:"Required"`
}

func (request *CreateBook) Bind(body []byte) {
	utils.BindJson(body, request)
}
func (request *UpdateBook) Bind(body []byte) {
	utils.BindJson(body, request)
}
func (request *UpdateBook) Verify() (b bool, msg string) {
	return validate(request)

}
func (request *CreateBook) Verify() (b bool, msg string) {
	return validate(request)
}
func validate(request interface{}) (b bool, msg string) {
	valid := validation.Validation{}
	verify, err := valid.Valid(request)
	if err != nil {
		panic(err)
	}
	if !verify {
		for _, err := range valid.Errors {
			return false, err.Field + " " + err.Message
		}
	}
	return true, ""
}
