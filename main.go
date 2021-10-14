package main

import (
	"io/ioutil"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"lin-cms-go/pkg/core"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gopkg.in/yaml.v2"
)

func initApp(c *conf.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			if e, ok := err.(core.IError); ok {
				if e.Err == nil {
					return e.HttpError(c)
				}

			}
			if e, ok := err.(error); ok {
				return core.InvalidParamsError(c, e.Error())

			}
			return core.ServerError(c, err)

		}})
	app.Use(recover.New(), logger.New())
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
	log.Fatal(app.Listen(c.Server.Http.Addr))
}
