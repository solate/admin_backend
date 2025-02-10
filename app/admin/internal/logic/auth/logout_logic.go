package auth

import (
	"context"

	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/ent/generated/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp bool, err error) {

	// 1. 查找用户
	user, err := l.svcCtx.DB.User.Query().
		Where(user.UserID(req.UserID)).
		Where(user.DeletedAt(0)).
		Only(l.ctx)

	if err != nil {
		l.Error("Logout User.Query Error:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCodeMsg(xerr.DbError, err.Error())
	}

	// 更新user
	_, err = l.svcCtx.DB.User.UpdateOne(user).
		SetToken("").
		Save(l.ctx)
	if err != nil {

		l.Error("Logout User.UpdateOne Error:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, err.Error())
	}

	// 记录登出日志
	_, err = l.svcCtx.DB.SystemLog.Create().
		SetUserID(user.UserID).
		Save(l.ctx)

	return true, nil
}

// func (l *LogoutLogic) addLog(uid int) error {
// 	logDao := userRepo.NewSystemLogDao(l.ctx, l.svcCtx)
// 	_, err := logDao.Create(userStatus.SystemLogin, "登录了系统", userStatus.Login, uid)
// 	if err != nil {
// 		l.Info(fmt.Sprintf("%s - %s: %d", userStatus.SystemLogin, userStatus.Login, uid), err)
// 		return err
// 	}
// 	return nil
// }
