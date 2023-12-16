package postgres

import "github.com/ProxySafe/site-backend/src/modules/db"

type pgConfigurator struct {
	master []string
	slave  []string
}

func NewPgConfigurator(config *DbConfig) db.IConfigurator {
	return &pgConfigurator{
		master: config.PGReadNodes,
		slave:  config.PGWriteNodes,
	}
}

func (c *pgConfigurator) Configure(cl db.ICluster) error {
	for _, addr := range c.master {
		if err := cl.AddNode(db.Master, addr); err != nil {
			return err
		}
	}

	for _, addr := range c.slave {
		if err := cl.AddNode(db.LightSlave, addr); err != nil {
			return err
		}
	}
	return nil
}

func (c *pgConfigurator) DriverName() string {
	return "postgres"
}
