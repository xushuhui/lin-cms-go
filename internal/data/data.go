package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewLinUserRepo)

// Data .
type Data struct {
	// warpped database client
	db *ent.Client
}

func NewDB(conf *conf.Data) *ent.Client {

	opt := ent.Debug()
	client, err := ent.Open(conf.Database.Driver, conf.Database.Source, opt)
	if err != nil {
		panic(err)
	}

	return client
}

// NewData .
func NewData(c *conf.Data, client *ent.Client) (*Data, func(), error) {
	cleanup := func() {
		//if err := client.Close(); err != nil {
		//	log.NewHelper(logger).Error(err)
		//}
		//log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		client,
	}, cleanup, nil
}
