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
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gsssssn")
	})

	router.POST("/login", api.Login)
	router.Use(middleware.Auth())
	//需要登录的接口
	return router
}
