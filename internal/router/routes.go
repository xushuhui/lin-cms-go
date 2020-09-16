package router

import (
	"github.com/gin-gonic/gin"
	"lin-cms-go/api"
	"lin-cms-go/internal/middleware"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	router.StaticFS("/upload", http.Dir("storage/upload"))

	router.Use(middleware.Cors(), middleware.AccessLog(), middleware.ErrorHandle())

	cms := router.Group("/cms")
	cms.POST("/user/login", api.Login)
	cms.Use(middleware.Auth())
	{
		userRouter := cms.Group("/user")
		adminRouter := cms.Group("/admin")
		logRouter := cms.Group("/log")
		userRouter.POST("/register", api.Register)
		userRouter.PUT("", api.UpdateMe)
		userRouter.PUT("/change_password", api.ChangeMyPassword)
		userRouter.GET("/permissions", api.GetMyPermissions)
		userRouter.GET("/information", api.GetMyInfomation)

		adminRouter.GET("/permissions", api.GetAllPermissions)
		adminRouter.GET("/users", api.GetUsers)
		adminRouter.PUT("/user/:id/password", api.ChangeUserPassword)
		adminRouter.DELETE("/user/:id", api.DeleteUser)
		adminRouter.PUT("/user/:id", api.UpdateUser)

		adminRouter.GET("/group/:id", api.GetGroup)
		adminRouter.PUT("/group/:id", api.UpdateGroup)
		adminRouter.DELETE("/group/:id", api.DeleteGroup)
		adminRouter.GET("/groups", api.GetGroups)
		adminRouter.POST("/group", api.CreateGroup)

		adminRouter.POST("/permission/dispatch", api.DispatchPermission)
		adminRouter.POST("/permissions/dispatch", api.DispatchPermissions)
		adminRouter.POST("/permissions/remove", api.RemovePermissions)

		cms.POST("/file", api.Upload)
		logRouter.GET("", api.GetLogs)
		logRouter.GET("/search", api.SearchLogs)
		logRouter.GET("/users", api.GetLogUsers)
	}

	//需要登录的接口
	return router
}
