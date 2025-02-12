package useragent

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIP 返回客户端 IP
// 优先级：X-Forwarded-For > X-Real-IP > RemoteAddr
func GetClientIP(r *http.Request) string {
	// 1. 检查 X-Forwarded-For 头部
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，取第一个非空的 IP
		ips := strings.Split(xff, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if ip != "" {
				return ip
			}
		}
	}

	// 2. 检查 X-Real-IP 头部
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// 3. 最后使用 RemoteAddr
	// RemoteAddr 的格式通常是 "IP:Port"，需要去掉端口号
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// 如果解析失败，直接返回 RemoteAddr
		return r.RemoteAddr
	}

	return ip
}
