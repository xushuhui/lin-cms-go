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
		ListBook(ctx context.Context, page, size int) ([]*Book, int64, error)
		CreateBook(ctx context.Context, req *api.CreateBookRequest) error
		UpdateBook(ctx context.Context, req *api.UpdateBookRequest) error
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
		Id:      uint32(b.ID),
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

func (u *BookUsecase) ListBook(ctx context.Context, page, size int) ([]*api.Book, int64, error) {
	books, total, err := u.br.ListBook(ctx, int(page), int(size))
	if err != nil {
		return nil, 0, err
	}

	return outBooks(books), total, nil
}

func (u *BookUsecase) UpdateBook(ctx context.Context, req *api.UpdateBookRequest) error {
	_, err := u.br.GetBook(ctx, req.Id)
	if err != nil {
		return errors.Wrap(err, "GetBookError")
	}

	err = u.br.UpdateBook(ctx, req)
	return err
}

func (u *BookUsecase) CreateBook(ctx context.Context, req *api.CreateBookRequest) error {
	
	err := u.br.CreateBook(ctx, req)
	return err
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
