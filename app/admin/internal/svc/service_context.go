package svc

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/solate/admin_backend/app/admin/internal/config"
	"github.com/solate/admin_backend/app/admin/internal/middleware"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/migrate"
	"github.com/solate/admin_backend/pkg/utils/cache"
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
	}
}

// initOrm
func initOrm(c config.Config) *generated.Client {
	ops := make([]generated.Option, 0)
	if c.ShowSQL {
		ops = append(ops, generated.Debug())
	}
	client, err := generated.Open(dialect.Postgres, c.DataSource, ops...)
	if err != nil {
		logx.Errorf("ent.Open error: %v", err)
		panic(err)
	}
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
	); err != nil {
		logx.Errorf("ent.Schema.Create error: %v", err)
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
