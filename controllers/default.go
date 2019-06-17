package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lin-cms-beego/core"
	"lin-cms-beego/models"
	"lin-cms-beego/utils"
)

type MainController struct {
	beego.Controller
}
type user struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (c *MainController) Get() {
	c.Data["Website"] = c.GetString("name")
	c.Data["Email"] = "123@qq.com"
}
func (c *MainController) Post() {
	var u *user
	utils.BindJson(c.Ctx.Input.RequestBody, &u)
	userModel, _ := models.GetUserById(1)

	fmt.Println(userModel)
	fmt.Println(beego.AppConfig)
	c.Data["json"] = core.SetData(userModel)
	c.ServeJSON()
}
