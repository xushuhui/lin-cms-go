package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"lin-cms-beego/core"
	"lin-cms-beego/models"
	"lin-cms-beego/service"
	"lin-cms-beego/utils"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

type user struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

//登录
func (c *UserController) Login() {
	var u *user
	utils.BindJson(c.Ctx.Input.RequestBody, &u)
	userModel, _ := models.GetUserByNickname(u.Nickname)
	if userModel == nil {
		c.Data["json"] = core.Fail(core.CodeNoUser)
		c.ServeJSON()
	}

	fmt.Printf(" %T\n", userModel.Id)
	if !utils.ValidatePassword(u.Password, "test", userModel.Password) {
		c.Data["json"] = core.Fail(core.CodeErrorPassword)
		c.ServeJSON()
	}
	tokenString := service.MakeToken(userModel)
	usertoken := map[string]interface{}{"token": tokenString}
	c.Data["json"] = core.SetData(usertoken)
	c.ServeJSON()
}
func (c *UserController) GetInfo() {

	authString := c.Ctx.Input.Header("Authorization")
	token, err := service.ParseToken(authString)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int(claims["user_id"].(float64))
		user, _ := models.GetUserById(userId)
		c.Data["json"] = core.SetData(user)
		c.ServeJSON()
	} else {
		fmt.Println(err)
	}
}
