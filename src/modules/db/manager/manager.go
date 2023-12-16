package manager

import (
	"log"

	"github.com/ProxySafe/site-backend/src/modules/db"
)

type manager struct {
	cluster db.ICluster
}

func NewManager(cl db.ICluster) db.IDBManager {
	return &manager{
		cluster: cl,
	}
}

func (m *manager) ReadDB() db.ISQLExecutor {
	ex, err := m.cluster.Next(db.LightSlave)
	if err != nil {
		log.Fatal(err)
	}

	return ex
}

func (m *manager) WriteDB() db.ISQLExecutor {
	ex, err := m.cluster.Next(db.Master)
	if err != nil {
		log.Fatal(err)
	}

	return ex
}
