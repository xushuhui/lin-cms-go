package api

import (
	"github.com/gogf/gf/net/ghttp"
	"lin-cms-go/app/request"
	"lin-cms-go/app/service"
)

type UserController struct{}

func NewAdminController() *UserController {
	return &UserController{}
}
func (c *UserController) Login(r *ghttp.Request) {
	service.Login(request.Login(r))
}
func (c *UserController) Register(r *ghttp.Request) {
	service.Register(request.Register(r))
}

func (c *UserController) Permission(r *ghttp.Request) {

}
func (c *UserController) ChangeUserPassword(r *ghttp.Request) {

}
func (c *UserController) DeleteUser(r *ghttp.Request) {

}
func (c *UserController) UpdateUser(r *ghttp.Request) {

}
func (c *UserController) UpdateSelfUser(r *ghttp.Request) {

}
func (c *UserController) GetSelfInformation(r *ghttp.Request) {

}

func (c *UserController) UpdateSelfPassword(r *ghttp.Request) {

}
func (c *UserController) Refresh(r *ghttp.Request) {

}
