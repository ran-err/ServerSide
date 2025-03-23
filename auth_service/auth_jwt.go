package auth_service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func (c *JWTClaims) GetUserID() string {
	return c.UserID
}

type AuthServiceJWT struct {
	secretKey string
}

func NewAuthServiceJWT(secretKey string) AuthService {
	return &AuthServiceJWT{
		secretKey: secretKey,
	}
}

func (s *AuthServiceJWT) GenerateToken(claims Claims) (string, error) {
	jwtClaims := &JWTClaims{
		UserID: claims.GetUserID(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthServiceJWT) ParseToken(tokenStr string) (Claims, error) {
	jwtClaims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return jwtClaims, nil
}
