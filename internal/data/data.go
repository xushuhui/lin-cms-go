package data

import (
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data/ent"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDBClient)

// Data .
type DataSource struct {
	// warpped database client
	db *ent.Client
}

var once sync.Once
var db *ent.Client

func GetDBClient() *ent.Client {
	return db
}

func NewDBClient(conf *conf.Data) *ent.Client {
	var err error
	once.Do(func() {
		opt := ent.Debug()
		db, err = ent.Open(conf.Database.Driver, conf.Database.Source, opt)
		if err != nil {
			panic(err)
		}
	})
	return db
}

// NewData .
func NewData(c *conf.Data, client *ent.Client) (*DataSource, func(), error) {
	cleanup := func() {
		//if err := client.Close(); err != nil {
		//	log.NewHelper(logger).Error(err)
		//}
		//log.NewHelper(logger).Info("closing the data resources")
	}

	return &DataSource{
		client,
	}, cleanup, nil
}
