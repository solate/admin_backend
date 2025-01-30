package user

import (
	"context"
	"time"

	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
	"github.com/solate/admin_backend/pkg/utils/passwordgen"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// 1. 参数验证
	if req.Phone == "" || req.Name == "" {
		return nil, xerr.NewErrCode(xerr.ParamError)
	}

	// 2. 检查手机号是否已存在
	exist, err := l.svcCtx.Orm.User.Query().Where(user.Phone(req.Phone)).Exist(l.ctx)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	if exist {
		return nil, xerr.NewErrMsg("手机号已存在")
	}

	// 3. 生成密码盐和加密密码
	salt, err := passwordgen.GenerateSalt()
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}
	hashedPassword, err := passwordgen.Argon2Hash(req.PwdHashed, salt)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	// 4. 创建用户
	now := time.Now().UnixMilli()
	user, err := l.svcCtx.Orm.User.Create().
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetPhone(req.Phone).
		SetUserName(req.Name).
		SetPassword(hashedPassword).
		SetSalt(string(salt)).
		SetStatus(1). // 默认启用
		SetNickName(req.Name).
		SetEmail(req.Email).
		SetSex(req.Gender).
		Save(l.ctx)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 5. 返回结果
	return &types.CreateUserResp{
		UserID: int(user.UserID),
	}, nil
}
