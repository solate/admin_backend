package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/casbinrulerepo"
	"admin_backend/app/admin/internal/repository/rolerepo"
	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/role"
	"admin_backend/pkg/ent/generated/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	userRepo   *userrepo.UserRepo
	roleRepo   *rolerepo.RoleRepo
	casbinRepo *casbinrulerepo.CasbinRuleRepo
}

// 获取用户列表
func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		userRepo:   userrepo.NewUserRepo(svcCtx.DB),
		roleRepo:   rolerepo.NewRoleRepo(svcCtx.DB),
		casbinRepo: casbinrulerepo.NewCasbinRuleRepo(svcCtx.DB),
	}
}

func (l *ListUserLogic) ListUser(req *types.UserListReq) (resp *types.UserListResp, err error) {

	// 1. 构建查询条件
	where := []predicate.User{}

	if req.Name != "" {
		where = append(where, user.NameContains(req.Name))
	}

	if req.Phone != "" {
		where = append(where, user.Phone(req.Phone))
	}

	if req.Status != 0 {
		where = append(where, user.StatusEQ(req.Status))
	}

	list, total, err := l.userRepo.PageList(l.ctx, req.Current, req.PageSize, where)
	if err != nil {
		l.Error("ListUser Logic PageList err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "list user page err.")
	}

	// 获取casbin 中 具有的角色
	// sql := "SELECT * FROM casbin_rule WHERE ptype = 'g' AND v0 IN (?);"
	var userIDs []any
	for _, user := range list {
		userIDs = append(userIDs, user.UserID)
	}
	rules, err := l.casbinRepo.QueryBySQL(l.ctx, userIDs...)
	if err != nil {
		l.Error("ListUser Logic QueryBySQL err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "list user page err.")
	}

	roleCodeMap := make(map[string]bool)
	for _, rule := range rules {
		roleCodeMap[rule.V1] = true
	}
	var roleCodes []string
	for code := range roleCodeMap {
		roleCodes = append(roleCodes, code)
	}

	rolesWhere := []predicate.Role{
		role.CodeIn(roleCodes...),
	}
	roles, err := l.roleRepo.List(l.ctx, rolesWhere)
	if err != nil {
		l.Error("ListUser Logic List err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "list user page err.")
	}

	// 4. 构建返回结果
	// userList := make([]*types.UserInfo, 0)
	// for _, user := range list {
	// 	userList = append(userList, &types.UserInfo{
	// 		UserID:   user.UserID,
	// 		UserName: user.UserName,
	// 		Phone:    user.Phone,
	// 		Email:    user.Email,
	// 		Status:   user.Status,
	// 		RoleList: []*types.RoleListInfo{},
	// 	})
	// }
	userList := constructResponse(list, roles, rules)

	return &types.UserListResp{
		List: userList,
		Page: &types.PageResponse{
			Total:           total,
			PageSize:        len(list),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}

// 构造返回值
func constructResponse(userList []*generated.User, roleList []*generated.Role, casbinRules []*generated.CasbinRule) []*types.UserInfo {

	ruleMap := make(map[string][]string) // map[userID][]roleCode
	for _, rule := range casbinRules {
		ruleMap[rule.V0] = append(ruleMap[rule.V0], rule.V1)
	}

	roleMap := make(map[string]*generated.Role)
	for _, role := range roleList {
		roleMap[role.Code] = role
	}

	resp := make([]*types.UserInfo, 0)
	for _, user := range userList {
		roleCodeList := ruleMap[user.UserID]

		// 获取角色信息
		tmpRoleList := make([]*types.RoleListInfo, 0)
		for _, roleCode := range roleCodeList {
			role := roleMap[roleCode]
			tmpRoleList = append(tmpRoleList, &types.RoleListInfo{
				RoleID: role.RoleID,
				Code:   role.Code,
				Name:   role.Name,
				Sort:   role.Sort,
			})
		}

		// 构建返回值
		userInfo := &types.UserInfo{
			UserID:   user.UserID,
			UserName: user.UserName,
			Name:     user.Name,
			Phone:    user.Phone,
			Email:    user.Email,
			Status:   user.Status,
			RoleList: tmpRoleList,
		}

		resp = append(resp, userInfo)
	}

	return resp
}
