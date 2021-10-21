package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
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
		err = core.NewErrorCode(errcode.BookNotFound)
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
		return err
	}
	if book != nil {
		err = core.NewErrorCode(errcode.BookTitleRepetition)
		return
	}
	err = data.CreateBook(ctx, req.Title, req.Author, req.Summary, req.Image)
	return
}

func DeleteBook(ctx context.Context, id int) (err error) {
	_, err = data.GetBookById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.BookNotFound)
		return
	}
	err = data.DeleteBook(ctx, id)
	return
}

func GetBook(ctx context.Context, id int) (res interface{}, err error) {
	book, err := data.GetBookById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.BookNotFound)
		return
	}
	res = book
	return
}
