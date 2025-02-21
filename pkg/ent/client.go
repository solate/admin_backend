package ent

import (
	"context"
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"

	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/ent/generated/migrate"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
)

type Client struct {
	*generated.Client
	*sql.DB
}

func NewClient(ctx context.Context, dataSource string, ops ...generated.Option) (*Client, error) {
	logx.Infof("Initializing new client with data source: %s", dataSource)

	drv, err := entsql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Client: generated.NewClient(append(ops, generated.Driver(drv))...),
		DB:     drv.DB(),
	}

	// client, err := generated.Open("postgres", dataSource, ops...)
	// if err != nil {
	// 	logx.Errorf("ent.Open error: %v", err)
	// 	return nil, err
	// }
	// 打印客户端信息
	logx.Info("Created client:", client)

	// 运行数据库迁移
	logx.Info("Running schema migration...")
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil { // 移除了 WithDropIndex 选项
		logx.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}

	logx.Info("Schema migration completed successfully.")
	return client, nil
}
