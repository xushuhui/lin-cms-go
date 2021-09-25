package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
)

func GetGroups(c *fiber.Ctx) error {
	data, err := services.GetGroups()
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
func GetGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	data, err := services.GetGroup(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
func CreateGroup(c *fiber.Ctx) error {
	var req request.CreateGroup
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.CreateGroup(req)
	if err != nil {
		c.Error(err)
		return
	}
	return
}
func UpdateGroup(c *fiber.Ctx) error {
	var req request.UpdateGroup
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.UpdateGroup(req)
	if err != nil {
		c.Error(err)
		return
	}
	return
}
func DeleteGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	err = services.DeleteGroup(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
