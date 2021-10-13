package biz

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
)

type LinUserUsecase struct {
	linUserRepo LinUserRepo
	linUserGroupRepo LinUserGroupRepo
	linUserPermissionRepo LinUserPermissionRepo
}
type LinUserGroupRepo interface {
	
}
type LinUserPermissionRepo interface{

}
type LinUserRepo interface {
	GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (*ent.LinUserIdentiy, error)
	CreateLinUser(ctx context.Context, username, password, email string, groupId int) error
	GetLinUserById(ctx context.Context, uid int) (*ent.LinUser, error)
	UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) error
	UpdateLinUserIdentityPassword(ctx context.Context, username, password string) error
}

func NewLinUserUsecase(linUserRepo LinUserRepo,linUserGroupRepo LinUserGroupRepo,linUserPermissionRepo LinUserPermissionRepo) *LinUserUsecase {
	return &LinUserUsecase{
		linUserRepo:linUserRepo,
		linUserGroupRepo: linUserGroupRepo,
		linUserPermissionRepo: linUserPermissionRepo,
	}
}
func  Login(ctx context.Context, username, password string) (data map[string]interface{}, err error) {

	// 正确密码验证
	userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
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
	//TODO token
	//token, err := lib.GenerateToken(userIdentityModel.UserID)
	//if err != nil {
	//	return
	//}
	//data = make(map[string]interface{})
	//data["access_token"] = token
	return
}
func (uc *LinUserUsecase) Register(ctx context.Context, req request.Register) (err error) {
	userIdentityModel, err := uc.linUserRepo.GetLinUserIdentityByIdentifier(ctx, req.Username)
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
	err = uc.linUserRepo.CreateLinUser(ctx, req.Username, string(hash), req.Email, req.GroupId)
	return

}
func (uc *LinUserUsecase) UpdateMe(ctx context.Context, req request.UpdateMe, uid int) (err error) {
	_, err = uc.linUserRepo.GetLinUserById(ctx, uid)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	err = uc.linUserRepo.UpdateLinUser(ctx, uid, req.Avatar, req.Nickname, req.Email)
	return
}
func (uc *LinUserUsecase) ChangeMyPassword(ctx context.Context, req request.ChangeMyPassword, uid int) (err error) {
	var username string
	//todo jwt username
	userIdentityModel, err := uc.linUserRepo.GetLinUserIdentityByIdentifier(ctx, username)
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
	err = uc.linUserRepo.UpdateLinUserIdentityPassword(ctx, username, string(hash))
	if err != nil {
		return
	}

	return
}
func (uc *LinUserUsecase) GetMyPermissions(ctx context.Context,uid int) (data map[string]interface{}, err error) {
	user, err := uc.linUserRepo.GetLinUserById(ctx, uid)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	groupModel,err :=  user.QueryLinGroup().First(ctx)
	if err != nil {
		return
	}
	var isRoot bool
	if groupModel.Level == 1 {
		isRoot = true
	}
	fmt.Println(isRoot)

	//data = utils.Struct2MapJson(userModel)
	//data["is_admin"] = isAdmin
	//data["permissions"] = permissions
	return
}
func (uc *LinUserUsecase) GetMyInfomation(ctx context.Context,uid int) (data map[string]interface{}, err error) {
		_, err = uc.linUserRepo.GetLinUserById(ctx, uid)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	return
}
