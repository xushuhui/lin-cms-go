package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"lin-cms-go/api"

	"github.com/go-kratos/kratos/v2/log"
)

type (
	LessonRepo interface {
		GetLesson(ctx context.Context, id int64) (*Lesson, error)
		ListLesson(ctx context.Context, page, size int) ([]*Lesson, int64, error)
		CreateLesson(ctx context.Context, req *api.CreateLessonRequest) error
		UpdateLesson(ctx context.Context, req *api.UpdateLessonRequest) error
		DeleteLesson(ctx context.Context, id int64) error
	}
	Lesson struct {
		ID         int64
		CreatedAt  time.Time
		Title      string
	}
)

type LessonUsecase struct {
	log *log.Helper
	br  LessonRepo
}

func NewLessonUsecase(br LessonRepo, logger log.Logger) *LessonUsecase {
	return &LessonUsecase{
		log: log.NewHelper(logger),
		br:  br,
	}
}

func outLesson(b *Lesson) *api.Lesson {
	return &api.Lesson{
		Id:    uint32(b.ID),
		Title: b.Title,
	}
}

func outLessons(bs []*Lesson) []*api.Lesson {
	res := make([]*api.Lesson, len(bs))
	for i := range bs {
		res[i] = outLesson(bs[i])
	}
	return res
}

func (u *LessonUsecase) ListLesson(ctx context.Context, page, size int) ([]*api.Lesson, int64, error) {
	lessons, total, err := u.br.ListLesson(ctx, int(page), int(size))
	if err != nil {
		return nil, 0, err
	}

	return outLessons(lessons), total, nil
}

func (u *LessonUsecase) UpdateLesson(ctx context.Context, req *api.UpdateLessonRequest) error {
	_, err := u.br.GetLesson(ctx, req.Id)
	if err != nil {
		return errors.Wrap(err, "GetLessonError")
	}

	err = u.br.UpdateLesson(ctx, req)
	return err
}

func (u *LessonUsecase) CreateLesson(ctx context.Context, req *api.CreateLessonRequest) error {
	err := u.br.CreateLesson(ctx, req)
	return err
}

func (u *LessonUsecase) DeleteLesson(ctx context.Context, id int64) error {
	_, err := u.br.GetLesson(ctx, id)
	if err != nil {
		return errors.Wrap(err, "GetLessonError")
	}
	err = u.DeleteLesson(ctx, id)
	if err != nil {
		return errors.Wrap(err, "DeleteLessonError")
	}
	return nil
}

func (u *LessonUsecase) GetLesson(ctx context.Context, id int64) (*api.Lesson, error) {
	lesson, err := u.br.GetLesson(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "GetLessonError")
	}

	return outLesson(lesson), nil
}
