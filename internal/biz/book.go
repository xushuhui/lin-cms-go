package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
)

func GetBookAll(ctx context.Context) (res map[string]interface{}, err error) {
	bookModel, err := data.GetBookAll(ctx)
	if err != nil {
		return
	}
	res = make(map[string]interface{})
	res["books"] = bookModel
	return
}

func UpdateBook(ctx context.Context, req request.UpdateBook) (err error) {
	_, err = data.GetBookById(ctx, req.Id)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.BookNotFound)
		return
	}
	if err != nil {
		return
	}
	err = data.UpdateBook(ctx, req.Id, req.Title, req.Author, req.Summary, req.Image)
	return
}

func CreateBook(ctx context.Context, req request.CreateBook) (err error) {
	err = data.CreateBook(ctx, req.Title, req.Author, req.Summary, req.Image)
	return
}

func DeleteBook(ctx context.Context, req request.DeleteBook) (err error) {
	_, err = data.GetBookById(ctx, req.Id)
	if ent.IsNotFound(err) {
		err = core.NewErrorCode(errcode.BookNotFound)
		return
	}
	err = data.DeleteBook(ctx, req.Id)
	return
}
