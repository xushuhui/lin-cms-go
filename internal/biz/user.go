package biz

import (
	"context"
	"time"

	"lin-cms-go/api"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Userusecase struct {
	log *log.Helper
	ur  LinUserRepo
}
type UserClaims struct {
	*jwtv4.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func (u *Userusecase) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginReply, error) {
	userIdentity, err := u.ur.GetUserIdentityByIdentifier(ctx, req.Username)
	if err != nil {
		return nil, errors.Unauthorized("USER_NOT_FOUND", "用户不存在")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userIdentity.Password), []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.Unauthorized("PASSWORD_INCORRECT", "密码错误")
		}
		return nil, errors.InternalServer("INTERNAL_ERROR", "密码验证时发生内部错误")
	}

	accessToken, refreshToken, err := generateToken(ctx, userIdentity.ID)
	if err != nil {
		return nil, err
	}
	return &api.LoginReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(ctx context.Context, userID int64) (string, string, error) {
	secret := "lincms"
	claims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, &UserClaims{
		UserID: userID,
		RegisteredClaims: &jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Issuer:    "lincms",
			Subject:   "user",
		},
	})

	accessToken, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", "", errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}
	refreshToken, err := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, jwtv4.RegisteredClaims{
		ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		Issuer:    secret,
	}).SignedString(secret)
	if err != nil {
		return "", "", errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}

	return accessToken, refreshToken, nil
}

func (u *Userusecase) Register(ctx context.Context, req *api.CreateUserRequest) error {
	userIdentity, err := u.ur.GetUserIdentityByIdentifier(ctx, req.Username)
	if err != nil {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}
	if userIdentity.ID > 0 {
		return errors.InternalServer("INTERNAL_ERROR", "用户已经存在")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}
	userIdentity = &LinUser{
		Username: req.Username,
		Password: string(hash),
		Phone:    req.Phone,
	}
	err = u.ur.CreateUser(ctx, userIdentity)
	if err != nil {
		return errors.InternalServer("INTERNAL_ERROR", "服务器内部错误")
	}

	return nil
}

func UpdateMe(ctx context.Context, req *api.UpdateMeRequest, uid int) (err error) {
	// _, err = data.GetLinUserById(ctx, uid)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinUserById ")
	// 	return
	// }
	// err = data.UpdateLinUser(ctx, uid, req.Avatar, req.Nickname, req.Email)
	// if err != nil {
	// 	err = errors.Wrap(err, "UpdateLinUser ")
	// 	return
	// }
	return
}

func ChangeMyPassword(ctx context.Context, req *api.UpdateMyPasswordRequest, username string) (err error) {
	// userIdentityModel, err := data.GetLinUserIdentityByIdentifier(ctx, username)
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinUserIdentityByIdentifier ")
	// 	return err
	// }
	// err = bcrypt.CompareHashAndPassword([]byte(userIdentityModel.Credential), []byte(req.OldPassword))
	// if err == bcrypt.ErrMismatchedHashAndPassword {
	// 	err = core.ParamsError(errcode.ErrorPassWord)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "CompareHashAndPassword ")
	// 	return
	// }
	// hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	err = errors.Wrap(err, "GenerateFromPassword ")
	// 	return
	// }
	// err = data.UpdateLinUserIdentityPassword(ctx, username, string(hash))
	// if err != nil {
	// 	err = errors.Wrap(err, "UpdateLinUserIdentityPassword ")
	// 	return
	// }

	return
}

func GetMyPermissions(ctx context.Context, uid int) (res map[string]interface{}, err error) {
	// user, err := data.GetLinUserById(ctx, uid)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinUserById ")
	// 	return
	// }
	// groupModel, err := user.QueryLinGroup().First(ctx)
	// if err != nil {
	// 	err = errors.Wrap(err, "user.QueryLinGroup ")
	// 	return
	// }
	// var isRoot bool
	// if groupModel.Level == 1 {
	// 	isRoot = true
	// }
	//
	// res = make(map[string]interface{})
	// res["is_admin"] = isRoot
	// data["permissions"] = permissions
	return
}

func GetMyInfomation(ctx context.Context, uid int) (res LinUser, err error) {
	// usermodel, err := data.GetLinUserById(ctx, uid)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.UserNotFound)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinUserById ")
	// 	return
	// }
	// res = LinUser{*usermodel}
	return
}
