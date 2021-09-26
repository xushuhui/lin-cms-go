package api

import (
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"

	"lin-cms-go/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	uc *biz.LinUserUsecase
}

func NewUser(uc *biz.LinUserUsecase) *User {
	return &User{
		uc: uc,
	}
}
func (u *User) Hello(ctx *fiber.Ctx) error {
	return ctx.JSON("Hello")
}
func (u *User) Login(c *fiber.Ctx) error {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	data, err := u.uc.Login(c.Context(), req.Username, req.Password)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}

func (u *User) Register(c *fiber.Ctx) error {
	var req request.Register
	if err := core.ParseRequest(c, &req); err != nil {

		return err
	}

	err := biz.Register(req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func (u *User) UpdateMe(c *fiber.Ctx) error {
	var req request.UpdateMe
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	uid, err := uid(c)
	if err != nil {
		return err
	}

	err = biz.UpdateMe(req, uid)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func (u *User) ChangeMyPassword(c *fiber.Ctx) error {

	var req request.ChangeMyPassword
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}

	uid, err := uid(c)
	if err != nil {
		return err
	}

	err = biz.ChangeMyPassword(req, uid)
	if err != nil {

		return err
	}
	return core.SuccessResp(c)

}
func uid(c *fiber.Ctx) (uid uint, err error) {
	u := c.Get("uid", "")

	uid, err = utils.ToUint(u)
	return
}

func (u *User) GetMyPermissions(c *fiber.Ctx) error {
	uid, err := uid(c)
	if err != nil {
		return err
	}
	data, err := biz.GetMyPermissions(uid)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}

func (u *User) GetMyInfomation(c *fiber.Ctx) error {
	uid, err := uid(c)
	if err != nil {

		return err
	}
	data, err := biz.GetMyInfomation(uid)
	if err != nil {

		return err
	}
	return core.SetData(c, data)

}

//TODO 对cms意义并不大，先不实现
func RefreshToken(c *fiber.Ctx) error {
	return nil
}
