package controllers

import (
	"github.com/astaxie/beego"
	"lin-cms-beego/core"
	"lin-cms-beego/models"
	"lin-cms-beego/request"
)

type BookController struct {
	beego.Controller
}

func (b *BookController) CreateBook() {
	bookRequest := request.CreateBook{}
	bookRequest.Bind(b.Ctx.Input.RequestBody)
	if isCorrect, errmsg := bookRequest.Verify(); !isCorrect {
		b.Data["json"] = core.ParmsError(errmsg)
		b.ServeJSON()
	}
	b.Data["json"] = bookRequest
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

	if book, _ := models.GetBookById(r.Id); book == nil {
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

	if err := models.UpdateBookById(&bookModel); err != nil {
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
	if book, _ := models.GetBookById(id); book == nil {
		b.Data["json"] = core.Fail(core.CodeNoBook)
		b.ServeJSON()
	}
	if err := models.SoftDeleteBook(id); err != nil {
		panic(err)
	}
	b.Data["json"] = core.Succeed()
	b.ServeJSON()
}
