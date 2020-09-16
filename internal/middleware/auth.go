package middleware

import (
	"github.com/gin-gonic/gin"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/lib"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			core.InvalidParamsResp(c, "empty Authorization")
			return
		}

		claims, err := lib.ParseToken(Authorization)
		if err != nil {
			core.FailResp(c, errcode.ErrorAuthToken)
			return

		}
		if time.Now().Unix() > claims.ExpiresAt {
			core.FailResp(c, errcode.TimeoutAuthToken)
			return
		}
		c.Set("uid", claims.Uid)
		c.Next()
	}

}
