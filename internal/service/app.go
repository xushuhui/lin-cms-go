package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	pb "lin-cms-go/api"
	"lin-cms-go/internal/biz"
)

type AppService struct {
	pb.UnimplementedAppServer
	bu *biz.BookUsecase
	lu *biz.LessonUsecase
	tu *biz.TeacherUsecase
}

func NewAppService(bu *biz.BookUsecase, lu *biz.LessonUsecase, tu *biz.TeacherUsecase) *AppService {
	return &AppService{
		bu: bu,
		lu: lu,
		tu: tu,
	}
}

func (s *AppService) CreateLesson(ctx context.Context, req *pb.CreateLessonRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *AppService) ListLesson(ctx context.Context, req *pb.PageRequest) (*pb.ListLessonReply, error) {
	return &pb.ListLessonReply{
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}

func (s *AppService) GetLesson(ctx context.Context, req *pb.IDRequest) (*pb.GetLessonReply, error) {
	return &pb.GetLessonReply{}, nil
}

func (s *AppService) UpdateLesson(ctx context.Context, req *pb.UpdateLessonRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *AppService) DeleteLesson(ctx context.Context, req *pb.IDRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *AppService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *AppService) ListTeacher(ctx context.Context, req *pb.PageRequest) (*pb.ListTeacherReply, error) {
	page, size := defaultPageRequest(req.Page, req.Size)
	list, total, err := s.tu.ListTeacher(ctx, page, size)
	if err != nil {
		return nil, fmt.Errorf("ListTeacher failed %w ", err)
	}
	return &pb.ListTeacherReply{
		List:  list,
		Total: uint32(total),
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}

func (s *AppService) GetTeacher(ctx context.Context, req *pb.IDRequest) (*pb.GetTeacherReply, error) {
	teacher, err := s.tu.GetTeacher(ctx, req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound(err.Error(), "teacher not found")
	}
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (s *AppService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *AppService) DeleteTeacher(ctx context.Context, req *pb.IDRequest) (*emptypb.Empty, error) {
	err := s.tu.DeleteTeacher(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
