package initialize

import (
	"context"
	"fmt"
	"time"

	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"
)

// InitCategories 初始化商品分类数据
func InitCategories(ctx context.Context, tx *generated.Tx, tenantCode string) error {
	// 一级分类数据
	level1Categories := []struct {
		name        string
		description string
		icon        string
		keywords    string
		unit        string
	}{
		{name: "电子产品", description: "各类电子设备和数码产品", icon: "electronics", keywords: "手机,电脑,数码", unit: "件"},
		{name: "服装服饰", description: "男女服装、鞋帽、箱包", icon: "clothing", keywords: "衣服,鞋子,包包", unit: "件"},
		{name: "家居用品", description: "家具、家纺、生活日用", icon: "home", keywords: "家具,床上用品,收纳", unit: "件"},
		{name: "食品生鲜", description: "食品、饮料、生鲜", icon: "food", keywords: "零食,饮料,生鲜", unit: "件"},
		{name: "美妆个护", description: "化妆品、个人护理", icon: "beauty", keywords: "化妆品,护肤,美妆", unit: "件"},
	}

	// 二级分类数据映射
	level2Categories := map[string][]string{
		"电子产品": {"手机", "电脑", "相机", "智能设备", "配件"},
		"服装服饰": {"男装", "女装", "童装", "鞋靴", "箱包"},
		"家居用品": {"家具", "家纺", "收纳", "灯具", "家装"},
		"食品生鲜": {"零食", "饮料", "生鲜", "粮油", "调味品"},
		"美妆个护": {"面部护理", "彩妆", "香水", "个人护理", "美妆工具"},
	}

	// 三级分类数据映射
	level3Categories := map[string][]string{
		"手机":   {"智能手机", "功能机", "手机壳", "充电器", "手机支架"},
		"电脑":   {"笔记本", "台式机", "平板电脑", "显示器", "电脑配件"},
		"男装":   {"T恤", "衬衫", "夹克", "西装", "休闲裤"},
		"女装":   {"连衣裙", "上衣", "裤子", "外套", "半身裙"},
		"家具":   {"沙发", "床", "桌椅", "柜子", "架子"},
		"零食":   {"饼干", "糖果", "坚果", "膨化食品", "肉干"},
		"面部护理": {"洁面", "爽肤水", "精华", "面霜", "面膜"},
	}

	// 创建一级分类
	for _, cat1 := range level1Categories {
		categoryID, err := idgen.GenerateUUID()
		if err != nil {
			return fmt.Errorf("生成一级分类ID失败: %v", err)
		}

		now := time.Now().UnixMilli()
		level1, err := tx.Category.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantCode(tenantCode).
			SetCategoryID(categoryID).
			SetName(cat1.name).
			SetDescription(cat1.description).
			SetIcon(cat1.icon).
			SetKeywords(cat1.keywords).
			SetUnit(cat1.unit).
			SetLevel(1).
			SetStatus(1).
			SetSort(0).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("创建一级分类失败: %v", err)
		}

		// 创建二级分类
		if level2Names, ok := level2Categories[cat1.name]; ok {
			for _, cat2Name := range level2Names {
				categoryID, err := idgen.GenerateUUID()
				if err != nil {
					return fmt.Errorf("生成二级分类ID失败: %v", err)
				}

				level2, err := tx.Category.Create().
					SetCreatedAt(now).
					SetUpdatedAt(now).
					SetTenantCode(tenantCode).
					SetCategoryID(categoryID).
					SetName(cat2Name).
					SetParentID(level1.CategoryID).
					SetLevel(2).
					SetStatus(1).
					SetSort(0).
					Save(ctx)

				if err != nil {
					return fmt.Errorf("创建二级分类失败: %v", err)
				}

				// 创建三级分类
				if level3Names, ok := level3Categories[cat2Name]; ok {
					for _, cat3Name := range level3Names {
						categoryID, err := idgen.GenerateUUID()
						if err != nil {
							return fmt.Errorf("生成三级分类ID失败: %v", err)
						}

						_, err = tx.Category.Create().
							SetCreatedAt(now).
							SetUpdatedAt(now).
							SetTenantCode(tenantCode).
							SetCategoryID(categoryID).
							SetName(cat3Name).
							SetParentID(level2.CategoryID).
							SetLevel(3).
							SetStatus(1).
							SetSort(0).
							Save(ctx)

						if err != nil {
							return fmt.Errorf("创建三级分类失败: %v", err)
						}
					}
				}
			}
		}
	}

	return nil
}
