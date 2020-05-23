package request

import (
	"lin-cms-go/core"

	"github.com/gogf/gf/net/ghttp"
)

type UserLoginRequest struct {
	username string `v:"required#账号不能为空"`
	password string `v:"required#密码不能为空"`
}
func init(){
	
}
func load(r *ghttp.Request)  {
	var data *UserLoginRequest
    // 这里没有使用Parse而是仅用GetStruct获取对象，
    // 数据校验交给后续的service层统一处理。
    if err := r.GetStruct(&data); err != nil {
        core.JsonExit(r, 1, err.Error())
    }
}