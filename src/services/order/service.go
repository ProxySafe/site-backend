package order

import (
	"context"
	"sync"
	"time"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
	"github.com/ProxySafe/site-backend/src/services"
)

type service struct {
	mu        *sync.Mutex
	repo      repositories.IOrderRepository
	proxyRepo repositories.IProxyRepository
}

func NewService(
	repo repositories.IOrderRepository,
	proxyRepo repositories.IProxyRepository,
) services.IOrderService {
	return &service{
		mu:        &sync.Mutex{},
		repo:      repo,
		proxyRepo: proxyRepo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]entities.Order, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) GetByAccount(ctx context.Context, accountId int64) ([]entities.Order, error) {
	return s.repo.FindByAccountId(ctx, accountId)
}

func (s *service) CreateOrderByProxies(ctx context.Context, period, accountId int, proxies []entities.Proxy) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	orderWillLast := time.Hour * 24 * 30 * time.Duration(period)
	newOrder := &entities.Order{
		OrderDate:           now,
		AccountId:           accountId,
		OrderExpirationDate: now.Add(orderWillLast),
	}

	orderId, err := s.repo.Add(ctx, newOrder)
	if err != nil {
		return err
	}

	return s.proxyRepo.SetToOrder(ctx, int(orderId), proxies)
}
