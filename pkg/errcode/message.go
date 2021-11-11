package errcode

import "github.com/xushuhui/goal/core"

const (
	AuthCheckTokenFail       = 10000
	AuthCheckTokenTimeout    = 10001
	ErrorAuthToken           = 10002
	TimeoutAuthToken         = 10003
	ErrorPassWord            = 10004
	UserFound                = 10005
	UserNotFound             = 10006
	BookNotFound             = 20000
	BookTitleRepetition      = 20001
	GroupNotFound            = 30000
	GroupFound               = 30001
	RootGroupNotAllowDelete  = 30002
	GuestGroupNotAllowDelete = 30003
	PermissionNotFound       = 30004
)

func init() {
	core.CodeMapping = map[int]string{
		AuthCheckTokenFail:       "Token鉴权失败",
		AuthCheckTokenTimeout:    "Token已超时",
		ErrorAuthToken:           "Token错误",
		ErrorPassWord:            "密码错误",
		UserFound:                "用户已存在",
		UserNotFound:             "用户不存在",
		BookNotFound:             "书籍不存在",
		BookTitleRepetition:      "书籍标题重复",
		GroupNotFound:            "分组不存在",
		RootGroupNotAllowDelete:  "root分组不允许删除",
		GuestGroupNotAllowDelete: "guest分组不允许删除",
		GroupFound:               "分组已存在",
		PermissionNotFound:       "权限不存在",
	}
}
