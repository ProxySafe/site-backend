package resources

import (
	"github.com/ProxySafe/site-backend/src/app/config"
	"github.com/ProxySafe/site-backend/src/modules/db"
	"github.com/ProxySafe/site-backend/src/modules/db/configurator/postgres"
	"github.com/ProxySafe/site-backend/src/modules/db/core"
	"github.com/ProxySafe/site-backend/src/modules/db/manager"
	_ "github.com/lib/pq"
)

type Resources struct {
	DBManager  db.IDBManager
	SigningKey string
	TokenTTL   int64
}

func NewResources(cfg *config.Config) *Resources {
	r := &Resources{
		SigningKey: cfg.SigningKey,
		TokenTTL:   cfg.TokenTTL,
	}
	r.initDBManager(cfg)
	return r
}

func (r *Resources) initDBManager(cfg *config.Config) {
	configurator := postgres.NewPgConfigurator(cfg.DB.Base)
	cluster, err := core.NewCluster(configurator)
	if err != nil {
		panic(err)
	}

	r.DBManager = manager.NewManager(cluster)
}
