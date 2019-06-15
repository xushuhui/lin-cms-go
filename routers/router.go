package routers

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
