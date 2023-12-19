package services

import (
	"context"

	"github.com/ProxySafe/site-backend/src/domains/entities"
)

type IAccountService interface {
	GetAll(ctx context.Context) ([]entities.Account, error)
	GetByUsername(ctx context.Context, userName string) (*entities.Account, error)
	CreateAccount(ctx context.Context, userName, email, password string, telephone *string) (*entities.Account, error)
}

type IAuthService interface {
	GenerateAccessToken(ctx context.Context, userName string) (string, error)
	GenerateRefreshToken(
		ctx context.Context,
		accountId int64,
		fingerprint entities.Fingerprint,
	) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, bool, error)
	RefreshAccessToken(
		ctx context.Context,
		oldAccessToken,
		refreshToken string,
		fingerprint entities.Fingerprint,
	) (string, error)
	RemoveRefreshToken(ctx context.Context, accessToken string) error
}

type IEmailService interface {
}
