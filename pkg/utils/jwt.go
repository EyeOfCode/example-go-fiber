package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	secretKey        string
	refreshSecretKey string
	expiresIn        string
	refreshExpiresIn string
}

type JWTClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthHandler(secretKey, refreshSecretKey, expiresIn, refreshExpiresIn string) *AuthHandler {
	return &AuthHandler{
		secretKey:        secretKey,
		refreshSecretKey: refreshSecretKey,
		expiresIn:        expiresIn,
		refreshExpiresIn: refreshExpiresIn,
	}
}

func (s *AuthHandler) GenerateToken(userID string, roles []string) (string, error) {
	expDuration, _ := time.ParseDuration(s.expiresIn)
	claims := JWTClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthHandler) GenerateRefreshToken(userID string, roles []string) (string, error) {
	expDuration, _ := time.ParseDuration(s.refreshExpiresIn)
	claims := JWTClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.refreshSecretKey))
}

func (s *AuthHandler) GenerateTokenPair(userID string, roles []string) (*TokenPair, error) {
	accessToken, err := s.GenerateToken(userID, roles)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefreshToken(userID, roles)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthHandler) ValidateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func (s *AuthHandler) ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.refreshSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
