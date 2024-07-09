package data

import (
	"lin-cms-go/internal/conf"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewBookRepo,NewLinUserRepo,NewLessonRepo,NewTeacherRepo)

type Data struct {
	db *gorm.DB
}

func NewData(conf *conf.Data) (*Data, func(), error) {
	d := &Data{
		db: NewDB(conf),
	}
	main, _ := d.db.DB()

	return d, func() {
		main.Close()
	}, nil
}

func NewDB(conf *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Main), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.Debug()
}
