package main

import (
	"context"
	"log"

	"admin_backend/cmd/init_db/initialize"
	"admin_backend/pkg/ent"
	"admin_backend/pkg/ent/generated"
)

func main() {
	/*
		DROP DATABASE IF EXISTS testdb;
		CREATE DATABASE testdb WITH ENCODING 'UTF8' LC_COLLATE='zh_CN.UTF-8' LC_CTYPE='zh_CN.UTF-8' TEMPLATE=template0;
	*/

	ctx := context.Background()
	// 初始化DB
	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable client_encoding=UTF8"
	client, err := ent.NewClient(ctx, dataSource)
	if err != nil {
		panic(err)
	}

	// 运行自动迁移工具
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ent.WithTx(ctx, client.Client, func(tx *generated.Tx) error {

		// 初始化租户
		_, err = initialize.InitTenant(ctx, tx)
		if err != nil {
			return err
		}

		// 初始化用户
		_, err = initialize.InitUser(ctx, tx, initialize.TenantCode)
		if err != nil {
			return err
		}

		// 初始化菜单
		if err := initialize.InitMenus(ctx, tx, initialize.TenantCode); err != nil {
			return err
		}

		// 初始化商品分类
		if err := initialize.InitCategories(ctx, tx, initialize.TenantCode); err != nil {
			return err
		}

		// 初始化字典类型
		if err := initialize.InitDictTypes(ctx, tx, initialize.TenantCode); err != nil {
			return err
		}

		// 初始化字典项
		if err := initialize.InitDictItems(ctx, tx, initialize.TenantCode); err != nil {
			return err
		}

		return nil
	})
}
