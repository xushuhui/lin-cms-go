/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package models

import (
	"fmt"
	"time"
)

type Book struct {
	ID        int    `json:"id"`
	Title     string `orm:"size(128)" json:"title"`
	Author    string `orm:"size(128)" json:"author"`
	Summary   string `orm:"size(128)" json:"summary"`
	Image     string `orm:"size(128)" json:"image"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`
	DeletedAt int64  `json:"-"`
}

