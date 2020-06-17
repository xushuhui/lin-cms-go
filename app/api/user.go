package api

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"lin-cms-go/app/request"
)

type UserController struct{}

func NewAdminController() *UserController {
	return &UserController{}
}
func (c *UserController) Login(r *ghttp.Request) {
	param := request.Login(r)
	fmt.Println(param)
}
func (c *UserController) Users(r *ghttp.Request) {

}
func (c *UserController) Permission() {

}
func (c *UserController) ChangeUserPassword() {

}
func (c *UserController) DeleteUser() {

}
func (c *UserController) UpdateUser() {

}
func (c *UserController) GetInformation() {

}
