package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/pkg/enum"
	"lin-cms-go/pkg/errcode"

	"github.com/xushuhui/goal/core"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xushuhui/goal/utils"
)

func LocalUser(c *fiber.Ctx) (user model.LinUser) {
	local := c.Locals("user")
	if local == nil {
		return
	}
	jwtToken := local.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	bytes, _ := utils.JSONEncode(claims["user"])

	utils.JSONDecode(bytes, &user)
	return
}

func IsAdmin(c *fiber.Ctx) (is bool, err error) {
	user := LocalUser(c)
	if user.ID == 0 {
		err = core.UnAuthenticatedError(errcode.ErrorAuthToken)
		return
	}
	u, err := data.GetLinUserWithGroupById(c.Context(), user.ID)
	if err != nil {
		return
	}

	for _, v := range u.Edges.LinGroup {
		if v.Level == enum.ROOT {
			is = true
		}
	}
	return
}

type Permission struct {
	Name   string `json:"name"`
	Module string `json:"module"`
}

func UserHasPermission(c *fiber.Ctx) (has bool, err error) {
	user := LocalUser(c)

	u, err := data.GetLinUserWithGroupById(context.Background(), user.ID)
	if model.IsNotFound(err) {
		err = core.NotFoundError(errcode.UserNotFound)
		return
	}

	local := c.Locals("permission")
	if local == nil {
		return false, core.UnAuthenticatedError(errcode.UserPermissionRequired)
	}
	userPermission := local.(Permission)
	var ps []model.LinPermission
	for _, v := range u.Edges.LinGroup {
		for _, p := range v.Edges.LinPermission {
			ps = append(ps, *p)
		}
	}
	for _, p := range ps {
		if p.Module == userPermission.Module && p.Name == userPermission.Name {
			has = true
		}
	}
	return
}
