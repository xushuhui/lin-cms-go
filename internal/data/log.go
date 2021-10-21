package data

import (
	"context"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/linlog"
	"time"
)

func (p *Paging) FindLogsByUsernameAndRange(ctx context.Context, name string, start time.Time, end time.Time) (model []*ent.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(name)).Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) GetLogAll(ctx context.Context) (model []*ent.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByUsername(ctx context.Context, name string) (model []*ent.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(name)).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) GetLogUsers(ctx context.Context) (model []string, err error) {
	model, err = GetDB().LinLog.Query().Limit(p.Size).Offset(p.Offset).GroupBy(linlog.FieldUsername).Strings(ctx)
	return
}

func GetLogUsersTotal(ctx context.Context) (total int, err error) {
	total, err = GetDB().LinLog.Query().Select(linlog.FieldUsername).Count(ctx)
	return
}

func GetLosTotal(ctx context.Context) (total int) {
	total, _ = GetDB().LinLog.Query().Count(ctx)
	return
}

func GetLogsByUsernameTotal(ctx context.Context, name string) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.Username(name)).Count(ctx)
	return
}

func GetLogsByUsernameAndRangeTotal(ctx context.Context, name string, start time.Time, end time.Time) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.Username(name)).Where(linlog.CreateTimeGT(start)).Where(linlog.CreateTimeLT(end)).Count(ctx)
	return
}
