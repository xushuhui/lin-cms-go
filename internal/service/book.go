package service

import (
	"context"

	"lin-cms-go/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CmsService) CreateBook(ctx context.Context, req *api.CreateBookRequest) (*emptypb.Empty, error) {
	if err := s.bu.CreateBook(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *CmsService) ListBook(ctx context.Context, req *api.PageRequest) (*api.ListBookReply, error) {
	// TODO 权限判断
	list, total, err := s.bu.ListBook(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &api.ListBookReply{
		List: list,
		Total: total,
	}, nil
	
}

func (s *CmsService) GetBook(ctx context.Context, req *api.IDRequest) (*api.GetBookReply, error) {
	// TODO 权限判断
	book, err := s.bu.GetBook(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.GetBookReply{
		Book: book,
	}, nil
	
}

func (s *CmsService) UpdateBook(ctx context.Context, req *api.UpdateBookRequest) (*emptypb.Empty, error) {
	if err := s.bu.UpdateBook(ctx, req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *CmsService) DeleteBook(ctx context.Context, req *api.IDRequest) (*emptypb.Empty, error) {
	if err := s.bu.DeleteBook(ctx, req.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// func GetBooks(c *fiber.Ctx) error {
// 	// TODO book 接口少了权限判断

// 	size := core.GetSize(c)
// 	page := core.GetPage(c)
// 	data, err := biz.GetBookAll(c.Context(), page, size)
// 	if err != nil {
// 		return err
// 	}
// 	total, err := biz.GetBookTotal(c.Context())
// 	if err != nil {
// 		return err
// 	}
// 	core.SetPage(c, data, total)
// 	return nil
// }

// func UpdateBook(c *fiber.Ctx) error {
// 	var req api.UpdateBookRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}

// 	err = biz.UpdateBook(c.Context(), id, &req)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func CreateBook(c *fiber.Ctx) error {
// 	var req api.CreateBookRequest
// 	if err := core.ParseRequest(c, &req); err != nil {
// 		return err
// 	}
// 	err := biz.CreateBook(c.Context(), &req)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func DeleteBook(c *fiber.Ctx) error {
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}
// 	err = biz.DeleteBook(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}
// 	return core.SuccessResp(c)
// }

// func GetBook(c *fiber.Ctx) error {
// 	id, err := utils.StringToInt(c.Params("id"))
// 	if err != nil {
// 		return err
// 	}
// 	data, err := biz.GetBook(c.Context(), id)
// 	if err != nil {
// 		return err
// 	}
// 	core.SetData(c, data)
// 	return nil
// }
