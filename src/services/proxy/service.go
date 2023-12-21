package proxy

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
	"github.com/ProxySafe/site-backend/src/services"
)

type service struct {
	repo repositories.IProxyRepository
}

func NewService(repo repositories.IProxyRepository) services.IProxyService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]entities.Proxy, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) GetNoBusy(ctx context.Context) ([]entities.Proxy, error) {
	return s.repo.FindNoBusy(ctx)
}

func (s *service) GetByAccount(ctx context.Context, accountId int64) ([]entities.Proxy, error) {
	return s.repo.FindByAccountId(ctx, accountId)
}

func (s *service) GetProxiesByAmount(ctx context.Context, amount int) ([]entities.Proxy, error) {
	availableProxies, err := s.GetNoBusy(ctx)
	if err != nil {
		return nil, err
	}

	if len(availableProxies) < amount {
		return nil, &ErrNoAvailableProxies{
			Amount: amount,
		}
	}

	return availableProxies[:amount], nil
}
