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
		var err error
		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			err = core.NewInvalidParamsError("empty Authorization")
			return
		}

		claims, err := lib.ParseToken(Authorization)
		if err != nil {
			err = core.NewError(errcode.ErrorAuthToken)
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			err = core.NewError(errcode.TimeoutAuthToken)
			return
		}

		c.Next()
	}

}
