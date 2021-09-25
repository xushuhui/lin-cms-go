package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unrolled/secure"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *fiber.Ctx) error {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
