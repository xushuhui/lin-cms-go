package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/biz"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
)

func GetGroups(c *fiber.Ctx) error {
	data, err := biz.GetGroups()
	if err != nil {

		return err
	}
	return core.SetData(c, data)


}
func GetGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {

		return err
	}
	data, err := biz.GetGroup(id)
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

	err := biz.CreateGroup(req)

	return err
}
func UpdateGroup(c *fiber.Ctx) error {
	var req request.UpdateGroup
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	err := biz.UpdateGroup(req)

	return err
}
func DeleteGroup(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {

		return err
	}
	err = biz.DeleteGroup(id)
	if err != nil {

		return err
	}
	return 	core.SuccessResp(c)

}
