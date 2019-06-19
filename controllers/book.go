package controllers

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/core"
	"lin-cms-beego/models"
	"lin-cms-beego/request"
	"lin-cms-beego/utils"
)

type BookController struct {
	beego.Controller
}

func (b *BookController) CreateBook() {
	var bookRequest *request.CreateBook
	utils.BindJson(b.Ctx.Input.RequestBody, &bookRequest)

	b.Data["json"] = ""
	b.ServeJSON()
}
func (b *BookController) GetBook() {

	id, err := b.GetInt(":id")
	if err != nil {
		panic(err)
	}
	book, _ := models.GetBookById(id)
	b.Data["json"] = core.SetData(book)
	b.ServeJSON()
}
func (b *BookController) GetBooks() {
	book, _ := models.GetBooks()
	b.Data["json"] = core.SetData(book)
	b.ServeJSON()
}
func (b *BookController) UpdateBook() {
	var bookRequest *request.GetBook
	utils.BindJson(b.Ctx.Input.RequestBody, &bookRequest)
	b.Data["json"] = ""
	b.ServeJSON()
}
func (b *BookController) DeleteBook() {
	id, err := b.GetInt(":id")
	if err != nil {
		panic(err)
	}
	b.Data["json"] = id
	b.ServeJSON()
}
