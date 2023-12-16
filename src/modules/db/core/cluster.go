package core

import (
	"sync"

	"github.com/ProxySafe/site-backend/src/modules/db"
)

const (
	maxConnsCounterValue uint64 = 1e9 * 5
)

type cluster struct {
	n          uint64
	driverName string
	mu         *sync.RWMutex
	conn       map[db.NodeRole]*db.ExecutorsMap
}

func NewCluster(dbConf db.IConfigurator) (db.ICluster, error) {
	cl := &cluster{
		mu:         &sync.RWMutex{},
		driverName: dbConf.DriverName(),
		conn:       make(map[db.NodeRole]*db.ExecutorsMap),
	}

	if err := dbConf.Configure(cl); err != nil {
		return nil, err
	}
	return cl, nil
}

func (c *cluster) AddNode(role db.NodeRole, addr string) error {
	ex, err := NewExecutor(c.driverName, addr)
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	switch executors, ok := c.conn[role]; ok {
	case true:
		_, addrExists := (*executors)[addr]
		if addrExists {
			return &ErrAddrAlreadyExists{}
		}
		(*executors)[addr] = ex
	case false:
		e := &db.ExecutorsMap{
			addr: ex,
		}
		c.conn[role] = e
	}
	return nil
}

func (c *cluster) RemoveNode(role db.NodeRole, addr string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	executors, ok := c.conn[role]
	if !ok {
		return &ErrNoSuchRole{
			Role: role,
		}
	}

	delete(*executors, addr)
	return nil
}

func (c *cluster) Next(role db.NodeRole) (db.ISQLExecutor, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	executors, ok := c.conn[role]
	if !ok {
		return nil, &ErrNoSuchRole{
			Role: role,
		}
	}

	for _, ex := range *executors {
		return ex, nil
	}
	return nil, nil
}
