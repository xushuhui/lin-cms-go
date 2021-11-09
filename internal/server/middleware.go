package server

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data"
)

func UserLog(c *fiber.Ctx) error {

	c.Next()

	user := biz.LocalUser(c)

	msg := c.Locals("logMessage")
	if msg == nil {
		return nil
	}
	err := data.CreateLog(c.Context(), c.Response().StatusCode(), user.ID, user.Username, user.Username+msg.(string), c.Method(), c.Path(), "")

	return err
}
