package logic

import (
	"context"
	"mixbench/service/v1/luckdraw/rpc/internal/svc"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentifyLogic {
	return &IdentifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** Identify 抽奖凭证校验，验证用户身份，颁发jwt
 * 传入的password可能是line或者Phone，或者passport
 * 先匹配passport，然后再去匹配 email ,再去匹配 phone
 */
func (l *IdentifyLogic) Identify(in *luckdraw_pb.IdentifyReq) (*luckdraw_pb.IdentifyResp, error) {
	mdUser, err := SearchUser(l.ctx, l.svcCtx.UserModel, in.Passport)
	if err != nil {
		return nil, err
	}
	//找不到
	if mdUser == nil {
		return &luckdraw_pb.IdentifyResp{
			Status: false,
		}, nil
	}
	//无效用户
	if mdUser.Status != 1 {
		return &luckdraw_pb.IdentifyResp{
			Status: false,
		}, nil
	}
	generate := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	resp, err := generate.GenerateToken(&luckdraw_pb.GenerateTokenReq{
		JwtData: &luckdraw_pb.JwtData{
			UserId:   mdUser.Id,
			Ip:       in.JwtData.Ip,
			Platform: in.JwtData.Platform,
		},
	})
	if err != nil {
		return nil, err
	}
	return &luckdraw_pb.IdentifyResp{
		Status:   true,
		JwtToken: resp.JwtToken,
	}, nil
}
