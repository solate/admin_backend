package user

import (
	"context"
	"time"

	"admin_backend/app/admin/internal/repository/login_log_repo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/context_util"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated/loginlog"
	"admin_backend/pkg/ent/generated/predicate"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLoginLogLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	loginLogRepo *login_log_repo.LoginLogRepo
}

// 查询登录记录
func NewListLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLoginLogLogic {
	return &ListLoginLogLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		loginLogRepo: login_log_repo.NewLoginLogRepo(svcCtx.DB),
	}
}

func (l *ListLoginLogLogic) ListLoginLog(req *types.LoginLogListReq) (resp *types.LoginLogListResp, err error) {
	// 获取租户编码
	tenantCode, err := context_util.GetTenantCodeFromCtx(l.ctx)
	if err != nil {
		l.Error("ListLoginLog context_util.GetTenantCodeFromCtx err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "获取租户编码失败")
	}

	// 构建查询条件
	where := []predicate.LoginLog{
		loginlog.TenantCode(tenantCode),
	}

	// 添加过滤条件
	if req.UserName != "" {
		where = append(where, loginlog.UserNameContains(req.UserName))
	}
	if req.IP != "" {
		where = append(where, loginlog.IPContains(req.IP))
	}

	// 时间范围查询
	if req.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			where = append(where, loginlog.LoginTimeGTE(startTime.UnixMilli()))
		}
	}
	if req.EndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			where = append(where, loginlog.LoginTimeLTE(endTime.UnixMilli()))
		}
	}

	// 查询列表
	list, total, err := l.loginLogRepo.PageList(l.ctx, req.Current, req.PageSize, where)
	if err != nil {
		l.Error("ListLoginLog loginLogRepo.PageList err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "查询登录日志失败")
	}

	// 构建返回结果
	loginLogs := make([]*types.LoginLogInfo, 0)
	for _, log := range list {
		loginLogs = append(loginLogs, &types.LoginLogInfo{
			ID:        log.LogID,
			UserID:    log.UserID,
			UserName:  log.UserName,
			IP:        log.IP,
			UserAgent: log.UserAgent,
			Message:   log.Message,
			CreatedAt: time.UnixMilli(log.CreatedAt).Format("2006-01-02 15:04:05"),
		})
	}

	return &types.LoginLogListResp{
		List: loginLogs,
		Page: &types.PageResponse{
			Total:           total,
			PageSize:        len(list),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}
