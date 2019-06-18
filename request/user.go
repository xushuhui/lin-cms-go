package request

type UserLogin struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
type UserRegister struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
