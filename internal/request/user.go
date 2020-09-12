package request

type Login struct {
	Phone    string `json:"phone" binding:"required,mobile" comment:"手机号"`
	Password string `json:"password" binding:"required" comment:"密码"`
}
