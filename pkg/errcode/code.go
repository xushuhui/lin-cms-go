package errcode

const (
	SUCCESS = iota
	InvalidParams
	ServerError
	AuthCheckTokenFail
	AuthCheckTokenTimeout
	ErrorAuthToken
	TimeoutAuthToken
	ErrorPassWord
)
