package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
)

func Upload(c *fiber.Ctx) error {
	return core.SuccessResp(c)
}
func GetLogs(c *fiber.Ctx) error {
	var req request.GetLogs
	page := core.GetPage(c)
	size := core.GetSize(c)
	if err := core.ParseQuery(c, &req); err != nil {
		return err
	}
	data, total, err := biz.GetLogs(c.Context(), req, page, size)
	if err != nil {
		return err
	}
	return core.SetPage(c, data, total)
}

func SearchLogs(c *fiber.Ctx) error {
	var req request.SearchLogs
	if err := core.ParseQuery(c, &req); err != nil {
		return err
	}
	page := core.GetPage(c)
	size := core.GetSize(c)

	data, total, err := biz.SearchLogs(c.Context(), req, page, size)
	if err != nil {
		return err
	}
	return core.SetPage(c, data, total)
}

func GetLogUsers(c *fiber.Ctx) error {
	page := core.GetPage(c)
	size := core.GetSize(c)
	data, total, err := biz.GetLogUsers(c.Context(), page, size)
	if err != nil {
		return err
	}
	return core.SetPage(c, data, total)
}
