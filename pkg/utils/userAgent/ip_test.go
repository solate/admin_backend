package userAgent

import (
	"net/http/httptest"
	"testing"
)

func TestGetClientIP(t *testing.T) {
	tests := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		expectedIP string
	}{
		{
			name: "X-Forwarded-For exists",
			headers: map[string]string{
				"X-Forwarded-For": "192.168.1.1, 10.0.0.1",
			},
			expectedIP: "192.168.1.1",
		},
		{
			name: "X-Real-IP exists, X-Forwarded-For does not exist",
			headers: map[string]string{
				"X-Real-IP": "192.168.1.2",
			},
			expectedIP: "192.168.1.2",
		},
		{
			name:       "Only RemoteAddr exists",
			remoteAddr: "127.0.0.1:12345",
			expectedIP: "127.0.0.1",
		},
		{
			name: "X-Forwarded-For with multiple IPs",
			headers: map[string]string{
				"X-Forwarded-For": "192.168.1.3, 10.0.0.2, 10.0.0.3",
			},
			expectedIP: "192.168.1.3",
		},
		{
			name:       "No headers, only RemoteAddr",
			remoteAddr: "[::1]:12345", // IPv6 address
			expectedIP: "::1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}
			if tt.remoteAddr != "" {
				req.RemoteAddr = tt.remoteAddr
			}

			ip := GetClientIP(req)
			if ip != tt.expectedIP {
				t.Errorf("expected IP %s, got %s", tt.expectedIP, ip)
			}
		})
	}
}
