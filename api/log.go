package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
)

func Upload(c *fiber.Ctx) error {
	core.SuccessResp(c)
	return
}
func GetLogs(c *fiber.Ctx) error {
	var req request.GetLogs
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.GetLogs(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
func SearchLogs(c *fiber.Ctx) error {
	var req request.SearchLogs
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.SearchLogs(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
func GetLogUsers(c *fiber.Ctx) error {
	var req request.GetLogUsers
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.GetLogUsers(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
