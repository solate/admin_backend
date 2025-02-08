package jwt

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {
	token, err := GenerateToken(1, "test", JWTConfig{
		AccessExpire: 3600,
		AccessSecret: []byte("test"),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("token:", token)
	claims, err := ParseToken(token, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims.UserID, claims.TenantCode, claims.RoleCode)
	fmt.Println("end success")
}
