package data


import (
	"context"

	"lin-cms-go/api"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type lessonRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewLessonRepo(data *Data, logger log.Logger) biz.LessonRepo {
	return &lessonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *lessonRepo) GetLesson(ctx context.Context, id int64) (*biz.Lesson, error) {
	var lesson model.Lesson
	err := r.data.db.Where("id = ?", id).First(&lesson).Error
	if err != nil {
		return nil, err
	}
	return toLesson(&lesson), nil
}

func (r *lessonRepo) ListLesson(ctx context.Context, page , size int) ([]*biz.Lesson, int64, error) {
	var count int64
	err := r.data.db.Find(&model.Lesson{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	var lessons []*model.Lesson
	err = r.data.db.Offset((page - 1) * size).Limit(size).Find(&lessons).Error
	if err != nil {
		return nil, 0, err
	}
	return toLessons(lessons), count, nil
}

func (r *lessonRepo) CreateLesson(ctx context.Context, req *api.CreateLessonRequest) error {
	err := r.data.db.Create(&model.Lesson{
		Title:      req.GetTitle(),
	
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *lessonRepo) UpdateLesson(ctx context.Context,req *api.UpdateLessonRequest) error {
	err := r.data.db.Model(&model.Lesson{}).Where("id = ?", req.GetId()).Updates(map[string]interface{}{
		"title":      req.GetTitle(),
	
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *lessonRepo) DeleteLesson(ctx context.Context, id int64) error {
	err := r.data.db.Delete(&model.Lesson{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func toLesson(model *model.Lesson) *biz.Lesson {
	return &biz.Lesson{
		ID:         model.ID,
		Title:      model.Title,
		
		CreateTime: model.CreatedAt,
		UpdateTime: model.UpdatedAt,
	}
}
func toLessons(models []*model.Lesson) []*biz.Lesson {
	var lessons = make([]*biz.Lesson, len(models))
	for i := range models {
		lessons[i] = toLesson(models[i])
	}
	return lessons
}




