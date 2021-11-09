package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/grestful/logs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"lin-cms-go/pkg/core"
)

func errorHandler(c *fiber.Ctx, err error) error {
	// response err handle
	if e, ok := err.(core.HttpError); ok {
		if e.Err == nil {
			return e.HandleHttpError(c)
		}
	}

	return core.HandleServerError(c, err)
}
func initLog(c *conf.Config) {
	log.SetFile(log.FileConfig{

		Category: "file",
		Level:    "DEBUG",
		Filename: c.Log.Path + "/d.log",
		Rotate:   true,
		Daily:    true,
	})
	writer := log.GetLogger("file")
	log.SetDefaultLog(writer)

}
func initApp(c *conf.Config) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	app.Use(logger.New())
	app.Use(cors.New(), server.UserLog)
	initLog(c)
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
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	app := initApp(c)
	app.Listen(c.Server.Http.Addr)

}
