package service

import (
	"fmt"
	"lin-cms-go/app/request"
)

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RegisterResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(param *request.UserLoginRequest) {
	fmt.Println(param)
}
func Register(param *request.UserRegRequest) {
	fmt.Println(param)
}
