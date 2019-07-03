/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package routers

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//初始化 namespace
	cmsRoute := beego.NewNamespace("/cms",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/info", &controllers.UserController{}, "get:GetInfo"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		),
	)
	//注册 namespace
	//beego.AddNamespace(ns)
	bookRoute := beego.NewNamespace("/v1",
		beego.NSRouter("/book", &controllers.BookController{}, "post:CreateBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "get:GetBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "put:UpdateBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "delete:DeleteBook"),
		beego.NSRouter("/books", &controllers.BookController{}, "get:GetBooks"),
	)
	beego.AddNamespace(cmsRoute, bookRoute)
}
