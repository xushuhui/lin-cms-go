package main

import (
	_ "lin-cms-go/boot"
	_ "lin-cms-go/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
