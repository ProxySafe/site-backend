package auth

import (
	"context"
	"time"

	"github.com/ProxySafe/site-backend/src/domains/entities"
	"github.com/ProxySafe/site-backend/src/domains/repositories"
	"github.com/ProxySafe/site-backend/src/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type service struct {
	tokenTTL   int64
	signingKey string
	repo       repositories.IRefreshTokenRepository
}

func NewService(
	signingKey string,
	tokenTTL int64,
	repo repositories.IRefreshTokenRepository,
) services.IAuthService {
	return &service{
		signingKey: signingKey,
		repo:       repo,
	}
}

func (s *service) GenerateAccessToken(ctx context.Context, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(s.tokenTTL)).Unix(),
		Subject:   userName,
	})

	return token.SignedString([]byte(s.signingKey))
}

func (s *service) GenerateRefreshToken(
	ctx context.Context,
	accountId int64,
	fingerprint entities.Fingerprint,
) (string, error) {
	// TODO: think about how to make it more complicated
	token := uuid.NewString()

	refreshToken := &entities.RefreshToken{
		Token:       token,
		AccountId:   accountId,
		Expires:     time.Now().Add(time.Hour * 40),
		Fingerprint: fingerprint.Fingerprint,
		Os:          fingerprint.Os,
		UserAgent:   fingerprint.UserAgent,
	}

	if err := s.repo.Add(ctx, refreshToken); err != nil {
		return "", err
	}

	return token, nil
}
