package main

import (
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/log"
)

func errorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	if e, ok := err.(core.IError); ok {
		if e.Err == nil {
			return e.HttpError(c)
		}

	}

	return core.ServerError(c, err)
}
func initApp(c *conf.Config) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: log.Defaultlogger,
		Fields: []string{"method", "url", "latency", "status", "reqbody", "body", "error"},
	}))
	app.Use(recover.New(), cors.New())
	data.NewDataSource(&c.Data)
	//TODO clean source

	core.InitValidate()
	server.InitRoute(app)

	return app
}
func main() {

	c := new(conf.Config)
	yamlFile, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}

	app := initApp(c)
	app.Listen(c.Server.Http.Addr)
}
