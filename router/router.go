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
	c := api.NewAdminController()
	s.Group("/cms", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)

		s.Group("/admin", func(group *ghttp.RouterGroup) {

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
		group.GET("/log/users", c.GetLogUsers)
		group.POST("/user/login", c.Login)
		group.POST("/user/register", c.Register)
		group.PUT("/user", c.UpdateSelfUser)
		group.PUT("/user/change_password", c.UpdateSelfPassword)
		group.GET("/user/refresh", c.Refresh)
		group.GET("/user/permissions", c.GetSelfPermissions)
		group.GET("/user/information", c.GetSelfInformation)
	})

}
