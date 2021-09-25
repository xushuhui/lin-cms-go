package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/biz"
	"lin-cms-go/pkg/core"
)

func Upload(c *fiber.Ctx) error {

	return core.SuccessResp(c)
}
func GetLogs(c *fiber.Ctx) error {
	var req request.GetLogs
	if err := core.ParseRequest(c, &req); err != nil {
	
		return err
	}

	data, err := biz.GetLogs(req)
	if err != nil {
		return err
	}
	return core.SetData(c, data)

}
func SearchLogs(c *fiber.Ctx) error {
	var req request.SearchLogs
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	data, err := biz.SearchLogs(req)
	if err != nil {
		return err
	}

	return core.SetData(c, data)
}
func GetLogUsers(c *fiber.Ctx) error {
	var req request.GetLogUsers
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	data, err := biz.GetLogUsers(req)
	if err != nil {
		return err
	}

	return core.SetData(c, data)
}
