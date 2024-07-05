package biz

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"lin-cms-go/api"

	"github.com/go-kratos/kratos/v2/log"
)

type (
	TeacherRepo interface {
		GetTeacher(ctx context.Context, id int64) (*Teacher, error)
		ListTeacher(ctx context.Context, page, size int) ([]*Teacher, int64, error)
		CreateTeacher(ctx context.Context, req *api.CreateTeacherRequest) error
		UpdateTeacher(ctx context.Context, req *api.UpdateTeacherRequest) error
		DeleteTeacher(ctx context.Context, id int64) error
	}
	Teacher struct {
		ID         int64
		CreateTime time.Time
		UpdateTime time.Time
		Name      string
		
	}
)

type TeacherUsecase struct {
	log *log.Helper
	br  TeacherRepo
}

func NewTeacherUsecase(br TeacherRepo, logger log.Logger) *TeacherUsecase {
	return &TeacherUsecase{
		log: log.NewHelper(logger),
		br:  br,
	}
}

func outTeacher(b *Teacher) *api.Teacher {
	return &api.Teacher{
		Id:      uint32(b.ID),
		Name:    b.Name,
	
	}
}

func outTeachers(bs []*Teacher) []*api.Teacher {
	res := make([]*api.Teacher, len(bs))
	for i := range bs {
		res[i] = outTeacher(bs[i])
	}
	return res
}

func (u *TeacherUsecase) ListTeacher(ctx context.Context, page, size int) ([]*api.Teacher, int64, error) {
	teachers, total, err := u.br.ListTeacher(ctx, int(page), int(size))
	if err != nil {
		return nil, 0, err
	}

	return outTeachers(teachers), total, nil
}

func (u *TeacherUsecase) UpdateTeacher(ctx context.Context, req *api.UpdateTeacherRequest) error {
	_, err := u.br.GetTeacher(ctx, req.Id)
	if err != nil {
		return errors.Wrap(err, "GetTeacherError")
	}

	err = u.br.UpdateTeacher(ctx, req)
	return err
}

func (u *TeacherUsecase) CreateTeacher(ctx context.Context, req *api.CreateTeacherRequest) error {
	
	err := u.br.CreateTeacher(ctx, req)
	return err
}

func (u *TeacherUsecase) DeleteTeacher(ctx context.Context, id int64) error {
	_, err := u.br.GetTeacher(ctx, id)
	if err != nil {
		return errors.Wrap(err, "GetTeacherError")
	}
	err = u.DeleteTeacher(ctx, id)
	if err != nil {
		return errors.Wrap(err, "DeleteTeacherError")
	}
	return nil
}

func (u *TeacherUsecase) GetTeacher(ctx context.Context, id int64) (*api.Teacher, error) {
	teacher, err := u.br.GetTeacher(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "GetTeacherError")
	}

	return outTeacher(teacher), nil
}
