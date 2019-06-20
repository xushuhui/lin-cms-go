package main

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/core"
	_ "lin-cms-beego/routers"
)

func main() {
	core.InitEnv()

	beego.Run()
}
