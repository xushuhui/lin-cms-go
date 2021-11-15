package biz

import (
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/pkg/enum"

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

func UserHasPermission(userId int) {
}
