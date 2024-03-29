package request

type GetUsers struct {
	GroupId int `json:"group_id" `
}

type Pages struct {
	Count int `json:"count" validate:"required" comment:"数量"`
	Page  int `json:"page" validate:"required" comment:"页码"`
}

type ChangeUserPassword struct {
	Id              int    `json:"uid"`
	NewPassword     string `json:"new_password" validate:"required" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" validate:"required，eqcsfield=NewPassword" comment:"确认密码"`
}

type UpdateUser struct {
	Id       int   `json:"id" validate:"required"`
	GroupIds []int `json:"group_ids" validate:"required"`
}

type CreateGroup struct {
	Name          string `json:"name" validate:"required" comment:"分组名称"`
	Info          string `json:"info" validate:"required" comment:"分组信息"`
	PermissionIds []int  `json:"permission_ids"`
}

type UpdateGroup struct {
	Name string `json:"name" validate:"required" comment:"分组名称" `
	Info string `json:"info" validate:"required" comment:"分组信息" `
}
