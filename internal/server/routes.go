package server

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"lin-cms-go/api"
)

func InitRoute(app *fiber.App) {
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	app.Static("/upload", "storage/upload")
	app.Get("/", api.Hello)
	cms := app.Group("/cms")
	v1 := app.Group("/v1")
	cms.Post("/file", api.Upload)
	cms.Post("/user/login", api.Login)
	cms.Post("/user/register", api.Register)
	//cms.Use(jwtware.New(jwtware.Config{
	//	SigningKey: []byte("secret"),
	//}))

	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	{
		userRouter := cms.Group("/user")
		adminRouter := cms.Group("/admin")
		logRouter := cms.Group("/log")
		bookRouter := v1.Group("/book")

		userRouter.Put("/", api.UpdateMe)
		userRouter.Put("/change_password", api.ChangeMyPassword)
		userRouter.Get("/permissions", api.GetMyPermissions)
		userRouter.Get("/information", api.GetMyInfomation)

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

		logRouter.Get("/", api.GetLogs)
		logRouter.Get("/search", api.SearchLogs)
		logRouter.Get("/users", api.GetLogUsers)

		bookRouter.Get("/", api.GetBooks)
		bookRouter.Get("/:id", api.GetBook)
		bookRouter.Put("/:id", api.UpdateBook)
		bookRouter.Post("/", api.CreateBook)
		bookRouter.Delete("/:id", api.DeleteBook)
	}

}
