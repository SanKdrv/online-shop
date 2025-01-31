package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

type UsersService struct {
	repo repository.Users
}

func (s *UsersService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	claimsAccess := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(tokenTTL).Unix(),
		"iat":     time.Now().Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	accessTokenString, err := accessToken.SignedString("dsfsdffdssd")
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}

func (s *UsersService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(user domain.User) (int64, error) {
	if (user.Username == "") || (user.PasswordHash == "") || (user.Email == "") {
		return 0, errors.New("required fields are empty or invalid")
	}
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UsersService) VerifyUser(identifier, password, verifyBy string) (domain.User, error) {
	var user domain.User
	var err error

	switch verifyBy {
	case "username":
		user, err = s.repo.GetUserByUsername(identifier, generatePasswordHash(password))
	case "email":
		user, err = s.repo.GetUserByEmail(identifier, generatePasswordHash(password))
	default:
		return domain.User{}, errors.New("unsupported verification method")
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, err // User not found
		}
		return domain.User{}, err // Other error
	}

	return user, nil // User verified
}
