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

func (s *service) GetAll(ctx context.Context) ([]entities.Account, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) GetByUsername(ctx context.Context, userName string) (*entities.Account, error) {
	return s.repo.FindByUsername(ctx, userName)
}
