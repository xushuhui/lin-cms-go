package data

import (
	"context"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/book"
	"time"
)

func GetBookById(ctx context.Context, id int) (model *model.Book, err error) {
	model, err = GetDB().Book.Query().Where(book.ID(id)).First(ctx)
	return
}

func GetBookByTitle(ctx context.Context, title string) (model *model.Book, err error) {
	model, err = GetDB().Book.Query().Where(book.Title(title)).First(ctx)
	return
}

func UpdateBook(ctx context.Context, id int, title string, author string, summary string, image string) (err error) {
	_, err = GetDB().Book.Update().Where(book.ID(id)).SetTitle(title).SetAuthor(author).SetSummary(summary).SetImage(image).Save(ctx)
	return
}

func CreateBook(ctx context.Context, title string, author string, summary string, image string) (err error) {
	_, err = GetDB().Book.Create().SetTitle(title).SetAuthor(author).SetSummary(summary).SetImage(image).Save(ctx)
	return
}

func SoftDeleteBook(ctx context.Context, id int) (err error) {
	_, err = GetDB().Book.Update().Where(book.ID(id)).SetDeleteTime(time.Now()).Save(ctx)
	return
}

func GetBookCount(ctx context.Context) (count int, err error) {
	count, err = GetDB().Book.Query().Count(ctx)
	return
}

func (p *Paging) ListBook(ctx context.Context) (model []*model.Book, err error) {
	model, err = GetDB().Book.Query().Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}
