package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("your-secret-key")

type TenantClaims struct {
	TenantID string `json:"tenantID"`
	jwt.RegisteredClaims
}

// 生成带租户ID的Token
func GenerateToken(tenantID string, userID string) (string, error) {
	claims := TenantClaims{
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// 解析并验证Token
func ParseToken(tokenString string) (*TenantClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TenantClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(*TenantClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
