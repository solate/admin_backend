package middleware

import (
	"net/http"
)

// CorsMiddleware 是一个自定义的跨域中间件
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源访问（可以根据需求限制特定域名）
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的 HTTP 方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的自定义请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 如果是预检请求（OPTIONS 方法），直接返回 200
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// 调用下一个处理器
		next(w, r)
	}
}
