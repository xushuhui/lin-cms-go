package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/utils"
)

func GetLogs(ctx context.Context, req request.GetLogs) (res interface{}, total int, err error) {
	var logs []*ent.LinLog
	paging := data.NewPaging(req.Page, req.Count)

	if req.Name == "" && req.Start == "" && req.End == "" {
		logs, err = paging.GetLogAll(ctx)
		total = data.GetLosTotal(ctx)
	}
	if req.Name != "" && req.Start == "" && req.End == "" {
		logs, err = paging.FindLogsByUsername(ctx, req.Name)
		total = data.GetLogsByUsernameTotal(ctx, req.Name)
	}
	if req.Name != "" && req.Start != "" && req.End != "" {
		start := utils.String2time(req.Start)
		end := utils.String2time(req.End)
		logs, err = paging.FindLogsByUsernameAndRange(ctx, req.Name, start, end)
		total = data.GetLogsByUsernameAndRangeTotal(ctx, req.Name, start, end)
	}
	if err != nil {
		return
	}
	res = logs
	return
}

func SearchLogs(req request.SearchLogs) (res map[string]interface{}, err error) {
	return
}
func GetLogUsers(ctx context.Context, page, size int) (res interface{}, total int, err error) {
	paging := data.NewPaging(page, size)
	res, err = paging.GetLogUsers(ctx)
	if err != nil {
		return
	}
	total, _ = data.GetLogUsersTotal(ctx)
	return
}
