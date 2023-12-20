package auth

import (
	"context"
	"fmt"
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

func (s *service) RefreshAccessToken(
	ctx context.Context,
	oldAccessToken,
	refreshToken string,
	fingerprint entities.Fingerprint,
) (string, error) {
	username, _, _ := s.ParseToken(ctx, oldAccessToken)
	if username == "" {
		return "", fmt.Errorf("invalid old access token")
	}

	refreshTokenStruct, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if refreshTokenStruct.Fingerprint != fingerprint.Fingerprint ||
		refreshTokenStruct.Os != fingerprint.Os || refreshTokenStruct.UserAgent != fingerprint.UserAgent {
		// TODO: make type for error
		return "", fmt.Errorf("incorrect fingerprint")
	}

	if refreshTokenStruct.Token != refreshToken {
		// TODO: make type for error
		return "", fmt.Errorf("invalid refresh token")
	}

	if refreshTokenStruct.Expires.Before(time.Now()) {
		// TODO: make type for error
		return "", fmt.Errorf("refresh token has expired")
	}

	return s.GenerateAccessToken(ctx, username)
}

func (s *service) ParseToken(ctx context.Context, accessToken string) (string, bool, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signingKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false, fmt.Errorf("error get user claims from token")
	}

	username := claims["sub"].(string)
	if err != nil {
		return username, false, err
	}

	if !token.Valid {
		return username, false, fmt.Errorf("invalid access token")
	}

	return claims["sub"].(string), claims.VerifyExpiresAt(time.Now().Unix(), true), nil
}

func (s *service) RemoveRefreshToken(
	ctx context.Context,
	accessToken string,
	fp *entities.Fingerprint,
) error {
	username, _, _ := s.ParseToken(ctx, accessToken)
	if username == "" {
		// TODO: make custom type for error
		return fmt.Errorf("invalid access token")
	}

	refreshToken, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}

	if refreshToken.Fingerprint != fp.Fingerprint ||
		refreshToken.Os != fp.Os || refreshToken.UserAgent != fp.UserAgent {
		// TODO: make custom error
		return fmt.Errorf("cannot delete refresh token")
	}
	return s.repo.Remove(ctx, refreshToken)
}
