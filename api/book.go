package api

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
)

func GetBooks(c *fiber.Ctx) error {
	data, err := biz.GetBookAll(c.Context())
	if err != nil {
		return err
	}
	core.SetData(c, data)
	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	var req request.UpdateBook
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	err := biz.UpdateBook(c.Context(), req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func CreateBook(c *fiber.Ctx) error {
	var req request.CreateBook
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	err := biz.CreateBook(c.Context(), req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}

func DeleteBook(c *fiber.Ctx) error {
	var req request.DeleteBook
	if err := core.ParseRequest(c, &req); err != nil {
		return err
	}
	err := biz.DeleteBook(c.Context(), req)
	if err != nil {
		return err
	}
	return core.SuccessResp(c)
}
