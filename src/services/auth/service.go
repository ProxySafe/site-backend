package auth

import "github.com/ProxySafe/site-backend/src/services"

type service struct {
}

func NewService() services.IAuthService {
	return &service{}
}
