package data

import (
	"context"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/book"
)

func GetBookAll(ctx context.Context) (model []*ent.Book, err error) {
	model, err = GetDB().Book.Query().All(ctx)
	return
}

func GetBookById(ctx context.Context, id int) (model *ent.Book, err error) {
	model, err = GetDB().Book.Query().Where(book.ID(id)).First(ctx)
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

func DeleteBook(ctx context.Context, id int) (err error) {
	err = GetDB().Book.DeleteOneID(id).Exec(ctx)
	return
}
