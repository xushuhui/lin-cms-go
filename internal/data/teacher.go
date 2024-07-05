package data

import (
	"context"

	"lin-cms-go/api"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type teacherRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewTeacherRepo(data *Data, logger log.Logger) biz.TeacherRepo {
	return &teacherRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *teacherRepo) GetTeacher(ctx context.Context, id int64) (*biz.Teacher, error) {
	var teacher model.Teacher
	err := r.data.db.Where("id = ?", id).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return toTeacher(&teacher), nil
}

func (r *teacherRepo) ListTeacher(ctx context.Context, page, size int) ([]*biz.Teacher, int64, error) {
	var count int64
	err := r.data.db.Find(&model.Teacher{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	var teachers []*model.Teacher
	err = r.data.db.Offset((page - 1) * size).Limit(size).Find(&teachers).Error
	if err != nil {
		return nil, 0, err
	}
	return toTeachers(teachers), count, nil
}

func (r *teacherRepo) CreateTeacher(ctx context.Context, req *api.CreateTeacherRequest) error {
	err := r.data.db.Create(&model.Teacher{
		Name: req.Name,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *teacherRepo) UpdateTeacher(ctx context.Context, req *api.UpdateTeacherRequest) error {
	err := r.data.db.Model(&model.Teacher{}).Where("id = ?", req.GetId()).Updates(map[string]interface{}{
		"name": req.Name,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *teacherRepo) DeleteTeacher(ctx context.Context, id int64) error {
	err := r.data.db.Delete(&model.Teacher{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func toTeacher(model *model.Teacher) *biz.Teacher {
	return &biz.Teacher{
		ID:   model.ID,
		Name: model.Name,

		CreateTime: model.CreatedAt,
		UpdateTime: model.UpdatedAt,
	}
}

func toTeachers(models []*model.Teacher) []*biz.Teacher {
	teachers := make([]*biz.Teacher, len(models))
	for i := range models {
		teachers[i] = toTeacher(models[i])
	}
	return teachers
}
