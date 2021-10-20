package data

import (
	"github.com/grestful/logs"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data/ent"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var ds *DataSource
var once sync.Once

// DataSource .
type DataSource struct {
	Db map[string]*ent.Client
}

func NewDataSource(conf *conf.Data) {
	ds = &DataSource{
		Db: make(map[string]*ent.Client),
	}
	db := NewDBClient(conf)
	ds.Db["default"] = db
}

func NewDBClient(conf *conf.Data) (db *ent.Client) {

	var err error
	once.Do(func() {

		db, err = ent.Open(conf.Database.Driver, conf.Database.Source, ent.Debug(), ent.Log(func(i ...interface{}) {
			logs.Debug("sql", i, 111111111111)
		}))
		if err != nil {
			panic(err)
		}
	})

	return db
}

func (d *DataSource) GetDb(name string) *ent.Client {
	if db, ok := d.Db[name]; ok {
		return db
	}
	panic(name + " db not exists")
}

func GetDB() *ent.Client {
	return ds.GetDb("default")
}
