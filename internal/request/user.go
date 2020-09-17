package request

type Login struct {
	Username string `json:"username" binding:"required" comment:"用户名"`
	Password string `json:"password" binding:"required" comment:"密码"`
}
type Register struct {
	Username        string `json:"username" binding:"required" comment:"用户名"`
	Email           string `json:"email" binding:"required" comment:"邮箱"`
	Password        string `json:"password" binding:"required" comment:"密码"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqcsfield=Password" comment:"确认密码"`
	GroupId         uint   `json:"group_id" binding:"required" comment:"分组id"`
}
type UpdateMe struct {
	Nickname string `json:"nickname" binding:"required" comment:"昵称"`
	Avatar   string `json:"avatar" binding:"required" comment:"头像"`
	Email    string `json:"email" binding:"required" comment:"邮箱"`
}
type ChangeMyPassword struct {
	NewPassword     string `json:"new_password" binding:"required" comment:"新密码"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqcsfield=NewPassword" comment:"确认密码"`
	OldPassword     string `json:"old_password" binding:"required" comment:"旧密码"`
}
