package menu

import (
	"context"

	"admin_backend/app/admin/internal/repository/menurepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	menuRepo *menurepo.MenuRepo
}

// 删除菜单
func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		menuRepo: menurepo.NewMenuRepo(svcCtx.DB),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp bool, err error) {
	// 1. 检查菜单是否存在
	menu, err := l.menuRepo.GetByMenuID(l.ctx, req.MenuID)
	if err != nil {
		l.Error("DeleteMenu menuRepo.GetByMenuID err:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("菜单不存在")
		}
		return false, xerr.NewErrCodeMsg(xerr.DbError, "查询菜单失败")
	}

	// 2. 软删除菜单
	_, err = l.menuRepo.Delete(l.ctx, menu.MenuID)
	if err != nil {
		l.Error("DeleteMenu menuRepo.Delete err:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, "删除菜单失败")
	}

	return true, nil
}
