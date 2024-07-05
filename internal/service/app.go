package service

import (
	"context"

	pb "lin-cms-go/api"
)

type AppService struct {
	pb.UnimplementedAppServer
}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) CreateLesson(ctx context.Context, req *pb.CreateLessonRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *AppService) ListLesson(ctx context.Context, req *pb.PageRequest) (*pb.ListLessonReply, error) {
	return &pb.ListLessonReply{}, nil
}
func (s *AppService) GetLesson(ctx context.Context, req *pb.IDRequest) (*pb.GetLessonReply, error) {
	return &pb.GetLessonReply{}, nil
}
func (s *AppService) UpdateLesson(ctx context.Context, req *pb.UpdateLessonRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *AppService) DeleteLesson(ctx context.Context, req *pb.IDRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *AppService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *AppService) ListTeacher(ctx context.Context, req *pb.PageRequest) (*pb.ListTeacherReply, error) {
	return &pb.ListTeacherReply{}, nil
}
func (s *AppService) GetTeacher(ctx context.Context, req *pb.IDRequest) (*pb.GetTeacherReply, error) {
	return &pb.GetTeacherReply{}, nil
}
func (s *AppService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
func (s *AppService) DeleteTeacher(ctx context.Context, req *pb.IDRequest) (*pb.google_protobuf_Empty, error) {
	return &pb.google_protobuf_Empty{}, nil
}
