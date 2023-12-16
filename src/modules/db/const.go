package db

type NodeRole int

func (r NodeRole) String() string {
	switch r {
	case Master:
		return `master`
	case LightSlave:
		return `light_slave`
	case HeavySlave:
		return `heavy_slave`
	}
	return `unknown`
}

const (
	Master NodeRole = iota
	LightSlave
	HeavySlave
	LastRole
)
