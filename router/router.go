package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"lin-cms-go/app/api"
	"lin-cms-go/app/api/hello"
	"lin-cms-go/app/middleware"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		c := api.NewAdminController()
		group.POST("/", c.GetUser)
	})
}
