package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Vis7044/GinCrud2/models"
	"github.com/Vis7044/GinCrud2/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(r *repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Register(ctx,user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
    user, err := s.repo.FindByEmail(ctx, email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return "", errors.New("invalid credentials")
    }

    claims := jwt.MapClaims{
        "userId": user.Id.Hex(),
        "email":  user.Email,
        "exp":    time.Now().Add(time.Hour * 24).Unix(),
        "iat":    time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err = godotenv.Load()
	if err != nil {
		return "", errors.New("not secret token")
	}
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        return "", errors.New("jwt secret not found")
    }

    return token.SignedString([]byte(jwtSecret))
}
