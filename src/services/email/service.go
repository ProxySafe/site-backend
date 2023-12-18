package email

import "github.com/ProxySafe/site-backend/src/services"

type service struct {
}

func NewService() services.IEmailService {
	return &service{}
}
