package core

import (
	"github.com/gofiber/fiber/v2"
	"lin-cms-go/global"
	"lin-cms-go/pkg/utils"
)

type Pager struct {
	List interface{} `json:"list"`
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

func GetPage(c *fiber.Ctx) error int {
	page, _ := utils.StringToInt(c.Query("page"))
	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *fiber.Ctx) error int {
	pageSize, _ := utils.StringToInt(c.Query("page_size"))
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
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
