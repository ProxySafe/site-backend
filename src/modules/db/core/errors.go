package core

import (
	"fmt"

	"github.com/ProxySafe/site-backend/src/modules/db"
)

type ErrNoSuchRole struct {
	Role db.NodeRole
}

func (e *ErrNoSuchRole) Error() string {
	return fmt.Sprintf("no such role: %s", e.Role.String())
}

type ErrAddrAlreadyExists struct {
	Addr string
}

func (e ErrAddrAlreadyExists) Error() string {
	return fmt.Sprintf("addr already exists: %s", e.Addr)
}
