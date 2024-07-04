package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"lin-cms-go/api"

	"github.com/go-kratos/kratos/v2/log"
)

type (
	BookRepo interface {
		GetBook(ctx context.Context, id int64) (*Book, error)
		ListBook(ctx context.Context, page, size int32) ([]*Book, int64, error)
		CreateBook(ctx context.Context, book *Book) error
		UpdateBook(ctx context.Context, book *Book) error
		DeleteBook(ctx context.Context, id int64) error
	}
	Book struct {
		ID         int64
		CreateTime time.Time
		UpdateTime time.Time
		Title      string
		Author     string
		Summary    string
		Image      string
	}
)


type BookUsecase struct {
	log *log.Helper
	br  BookRepo
}

func NewBookUsecase(br BookRepo, logger log.Logger) *BookUsecase {
	return &BookUsecase{
		log: log.NewHelper(logger),
		br:  br,
	}
}

func outBook(b *Book) *api.Book {
	return &api.Book{
		Id:      b.ID,
		Title:   b.Title,
		Author:  b.Author,
		Summary: b.Summary,
		Image:   b.Image,
	}
}

func outBooks(bs []*Book) []*api.Book {
	res := make([]*api.Book, len(bs))
	for i := range bs {
		res[i] = outBook(bs[i])
	}
	return res
}

func (u *BookUsecase) ListBook(ctx context.Context, page, size int32) ([]*api.Book, int64, error) {
	books, total, err := u.br.ListBook(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	return outBooks(books), total, nil
}

func (u *BookUsecase) UpdateBook(ctx context.Context, id int, req *api.UpdateBookRequest) (err error) {
	// if err != nil {
	// 	return
	// }
	// err = data.UpdateBook(ctx, id, req.Title, req.Author, req.Summary, req.Image)
	return
}

func (u *BookUsecase) CreateBook(ctx context.Context, req *api.CreateBookRequest) (err error) {
	// err = data.CreateBook(ctx, req.Title, req.Author, req.Summary, req.Image)
	// if err != nil {
	// 	err = errors.Wrap(err, "createBookError")
	// }

	return
}

func (u *BookUsecase) DeleteBook(ctx context.Context, id int64) error {
	_, err := u.br.GetBook(ctx, id)
	if err != nil {
		return errors.Wrap(err, "GetBookError")
	}
	err = u.DeleteBook(ctx, id)
	if err != nil {
		return errors.Wrap(err, "DeleteBookError")
	}
	return nil
}

func (u *BookUsecase) GetBook(ctx context.Context, id int64) (*api.Book, error) {
	book, err := u.br.GetBook(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "GetBookError")
	}

	return outBook(book), nil
}
