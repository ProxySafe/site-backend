package account

import (
	"context"
	"time"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/ProxySafe/site-backend/src/utils"
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

func (s *service) CreateAccount(ctx context.Context, userName, email, password string, telephone *string) (*entities.Account, error) {
	a := &entities.Account{
		Name:           userName,
		Email:          email,
		HashedPassword: utils.GetPasswordHash(password),
		CreatedAt:      time.Now(),
		Enabled:        true,
		Telephone:      telephone,
	}

	return a, s.repo.Add(ctx, a)
}
