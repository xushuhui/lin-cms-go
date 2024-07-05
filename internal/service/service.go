package service

import (
	"context"

	"github.com/google/wire"
	"github.com/spf13/cast"
)

var ProviderSet = wire.NewSet(NewCmsService)

func CurrentUserId(ctx context.Context) string {
	// if ginCtx, ok := ctx.(*gin.Context); ok {
	// 	return ginCtx.GetString("userId")
	// }

	return cast.ToString(ctx.Value("userID"))
}

func CurrentIP(ctx context.Context) string {
	// if ginCtx, ok := ctx.(*gin.Context); ok {
	// 	return ginCtx.GetString("ip")
	// }
	return cast.ToString(ctx.Value("ip"))
}

func getTotalPages(total int64, size int32) int32 {
	if total == 0 || size == 0 {
		return 0
	}
	if total%int64(size) == 0 {
		return int32(total / int64(size))
	}
	return int32(total/int64(size)) + 1
}

func defaultPageRequest(page, size int32) (int, int) {

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return int(page), int(size)
}
