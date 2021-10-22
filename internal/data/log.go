package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"lin-cms-go/internal/data/model/predicate"

	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/linlog"
	"time"
)

func (p *Paging) FindLogsByUsernameAndRange(ctx context.Context, name string, start time.Time, end time.Time) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(name)).Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) GetLogAll(ctx context.Context) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByUsername(ctx context.Context, name string) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(name)).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByRange(ctx context.Context, start time.Time, end time.Time) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}
func WithKeyword(keyword string) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}
}
func WithUsername(name string) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		s.Where(sql.EQ(linlog.FieldUsername, name))
	}
}
func (p *Paging) Search(ctx context.Context, query []predicate.LinLog) (logs []*model.LinLog, err error) {
	logs, err = GetDB().LinLog.Query().Where(query...).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}
func (p *Paging) FindLogsByUsernameAndRangeAndKeyword(ctx context.Context, name string, keyword string, start time.Time, end time.Time) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(name)).Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByKeyword(ctx context.Context, keyword string) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByUsernameAndKeyword(ctx context.Context, username string, keyword string) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.Username(username)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Limit(p.Size).Offset(p.Offset).All(ctx)
	return
}

func (p *Paging) FindLogsByRangeAndKeyword(ctx context.Context, keyword string, start time.Time, end time.Time) (model []*model.LinLog, err error) {
	model, err = GetDB().LinLog.Query().Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Limit(p.Size).Offset(p.Offset).All(ctx)
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

func GetLogsByRangeTotal(ctx context.Context, start time.Time, end time.Time) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.CreateTimeGT(start)).Where(linlog.CreateTimeLT(end)).Count(ctx)
	return
}

func GetLogsByUsernameAndRangeAndKeywordTotal(ctx context.Context, name string, keyword string, start time.Time, end time.Time) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.Username(name)).Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Count(ctx)
	return
}

func GetLogsByKeywordTotal(ctx context.Context, keyword string) (total int) {
	total, _ = GetDB().LinLog.Query().Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Count(ctx)
	return
}

func GetLogsByUsernameAndKeywordTotal(ctx context.Context, username string, keyword string) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.Username(username)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Count(ctx)
	return
}

func GetLogsByRangeAndKeywordTotal(ctx context.Context, keyword string, start time.Time, end time.Time) (total int) {
	total, _ = GetDB().LinLog.Query().Where(linlog.And(
		linlog.CreateTimeGT(start),
		linlog.CreateTimeLT(end),
	)).Where(func(s *sql.Selector) {
		s.Where(sql.Like(linlog.FieldMessage, "%"+keyword+"%"))
	}).Count(ctx)
	return
}
