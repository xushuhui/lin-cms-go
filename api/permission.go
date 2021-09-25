package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/biz"
	"lin-cms-go/pkg/core"
)

func GetAllPermissions(c *fiber.Ctx) error {
	data, err := biz.GetAllPermissions()
	if err != nil {
		
		return err
	}
	core.SetData(c, data)
	return
}

//TODO: 没找到必须实现该接口的业务场景，而且功能和分配多个权限重复，开发待定
func DispatchPermission(c *fiber.Ctx) error {

}

func DispatchPermissions(c *fiber.Ctx) error {
	var req request.DispatchPermissions
	if err := core.ParseRequest(c, &req); err != nil {
 		return err
	}
	err := biz.DispatchPermissions(req)
	if err != nil {
		return err
	}
	core.SuccessResp(c)
	return
}

func RemovePermissions(c *fiber.Ctx) error {
	var req request.RemovePermissions
	if err := core.ParseRequest(c, &req); err != nil {
		
		return err
	}
	err := biz.RemovePermissions(req)
	if err != nil {
		
		return err
	}
	core.SuccessResp(c)
	return
}
