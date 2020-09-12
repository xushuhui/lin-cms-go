package services

import (
	"golang.org/x/crypto/bcrypt"
	"lin-cms-go/internal/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/lib"
)

func Login(req request.Login) (data string, err error) {

	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	if err != nil {
		return
	}
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		err = core.NewError(errcode.ErrorPassWord)
		return
	}
	data, err = lib.GenerateToken(userModel.Id)
	if err != nil {
		return
	}

	return
}
