package data

import (
	"context"

	"lin-cms-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	"lin-cms-go/internal/data/model"
)

type logRepo struct {
	data *Data
	log  *log.Helper
}

func NewLogRepo(data *Data, logger log.Logger) biz.LogRepo {
	return &logRepo{data: data, log: log.NewHelper(logger)}
}

func (r *logRepo) Search(ctx context.Context, page, size int, keyword, name string) (logs []*model.LinLog, err error) {
	db := r.data.db.Model(&model.LinLog{})
	if keyword != "" {
		db = db.Where("message LIKE ?", "%"+keyword+"%")
	}
	if name != "" {
		db = db.Where("username = ?", name)
	}
	err = db.Order("created_at desc").Offset((page - 1) * size).Limit(size).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return
}

// func GetSearchTotal(ctx context.Context, query []predicate.LinLog) (total int) {
// 	total, _ = GetDB().LinLog.Query().Where(query...).Count(ctx)
// 	return
// }

// func (p *Paging) GetLogUsers(ctx context.Context) (model []string, err error) {
// 	model, err = GetDB().LinLog.Query().Limit(p.Size).Offset(p.Offset).GroupBy(linlog.FieldUsername).Strings(ctx)
// 	return
// }

// func GetLogUsersTotal(ctx context.Context) (total int, err error) {
// 	total, err = GetDB().LinLog.Query().Select(linlog.FieldUsername).Count(ctx)
// 	return
// }

// func CreateLog(ctx context.Context, statusCode, userId int, username, message, method, path, permission string) (err error) {
// 	_, err = GetDB().LinLog.Create().SetStatusCode(statusCode).SetUserID(userId).SetUsername(username).SetMessage(message).
// 		SetMethod(method).SetPath(path).SetPermission(permission).Save(ctx)
// 	return
// }
