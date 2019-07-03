/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func BindJson(body []byte, obj interface{}) {
	var err error
	if err = json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
	}
}
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(Md5Encode(plainpwd) + Md5Encode(salt))
}
func ValidatePassword(plainpwd, salt, password string) bool {
	return MakePassword(plainpwd, salt) == password
}
