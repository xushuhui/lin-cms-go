package api

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"lin-cms-go/core"
)

type UserLoginRequest struct {
	Username string `v:"required#账号不能为空" p:"username"`
	Password string `v:"required#密码不能为空" p:"password"`
}

func load(r *ghttp.Request) *UserLoginRequest {
	var data *UserLoginRequest

	// 数据校验交给后续的service层统一处理。

	if err := r.Parse(&data); err != nil {
		// Validation error.
		if v, ok := err.(*gvalid.Error); ok {

			core.JsonExit(r, 1, v.FirstString())
		}
		// Other error.

		core.JsonExit(r, 1, err.Error())
	}
	return data
}

type UserController struct{}

func NewAdminController() *UserController {
	return &UserController{}
}
func (c *UserController) Users(r *ghttp.Request) {
	req := load(r)
	fmt.Println(req.Username, req.Password)
}
func (c *UserController) Permission() {

}
func (c *UserController) ChangeUserPassword() {

}
func (c *UserController) DeleteUser() {

}
func (c *UserController) UpdateUser() {

}
func (c *UserController) GetAllGroup() {

}
