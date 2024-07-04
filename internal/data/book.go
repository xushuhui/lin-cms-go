package data

import (
	"context"

	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type bookRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBookRepo(data *Data, logger log.Logger) biz.BookRepo {
	return &bookRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *bookRepo) GetBook(ctx context.Context, id int64) (*biz.Book, error) {
	panic("not implemented") // TODO: Implement
}

func (r *bookRepo) ListBook(ctx context.Context, page int32, size int32) ([]*biz.Book, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *bookRepo) CreateBook(ctx context.Context, book *biz.Book) error {
	panic("not implemented") // TODO: Implement
}

func (r *bookRepo) UpdateBook(ctx context.Context, book *biz.Book) error {
	panic("not implemented") // TODO: Implement
}

func (r *bookRepo) DeleteBook(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}

func ToBook(model *model.Book) *biz.Book {
	return &biz.Book{
		ID:         model.ID,
		Title:      model.Title,
		Author:     model.Author,
		Summary:    model.Summary,
		Image:      model.Image,
		CreateTime: model.CreatedAt,
		UpdateTime: model.UpdatedAt,
	}
}

// func GetBookById(ctx context.Context, id int) (*biz.Book, error) {
// 	model, err = GetDB().Book.Query().Where(book.ID(id)).First(ctx)
// 	return
// }

// func GetBookByTitle(ctx context.Context, title string) (model *model.Book, err error) {
// 	model, err = GetDB().Book.Query().Where(book.Title(title)).First(ctx)
// 	return
// }

// func UpdateBook(ctx context.Context, id int, title string, author string, summary string, image string) (err error) {
// 	_, err = GetDB().Book.Update().Where(book.ID(id)).SetTitle(title).SetAuthor(author).SetSummary(summary).SetImage(image).Save(ctx)
// 	return
// }

// func CreateBook(ctx context.Context, title string, author string, summary string, image string) (err error) {
// 	_, err = GetDB().Book.Create().SetTitle(title).SetAuthor(author).SetSummary(summary).SetImage(image).Save(ctx)
// 	return
// }

// func SoftDeleteBook(ctx context.Context, id int) (err error) {
// 	_, err = GetDB().Book.Update().Where(book.ID(id)).SetDeleteTime(time.Now()).Save(ctx)
// 	return
// }

// func GetBookCount(ctx context.Context) (count int, err error) {
// 	count, err = GetDB().Book.Query().Count(ctx)
// 	return
// }

// func (p *Paging) ListBook(ctx context.Context) (model []*model.Book, err error) {
// 	model, err = GetDB().Book.Query().Limit(p.Size).Offset(p.Offset).All(ctx)
// 	return
// }
