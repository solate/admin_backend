package ent

import (
	"context"

	"admin_backend/pkg/ent/generated"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewClient(ctx context.Context, dataSource string) (*generated.Client, error) {
	logx.Infof("Initializing new client with data source: %s", dataSource)

	client, err := generated.Open("postgres", dataSource)
	if err != nil {
		logx.Errorf("ent.Open error: %v", err)
		return nil, err
	}

	// 打印客户端信息
	logx.Info("Created client:", client)

	// 运行数据库迁移
	logx.Info("Running schema migration...")
	if err := client.Schema.Create(ctx); err != nil { // 移除了 WithDropIndex 选项
		logx.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}

	logx.Info("Schema migration completed successfully.")
	return client, nil
}
