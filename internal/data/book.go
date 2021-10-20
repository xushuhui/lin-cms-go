package data

import (
	"context"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/book"
	"lin-cms-go/pkg/core"
)

func GetBookAll(ctx context.Context, offset int, size int) (model []*ent.Book, err error) {
	model, err = GetDB().Book.Query().Limit(size).Offset(offset).All(ctx)
	return
}

func GetBookById(ctx context.Context, id int) (model *ent.Book, err error) {
	model, err = GetDB().Book.Query().Where(book.ID(id)).First(ctx)
	return
}

func GetBookByTitle(ctx context.Context, title string) (model *ent.Book, err error) {
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

func DeleteBook(ctx context.Context, id int) (err error) {
	err = GetDB().Book.DeleteOneID(id).Exec(ctx)
	return
}

func GetBookCount(ctx context.Context) (count int, err error) {
	count, err = GetDB().Book.Query().Count(ctx)

	return
}

type Paging struct {
	Page   int
	Size   int
	Offset int
}

func NewPaging(page, pageSize int) *Paging {
	return &Paging{
		Page:   page,
		Size:   pageSize,
		Offset: core.GetPageOffset(page, pageSize),
	}
}
func (p *Paging) ListBook(ctx context.Context) (model []*ent.Book, err error) {
	model, err = GetDB().Book.Query().Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}
