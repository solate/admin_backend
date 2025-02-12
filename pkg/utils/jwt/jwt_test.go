package jwt

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	// 生成token
	token, err := GenerateToken("1", "test", JWTConfig{
		AccessExpire: 3600,
		AccessSecret: []byte("test"),
	})
	if err != nil {
		t.Fatal(err)
		fmt.Println("token:", token)
	}

	// 测试解析带Bearer前缀的token
	claims, err := ParseToken("Bearer "+token, []byte("test"))
	if err != nil {
		t.Fatal(err)
		t.Log("解析带Bearer前缀的token成功:", claims.UserID, claims.TenantCode)
	}
	t.Log("解析带Bearer前缀的token成功:", claims.UserID, claims.TenantCode)

	claims, err = ParseToken(token, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("解析不带Bearer前缀的token成功:", claims.UserID, claims.TenantCode)

	// 测试无效的token
	_, err = ParseToken("invalid_token", []byte("test"))
	if err == nil {
		t.Fatal("期望解析无效token时返回错误，但没有")
	}
	t.Log("无效token测试通过")
}
