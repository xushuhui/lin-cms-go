package server

import (
	"fmt"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data"
	"lin-cms-go/pkg/errcode"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/xushuhui/goal/core"

	"github.com/gofiber/fiber/v2"
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

type Permission struct {
	Name   string `json:"name"`
	Module string `json:"module"`
}

func SetPermission(name, module string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("permission", Permission{Name: name, Module: module})
		return c.Next()
	}
}

func AdminRequired(c *fiber.Ctx) error {
	fmt.Println("-------------------admin require middle+")
	fmt.Println(c.Locals("permission"))
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

	local := c.Locals("permission")
	if local == nil {
		return core.NewErrorCode(errcode.UserPermissionRequired)
	}
	p := local.(Permission)
	fmt.Println(p)
	return c.Next()
}

//
//func (a *Auth) GroupRequired(c *gin.Context) {
//	user, _ := c.Get("currentUser")
//	userId := user.(model.User).ID
//	// admin直接通过
//	admin, _ := a.UserService.IsAdmin(userId)
//	if admin {
//		c.Next()
//	} else {
//		meta, ok := c.Get("meta")
//		if !ok {
//			return
//		}
//		routeMeta := meta.(router.Meta)
//		if !routeMeta.Mount {
//			c.Next()
//		} else {
//			hasPermission := a.GroupService.GetUserHasPermission(userId, routeMeta)
//			if !hasPermission {
//				_ = c.Error(response.UnifyResponse(10001, c))
//				c.Abort()
//				return
//			} else {
//				c.Next()
//			}
//		}
//	}
//}

//
//func (a *Auth) AdminRequired(c *gin.Context) {
//	if err := a.mountUser(c); err != nil {
//		_ = c.Error(err)
//		c.Abort()
//		return
//	}
//	user, _ := c.Get("currentUser")
//	currentUser := user.(model.User)
//	admin, err := a.UserService.IsAdmin(currentUser.ID)
//	if err != nil {
//		_ = c.Error(response.UnifyResponse(10021, c))
//		c.Abort()
//		return
//	}
//	if admin {
//		c.Next()
//	} else {
//		_ = c.Error(response.UnifyResponse(10001, c))
//		c.Abort()
//		return
//	}
//}
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
