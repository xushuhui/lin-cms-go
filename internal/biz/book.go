package biz

import (
	"context"
	"github.com/pkg/errors"

	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/errcode"

	"github.com/xushuhui/goal/core"
)

func GetBookAll(ctx context.Context, page int, size int) (res interface{}, err error) {
	books, err := data.NewPaging(page, size).ListBook(ctx)
	if err != nil {
		return
	}
	res = books
	return
}

func GetBookTotal(ctx context.Context) (total int, err error) {
	total, err = data.GetBookCount(ctx)
	return
}

func UpdateBook(ctx context.Context, id int, req request.UpdateBook) (err error) {
	_, err = data.GetBookById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NotFoundError(errcode.BookNotFound)
		return
	}
	if err != nil {
		return
	}
	err = data.UpdateBook(ctx, id, req.Title, req.Author, req.Summary, req.Image)
	return
}

func CreateBook(ctx context.Context, req request.CreateBook) (err error) {
	book, err := data.GetBookByTitle(ctx, req.Title)
	if model.MaskNotFound(err) != nil {
		return errors.Wrap(err, "GetBookByTitleError")
	}
	if book != nil {
		err = core.ParamsError(errcode.BookTitleRepetition)
		return
	}
	err = data.CreateBook(ctx, req.Title, req.Author, req.Summary, req.Image)
	if err != nil {
		err = errors.Wrap(err, "createBookError")
	}

	return
}

func DeleteBook(ctx context.Context, id int) (err error) {
	_, err = data.GetBookById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NotFoundError(errcode.BookNotFound)
		return
	}
	if err != nil {

		return errors.Wrap(err, "GetBookByIdError")
	}
	err = data.SoftDeleteBook(ctx, id)
	return
}

func GetBook(ctx context.Context, id int) (res interface{}, err error) {
	book, err := data.GetBookById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NotFoundError(errcode.BookNotFound)
		return
	}
	if err != nil {
		err = errors.Wrap(err, "GetBookError")
		return
	}
	res = book
	return
}
