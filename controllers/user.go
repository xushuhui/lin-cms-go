/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"lin-cms-beego/core"
	"lin-cms-beego/models"
	"lin-cms-beego/request"
	"lin-cms-beego/service"
	"lin-cms-beego/utils"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

//登录
func (c *UserController) Login() {
	var u *request.UserLogin
	utils.BindJson(c.Ctx.Input.RequestBody, &u)
	userModel, _ := models.GetUserByNickname(u.Nickname)
	if userModel == nil {
		c.Data["json"] = core.Fail(core.CodeNoUser)
		c.ServeJSON()
	}

	//fmt.Printf(" %T\n", userModel.Id)
	if !utils.ValidatePassword(u.Password, "test", userModel.Password) {
		c.Data["json"] = core.Fail(core.CodeErrorPassword)
		c.ServeJSON()
	}
	tokenString := service.MakeToken(userModel)
	usertoken := map[string]interface{}{"token": tokenString}
	c.Data["json"] = core.SetData(usertoken)
	c.ServeJSON()
}

//用户信息
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

//注册
func (c *UserController) Register() {
	var u *request.UserRegister
	utils.BindJson(c.Ctx.Input.RequestBody, &u)
	userModel, _ := models.GetUserByNickname(u.Nickname)
	if userModel != nil {
		c.Data["json"] = core.Fail(core.CodeUserExist)
		c.ServeJSON()
	}
	userModel, _ = models.GetUserByEmail(u.Nickname)
	if userModel != nil {
		c.Data["json"] = core.Fail(core.CodeUserExist)
		c.ServeJSON()
	}

	user := models.User{
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: utils.MakePassword(u.Nickname, "test"),
	}

	_, err := models.AddUser(&user)
	if err != nil {
		c.Data["json"] = core.Fail(core.CodeSqlError)
		c.ServeJSON()
	}
	c.Data["json"] = core.Succeed()
	c.ServeJSON()
}
