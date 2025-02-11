package auth

import (
	"context"

	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) GetCaptcha() (resp *types.CaptchaResp, err error) {
	// 生成验证码
	id, b64s, _, err := l.svcCtx.CaptchaManager.Generate()
	if err != nil {
		l.Error("GetCaptcha Generate err:", err.Error())
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	return &types.CaptchaResp{
		CaptchaId:  id,
		CaptchaUrl: b64s,
	}, nil
}
