package core

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/pkg/errcode"
)

// Error 数据返回通用 JSON 数据结构
type IError struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
	Err     error       `json:"-"`
}

func (err IError) Error() (re string) {
	return fmt.Sprintf("code=%v, Message=%v", err.Code, err.Message)
}

func NewErrorCode(code int) (err error) {
	err = IError{
		Code:    code,
		Message: errcode.GetMsg(code),
	}
	return
}
func NewErrorMessage(code int, message string) (err error) {
	err = IError{
		Code:    code,
		Message: message,
	}
	return
}
func NewInvalidParamsError(message string) (err error) {
	return NewErrorMessage(errcode.InvalidParams, message)
}

func ValidateRequest(obj interface{}) error {

	err := validate.Struct(obj)

	if err != nil {
		s := Translate(err.(validator.ValidationErrors))
		return NewInvalidParamsError(s)
	}
	return nil
}
func ParseQuery(c *fiber.Ctx, request interface{}) (err error) {
	err = c.QueryParser(request)

	if err != nil {
		return err
	}

	err = ValidateRequest(request)
	return
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
	return c.JSON(IError{
		Code:    code,
		Message: errcode.GetMsg(code),
	})
}

func ErrorResp(c *fiber.Ctx, code int, msg string) error {
	return c.JSON(IError{
		Code:    code,
		Message: msg,
	})

}
func InvalidParamsError(c *fiber.Ctx, msg string) error {

	return c.JSON(IError{
		Code:    errcode.InvalidParams,
		Message: errcode.GetMsg(errcode.InvalidParams),
	})

}

func SuccessResp(c *fiber.Ctx) error {
	return c.JSON(IError{
		Code:    0,
		Message: errcode.GetMsg(0),
	})
}
func SetData(c *fiber.Ctx, data interface{}) error {
	return c.JSON(IError{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data:    data,
	})
}

func SetPage(c *fiber.Ctx, list interface{}, totalRows int) error {
	return c.JSON(IError{
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

	return c.JSON(IError{
		Code:    errcode.ServerError,
		Message: errcode.GetMsg(errcode.ServerError),
		Err:     err,
	})
}
func (err *IError) HttpError(c *fiber.Ctx) error {
	return c.JSON(IError{
		Code:    err.Code,
		Message: err.Message,
	})
}
