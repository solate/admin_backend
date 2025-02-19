package menu

import (
	"context"

	"admin_backend/app/admin/internal/repository/menurepo"
	"admin_backend/app/admin/internal/repository/permissionrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/ent/generated/menu"
	"admin_backend/pkg/ent/generated/permission"
	"admin_backend/pkg/ent/generated/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuTreeLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	menuRepo       *menurepo.MenuRepo
	permissionRepo *permissionrepo.PermissionRepo
}

// 获取菜单树
func NewGetMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuTreeLogic {
	return &GetMenuTreeLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		menuRepo:       menurepo.NewMenuRepo(svcCtx.DB),
		permissionRepo: permissionrepo.NewPermissionRepo(svcCtx.DB),
	}
}

func (l *GetMenuTreeLogic) GetMenuTree() (resp *types.MenuTreeResp, err error) {
	// 1. 获取租户编码
	tenantCode := contextutil.GetTenantCodeFromCtx(l.ctx)
	// 2. 获取用户角色编码
	roleCode := contextutil.GetRoleCodeFromCtx(l.ctx)

	// 3. 获取角色权限
	pm := l.svcCtx.CasbinManager
	permissions, err := pm.GetRolePermissions(roleCode, tenantCode)
	if err != nil {
		l.Errorf("获取角色权限失败: %v", err)
		return nil, err
	}

	// 4. 获取权限编码
	var resources []string
	for _, permission := range permissions {
		resource := permission[2]
		resources = append(resources, resource)
	}
	where := []predicate.Permission{
		permission.ResourceIn(resources...),
	}
	permissionList, err := l.permissionRepo.List(l.ctx, where)
	if err != nil {
		l.Errorf("获取权限失败: %v", err)
		return nil, err
	}

	// 5. 获取菜单ID
	var menuIDList []string
	for _, v := range permissionList {
		menuIDList = append(menuIDList, v.MenuID)
	}
	menuWhere := []predicate.Menu{
		menu.MenuIDIn(menuIDList...),
	}
	// 6. 获取菜单列表
	menus, err := l.menuRepo.List(l.ctx, menuWhere)
	if err != nil {
		l.Errorf("获取菜单列表失败: %v", err)
		return nil, err
	}

	// 7. 过滤
	var authorizedMenus []*types.MenuTree
	for _, m := range menus {
		authorizedMenus = append(authorizedMenus, &types.MenuTree{
			ParentID:  m.ParentID,
			Name:      m.Name,
			Code:      m.Code,
			Icon:      m.Icon,
			Path:      m.Path,
			Component: m.Component,
			Sort:      m.Sort,
			Status:    m.Status,
			Type:      m.Type,
		})
	}

	// 8. 构建菜单树
	tree := buildMenuTree(authorizedMenus, "") // 从根节点(parentId="")开始构建

	return &types.MenuTreeResp{
		List: tree,
	}, nil
}

// buildMenuTree 构建菜单树
func buildMenuTree(menus []*types.MenuTree, parentID string) []*types.MenuTree {
	var tree []*types.MenuTree
	for _, menu := range menus {
		if menu.ParentID == parentID {
			// 递归获取子菜单
			menu.Children = buildMenuTree(menus, menu.MenuID)
			tree = append(tree, menu)
		}
	}
	return tree
}
