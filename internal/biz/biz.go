package biz

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xushuhui/goal/utils"
	"lin-cms-go/internal/data/model"
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
