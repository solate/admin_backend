package captcha

import (
	"math/rand"
	"time"
)

// GenerateCaptcha 生成一个指定长度的随机验证码
func GenerateCaptcha(length int) string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
