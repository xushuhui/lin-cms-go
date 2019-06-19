package routers

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//初始化 namespace
	ns := beego.NewNamespace("/cms",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/info", &controllers.UserController{}, "get:GetInfo"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		),
	)
	//注册 namespace
	beego.AddNamespace(ns)
	bookRoute := beego.NewNamespace("/v1",
		beego.NSRouter("/book", &controllers.BookController{}, "post:CreateBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "get:GetBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "put:UpdateBook"),
		beego.NSRouter("/book/:id", &controllers.BookController{}, "delete:DeleteBook"),
		beego.NSRouter("/books", &controllers.BookController{}, "get:GetBooks"),
	)
	beego.AddNamespace(bookRoute)
}
