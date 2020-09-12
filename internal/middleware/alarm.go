package middleware

import (
	"github.com/gin-gonic/gin"
	"lin-cms-go/pkg/core"
	"runtime/debug"
	"strings"
)

func Alarm() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				//subject := fmt.Sprintf("【重要错误】%s 项目出错了！", config.AppName)
				//
				//body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				//body  = strings.ReplaceAll(body, "{RequestTime}", utils.GetCurrentDate())
				//body  = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method + "  " + c.Request.Host + c.Request.RequestURI)
				//body  = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				//body  = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				//body  = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				//_ = util.SendMail("", subject, body)

				core.ServerError(c)
			}
		}()
		c.Next()
	}
}
