package dict

import (
	"context"

	"admin_backend/app/admin/internal/repository/dictrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictItemsByTypeLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	dictRepo *dictrepo.DictItemRepo
}

// 获取指定类型的字典数据列表
func NewGetDictItemsByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictItemsByTypeLogic {
	return &GetDictItemsByTypeLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		dictRepo: dictrepo.NewDictItemRepo(svcCtx.DB),
	}
}

func (l *GetDictItemsByTypeLogic) GetDictItemsByType(req *types.GetDictItemsByTypeReq) (resp *types.DictItemListResp, err error) {
	// 1. 获取字典数据列表
	list, err := l.dictRepo.GetByTypeCode(l.ctx, req.TypeCode)
	if err != nil {
		l.Error("GetDictItemsByType l.dictRepo.GetByTypeCode err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "获取字典数据列表失败")
	}

	// 2. 构建返回结果
	var dictItems []*types.DictItemInfo
	for _, item := range list {
		dictItems = append(dictItems, &types.DictItemInfo{
			ItemID:      item.ItemID,
			TypeCode:    item.TypeCode,
			Label:       item.Label,
			Value:       item.Value,
			Sort:        item.Sort,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
		})
	}

	return &types.DictItemListResp{
		List: dictItems,
	}, nil
}
