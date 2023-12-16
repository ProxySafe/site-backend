package repositories

import (
	"github.com/ProxySafe/site-backend/src/app/resources"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
)

type Repositories struct {
	AccountRepository repositories.IAccountRepository
}

func NewRepositories(res *resources.Resources) *Repositories {
	db := res.DBManager
	return &Repositories{
		AccountRepository: repositories.NewAccountRepository(db),
	}
}
