package biz

import (
	"context"

	"lin-cms-go/api"
)

type LogRepo interface {
	ListLog(ctx context.Context, page, size int) ([]*Log, int64, error)
	SearchLog(ctx context.Context, req api.SearchLogRequest) ([]*Log, int64, error)
}
type Log struct{}

func GetLogs(ctx context.Context, req *api.ListLogRequest) (res interface{}, total int, err error) {
	// var logs []*model.LinLog
	//
	// paging := data.NewPaging(page, size)
	// var query []predicate.LinLog
	// if req.Name != "" {
	// 	query = append(query, data.WithUsername(req.Name))
	// }
	// if req.Start != "" && req.End != "" {
	// 	start := utils.String2time(req.Start)
	// 	end := utils.String2time(req.End)
	// 	q := linlog.And(linlog.CreatedAtGT(start), linlog.CreatedAtLT(end))
	// 	query = append(query, q)
	// }
	// logs, err = paging.Search(ctx, query)
	// if err != nil {
	// 	err = errors.Wrap(err, "Search ")
	// 	return
	// }
	// total = data.GetSearchTotal(ctx, query)
	//
	// res = logs
	return
}

func SearchLogs(ctx context.Context, req api.SearchLogRequest, page int, size int) (res interface{}, total int, err error) {
	// var logs []*model.LinLog
	//
	// paging := data.NewPaging(page, size)
	// var query []predicate.LinLog
	// if req.Name != "" {
	// 	query = append(query, data.WithUsername(req.Name))
	// }
	// if req.Start != "" && req.End != "" {
	// 	start := utils.String2time(req.Start)
	// 	end := utils.String2time(req.End)
	// 	q := linlog.And(linlog.CreatedAtGT(start), linlog.CreatedAtLT(end))
	// 	query = append(query, q)
	// }
	// if req.Keyword != "" {
	// 	query = append(query, data.WithKeyword(req.Name))
	// }
	// logs, err = paging.Search(ctx, query)
	// if err != nil {
	// 	err = errors.Wrap(err, "Search ")
	// 	return
	// }
	// total = data.GetSearchTotal(ctx, query)
	// res = logs
	return
}

func GetLogUsers(ctx context.Context, page, size int) (res interface{}, total int, err error) {
	// paging := data.NewPaging(page, size)
	// res, err = paging.GetLogUsers(ctx)
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLogUsers ")
	// 	return
	// }
	// total, err = data.GetLogUsersTotal(ctx)
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLogUsersTotal ")
	// 	return
	// }
	return
}
