package api

import (
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"

	"github.com/gin-gonic/gin"
)

/*

登录

URL:
	/login

POST参数：

	"phone":"2" //手机号
	"password": "1" //密码


返回值：

	"code": 0
	"message": "ok"

*/
func Login(c *gin.Context) {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.Login(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
