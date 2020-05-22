/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package core

const (
	OK int = iota
	DbConnectError
	SqlError
	InvaldParams
	SystemError
	NoUser
	ErrorPassword
	UserExist
	NoBook
)

var errormap = map[int]string{
	OK:             "OK",
	DbConnectError: "db connect error",
	SqlError:       "db sql error",
	InvaldParams:   "error params",
	SystemError:    "system error",
	NoUser:         "no user",
	ErrorPassword:  "pass error",
	UserExist:      "user exist",
	NoBook:         "book no exist",
}

func GetMessage(code int) string {
	return errormap[code]
}
