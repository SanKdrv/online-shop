package service

import (
	"backend/internal/config"
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type RefreshSessionService struct {
	repo repository.RefreshSession
}

func NewRefreshSessionService(repo repository.RefreshSession) *RefreshSessionService {
	return &RefreshSessionService{repo: repo}
}

//CreateRefreshSession

func (s *RefreshSessionService) CreateRefreshSession(user domain.User, obj domain.RefreshSession, cfg *config.Config) (string, string, error) {
	if session, err := s.repo.CreateRefreshSession(obj); err != nil {
		slog.Info("service failed to create refresh session", slog.String("error", err.Error()))
		return "", "", err
	} else {
		//access := domain.AccessToken{}
		claimsAccess := jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(cfg.Auth.AccessTokenTTL).Unix(),
			"iat":     time.Now().Unix(),
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
		accessTokenString, err := accessToken.SignedString([]byte(cfg.Auth.JWTSecret)[0:])
		refreshTokenString := session.RefreshToken
		slog.Info("service refresh session created", slog.String("access_token", accessTokenString), slog.String("refresh_token", refreshTokenString))
		if err != nil {
			return "", "", err
		}
		return accessTokenString, refreshTokenString, nil
	}
}

func (s *RefreshSessionService) UpdateRefreshSession(refreshToken domain.RefreshToken, req RefreshTokensRequest, cfg *config.Config) (domain.AccessToken, domain.RefreshToken, error) {
	// Удаляем старый токен из базы и содержимое сохраняем в переменную
	if refreshSession, err := s.repo.DeleteRefreshSession(refreshToken); err != nil {
		slog.Info("service failed to delete refresh session", slog.String("error", err.Error()))
		return domain.AccessToken{}, domain.RefreshToken{}, err
	} else {
		if refreshSession.ExpiresIn > time.Now().Unix() {
			// Создает новую рефреш-сессию и записывает ее в БД
			newRefrefreshSession := domain.RefreshSession{
				UserID:       refreshSession.UserID,
				RefreshToken: uuid.New().String(),
				UserAgent:    req.UserAgent,
				ExpiresIn:    time.Now().Add(time.Minute * cfg.Auth.RefreshTokenTTL).Unix(),
				CreatedAt:    time.Now(),
			}
			if session, err := s.repo.CreateRefreshSession(newRefrefreshSession); err != nil {
				slog.Info("service failed to create refresh session", slog.String("error", err.Error()))
				return domain.AccessToken{}, domain.RefreshToken{}, err
			} else {
				// Создаёт refresh_token
				newRefreshToken := domain.RefreshToken{
					RefreshToken: session.RefreshToken,
				}

				// Создаёт access_token
				claimsAccess := jwt.MapClaims{
					// Возьмём user_id из рефреш-сессии
					"user_id": session.UserID,
					"exp":     time.Now().Add(cfg.Auth.AccessTokenTTL).Unix(),
					"iat":     time.Now().Unix(),
				}
				accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
				accessTokenString, err := accessToken.SignedString([]byte(cfg.Auth.JWTSecret)[0:])
				if err != nil {
					return domain.AccessToken{}, domain.RefreshToken{}, err
				}
				newAccessToken := domain.AccessToken{
					AccessToken: accessTokenString,
				}

				// Возвращает access_token и refresh_token
				return newAccessToken, newRefreshToken, nil
			}
		} else {
			// Бросает ошибку TOKEN_EXPIRED/INVALID_REFRESH_SESSION
			return domain.AccessToken{}, domain.RefreshToken{}, errors.New("TOKEN_EXPIRED")
		}
	}
}

func (s *RefreshSessionService) GetRefreshSession(RefreshToken string) (domain.RefreshSession, error) {
	return domain.RefreshSession{}, nil
}

func (s *RefreshSessionService) DeleteRefreshSession(refreshToken string) error {
	refreshTokenObj := domain.RefreshToken{RefreshToken: refreshToken}
	if _, err := s.repo.DeleteRefreshSession(refreshTokenObj); err != nil {
		slog.Info("service failed to delete refresh session", slog.String("error", err.Error()))
		return err
	}
	return nil
}
