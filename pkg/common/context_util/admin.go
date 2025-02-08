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

// TenantCodeKey 是一个上下文键，用于存储租户ID
type TenantCodeKey struct{}

// SetTenantCodeToCtx 将租户ID设置到上下文中
func SetTenantCodeToCtx(ctx context.Context, tenantCode string) context.Context {
	return context.WithValue(ctx, TenantCodeKey{}, tenantCode)
}

// GetTenantCodeFromCtx 从上下文中获取租户ID
func GetTenantCodeFromCtx(ctx context.Context) (string, error) { // 获取租户ID
	tenantCode, ok := ctx.Value(TenantCodeKey{}).(string)
	if !ok {
		return "", errors.New("tenant code not found in context")
	}
	return tenantCode, nil
}

// RoleCodeKey 是一个上下文键，用于存储角色ID
type RoleCodeKey struct{}

// SetRoleCodeToCtx 将角色ID设置到上下文中
func SetRoleCodeToCtx(ctx context.Context, roleCode string) context.Context {
	return context.WithValue(ctx, RoleCodeKey{}, roleCode)
}

// GetRoleCodeFromCtx 从上下文中获取角色ID
func GetRoleCodeFromCtx(ctx context.Context) (string, error) {
	roleCode, ok := ctx.Value(RoleCodeKey{}).(string)
	if !ok {
		return "", errors.New("role code not found in context")
	}
	return roleCode, nil
}
