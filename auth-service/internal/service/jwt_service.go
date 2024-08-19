package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtService interface {
	GenerateJWT(interface{}) (string, error)
	ValidateJWT(string) (*Claims, error)
}

type jwtService struct {
	jwtKey []byte
}

type Claims struct {
	data interface{}
	jwt.RegisteredClaims
}

func NewJwtService() JwtService {
	return &jwtService{
		jwtKey: []byte(viper.GetString("authentication.jwt_secret")),
	}
}

func (k *jwtService) GenerateJWT(payload interface{}) (string, error) {
	claims := &Claims{
		data: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(k.jwtKey)
}

func (k *jwtService) ValidateJWT(token string) (*Claims, error) {
	claims := &Claims{}
	tk, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return k.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
