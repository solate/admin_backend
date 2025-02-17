package permission

import (
	"context"

	"admin_backend/app/admin/internal/repository/permissionrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/ent/generated/permission"
	"admin_backend/pkg/ent/generated/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionsLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	permissionRepo *permissionrepo.PermissionRepo
}

// 获取角色权限列表
func NewGetRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionsLogic {
	return &GetRolePermissionsLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		permissionRepo: permissionrepo.NewPermissionRepo(svcCtx.DB),
	}
}

func (l *GetRolePermissionsLogic) GetRolePermissions(req *types.GetRolePermissionsReq) (resp *types.GetRolePermissionsResp, err error) {
	tenantCode := contextutil.GetTenantCodeFromCtx(l.ctx)
	// 获取权限管理器
	pm := l.svcCtx.CasbinManager

	// 获取角色在该租户下的所有权限编码
	casbinPermissionList, err := pm.GetRolePermissions(req.RoleCode, tenantCode)
	if err != nil {
		l.Errorf("获取角色权限列表失败: %v", err)
		return nil, err
	}

	var permissionCodeList []string
	// [[admin default permissionCode GET api]]
	for _, inner := range casbinPermissionList {
		permissionCodeList = append(permissionCodeList, inner[2])
	}

	// 如果角色没有权限，则直接返回空列表
	if len(permissionCodeList) == 0 {
		return &types.GetRolePermissionsResp{
			List: make([]*types.PermissionInfo, 0),
		}, nil
	}

	// 根据权限编码获取完整的权限信息
	where := []predicate.Permission{
		permission.CodeIn(permissionCodeList...),
	}
	permissionList, err := l.permissionRepo.List(l.ctx, where)
	if err != nil {
		l.Errorf("获取权限信息失败, code: %v, err: %v", permissionCodeList, err)
		return nil, err
	}

	list := make([]*types.PermissionInfo, 0, len(permissionCodeList))
	for _, permissionInfo := range permissionList {
		list = append(list, &types.PermissionInfo{
			TenantCode:   permissionInfo.TenantCode,
			PermissionID: permissionInfo.PermissionID,
			Name:         permissionInfo.Name,
			Code:         permissionInfo.Code,
			Type:         permissionInfo.Type,
			Resource:     permissionInfo.Resource,
			Action:       permissionInfo.Action,
			ParentID:     permissionInfo.ParentID,
			Description:  permissionInfo.Description,
			Status:       permissionInfo.Status,
			MenuID:       permissionInfo.MenuID,
			CreatedAt:    permissionInfo.CreatedAt,
		})
	}

	resp = &types.GetRolePermissionsResp{
		List: list,
	}

	return
}
