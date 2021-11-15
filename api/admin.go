package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xushuhui/goal/core"
	"github.com/xushuhui/goal/utils"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
)

func GetUsers(c *fiber.Ctx) error {
	var req request.GetUsers
	if err := core.ParseRequest(c, &req); err != nil {

		return err
	}

	data, err := biz.GetUsers(req)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}

func ChangeUserPassword(c *fiber.Ctx) error {
	var req request.ChangeUserPassword
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	err := biz.ChangeUserPassword(c.Context(), req.Id, req.NewPassword)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)

}

func DeleteUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return err
	}
	err = biz.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)

}

func UpdateUser(c *fiber.Ctx) error {
	var req request.UpdateUser
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	err := biz.UpdateUser(req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)

}
