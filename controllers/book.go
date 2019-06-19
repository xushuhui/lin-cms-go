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

	r := request.UpdateBook{}
	r.Bind(b.Ctx.Input.RequestBody)
	if isCorrect, errmsg := r.Verify(); !isCorrect {
		b.Data["json"] = core.ParmsError(errmsg)
		b.ServeJSON()
	}

	book, _ := models.GetBookById(r.Id)
	if book == nil {
		b.Data["json"] = core.Fail(core.CodeNoBook)
		b.ServeJSON()
	}
	bookModel := models.Book{
		Id:      r.Id,
		Title:   r.Title,
		Author:  r.Author,
		Image:   r.Image,
		Summary: r.Summary,
	}
	err := models.UpdateBookById(&bookModel)
	if err != nil {
		panic(err)
	}
	b.Data["json"] = core.Succeed()
	b.ServeJSON()
}
func (b *BookController) DeleteBook() {
	id, err := b.GetInt(":id")
	if err != nil {
		panic(err)
	}
	err = models.DeleteBook(id)
	if err != nil {
		panic(err)
	}
	b.Data["json"] = core.Succeed()
	b.ServeJSON()
}
