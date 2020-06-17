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
	s.Use(middleware.CORS)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})

	s.Group("/cms", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		s.Group("/admin", func(group *ghttp.RouterGroup) {
			c := api.NewAdminController()
			group.GET("/users", c.Users)
			group.GET("/permission", c.Permission)
			group.PUT("/user/:id/password", c.ChangeUserPassword)
			group.DELETE("/user/:id", c.DeleteUser)
			group.PUT("/user/:id", c.UpdateUser)
			group.GET("/group/all", c.GetAllGroup)

		})
	})

}
