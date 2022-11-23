package logic

import (
	"context"
	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

/** SearchUser 搜索用户
 * 返回的error均为 errorx 错误
 * 先匹配passport，然后再去匹配 email ,再去匹配 phone
 */
func SearchUser(ctx context.Context, UserModel model.UserModel, passport string) (*model.User, error) {
	mdUser, err := UserModel.FindByPasspord(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return mdUser, nil
	}

	mdUser, err = UserModel.FindByEmail(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return mdUser, nil
	}

	mdUser, err = UserModel.FindByPhone(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		return mdUser, nil
	}

	if err == model.ErrNotFound {
		return nil, nil
	}
	return mdUser, nil
}

func convertActivity(mdActivity *model.Activity) *luckdraw_pb.Activity {
	if mdActivity == nil {
		return nil
	}
	return &luckdraw_pb.Activity{
		Id:        mdActivity.Id,
		Uuid:      mdActivity.Uuid,
		Name:      mdActivity.Name,
		ProjectId: mdActivity.ProjectId,
		Status:    utils.Int64ToBool(mdActivity.Status),
		Pattern:   mdActivity.Pattern,
		MaxCount:  mdActivity.MaxCount,
		MaxPrizes: mdActivity.MaxPrizes,
		StartTime: mdActivity.StartTime.Unix(),
		EndTime:   mdActivity.EndTime.Unix(),
	}
}

func convertPrize(mdPrize *model.Prize) *luckdraw_pb.Prize {
	if mdPrize == nil {
		return nil
	}
	return &luckdraw_pb.Prize{
		Id:     mdPrize.Id,
		Uuid:   mdPrize.Uuid,
		Name:   mdPrize.Name,
		Image:  mdPrize.Image,
		Price:  mdPrize.Price,
		Status: mdPrize.Status,
	}
}

func convertPrizesPool(mdPrizesPool *model.PrizesPool) *luckdraw_pb.PrizesPool {
	if mdPrizesPool == nil {
		return nil
	}
	return &luckdraw_pb.PrizesPool{
		Id:         mdPrizesPool.Id,
		Uuid:       mdPrizesPool.Uuid,
		ActivityId: mdPrizesPool.ActivityId,
		PrizeId:    mdPrizesPool.PrizeId,
		LuckValue:  mdPrizesPool.LuckValue,
		Status:     mdPrizesPool.Status,
		GradeId:    mdPrizesPool.GradeId,
		Number:     mdPrizesPool.Total,
		Remainder:  mdPrizesPool.Remainder,
	}
}

func convertUser(mdUser *model.User) *luckdraw_pb.User {
	if mdUser == nil {
		return nil
	}
	return &luckdraw_pb.User{
		Id:       mdUser.Id,
		Uuid:     mdUser.Uuid,
		Passport: mdUser.Passport,
		Email:    mdUser.Email,
		Phone:    mdUser.Phone,
		Name:     mdUser.Name,
		Address:  mdUser.Address,
		Score:    mdUser.Score,
		Status:   mdUser.Status,
	}
}

func convertLucker(mdLucker *model.Lucker) *luckdraw_pb.Lucker {
	if mdLucker == nil {
		return nil
	}
	return &luckdraw_pb.Lucker{
		Id:           mdLucker.Id,
		Uuid:         mdLucker.Uuid,
		UserId:       mdLucker.UserId,
		PrizesPoolId: mdLucker.PrizesPoolId,
		Status:       mdLucker.Status,
		LuckDate:     mdLucker.LuckDate.Unix(),
		Ip:           mdLucker.Ip,
	}
}

func convertPrizeGrade(in *model.PrizeGrade) *luckdraw_pb.PrizeGrade {
	if in == nil {
		return nil
	}
	return &luckdraw_pb.PrizeGrade{
		Id:    in.Id,
		Name:  in.Name,
		Iswin: in.IsWin,
	}
}

func convertPrizes(mPrizes []*model.Prize) (pbPrizes []*luckdraw_pb.Prize) {
	pbPrizes = make([]*luckdraw_pb.Prize, len(mPrizes))
	if len(mPrizes) == 0 {
		return
	}
	for index, mPrize := range mPrizes {
		lprize := convertPrize(mPrize)
		pbPrizes[index] = lprize
	}
	return
}

func convertLuckerList(mdLuckerList []*model.Lucker) (pbLuckerlist []*luckdraw_pb.Lucker) {
	pbLuckerlist = make([]*luckdraw_pb.Lucker, len(mdLuckerList))
	for index, mdLucker := range mdLuckerList {
		pbLucker := convertLucker(mdLucker)
		pbLuckerlist[index] = pbLucker
	}
	return
}

func convertPrizesPoolList(mdPrizesPoolList []*model.PrizesPool) (pbPrizesPoolList []*luckdraw_pb.PrizesPool) {
	pbPrizesPoolList = make([]*luckdraw_pb.PrizesPool, len(mdPrizesPoolList))
	for index, mdPrizesPool := range mdPrizesPoolList {
		pbPrizesPool := convertPrizesPool(mdPrizesPool)
		pbPrizesPoolList[index] = pbPrizesPool
	}
	return
}

func convertLuckerDetail(mdLuckerDetail *model.LuckerDetailsInfo) *luckdraw_pb.LuckerDetail {
	if mdLuckerDetail == nil {
		return nil
	}
	return &luckdraw_pb.LuckerDetail{
		Lucker:     convertLucker(mdLuckerDetail.Lucker),
		PrizesPool: nil, //不能对外公开
		Activity:   convertActivity(mdLuckerDetail.Activity),
		Prize:      convertPrize(mdLuckerDetail.Prize),
		PrizeGrade: convertPrizeGrade(mdLuckerDetail.PrizeGrade),
	}
}

func convertLuckerDetailsList(mdLuckerDetailsList []*model.LuckerDetailsInfo) (list []*luckdraw_pb.LuckerDetail) {
	list = make([]*luckdraw_pb.LuckerDetail, len(mdLuckerDetailsList))
	for index, mdluckerDetail := range mdLuckerDetailsList {
		pbLuckerDetail := convertLuckerDetail(mdluckerDetail)
		list[index] = pbLuckerDetail
	}
	return
}
