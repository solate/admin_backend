package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/solate/admin_backend/pkg/ent"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/utils/idgen"
	"github.com/solate/admin_backend/pkg/utils/passwordgen"
)

const (
	tenantCode = "default"
	password   = "admin@123"
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

	ent.WithTx(ctx, client, func(tx *generated.Tx) error {

		ids, err := idgen.GenerateUUIDs(3)
		if err != nil {
			return err
		}

		now := time.Now().UnixMilli()
		_, err = tx.Tenant.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantID(ids[0]).
			SetCode(tenantCode).
			SetName("默认租户").
			SetDescription("默认租户").
			SetStatus(1).
			Save(ctx)
		if err != nil {
			fmt.Println("failed creating schema tenant resources:", err)
			return err
		}

		fmt.Println("tenant:", ids[0], "code:", tenantCode, "name:", "默认租户", "description:", "默认租户", "status:", 1)

		salt, err := passwordgen.GenerateSalt()
		if err != nil {
			return err
		}
		hashedPassword, err := passwordgen.Argon2Hash(password, salt)
		if err != nil {
			return err
		}

		user, err := tx.User.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantCode(tenantCode).
			SetUserID(ids[1]).
			SetPhone("18888888888").
			SetUserName("admin").
			SetPwdHashed(hashedPassword).
			SetPwdSalt(salt).
			SetStatus(1).
			SetNickName("admin").
			SetEmail("admin123@qq.com").
			SetSex(1).
			Save(ctx)
		if err != nil {
			fmt.Println("failed creating schema user esources:", err)
			return err
		}

		fmt.Println("user:", user.UserName, "password:", password)

		return nil
	})

}
