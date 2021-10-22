package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/utils"
	"time"
)

func GetLogs(ctx context.Context, req request.GetLogs) (res interface{}, total int, err error) {
	var logs []*model.LinLog
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

func SearchLogs(ctx context.Context, req request.SearchLogs) (res interface{}, total int, err error) {
	var logs []*model.LinLog
	var start time.Time
	var end time.Time
	paging := data.NewPaging(req.Page, req.Count)
	if req.Start != "" && req.End != "" {
		start = utils.String2time(req.Start)
		end = utils.String2time(req.End)
	}
	if req.Name != "" && req.Start != "" && req.End != "" && req.Keyword != "" {
		logs, err = paging.FindLogsByUsernameAndRangeAndKeyword(ctx, req.Name, req.Keyword, start, end)
		total = data.GetLogsByUsernameAndRangeAndKeywordTotal(ctx, req.Name, req.Keyword, start, end)
	}
	if req.Name == "" && req.Start == "" && req.End == "" && req.Keyword == "" {
		logs, err = paging.GetLogAll(ctx)
		total = data.GetLosTotal(ctx)
	}
	if req.Keyword != "" && req.Start != "" && req.End != "" && req.Name == "" {
		logs, err = paging.FindLogsByRangeAndKeyword(ctx, req.Keyword, start, end)
		total = data.GetLogsByRangeAndKeywordTotal(ctx, req.Keyword, start, end)
	}
	if req.Keyword != "" && req.Name != "" && req.Start == "" && req.End == "" {
		logs, err = paging.FindLogsByUsernameAndKeyword(ctx, req.Name, req.Keyword)
		total = data.GetLogsByUsernameAndKeywordTotal(ctx, req.Name, req.Keyword)
	}
	if req.Keyword != "" && req.Name == "" && req.Start == "" && req.End == "" {
		logs, err = paging.FindLogsByKeyword(ctx, req.Keyword)
		total = data.GetLogsByKeywordTotal(ctx, req.Keyword)
	}
	if req.Keyword == "" && req.Name == "" && req.Start != "" && req.End != "" {
		logs, err = paging.FindLogsByRange(ctx, start, end)
		total = data.GetLogsByRangeTotal(ctx, start, end)
	}
	if req.Keyword == "" && req.Name != "" && req.Start == "" && req.End == "" {
		logs, err = paging.FindLogsByUsername(ctx, req.Name)
		total = data.GetLogsByUsernameTotal(ctx, req.Name)
	}
	if req.Keyword == "" && req.Name != "" && req.Start != "" && req.End != "" {
		logs, err = paging.FindLogsByUsernameAndRange(ctx, req.Name, start, end)
		total = data.GetLogsByUsernameAndRangeTotal(ctx, req.Name, start, end)
	}
	res = logs
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
