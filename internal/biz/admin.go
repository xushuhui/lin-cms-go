package biz

import (
	"context"
	"lin-cms-go/api"

	"github.com/go-kratos/kratos/v2/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LinUserRepo interface {
	ListUser(ctx context.Context, page, size int32, groupId int64) ([]*LinUser, int64, error)
	GetUser(ctx context.Context, userId int) (*LinUser, error)
	CreateUser(ctx context.Context, user *LinUser) error
	ChangeUserPassword(ctx context.Context, userId int, password string) error
	UpdateUser(ctx context.Context, user *LinUser) error
	GetUserByUsername(ctx context.Context, username string) (*LinUser, error)
	// userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
}
type LinUser struct {
	ID       int64
	Username string
	Password string
	Phone    string
}

func outLinUser(v *LinUser) *api.LinUser {
	return &api.LinUser{
		Id:       v.ID,
		Username: v.Username,
		Phone:    v.Phone,
	}

}
func outLinUsers(list []*LinUser) []*api.LinUser {
	items := make([]*api.LinUser, 0)
	for _, v := range list {
		items = append(items, outLinUser(v))
	}
	return items
}
func (u *LinUserusecase) ListUser(ctx context.Context, req *api.ListLinUserRequest) (*api.ListLinUserReply, error) {
	list, total, err := u.ur.ListUser(ctx, req.Page, req.Size, req.GroupId)
	if err != nil {
		return nil, err
	}

	return &api.ListLinUserReply{
		List:  outLinUsers(list),
		Total: uint32(total),
	}, nil

}
func (u *LinUserusecase) CreateUser(ctx context.Context, req *api.CreateLinUserRequest) error {
	userIdentity, err := u.ur.GetUserByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if userIdentity != nil && userIdentity.ID > 0 {
		return errors.InternalServer("INTERNAL_ERROR", "用户已经存在")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}
	user := &LinUser{
		Username: req.Username,
		Password: string(hash),
		Phone:    req.Phone,
	}
	err = u.ur.CreateUser(ctx, user)
	if err != nil {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}

	return nil
}

func ChangeUserPassword(ctx context.Context, userId int, password string) error {
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
	return nil
}

func DeleteUser(ctx context.Context, userId int) error {
	//_, err = data.GetLinUserById(ctx, userId)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	//err = data.SoftDeleteUser(ctx, userId)
	return nil
}

func UpdateUser(ctx context.Context, userId int, groupId []int32) (err error) {
	//_, err = data.GetLinUserById(ctx, userId)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	// err = data.AddLinUserGroupIDs(ctx, userId, groupId)
	return nil
}
