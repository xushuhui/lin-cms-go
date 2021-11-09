package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"lin-cms-go/internal/data/model/predicate"

	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/linlog"
)

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

func GetSearchTotal(ctx context.Context, query []predicate.LinLog) (total int) {
	total, _ = GetDB().LinLog.Query().Where(query...).Count(ctx)
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
func CreateLog(ctx context.Context, statusCode, userId int, username, message, method, path, permission string) (err error) {

	_, err = GetDB().LinLog.Create().SetStatusCode(statusCode).SetUserID(userId).SetUsername(username).SetMessage(message).
		SetMethod(method).SetPath(path).SetPermission(permission).Save(ctx)
	return
}
