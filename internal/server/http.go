package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"lin-cms-go/internal/conf"
)

func NewHTTPServer(c *conf.Server) *fiber.App {
	app := fiber.New()
	app.Use(logger.New(), recover.New())
	return app
}

//func httpListen() {
//	if err := HttpSrvHandler.ListenAndServe(); err != nil {
//		log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
//	}
//}
//func httpsListen() {
//	if err := HttpSrvHandler.ListenAndServeTLS("storage/cert.pem", "storage/key.pem"); err != nil {
//		log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
//	}
//}
//func HttpServerStop() {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
//		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
//	}
//	log.Printf(" [INFO] HttpServerStop stopped\n")
//}
