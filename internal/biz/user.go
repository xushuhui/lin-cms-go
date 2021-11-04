package biz

import (
	"context"
	"errors"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"

	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/lib"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

func Login(ctx context.Context, username, password string) (res map[string]interface{}, err error) {

	// 正确密码验证
	userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
	if model.IsNotFound(err) {
		err = core.NotFoundError(errcode.UserNotFound)
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
	user, err := data.GetLinUserById(ctx, userIdentityModel.UserID)
	if err != nil {
		return
	}
	// jwt
	claims := make(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, err := lib.GenerateJwt(claims)
	if err != nil {
		return
	}
	res = make(map[string]interface{})
	res["access_token"] = token
	return
}
func Register(ctx context.Context, req request.Register) (err error) {
	userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, req.Username)
	if model.MaskNotFound(err) != nil {
		return err
	}
	if userIdentityModel != nil && userIdentityModel.ID > 0 {
		err = errors.New("user is found")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	err = data.CreateLinUser(ctx, req.Username, string(hash), req.Email, req.GroupId)
	return

}
func UpdateMe(ctx context.Context, req request.UpdateMe, uid int) (err error) {
	_, err = data.GetLinUserById(ctx, uid)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	err = data.UpdateLinUser(ctx, uid, req.Avatar, req.Nickname, req.Email)
	return
}
func ChangeMyPassword(ctx context.Context, req request.ChangeMyPassword, username string) (err error) {

	userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
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
	err = data.UpdateLinUserIdentityPassword(ctx, username, string(hash))
	if err != nil {
		return
	}

	return
}
func GetMyPermissions(ctx context.Context, uid int) (res map[string]interface{}, err error) {
	user, err := data.GetLinUserById(ctx, uid)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	if err != nil {
		return
	}
	groupModel, err := user.QueryLinGroup().First(ctx)
	if err != nil {
		return
	}
	var isRoot bool
	if groupModel.Level == 1 {
		isRoot = true
	}

	res = make(map[string]interface{})
	res["is_admin"] = isRoot
	//data["permissions"] = permissions
	return
}
type LinUser struct {
	model.LinUser
}
func GetMyInfomation(ctx context.Context, uid int) (res LinUser, err error) {
	usermodel, err := data.GetLinUserById(ctx, uid)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.UserNotFound)
		return
	}
	res = LinUser{*usermodel}
	return
}
