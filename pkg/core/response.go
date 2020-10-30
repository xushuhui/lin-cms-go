package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"lin-cms-go/pkg/errcode"
)

// Error 数据返回通用 JSON 数据结构
type Error struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)

}

func (e Error) Error() (re string) {
	return fmt.Sprintf("code=%v, Message=%v", e.Code, e.Message)
}

func NewError(code int) (e Error) {
	e = Error{
		Code:    code,
		Message: errcode.GetMsg(code),
	}
	return
}
func NewErrorMessage(code int, message string) (e Error) {
	e = Error{
		Code:    code,
		Message: message,
	}
	return
}
func NewInvalidParamsError(message string) (e Error) {
	return NewErrorMessage(errcode.InvalidParams, message)
}
func ParseRequest(c *gin.Context, request interface{}) (err error) {
	err = c.ShouldBind(&request)

	if err != nil {
		msg := Translate(err.(validator.ValidationErrors))
		err = NewErrorMessage(errcode.InvalidParams, msg)
		return
	}
	return
}
func FailResp(c *gin.Context, code int) {
	c.AbortWithStatusJSON(200, Error{
		Code:    code,
		Message: errcode.GetMsg(code),
	})
	return
}

func ErrorResp(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(200, Error{
		Code:    code,
		Message: msg,
	})
	return
}
func InvalidParamsResp(c *gin.Context, msg string) {

	c.AbortWithStatusJSON(200, Error{
		Code:    errcode.InvalidParams,
		Message: msg,
	})
	return
}

func SuccessResp(c *gin.Context) {
	c.JSON(200, Error{
		Code:    0,
		Message: errcode.GetMsg(0),
	})
}
func SetData(c *gin.Context, data map[string]interface{}) {
	c.JSON(200, Error{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data:    data,
	})
}

func SetPage(c *gin.Context, list interface{}, totalRows int) {
	c.JSON(200, Error{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data: Pager{
			Page:      GetPage(c),
			PageSize:  GetPageSize(c),
			TotalRows: totalRows,
			List:      list,
		},
	})
}
func ServerError(c *gin.Context) {
	c.JSON(500, Error{
		Code:    errcode.ServerError,
		Message: errcode.GetMsg(errcode.ServerError),
	})
}
