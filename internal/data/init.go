package data

import (
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/pkg/core"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var ds *DataSource
var once sync.Once

type Paging struct {
	Page   int
	Size   int
	Offset int
}

func NewPaging(page, pageSize int) *Paging {
	return &Paging{
		Page:   page,
		Size:   pageSize,
		Offset: core.GetPageOffset(page, pageSize),
	}
}

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

		db, err = ent.Open(conf.Database.Driver, conf.Database.Source, ent.Debug())
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
