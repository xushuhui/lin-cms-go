package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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

	var bookRequest *request.UpdateBook

	utils.BindJson(b.Ctx.Input.RequestBody, &bookRequest)
	valid := validation.Validation{}

	verify, err := valid.Valid(bookRequest)
	if err != nil {
		panic(err)
	}
	if !verify {
		for _, err := range valid.Errors {
			b.Data["json"] = core.FailMsg(core.CodeInvaldParams, err.Field+" "+err.Message)
			b.ServeJSON()
		}
	}
	book, _ := models.GetBookById(bookRequest.Id)
	if book == nil {
		b.Data["json"] = core.Fail(core.CodeNoBook)
		b.ServeJSON()
	}
	bookModel := models.Book{
		Id:      bookRequest.Id,
		Title:   bookRequest.Title,
		Author:  bookRequest.Author,
		Image:   bookRequest.Image,
		Summary: bookRequest.Summary,
	}
	err = models.UpdateBookById(&bookModel)
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
