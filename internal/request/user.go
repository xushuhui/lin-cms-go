package request

type Login struct {
	Username string `json:"username" validate:"required,min=3,max=32" comment:"用户名"`
	Password string `json:"password" validate:"required" comment:"密码"`
}
type Register struct {
	Username        string `json:"username" validate:"required" comment:"用户名"`
	Email           string `json:"email" validate:"required" comment:"邮箱"`
	Password        string `json:"password" validate:"required" comment:"密码"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqcsfield=Password" comment:"确认密码"`
	GroupId         int    `json:"group_id" validate:"required" comment:"分组id"`
}
type UpdateMe struct {
	Nickname string `json:"nickname" validate:"required" comment:"昵称"`
	Avatar   string `json:"avatar" validate:"required" comment:"头像"`
	Email    string `json:"email" validate:"required" comment:"邮箱"`
}
type ChangeMyPassword struct {
	NewPassword     string `json:"new_password" validate:"required" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqcsfield=NewPassword" comment:"确认密码"`
	OldPassword     string `json:"old_password" validate:"required" comment:"旧密码"`
}
