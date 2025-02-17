package svc

import (
	"context"
	"fmt"

	"admin_backend/app/admin/internal/config"
	"admin_backend/app/admin/internal/middleware"
	"admin_backend/pkg/common/casbin"
	"admin_backend/pkg/ent"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/cache"
	"admin_backend/pkg/utils/captcha"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	DB              *generated.Client
	Redis           *redis.Redis
	AuthMiddleware  rest.Middleware
	PermissionCache *cache.PermissionCache
	CaptchaManager  *captcha.CaptchaManager // 验证码管理
	CasbinManager   *casbin.CasbinManager
}

func NewServiceContext(c config.Config) *ServiceContext {

	client := initOrm(c)
	rdb := initRedis(c)

	return &ServiceContext{
		Config:          c,
		DB:              client,
		Redis:           initRedis(c),
		AuthMiddleware:  middleware.NewAuthMiddleware(c).Handle,
		PermissionCache: cache.NewPermissionCache(rdb),
		CaptchaManager:  captcha.NewCaptchaManager(rdb),
		CasbinManager:   casbin.NewCasbinManager(client),
	}
}

// initOrm
func initOrm(c config.Config) *generated.Client {
	ops := make([]generated.Option, 0)
	if c.ShowSQL {
		ops = append(ops, generated.Debug())
	}
	client, err := ent.NewClient(context.Background(), c.DataSource, ops...)
	if err != nil {
		logx.Errorf("ent.Open error: %v", err)
		panic(err)
	}
	return client
}

func initRedis(c config.Config) *redis.Redis {
	return redis.New(fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port), func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
}
