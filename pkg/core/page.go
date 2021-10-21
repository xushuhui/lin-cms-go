package core

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/pkg/utils"
)

type Pager struct {
	Items interface{} `json:"items"`
	// 页码
	Page int `json:"page"`
	// 每页数量
	Size int `json:"size"`
	// 总行数
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

func GetPage(c *fiber.Ctx) int {
	page, _ := utils.StringToInt(c.Query("page"))
	if page <= 0 {
		return 1
	}

	return page
}

func GetSize(c *fiber.Ctx) int {
	pageSize, _ := utils.StringToInt(c.Query("count"))
	if pageSize <= 0 {
		return 10
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
