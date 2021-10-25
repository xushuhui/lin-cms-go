package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/linlog"
	"lin-cms-go/internal/data/model/predicate"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/utils"
	"time"
)

func GetLogs(ctx context.Context, req request.GetLogs, page int, size int) (res interface{}, total int, err error) {
	var logs []*model.LinLog
	var start time.Time
	var end time.Time
	paging := data.NewPaging(page, size)
	var query []predicate.LinLog
	if req.Name != "" {
		query = append(query, data.WithUsername(req.Name))
	}
	if req.Start != "" && req.End != "" {
		start = utils.String2time(req.Start)
		end = utils.String2time(req.End)
		q := linlog.And(linlog.CreateTimeGT(start), linlog.CreateTimeLT(end))
		query = append(query, q)
	}
	logs, err = paging.Search(ctx, query)
	total = data.GetSearchTotal(ctx, query)
	if err != nil {
		return
	}
	res = logs
	return
}

func SearchLogs(ctx context.Context, req request.SearchLogs, page int, size int) (res interface{}, total int, err error) {
	var logs []*model.LinLog
	var start time.Time
	var end time.Time
	paging := data.NewPaging(page, size)
	var query []predicate.LinLog
	if req.Name != "" {
		query = append(query, data.WithUsername(req.Name))
	}
	if req.Start != "" && req.End != "" {
		start = utils.String2time(req.Start)
		end = utils.String2time(req.End)
		q := linlog.And(linlog.CreateTimeGT(start), linlog.CreateTimeLT(end))
		query = append(query, q)
	}
	if req.Keyword != "" {
		query = append(query, data.WithKeyword(req.Name))
	}
	logs, err = paging.Search(ctx, query)
	total = data.GetSearchTotal(ctx, query)
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
