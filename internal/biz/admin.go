package biz

import (
	"context"
	"github.com/xushuhui/goal/core"
	"golang.org/x/crypto/bcrypt"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/errcode"
)

func GetUsers(req request.GetUsers) (res interface{}, err error) {
	return
}
func ChangeUserPassword(ctx context.Context, userId int, password string) (err error) {
	user, err := data.GetLinUserIdentityByUserId(ctx, userId)
	if err != nil {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	err = data.UpdateLinUserIdentityPassword(ctx, user.Identifier, string(hash))

	if err != nil {
		return
	}
	return
}
func DeleteUser(ctx context.Context, userId int) (err error) {
	_, err = data.GetLinUserById(ctx, userId)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	err = data.SoftDeleteUser(ctx, userId)
	return
}
func UpdateUser(req request.UpdateUser) (err error) {
	return
}
