package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckyDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyDetailsLogic {
	return &LuckyDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyDetailsLogic) LuckyDetails(req *types.LuckyDetailsReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return respx.ListResp("", []any{}, 0, req.PageSize, req.Page), nil
	}
	in := &luckydraw_pb.LuckyDetailsReq{
		ListReq: &luckydraw_pb.ListReq{PageSize: req.PageSize, Page: req.Page, Keyword: req.Keyword},
		Lucky: &luckydraw_pb.Lucky{
			Id:          0,
			Uuid:        "",
			UserId:      req.UserId,
			PrizePoolId: req.PrizePoolId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
		UserName:     req.UserName,
		ProjectId:    req.ProjectId,
		ActivityId:   req.ActivityId,
		ActivityName: req.ActivityName,
		PrizeId:      req.PeizeId,
		PrizeName:    req.PrizeName,
		PrizeGradeId: req.PeizeGradeId,
	}
	resp, err := l.svcCtx.LuckydrawRpc.GetLuckyDetailsList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	data := []*types.LuckyDetail{}
	for _, pbdetail := range resp.Data {
		apidetail := &types.LuckyDetail{
			Lucky:      *dataconv.PbLuckyToApiLucky(pbdetail.Lucky),
			User:       *dataconv.PbUserToApiUser(pbdetail.User),
			PrizePool:  *dataconv.PbPrizePoolToApiPrizePool(pbdetail.PrizePool),
			Prize:      *dataconv.PbPrizeToApiPrize(pbdetail.Prize),
			PrizeGrade: *dataconv.PbPrizeGradeToApiPrizeGrade(pbdetail.PrizeGrade),
			Activity:   *dataconv.PbActivityToApiActivity(pbdetail.Activity),
			Project:    *dataconv.PbProjectToApiProject(pbdetail.Project),
		}
		data = append(data, apidetail)
	}
	return respx.ListResp("", data, resp.Total, req.PageSize, req.Page), nil
}
