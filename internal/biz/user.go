package biz

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
)

type LinUserUsecase struct {
	repo LinUserRepo
}
type LinUserRepo interface {
	GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (*ent.LinUserIdentiy, error)
	CreateLinUser(ctx context.Context, username, password, email string, groupId int) error
	GetLinUserById(ctx context.Context, uid int) (*ent.LinUser, error)
	UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) error
	UpdateLinUserIdentityPassword(ctx context.Context, username, password string) error
}

func NewLinUserUsecase(repo LinUserRepo) *LinUserUsecase {
	return &LinUserUsecase{
		repo: repo,
	}
}
func (uc *LinUserUsecase) Login(ctx context.Context, username, password string) (data map[string]interface{}, err error) {

	// 正确密码验证
	userIdentityModel, err := uc.repo.GetLinUserIdentityByIdentifier(ctx, username)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userIdentityModel.Credential), []byte(password))
	if err != nil {
		err = core.NewErrorCode(errcode.ErrorPassWord)
		return
	}
	//token, err := lib.GenerateToken(userIdentityModel.UserID)
	//if err != nil {
	//	return
	//}
	//data = make(map[string]interface{})
	//data["access_token"] = token
	return
}
func (uc *LinUserUsecase) Register(ctx context.Context, req request.Register) (err error) {
	userIdentityModel, err := uc.repo.GetLinUserIdentityByIdentifier(ctx, req.Username)
	if ent.MaskNotFound(err) != nil {
		return err
	}
	if userIdentityModel.ID > 0 {
		err = errors.New("user is found")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	err = uc.repo.CreateLinUser(ctx, req.Username, string(hash), req.Email, req.GroupId)
	return
	//userModel = model.User{
	//	Username:  req.Username,
	//	Email:     req.Email,
	//	UserGroup: model.UserGroup{GroupID: req.GroupId},
	//	UserIdentity: model.UserIdentity{
	//		Identifier: req.Username,
	//		Credential: string(hash),
	//	},
	//}
	//global.DBEngine.Create(&userModel)

}
func (uc *LinUserUsecase) UpdateMe(ctx context.Context, req request.UpdateMe, uid int) (err error) {
	_, err = uc.repo.GetLinUserById(ctx, uid)

	if err != nil {
		return
	}
	err = uc.repo.UpdateLinUser(ctx, uid, req.Avatar, req.Nickname, req.Email)
	return
}
func (uc *LinUserUsecase) ChangeMyPassword(ctx context.Context, req request.ChangeMyPassword, uid int) (err error) {
	var username string
	//todo jwt username
	userIdentityModel, err := uc.repo.GetLinUserIdentityByIdentifier(ctx, username)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userIdentityModel.Credential), []byte(req.OldPassword))
	if err != nil {
		err = core.NewErrorCode(errcode.ErrorPassWord)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	err = uc.repo.UpdateLinUserIdentityPassword(ctx, username, string(hash))
	if err != nil {
		return
	}

	return
}
func (uc *LinUserUsecase) GetMyPermissions(uid uint) (data map[string]interface{}, err error) {
	//var userModel model.User
	//err = global.DBEngine.Preload("UserGroup.Group").First(&userModel, uid).Error
	//if err != nil {
	//	return
	//}
	//
	//var isAdmin bool
	//if userModel.UserGroup.Group.Level == model.GroupLevelRoot {
	//	isAdmin = true
	//}
	//
	//data = utils.Struct2MapJson(userModel)
	//data["is_admin"] = isAdmin
	//data["permissions"] = permissions
	return
}
func (uc *LinUserUsecase) GetMyInfomation(uid uint) (data map[string]interface{}, err error) {
	return
}
