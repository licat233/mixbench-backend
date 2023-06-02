package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PwdEncryptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPwdEncryptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PwdEncryptionLogic {
	return &PwdEncryptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PwdEncryptionLogic) PwdEncryption(req *types.PwdEncryptionReq) (any, error) {
	err := xcrypt.PasswordValidator(req.Password)
	if err != nil {
		return nil, errorx.New(err.Error())
	}
	hashPwd, err := xcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errorx.New(err.Error())
	}
	return respx.NewData(hashPwd), nil
}
