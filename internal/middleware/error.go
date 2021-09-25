package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/pkg/core"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *fiber.Ctx) error {

		c.Next()
		e := c.Errors.Last()
		if e == nil {
			return
		}
		err := e.Err

		switch err.(type) {
		case core.Error:
			codeErr := err.(core.Error)
			core.ErrorResp(c, codeErr.Code, codeErr.Message)

		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr := fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
			core.InvalidParamsResp(c, errStr)
		default:
			errStr := e.Err.Error()
			core.InvalidParamsResp(c, errStr)
		}

	}

}
