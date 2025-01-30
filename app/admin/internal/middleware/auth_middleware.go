package middleware

import (
	"context"
	"net/http"

	"github.com/solate/admin_backend/app/admin/internal/config"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/utils/jwt"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	TenantIDKey = "tenantID"
	UserIDKey   = "userID"
)

type AuthMiddleware struct {
	Config config.Config
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Config: c,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusBadRequest, "请先登录"))
			return
		}

		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusInternalServerError, "token无效"))
			return
		}

		// 将租户ID和用户ID存入context
		ctx := context.WithValue(r.Context(), TenantIDKey, claims.TenantID)
		ctx = context.WithValue(ctx, UserIDKey, claims.Subject)

		next(w, r.WithContext(ctx))
	}
}
