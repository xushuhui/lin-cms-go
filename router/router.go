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
		c := api.NewAdminController()
		s.Group("/admin", func(group *ghttp.RouterGroup) {

			group.GET("/users", c.Users)
			group.GET("/permission", c.Permission)
			group.PUT("/user/:id/password", c.ChangeUserPassword)
			group.DELETE("/user/:id", c.DeleteUser)
			group.PUT("/user/:id", c.UpdateUser)
			group.GET("/group/all", c.GetAllGroup)
			group.GET("/group/:id", c.GetGroup)
			group.POST("/group", c.CreateGroup)
			group.POST("/group/:id", c.UpdateGroup)
			group.DELETE("/group/:id", c.DeleteGroup)
			group.POST("/permission/dispatch", c.DispatchPermission)
			group.POST("/permission/dispatch/batch", c.DispatchPermissions)
			group.POST("/permission/remove", c.RemovePermissions)

		})
		group.POST("/file", c.Upload)
		group.GET("/log", c.GetLogs)
		group.GET("/log/search", c.SearchLogs)

	})

}
