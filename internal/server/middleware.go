package server

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/xushuhui/goal/core"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data"
	"lin-cms-go/pkg/errcode"
)

func UserLog(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}
	user := biz.LocalUser(c)

	msg := c.Locals("logMessage")
	if msg == nil {
		return nil
	}
	err = data.CreateLog(c.Context(), c.Response().StatusCode(), user.ID, user.Username, user.Username+msg.(string), c.Method(), c.Path(), "")

	return err
}

func SetPermission(name, module string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("permission", biz.Permission{Name: name, Module: module})
		return c.Next()
	}
}

func AdminRequired(c *fiber.Ctx) error {

	isAdmin, err := biz.IsAdmin(c)
	if err != nil {
		return err
	}
	if !isAdmin {
		return core.NewErrorCode(errcode.UserNoPermission)
	}
	return c.Next()
}

func LoginRequired(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodOptions {
		jwtware.New(jwtware.Config{
			SigningKey: []byte("secret"),
		})
		return c.Next()
	}
	return c.Next()
}

func GroupRequired(c *fiber.Ctx) error {
	isAdmin, err := biz.IsAdmin(c)
	if err != nil {
		return err
	}
	if isAdmin {
		return c.Next()
	}

	has, err := biz.UserHasPermission(c)
	if err != nil {
		return err
	}
	if !has {
		return core.NewErrorCode(errcode.UserNoPermission)
	}
	return c.Next()
}

//
//func (a *Auth) RefreshRequired(ctx *gin.Context) {
//	refreshToken, tokenErr := getHeaderToken(ctx)
//	if tokenErr != nil {
//		ctx.Error(tokenErr)
//		ctx.Abort()
//		return
//	}
//	payload, err := a.JWT.VerifyRefreshToken(refreshToken)
//	if err != nil {
//		ctx.Error(err)
//		ctx.Abort()
//	} else {
//		userId := payload.Identity
//		user, errs := a.UserService.GetUserById(userId)
//		if errs != nil {
//			ctx.Error(response.UnifyResponse(10021, ctx))
//			ctx.Abort()
//		} else {
//			ctx.Set("currentUser", user)
//			ctx.Next()
//		}
//	}
//}

//func getHeaderToken(ctx *gin.Context) (string, error) {
//	authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
//	if authorizationHeader == "" {
//		return "", response.UnifyResponse(10012, ctx)
//	}
//	fields := strings.Fields(authorizationHeader)
//	if fields[0] != AuthorizationTypeBearer {
//		return "", response.UnifyResponse(10013, ctx)
//	}
//	tokenString := fields[1]
//	return tokenString, nil
//}
