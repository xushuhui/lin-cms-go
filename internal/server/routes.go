package server

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/api"
)

func InitRoute(app *fiber.App, user *api.User) {
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	app.Static("/upload", "storage/upload")
	app.Get("/", user.Hello)
	cms := app.Group("/cms")
	cms.Post("/file", api.Upload)
	cms.Post("/user/login", user.Login)

	{
		userRouter := cms.Group("/user")
		adminRouter := cms.Group("/admin")
		logRouter := cms.Group("/log")
		userRouter.Post("/register", user.Register)
		userRouter.Put("/", user.UpdateMe)
		userRouter.Put("/change_password", user.ChangeMyPassword)
		userRouter.Get("/permissions", user.GetMyPermissions)
		userRouter.Get("/information", user.GetMyInfomation)

		adminRouter.Get("/permissions", api.GetAllPermissions)
		adminRouter.Get("/users", api.GetUsers)
		adminRouter.Put("/user/:id/password", api.ChangeUserPassword)
		adminRouter.Delete("/user/:id", api.DeleteUser)
		adminRouter.Put("/user/:id", api.UpdateUser)

		adminRouter.Get("/group/:id", api.GetGroup)
		adminRouter.Put("/group/:id", api.UpdateGroup)
		adminRouter.Delete("/group/:id", api.DeleteGroup)
		adminRouter.Get("/groups", api.GetGroups)
		adminRouter.Post("/group", api.CreateGroup)

		adminRouter.Post("/permission/dispatch", api.DispatchPermission)
		adminRouter.Post("/permissions/dispatch", api.DispatchPermissions)
		adminRouter.Post("/permissions/remove", api.RemovePermissions)

		logRouter.Get("", api.GetLogs)
		logRouter.Get("/search", api.SearchLogs)
		logRouter.Get("/users", api.GetLogUsers)
	}

}
