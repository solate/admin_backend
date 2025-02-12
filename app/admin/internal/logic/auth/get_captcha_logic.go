package auth

import (
	"context"
	"fmt"

	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取验证码
func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.CaptchaResp, err error) {
	fmt.Println("============1111=====")

	// 生成验证码
	id, b64s, _, err := l.svcCtx.CaptchaManager.Generate()
	if err != nil {
		l.Error("GetCaptcha Generate err:", err.Error())
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	fmt.Println("=================", id, b64s)

	return &types.CaptchaResp{
		CaptchaId:  id,
		CaptchaUrl: b64s,
	}, nil
}
