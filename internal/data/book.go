package data

import (
	"context"

	"lin-cms-go/api"
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
	var book model.Book
	err := r.data.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return toBook(&book), nil
}

func (r *bookRepo) ListBook(ctx context.Context, page, size int) ([]*biz.Book, int64, error) {
	var count int64
	err := r.data.db.Find(&model.Book{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	var books []*model.Book
	err = r.data.db.Offset((page - 1) * size).Limit(size).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return toBooks(books), count, nil
}

func (r *bookRepo) CreateBook(ctx context.Context, req *api.CreateBookRequest) error {
	err := r.data.db.Create(&model.Book{
		Title:   req.GetTitle(),
		Author:  req.GetAuthor(),
		Summary: req.GetSummary(),
		Image:   req.GetImage(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepo) UpdateBook(ctx context.Context, req *api.UpdateBookRequest) error {
	err := r.data.db.Model(&model.Book{}).Where("id = ?", req.GetId()).Updates(map[string]interface{}{
		"title":   req.GetTitle(),
		"author":  req.GetAuthor(),
		"summary": req.GetSummary(),
		"image":   req.GetImage(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepo) DeleteBook(ctx context.Context, id int64) error {
	err := r.data.db.Delete(&model.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func toBook(model *model.Book) *biz.Book {
	return &biz.Book{
		ID:         model.ID,
		Title:      model.Title,
		Author:     model.Author,
		Summary:    model.Summary,
		Image:      model.Image,
		CreatedAt:  model.CreatedAt,
	}
}

func toBooks(models []*model.Book) []*biz.Book {
	books := make([]*biz.Book, len(models))
	for i := range models {
		books[i] = toBook(models[i])
	}
	return books
}

// func GetBookByTitle(ctx context.Context, title string) (model *model.Book, err error) {
// 	model, err = GetDB().Book.Query().Where(book.Title(title)).First(ctx)
// 	return
// }
