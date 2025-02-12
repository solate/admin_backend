package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID     string `json:"user_id"`
	TenantCode string `json:"tenant_code"`
	RoleCode   string `json:"role_code"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	AccessSecret []byte // 密钥
	AccessExpire int64  // 过期时间
}

// 生成JWT Token
func GenerateToken(userID, tenantCode string, config JWTConfig) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:     userID,
		TenantCode: tenantCode,
		// RoleCode:   roleCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.AccessExpire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.AccessSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析JWT Token
func ParseToken(tokenString string, accessSecret []byte) (*Claims, error) {
	// 检查并去除Bearer前缀
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("token已过期")
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, fmt.Errorf("token格式错误")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return nil, fmt.Errorf("token签名无效")
		default:
			return nil, fmt.Errorf("token解析失败: %v", err)
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// 验证Token是否有效
func ValidateToken(tokenString string, accessSecret []byte) bool {
	_, err := ParseToken(tokenString, accessSecret)
	return err == nil
}
