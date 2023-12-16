package account

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
	"github.com/ProxySafe/site-backend/src/services"
)

type service struct {
	repo repositories.IAccountRepository
}

func NewService(repo repositories.IAccountRepository) services.IAccountService {
	return &service{
		repo: repo,
	}
}

func (s *service) FindAll(ctx context.Context) ([]entities.Account, error) {
	return s.repo.GetAll(ctx)
}
