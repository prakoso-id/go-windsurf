package services

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type AuthService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	secretKey string
	expiration time.Duration
}

func NewAuthService() AuthService {
	return &authService{
		secretKey: getSecretKey(),
		expiration: getJWTExpiration(),
	}
}

func getSecretKey() string {
	return os.Getenv("JWT_SECRET")
}

func getJWTExpiration() time.Duration {
	exp, _ := time.ParseDuration(os.Getenv("JWT_EXPIRATION"))
	return exp
}

func (s *authService) GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(s.expiration).Unix(),
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}
