package passwordgen

import (
	"strings"
	"testing"
)

func TestPasswordFlow(t *testing.T) {
	// 生成密码
	pwd, err := GeneratePassword(16)
	if err != nil {
		t.Fatalf("密码生成失败: %v", err)
	}
	t.Log(pwd)

	// 生成盐值
	salt, err := GenerateSalt()
	if err != nil {
		t.Fatalf("盐值生成失败: %v", err)
	}
	t.Log(salt)

	// 生成哈希
	hashed, err := Argon2Hash(pwd, salt)
	if err != nil {
		t.Fatalf("哈希生成失败: %v", err)
	}
	t.Log(hashed)

	// 验证正确密码
	if !VerifyPassword(pwd, hashed) {
		t.Error("正确密码验证失败")
	}

	// 验证错误密码
	if VerifyPassword("wrong_password", hashed) {
		t.Error("错误密码验证通过")
	}

	// 验证篡改哈希
	tampered := strings.Replace(hashed, "$", "#", 1)
	if VerifyPassword(pwd, tampered) {
		t.Error("篡改哈希验证通过")
	}
}
