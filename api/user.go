package api

import (
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

/*

登录

URL:
	POST /cms/user/login

参数：

	"username":"2" //用户名
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

/*

注册


URL:
	POST /cms/user/register

参数：

	"username": "范闲", //用户名
	"group_ids": [2], //分组id
  	"password": "123456", //密码
  	"confirm_password": "123456" //确认密码


返回值：

	"code": 0
	"message": "ok"

*/
func Register(c *gin.Context) {
	var req request.Register
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.Register(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

/*

更新当前登录用户信息


URL:
	PUT /cms/user

参数：

	"nickname": "范闲", //用户名
  	"avatar": "123456", //头像
  	"email": "123456@test.com", //邮箱


返回值：

	"code": 0
	"message": "ok"

*/
func UpdateMe(c *gin.Context) {
	var req request.UpdateMe
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	uid, err := uid(c)
	if err != nil {
		return
	}

	err = services.UpdateMe(req, uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

/*

修改当前登录用户


URL:
	PUT /cms/user/change_password

参数：

	"new_password": "147258",//新密码
  	"confirm_password": "147258", //确认密码
  	"old_password": "123456" //旧密码



返回值：

	"code": 0
	"message": "ok"

*/
func ChangeMyPassword(c *gin.Context) {

	var req request.ChangeMyPassword
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	uid, err := uid(c)
	if err != nil {
		return
	}

	err = services.ChangeMyPassword(req, uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
func uid(c *gin.Context) (uid uint, err error) {
	u, ok := c.Get("uid")

	if !ok {
		c.Error(core.NewErrorMessage(errcode.AuthCheckTokenFail, "uid错误"))
	}
	uid, err = utils.ToUint(u)
	return
}

/*

查询自己拥有的权限


URL:
	GET /cms/user/permissions

参数：



返回值：

	"code": 0
	"message": "ok"

*/
func GetMyPermissions(c *gin.Context) {
	uid, err := uid(c)
	if err != nil {
		c.Error(err)
		return
	}
	data, err := services.GetMyPermissions(uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}

/*

查询自己信息


URL:
	GET /cms/user/information

参数：



返回值：

	"code": 0
	"message": "ok"

*/
func GetMyInfomation(c *gin.Context) {
	uid, err := uid(c)
	if err != nil {
		c.Error(err)
		return
	}
	data, err := services.GetMyInfomation(uid)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}

//TODO: 对cms意义并不大，先不实现
func RefreshToken(c *gin.Context) {

}
