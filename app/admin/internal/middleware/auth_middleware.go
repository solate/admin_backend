package middleware

import (
	"net/http"

	"admin_backend/app/admin/internal/config"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/utils/jwt"

	"github.com/zeromicro/go-zero/rest/httpx"
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
			w.WriteHeader(http.StatusUnauthorized)
			httpx.Error(w, xerr.NewErrCodeMsg(xerr.TokenInvalid, "请先登录"))
			return
		}

		claims, err := jwt.ParseToken(tokenString, []byte(m.Config.JwtAuth.AccessSecret))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			httpx.Error(w, xerr.NewErrCodeMsg(xerr.TokenInvalid, err.Error()))
			return
		}

		// // 3. 租户隔离验证
		// if err := m.validateTenant(r, claims); err != nil {
		// 	httpx.Error(w, xerr.NewErrCodeMsg(http.StatusForbidden, err.Error()))
		// 	return
		// }

		// // 4. 权限验证（可选）
		// if len(m.requiredRoles) > 0 {
		// 	if err := m.validateRoles(claims); err != nil {
		// 		httpx.Error(w, xerr.NewErrCodeMsg(http.StatusForbidden, "权限不足"))
		// 		return
		// 	}
		// }

		if err := m.CheckAPIPermission(r, claims); err != nil {
			w.WriteHeader(http.StatusForbidden)
			httpx.Error(w, xerr.NewErrCodeMsg(http.StatusForbidden, "权限不足"))
			return
		}

		// 将租户ID和用户ID存入context
		ctx := r.Context()
		ctx = contextutil.SetUserIDToCtx(ctx, claims.UserID)
		ctx = contextutil.SetTenantCodeToCtx(ctx, claims.TenantCode)
		ctx = contextutil.SetRoleCodeToCtx(ctx, claims.RoleCode)

		next(w, r.WithContext(ctx))
	}
}

// // 租户隔离验证
// func (m *AuthMiddleware) validateTenant(r *http.Request, claims *jwt.Claims) error {
// 	reqTenant := r.Header.Get("x")
// 	if reqTenant != "" && claims.TenantID != reqTenant {
// 		return errors.New("租户身份不匹配")
// 	}
// 	return nil
// }

// 角色权限验证（可选）
func (m *AuthMiddleware) validateRoles(claims *jwt.Claims) error {
	// // 示例：从数据库或缓存中获取用户角色
	// userRoles := getUserRolesFromDB(claims.UserID)

	// // 检查用户是否具备所需角色
	// for _, requiredRole := range m.requiredRoles {
	// 	if !contains(userRoles, requiredRole) {
	// 		return error.New("用户角色不符合要求")
	// 	}
	// }

	return nil
}

func (m *AuthMiddleware) CheckAPIPermission(r *http.Request, claims *jwt.Claims) error {

	// userID := claims.UserID
	// tenantCode := claims.TenantCode
	// path := r.URL.Path
	// method := r.Method

	// hasPermission, err := casbin.PermissionManager.CheckPermission(userID, tenantCode, path, method, "api")

	// if err != nil || !hasPermission {

	// 	return xerr.NewErrCodeMsgByCode(xerr.ForbiddenError)
	// }

	return nil
}

func (m *AuthMiddleware) validatePermission(r *http.Request) error {

	// // 从上下文或请求头中获取用户信息
	// user := r.Header.Get("X-User-ID")
	// domain := r.Header.Get("X-Domain")

	// // 从请求路径构造资源标识
	// resource := strings.TrimPrefix(r.URL.Path, "/")
	// action := strings.ToLower(r.Method)

	// // 检查权限
	// hasPermission, err := m.pm.CheckPermission(user, domain, resource, action)
	// if err != nil {
	// 	fmt.Println("权限验证失败")
	// 	return err
	// }

	// if !hasPermission {
	// 	fmt.Println("无访问权限")
	// 	return err
	// }

	return nil

}

// 辅助函数：检查切片是否包含某个元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
