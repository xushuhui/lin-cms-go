package api

import (
	"github.com/gin-gonic/gin"
	"lin-cms-go/internal/request"
	"lin-cms-go/internal/services"
	"lin-cms-go/pkg/core"
)

/*

查询所有可分配的权限

URL:
GET	/cms/admin/permissions

参数：

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func GetAllPermissions(c *gin.Context) {
	data, err := services.GetAllPermissions()
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}

//TODO: 没找到必须实现该接口的业务场景，而且功能和分配多个权限重复，开发待定
func DispatchPermission(c *gin.Context) {

}

/*

分配多个权限

URL:
	POST /cms/admin/permissions/dispatch

参数：
	"group_id": 5, 分组 id
  	"permission_ids": [4, 5] 权限 id

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func DispatchPermissions(c *gin.Context) {
	var req request.DispatchPermissions
	if err := core.ParseRequest(c, req); err != nil {
		c.Error(err)
		return
	}
	err := services.DispatchPermissions(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

/*

删除多个权限

URL:
	POST /cms/admin/permissions/remove

参数：
	"group_id": 5, 分组 id
  	"permission_ids": [4, 5] 权限 id

返回值：

	"code": 0
	"message": "ok"
	"data": {"

	"}
*/
func RemovePermissions(c *gin.Context) {
	var req request.RemovePermissions
	if err := core.ParseRequest(c, req); err != nil {
		c.Error(err)
		return
	}
	err := services.RemovePermissions(req)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
