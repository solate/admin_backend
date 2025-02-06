package auth

import (
	"context"
	"time"

	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
	"github.com/solate/admin_backend/pkg/utils/idgen"
	"github.com/solate/admin_backend/pkg/utils/passwordgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 1. 检查手机号是否已存在
	exist, err := l.svcCtx.Orm.User.Query().Where(user.Phone(req.Phone)).Exist(l.ctx)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	if exist {
		return nil, xerr.NewErrMsg("手机号已存在")
	}

	// 2. 生成密码盐和加密密码
	salt, err := passwordgen.GenerateSalt()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}
	hashedPassword, err := passwordgen.Argon2Hash(req.Password, salt)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	// 3. 生成用户ID
	userID, err := idgen.GenerateUUID()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	// 4. 创建用户
	now := time.Now().UnixMilli()
	user, err := l.svcCtx.Orm.User.Create().
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetUserID(userID).
		SetPhone(req.Phone).
		SetUserName(req.UserName).
		SetPassword(hashedPassword).
		SetSalt(string(salt)).
		SetStatus(1). // 默认启用
		SetNickName(req.NickName).
		SetEmail(req.Email).
		SetSex(req.Sex).
		Save(l.ctx)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 5. 返回结果
	return &types.RegisterResp{
		UserID: user.UserID,
	}, nil
}
