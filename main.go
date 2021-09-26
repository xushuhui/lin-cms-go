package main

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lin-cms-go/api"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"log"
)

func initApp(c *conf.Config) *fiber.App {
	app := fiber.New()
	dbclient := data.NewDB(&c.Data)

	dataData, _, err := data.NewData(&c.Data, dbclient)
	if err != nil {
		panic(err)
	}
	repo := data.NewLinUserRepo(dataData)
	linUserUsecase := biz.NewLinUserUsecase(repo)
	userRoute := api.NewUser(linUserUsecase)
	server.InitRoute(app, userRoute)
	return app
}
func main() {
	//var c conf.Config
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
