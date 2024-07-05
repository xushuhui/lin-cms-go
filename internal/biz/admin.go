package biz

import (
	"context"
)

type LinUserRepo interface{
	ListUser(ctx context.Context, page, size int) ([]*LinUser, int64,error)
	GetUser(ctx context.Context, userId int) (*LinUser, error)
	CreateUser(ctx context.Context, user *LinUser) error
	ChangeUserPassword(ctx context.Context, userId int, password string) error
	UpdateUser(ctx context.Context,  user *LinUser) error
	GetUserIdentityByIdentifier(ctx context.Context, identifier string) (*LinUser, error)
	// userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
}
type LinUser struct {
	ID int64
	Username string
	Password string
	Phone string
}
func GetUsers(ctx context.Context, groupId int, page, size int) (res interface{}, err error) {
	// list, err := data.NewPaging(page, size).ListUserByGroupId(ctx, groupId)
	// res = list
	return
}

func ChangeUserPassword(ctx context.Context, userId int, password string) (err error) {
	// user, err := data.GetLinUserIdentityByUserId(ctx, userId)
	// if err != nil {
	// 	return
	// }
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return
	// }
	// err = data.UpdateLinUserIdentityPassword(ctx, user.Identifier, string(hash))
	// if err != nil {
	// 	return
	// }
	return
}

func DeleteUser(ctx context.Context, userId int) (err error) {
	//_, err = data.GetLinUserById(ctx, userId)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	//err = data.SoftDeleteUser(ctx, userId)
	return
}

func UpdateUser(ctx context.Context, userId int, groupId []int32) (err error) {
	//_, err = data.GetLinUserById(ctx, userId)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	// err = data.AddLinUserGroupIDs(ctx, userId, groupId)
	return
}
