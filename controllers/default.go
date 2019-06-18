package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = c.GetString("name")
	c.Data["Email"] = "123@qq.com"
}

//func (c *MainController) Post() {
//	var u *user
//	utils.BindJson(c.Ctx.Input.RequestBody, &u)
//	userModel, _ := models.GetUserById(1)
//
//	fmt.Println(userModel)
//	fmt.Println(beego.AppConfig)
//	c.Data["json"] = core.SetData(userModel)
//	c.ServeJSON()
//}
