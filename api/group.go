package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xushuhui/goal/core"
	"github.com/xushuhui/goal/utils"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
)

func GetGroups(c *fiber.Ctx) error {
	data, err := biz.GetGroups(c.Context())
	if err != nil {
		return err
	}
	return core.SetData(c, data)
}

func GetGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}
	data, err := biz.GetGroup(c.Context(), id)
	if err != nil {
		return err
	}
	return core.SetData(c, data)
}

func CreateGroup(c *fiber.Ctx) error {
	var req request.CreateGroup
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	err := biz.CreateGroup(c.Context(), req.Name, req.Info, req.PermissionIds)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func UpdateGroup(c *fiber.Ctx) error {
	var req request.UpdateGroup
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}

	err = biz.UpdateGroup(c.Context(), id, req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}
func DeleteGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}

	err = biz.DeleteGroup(c.Context(), id)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)

}
