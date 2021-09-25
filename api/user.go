package api

import (
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/biz"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/utils"

	"github.com/gofiber/fiber/v2"
)


func Login(c *fiber.Ctx) error {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := biz.Login(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}


func Register(c *fiber.Ctx) error {
	var req request.Register
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := biz.Register(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}


func UpdateMe(c *fiber.Ctx) error {
	var req request.UpdateMe
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	uid, err := uid(c)
	if err != nil {
		return
	}

	err = biz.UpdateMe(req, uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}


func ChangeMyPassword(c *fiber.Ctx) error {

	var req request.ChangeMyPassword
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	uid, err := uid(c)
	if err != nil {
		return
	}

	err = biz.ChangeMyPassword(req, uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
func uid(c *fiber.Ctx)  (uid uint, err error) {
	u, ok := c.Get("uid")

	if !ok {
		c.Error(core.NewErrorMessage(errcode.AuthCheckTokenFail, "uid错误"))
	}
	uid, err = utils.ToUint(u)
	return
}


func GetMyPermissions(c *fiber.Ctx) error {
	uid, err := uid(c)
	if err != nil {
		return err
	}
	data, err := biz.GetMyPermissions(uid)
	if err != nil {
		
		return err
	}
	core.SetData(c, data)
	return
}


func GetMyInfomation(c *fiber.Ctx) error {
	uid, err := uid(c)
	if err != nil {
	
		return err
	}
	data, err := biz.GetMyInfomation(uid)
	if err != nil {
		
		return err
	}
	core.SetData(c, data)
	return
}

//TODO 对cms意义并不大，先不实现
func RefreshToken(c *fiber.Ctx) error {

}
