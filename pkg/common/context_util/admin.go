package context_util

import (
	"context"
	"errors"
)

// UserIDKey 是一个上下文键，用于存储用户ID
type UserIDKey struct{}

// SetUserIDToCtx 将用户ID设置到上下文中
func SetUserIDToCtx(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, UserIDKey{}, userID)
}

// GetUserIDFromCtx 从上下文中获取用户ID
func GetUserIDFromCtx(ctx context.Context) (uint64, error) {
	userID, ok := ctx.Value(UserIDKey{}).(uint64)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID, nil
}

// TenantIDKey 是一个上下文键，用于存储租户ID
type TenantIDKey struct{}

// SetTenantIDToCtx 将租户ID设置到上下文中
func SetTenantIDToCtx(ctx context.Context, tenantID uint64) context.Context {
	return context.WithValue(ctx, TenantIDKey{}, tenantID)
}

// GetTenantIDFromCtx 从上下文中获取租户ID
func GetTenantIDFromCtx(ctx context.Context) (uint64, error) {
	tenantID, ok := ctx.Value(TenantIDKey{}).(uint64)
	if !ok {
		return 0, errors.New("tenant ID not found in context")
	}
	return tenantID, nil
}
