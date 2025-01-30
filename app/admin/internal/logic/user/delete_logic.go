package user

import (
	"context"
	"time"

	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// 1. 检查用户是否存在
	user, err := l.svcCtx.Orm.User.Query().Where(user.ID(req.UserID)).Only(l.ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, xerr.NewErrMsg("用户不存在")
		}
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 2. 软删除用户
	_, err = l.svcCtx.Orm.User.UpdateOne(user).
		SetDeletedAt(time.Now().UnixMilli()).
		Save(l.ctx)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	return &types.DeleteUserResp{
		Success: true,
	}, nil
}
