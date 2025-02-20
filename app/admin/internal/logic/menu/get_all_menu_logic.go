package menu

import (
	"context"

	"admin_backend/app/admin/internal/repository/menurepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllMenuLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	menuRepo *menurepo.MenuRepo
}

// 获取所有菜单
func NewGetAllMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllMenuLogic {
	return &GetAllMenuLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		menuRepo: menurepo.NewMenuRepo(svcCtx.DB),
	}
}

func (l *GetAllMenuLogic) GetAllMenu() (resp *types.MenuTreeResp, err error) {
	// 获取所有菜单列表
	menuList, err := l.menuRepo.List(l.ctx, nil)
	if err != nil {
		l.Errorf("获取菜单列表失败: %v", err)
		return nil, xerr.NewErrMsg("获取菜单列表失败")
	}

	// 构建响应数据
	var authorizedMenus []*types.MenuTree
	for _, menuInfo := range menuList {
		authorizedMenus = append(authorizedMenus, &types.MenuTree{
			MenuID:    menuInfo.MenuID,
			Code:      menuInfo.Code,
			ParentID:  menuInfo.ParentID,
			Name:      menuInfo.Name,
			Path:      menuInfo.Path,
			Component: menuInfo.Component,
			Redirect:  menuInfo.Redirect,
			Icon:      menuInfo.Icon,
			Sort:      menuInfo.Sort,
			Type:      menuInfo.Type,
			Status:    menuInfo.Status,
		})
	}

	// 构建菜单树
	tree := buildMenuTree(authorizedMenus, "") // 从根节点(parentId="")开始构建

	return &types.MenuTreeResp{
		List: tree,
	}, nil
}
