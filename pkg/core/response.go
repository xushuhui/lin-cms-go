package core

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/pkg/errcode"
)

// Error 数据返回通用 JSON 数据结构
type Error struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
	Err     error       `json:"-"`
}

func (e Error) Error() (re string) {
	return fmt.Sprintf("code=%v, Message=%v", e.Code, e.Message)
}

func NewErrorCode(code int) (e Error) {
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

func ValidateRequest(obj interface{}) error {

	err := validate.Struct(obj)

	if err != nil {
		s := Translate(err.(validator.ValidationErrors))
		return errors.New(s)
	}
	return nil
}
func ParseRequest(c *fiber.Ctx, request interface{}) (err error) {
	err = c.BodyParser(request)

	if err != nil {
		return err
	}
	err = ValidateRequest(request)
	return
}
func FailResp(c *fiber.Ctx, code int) error {
	return c.JSON(Error{
		Code:    code,
		Message: errcode.GetMsg(code),
	})
}

func ErrorResp(c *fiber.Ctx, code int, msg string) error {
	return c.JSON(Error{
		Code:    code,
		Message: msg,
	})

}
func InvalidParamsResp(c *fiber.Ctx, msg string) error {

	return c.JSON(Error{
		Code:    errcode.InvalidParams,
		Message: msg,
	})

}

func SuccessResp(c *fiber.Ctx) error {
	return c.JSON(Error{
		Code:    0,
		Message: errcode.GetMsg(0),
	})
}
func SetData(c *fiber.Ctx, data map[string]interface{}) error {
	return c.JSON(Error{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data:    data,
	})
}

func SetPage(c *fiber.Ctx, list interface{}, totalRows int) error {
	return c.JSON(Error{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data: Pager{
			Page:  GetPage(c),
			Size:  GetSize(c),
			Total: totalRows,
			Items: list,
		},
	})
}
func ServerError(c *fiber.Ctx, err error) error {

	return c.JSON(Error{
		Code:    errcode.ServerError,
		Message: errcode.GetMsg(errcode.ServerError),
		Err:     err,
	})
}
func (e Error) HttpError(c *fiber.Ctx) error {
	return c.JSON(Error{
		Code:    e.Code,
		Message: errcode.GetMsg(e.Code),
	})
}
