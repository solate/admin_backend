package menu

import (
	"context"

	"admin_backend/app/admin/internal/repository/menurepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	menuRepo *menurepo.MenuRepo
}

// 创建菜单
func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		menuRepo: menurepo.NewMenuRepo(svcCtx.DB),
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuReq) (resp *types.CreateMenuResp, err error) {
	// 1. 检查菜单编码是否已存在
	menu, err := l.menuRepo.GetByMenuCode(l.ctx, req.Code)
	if err != nil && !generated.IsNotFound(err) {
		l.Error("GetMenu l.menuRepo.GetByMenuCode err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "数据库查询错误")
	}
	if menu != nil {
		return nil, xerr.NewErrCodeMsg(xerr.DbRecordExist, "菜单编码已存在")
	}

	// 2. 生成菜单ID
	menuID, err := idgen.GenerateUUID()
	if err != nil {
		l.Error("CreateMenu GenerateUUID err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成菜单ID失败")
	}

	// 3. 设置默认值
	if req.Status == 0 {
		req.Status = 1 // 默认启用
	}

	// 4. 创建菜单
	newMenu := &generated.Menu{
		TenantCode: contextutil.GetTenantCodeFromCtx(l.ctx),
		MenuID:     menuID,
		Code:       req.Code,
		ParentID:   req.ParentID,
		Name:       req.Name,
		Path:       req.Path,
		Component:  req.Component,
		Redirect:   req.Redirect,
		Icon:       req.Icon,
		Sort:       req.Sort,
		Type:       req.Type,
		Status:     req.Status,
	}

	menu, err = l.menuRepo.Create(l.ctx, newMenu)
	if err != nil {
		l.Error("CreateMenu Create err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "创建菜单失败")
	}

	// 5. 返回结果
	return &types.CreateMenuResp{
		MenuID: newMenu.MenuID,
	}, nil
}
