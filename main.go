package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/log"
	"os"
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

	l := log.NewLogger(os.Stderr)
	w, _ := os.Create("./app.log")
	l.WithWriter(w).Debug("debug", "ssss")
	l.Error("err", "e")
	app.Use(recover.New(),
		logger.New(logger.Config{
			TimeFormat: "2006-01-02 15:04:05.000", Format: "[${time}] ${method} ${path} - ${status} ${latency} \n ${body} \n ${resBody} \n ${error}", Output: os.Stderr,
		}),
		cors.New())

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
