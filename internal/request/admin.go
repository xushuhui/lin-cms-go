package request

type GetUsers struct {
	GroupId int `json:"group_id" binding:"required"`
	Pages
}
type Pages struct {
	Count int `json:"count" binding:"required"`
	Page  int `json:"page" binding:"required"`
}
type ChangeUserPassword struct {
	Id              int    `json:"uid"`
	NewPassword     string `json:"new_password" binding:"required" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" binding:"required，eqcsfield=NewPassword" comment:"确认密码"`
}
type UpdateUser struct {
	Id       int   `json:"id" binding:"required"`
	GroupIds []int `json:"group_ids" binding:"required"`
}
type CreateGroup struct {
	Name          string `json:"name"`
	Info          string `json:"info"`
	PermissionIds []int  `json:"permission_ids"`
}
type UpdateGroup struct {
	Name string `json:"name"`
	Info string `json:"info"`
}
