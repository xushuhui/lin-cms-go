package service

import (
	"context"

	pb "lin-cms-go/api"
	"lin-cms-go/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CmsService struct {
	pb.UnimplementedCmsServer
	bu *biz.BookUsecase
}

func NewCmsService(bu *biz.BookUsecase) *CmsService {
	return &CmsService{
		bu: bu,
	}
}

func (s *CmsService) Ping(ctx context.Context, req *emptypb.Empty) (*pb.PingReply, error) {
	return &pb.PingReply{}, nil
}

func (s *CmsService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}

func (s *CmsService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) ListBook(ctx context.Context, req *pb.PageRequest) (*pb.ListBookReply, error) {
	return &pb.ListBookReply{}, nil
}

func (s *CmsService) GetBook(ctx context.Context, req *pb.IDRequest) (*pb.GetBookReply, error) {
	return &pb.GetBookReply{}, nil
}

func (s *CmsService) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) DeleteBook(ctx context.Context, req *pb.IDRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadReply, error) {
	return &pb.UploadReply{}, nil
}

func (s *CmsService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) UpdateMe(ctx context.Context, req *pb.UpdateMeRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) UpdateMyPassword(ctx context.Context, req *pb.UpdateMyPasswordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) ListMyPermission(ctx context.Context, req *emptypb.Empty) (*pb.ListMyPermissionReply, error) {
	return &pb.ListMyPermissionReply{}, nil
}

func (s *CmsService) GetMyInfomation(ctx context.Context, req *emptypb.Empty) (*pb.GetMyInfomationReply, error) {
	return &pb.GetMyInfomationReply{}, nil
}

func (s *CmsService) ListPermission(ctx context.Context, req *emptypb.Empty) (*pb.ListPermissionReply, error) {
	return &pb.ListPermissionReply{}, nil
}

func (s *CmsService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *CmsService) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) DeleteUser(ctx context.Context, req *pb.IDRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) GetUser(ctx context.Context, req *pb.IDRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}

func (s *CmsService) GetGroup(ctx context.Context, req *pb.IDRequest) (*pb.GetGroupReply, error) {
	return &pb.GetGroupReply{}, nil
}

func (s *CmsService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) DeleteGroup(ctx context.Context, req *pb.IDRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) ListGroup(ctx context.Context, req *emptypb.Empty) (*pb.ListGroupReply, error) {
	return &pb.ListGroupReply{}, nil
}

func (s *CmsService) DispatchPermission(ctx context.Context, req *pb.DispatchPermissionRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) DispatchPermissions(ctx context.Context, req *pb.DispatchPermissionsRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) RemovePermission(ctx context.Context, req *pb.RemovePermissionRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *CmsService) ListLog(ctx context.Context, req *pb.PageRequest) (*pb.ListLogReply, error) {
	return &pb.ListLogReply{}, nil
}

func (s *CmsService) SearchLog(ctx context.Context, req *pb.SearchLogRequest) (*pb.ListLogReply, error) {
	return &pb.ListLogReply{}, nil
}

func (s *CmsService) ListLogUser(ctx context.Context, req *pb.PageRequest) (*pb.ListLogUserReply, error) {
	return &pb.ListLogUserReply{}, nil
}
