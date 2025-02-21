package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 更新用户
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp bool, err error) {
	// 1. 检查用户是否存在
	user, err := l.userRepo.GetByUserID(l.ctx, req.UserID)
	if err != nil {
		l.Error("UpdateUser userRepo.GetByUserID err:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCodeMsg(xerr.DbError, "查询用户失败")
	}

	if user.UserName == "admin" && req.Status != 1 {
		return false, xerr.NewErrCodeMsg(xerr.ParamError, "管理员账号不能禁用")
	}

	// 2. 更新用户信息
	user.Name = req.Name
	user.Email = req.Email
	user.Status = req.Status
	user.Sex = req.Sex
	user.Avatar = req.Avatar

	_, err = l.userRepo.Update(l.ctx, user)
	if err != nil {
		l.Error("UpdateUser Update err:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, "更新用户失败")
	}

	// 3. 更新用户角色
	tenantCode := contextutil.GetTenantCodeFromCtx(l.ctx)
	err = l.UpdateUserRoles(req.UserID, req.RoleCodeList, tenantCode)
	if err != nil {
		return false, xerr.NewErrCodeMsg(xerr.ServerError, "更新用户角色失败")
	}

	return true, nil
}

// UpdateUserRoles 更新用户角色
func (l *UpdateUserLogic) UpdateUserRoles(userID string, roleCodeList []string, tenantCode string) error {
	pm := l.svcCtx.CasbinManager

	// 清空用户原有角色
	roles, err := pm.GetRolesForUser(userID, tenantCode)
	if err != nil {
		l.Errorf("获取用户角色失败: %v", err)
		return err
	}
	for _, role := range roles {
		err = pm.RemoveRoleForUser(userID, role, tenantCode)
		if err != nil {
			l.Errorf("删除用户角色失败: %v", err)
			return err
		}
	}

	// 用户添加新角色
	for _, code := range roleCodeList {
		err = pm.AddRoleForUser(userID, code, tenantCode)
		if err != nil {
			l.Errorf("添加用户角色失败: %v", err)
			return err
		}
	}

	return nil
}
