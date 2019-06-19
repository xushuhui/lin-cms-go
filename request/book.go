package request

import (
	"github.com/astaxie/beego/validation"
	"lin-cms-beego/utils"
)

type CreateBook struct {
	Title   string `json:"title"valid:"Required"`
	Author  string `json:"author"valid:"Required"`
	Summary string `json:"summary"valid:"Required"`
	Image   string `json:"image"valid:"Required"`
}

type UpdateBook struct {
	Id      int    `json:"id"valid:"Required"`
	Title   string `json:"title"valid:"Required"`
	Author  string `json:"author"valid:"Required"`
	Summary string `json:"summary"valid:"Required"`
	Image   string `json:"image"valid:"Required"`
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
