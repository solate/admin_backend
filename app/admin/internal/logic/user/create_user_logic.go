package user

import (
	"context"

	"github.com/solate/admin_backend/app/admin/internal/repository/user_repo"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/utils/idgen"
	"github.com/solate/admin_backend/pkg/utils/passwordgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *user_repo.UserRepo
}

// 创建用户
func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: user_repo.NewUserRepo(svcCtx.DB),
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// 1. 参数验证
	if req.Phone == "" || req.Name == "" {
		return nil, xerr.NewErrCode(xerr.ParamError)
	}

	// 2. 检查手机号是否已存在
	user, err := l.userRepo.GetByPhone(l.ctx, req.Phone)
	if err != nil {
		l.Error("CreateUser User.Exist err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "查询用户失败")
	}
	if user.Phone == req.Phone {
		return nil, xerr.NewErrMsg("手机号已存在")
	}

	// 3. 生成密码盐和加密密码
	salt, err := passwordgen.GenerateSalt()
	if err != nil {
		l.Error("CreateUser GenerateSalt err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成密码盐失败")
	}
	hashedPassword, err := passwordgen.Argon2Hash(req.PwdHashed, salt)
	if err != nil {
		l.Error("CreateUser Argon2Hash err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成密码加盐失败")
	}

	userID, err := idgen.GenerateUUID()
	if err != nil {
		l.Error("CreateUser GenerateUUID err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成用户ID失败")
	}

	// 4. 创建用户
	newUser := &generated.User{
		UserID:    userID,
		Phone:     req.Phone,
		UserName:  req.Name,
		PwdHashed: hashedPassword,
		PwdSalt:   string(salt),
		Status:    1, // 默认启用
		NickName:  req.Name,
		Email:     req.Email,
		Sex:       req.Sex,
	}
	user, err = l.userRepo.Create(l.ctx, newUser)
	if err != nil {
		l.Error("CreateUser Create err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "创建用户失败")
	}

	// 5. 返回结果
	return &types.CreateUserResp{
		UserID: user.UserID,
	}, nil
}
