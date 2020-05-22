/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package models



type User struct {
	ID       int    `json:"id"`
	Nickname string `orm:"size(128)" json:"nickname"`
	Password string `orm:"size(128)" json:"-"`
	Super    int    `json:"super"`
	Active   int    `json:"active"`
	Email    string `orm:"size(128)" json:"email"`
}

