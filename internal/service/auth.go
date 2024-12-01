package service

import (
	"errors"
	"fmt"
	"time"

	"auth-service/internal/model"
	"auth-service/internal/repository"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo      repository.UserRepository
	memcached     *memcache.Client
	logger        *zap.SugaredLogger
	jwtSigningKey []byte
	tokenTTL      int32 // seconds
}

func NewAuthService(repo repository.UserRepository, memcached *memcache.Client, logger *zap.SugaredLogger, secret []byte) AuthService {
	return &authService{
		userRepo:      repo,
		memcached:     memcached,
		logger:        logger,
		jwtSigningKey: secret,
		tokenTTL:      300, // 5 minutes
	}
}

func (s *authService) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("Failed to hash password:", err)
		return err
	}

	user := &model.User{Username: username, Password: string(hashedPassword)}
	if err := s.userRepo.CreateUser(user); err != nil {
		s.logger.Error("Failed to create user:", err)
		return err
	}
	return nil
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		s.logger.Error("Failed to find user:", err)
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Duration(s.tokenTTL) * time.Second).Unix(),
	})
	tokenString, err := token.SignedString(s.jwtSigningKey)
	if err != nil {
		s.logger.Error("Failed to generate token:", err)
		return "", err
	}

	err = s.memcached.Set(&memcache.Item{Key: "jwt:" + tokenString, Value: []byte(fmt.Sprint(user.ID)), Expiration: s.tokenTTL})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authService) ValidateToken(tokenString string) (uint, error) {
	item, err := s.memcached.Get("jwt:" + tokenString)
	if err != nil {
		s.logger.Warn("Token not found or expired:", err)
		return 0, errors.New("invalid or expired token")
	}

	return uint(item.Value[0]), nil
}
