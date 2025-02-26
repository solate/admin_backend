package initialize

import (
	"context"
	"fmt"
	"time"

	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"
)

// InitDictTypes 初始化字典类型数据
func InitDictTypes(ctx context.Context, tx *generated.Tx, tenantCode string) error {
	// 字典类型数据
	dictTypes := []struct {
		name        string
		code        string
		description string
	}{
		{name: "性别", code: "sex", description: "用户性别"},
		{name: "状态", code: "status", description: "数据状态"},
		{name: "是否", code: "bool", description: "是否类型"},
	}

	for _, dictType := range dictTypes {
		typeID, err := idgen.GenerateUUID()
		if err != nil {
			return fmt.Errorf("生成字典类型ID失败: %v", err)
		}

		now := time.Now().UnixMilli()
		_, err = tx.DictType.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantCode(tenantCode).
			SetTypeID(typeID).
			SetCode(dictType.code).
			SetName(dictType.name).
			SetDescription(dictType.description).
			SetStatus(1).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("创建字典类型失败: %v", err)
		}
	}

	return nil
}

// InitDictItems 初始化字典项数据
func InitDictItems(ctx context.Context, tx *generated.Tx, tenantCode string) error {
	// 字典项数据
	dictItems := []struct {
		typeCode    string
		label       string
		value       string
		description string
		sort        int
	}{
		{typeCode: "sex", label: "男", value: "1", description: "男性", sort: 1},
		{typeCode: "sex", label: "女", value: "2", description: "女性", sort: 2},
		{typeCode: "status", label: "启用", value: "1", description: "启用状态", sort: 1},
		{typeCode: "status", label: "禁用", value: "2", description: "禁用状态", sort: 2},
		{typeCode: "yes_no", label: "是", value: "1", description: "是", sort: 1},
		{typeCode: "yes_no", label: "否", value: "2", description: "否", sort: 2},
	}

	for _, item := range dictItems {
		itemID, err := idgen.GenerateUUID()
		if err != nil {
			return fmt.Errorf("生成字典项ID失败: %v", err)
		}

		now := time.Now().UnixMilli()
		_, err = tx.DictItem.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantCode(tenantCode).
			SetItemID(itemID).
			SetTypeCode(item.typeCode).
			SetLabel(item.label).
			SetValue(item.value).
			SetDescription(item.description).
			SetSort(item.sort).
			SetStatus(1).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("创建字典项失败: %v", err)
		}
	}

	return nil
}
