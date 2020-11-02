package api

import (
	"github.com/gin-gonic/gin"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/utils"
)

/*

查询所有用户


URL:
	GET	/cms/admin/users

参数：
	group_id : 0 //分组id
	count:10, //分页数
	page:1 //分页值

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func GetUsers(c *gin.Context) {
	var req request.GetUsers
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	data, err := services.GetUsers(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}

/*

修改用户密码

URL:
	PUT	/cms/admin/user/{id}/password

参数：

	"new_password": "147258", //新密码
	"confirm_password": "147258" //确认密码

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func ChangeUserPassword(c *gin.Context) {
	var req request.ChangeUserPassword
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.ChangeUserPassword(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

/*

删除用户

URL:
	DELETE 	/cms/admin/user/{id}

参数：


返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func DeleteUser(c *gin.Context) {

	id, err := utils.StringToInt(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	err = services.DeleteUser(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return

}

/*

管理员更新用户信息

URL:
	PUT  /cms/admin/user/{id}

参数：

	group_ids: [1,2] //分组 id，数组，如 [1,2,3]

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func UpdateUser(c *gin.Context) {
	var req request.UpdateUser
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := services.UpdateUser(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
