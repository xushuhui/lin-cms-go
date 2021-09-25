package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
)

func GetUsers(c *fiber.Ctx) error {
	var req request.GetUsers
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.GetUsers(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}

func ChangeUserPassword(c *fiber.Ctx) error {
	var req request.ChangeUserPassword
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.ChangeUserPassword(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	err = services.DeleteUser(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return

}

func UpdateUser(c *fiber.Ctx) error {
	var req request.UpdateUser
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.UpdateUser(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
