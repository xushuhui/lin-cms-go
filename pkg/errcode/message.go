package errcode

const (
	SUCCESS = iota
	InvalidParams
	DBError
	ServerError
	AuthCheckTokenFail
	AuthCheckTokenTimeout
	ErrorAuthToken
	TimeoutAuthToken
	ErrorPassWord
	UserExist
)

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	InvalidParams:         "请求参数错误",
	DBError:               "数据库错误",
	ServerError:           "系统异常，请联系管理员！",
	AuthCheckTokenFail:    "Token鉴权失败",
	AuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:        "Token错误",
	ErrorPassWord:         "密码错误",
	UserExist:             "用户已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InvalidParams]
}
