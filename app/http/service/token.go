/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"lin-cms-gin/models"
	"strings"
	"time"
)

//MakeToken 带权限创建令牌
func MakeToken(user *models.User) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 480).Unix(),
	})

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// ParseToken parse JWT token in http header.
func ParseToken(authString string) (t *jwt.Token, err error) {

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		//return nil
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		panic(err)
	}
	if !token.Valid {
		fmt.Println("vaild token")
	}

	return token, nil
}
