package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"
	"admin_backend/pkg/utils/passwordgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 创建用户
func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {

	// 1. 检查用户名是否已存在
	user, err := l.userRepo.GetByUserName(l.ctx, req.UserName)
	if err != nil && !generated.IsNotFound(err) {
		l.Error("GetUser l.userRepo.GetByUserName err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "数据库查询错误")
	}
	if user != nil {
		return nil, xerr.NewErrCodeMsg(xerr.DbRecordExist, "用户名已存在")
	}

	// 检查电话号是否已存在
	user, err = l.userRepo.GetByPhone(l.ctx, req.Phone)
	if err != nil && !generated.IsNotFound(err) {
		l.Error("GetUser l.userRepo.GetByPhone err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "数据库查询错误")
	}
	if user != nil {
		return nil, xerr.NewErrCodeMsg(xerr.DbRecordExist, "手机号已存在")
	}

	// 3. 生成密码盐和加密密码
	salt, err := passwordgen.GenerateSalt()
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成密码盐失败")
	}
	hashedPassword, err := passwordgen.Argon2Hash(req.Password, salt)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "密码加密失败")
	}

	userID, err := idgen.GenerateUUID()
	if err != nil {
		l.Error("CreateUser GenerateUUID err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成用户ID失败")
	}

	if req.Status == 0 {
		req.Status = 1 // 默认启用
	}

	// 4. 创建用户
	newUser := &generated.User{
		TenantCode: contextutil.GetTenantCodeFromCtx(l.ctx),
		UserID:     userID,
		Phone:      req.Phone,
		UserName:   req.UserName,
		Name:       req.Name,
		Email:      req.Email,
		Sex:        req.Sex,
		PwdHashed:  hashedPassword,
		PwdSalt:    salt,
		Status:     req.Status, // 默认启用
	}
	user, err = l.userRepo.Create(l.ctx, newUser)
	if err != nil {
		l.Error("CreateUser Create err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "创建用户失败")
	}

	// 5. 返回结果
	return &types.CreateUserResp{
		UserID: newUser.UserID,
	}, nil
}
