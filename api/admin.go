package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
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

	err := biz.ChangeUserPassword(req)
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
	err = biz.DeleteUser(id)
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
