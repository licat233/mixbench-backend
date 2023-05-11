package dataconv

import (
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"
	"time"
)

func PbActivityToMdActivity(pb *luckydraw_pb.Activity) *model.Activity {
	if pb == nil {
		return nil
	}
	return &model.Activity{
		Id:        pb.Id,
		Uuid:      pb.Uuid,
		Name:      pb.Name,
		ProjectId: pb.ProjectId,
		Status:    pb.Status,
		Pattern:   pb.Pattern,
		MaxCount:  pb.MaxCount,
		MaxPrizes: pb.MaxPrizes,
		StartTime: time.Unix(pb.StartTime, 0).Local(),
		EndTime:   time.Unix(pb.EndTime, 0).Local(),
	}
}

func MdActivityToPbActivity(md *model.Activity) *luckydraw_pb.Activity {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.Activity{
		Id:        md.Id,
		Uuid:      md.Uuid,
		Name:      md.Name,
		ProjectId: md.ProjectId,
		Status:    md.Status,
		Pattern:   md.Pattern,
		MaxCount:  md.MaxCount,
		MaxPrizes: md.MaxPrizes,
		StartTime: md.StartTime.Unix(),
		EndTime:   md.EndTime.Unix(),
	}
}

func PbActivities2MdActivities(pbList []*luckydraw_pb.Activity) []*model.Activity {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Activity, 0)
	for _, v := range pbList {
		data = append(data, PbActivityToMdActivity(v))
	}
	return data
}

func MdActivities2PbActivities(mdList []*model.Activity) []*luckydraw_pb.Activity {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.Activity, 0)
	for _, v := range mdList {
		data = append(data, MdActivityToPbActivity(v))
	}
	return data
}

func PbLuckyToMdLucky(pb *luckydraw_pb.Lucky) *model.Lucky {
	if pb == nil {
		return nil
	}
	return &model.Lucky{
		Id:          pb.Id,
		Uuid:        pb.Uuid,
		UserId:      pb.UserId,
		PrizePoolId: pb.PrizePoolId,
		Status:      pb.Status,
		LuckDate:    time.Unix(pb.LuckDate, 0).Local(),
		Ip:          pb.Ip,
	}
}

func MdLuckyToPbLucky(md *model.Lucky) *luckydraw_pb.Lucky {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.Lucky{
		Id:          md.Id,
		Uuid:        md.Uuid,
		UserId:      md.UserId,
		PrizePoolId: md.PrizePoolId,
		Status:      md.Status,
		LuckDate:    md.LuckDate.Unix(),
		Ip:          md.Ip,
	}
}

func PbLuckies2MdLuckies(pbList []*luckydraw_pb.Lucky) []*model.Lucky {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Lucky, 0)
	for _, v := range pbList {
		data = append(data, PbLuckyToMdLucky(v))
	}
	return data
}

func MdLuckies2PbLuckies(mdList []*model.Lucky) []*luckydraw_pb.Lucky {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.Lucky, 0)
	for _, v := range mdList {
		data = append(data, MdLuckyToPbLucky(v))
	}
	return data
}

func PbPrizeToMdPrize(pb *luckydraw_pb.Prize) *model.Prize {
	if pb == nil {
		return nil
	}
	return &model.Prize{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Image:  pb.Image,
		Price:  pb.Price,
		Status: pb.Status,
	}
}

func MdPrizeToPbPrize(md *model.Prize) *luckydraw_pb.Prize {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.Prize{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Image:  md.Image,
		Price:  md.Price,
		Status: md.Status,
	}
}

func PbPrizes2MdPrizes(pbList []*luckydraw_pb.Prize) []*model.Prize {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Prize, 0)
	for _, v := range pbList {
		data = append(data, PbPrizeToMdPrize(v))
	}
	return data
}

func MdPrizes2PbPrizes(mdList []*model.Prize) []*luckydraw_pb.Prize {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.Prize, 0)
	for _, v := range mdList {
		data = append(data, MdPrizeToPbPrize(v))
	}
	return data
}

func PbPrizeGradeToMdPrizeGrade(pb *luckydraw_pb.PrizeGrade) *model.PrizeGrade {
	if pb == nil {
		return nil
	}
	return &model.PrizeGrade{
		Id:    pb.Id,
		Uuid:  pb.Uuid,
		Name:  pb.Name,
		IsWin: pb.IsWin,
	}
}

func MdPrizeGradeToPbPrizeGrade(md *model.PrizeGrade) *luckydraw_pb.PrizeGrade {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.PrizeGrade{
		Id:    md.Id,
		Uuid:  md.Uuid,
		Name:  md.Name,
		IsWin: md.IsWin,
	}
}

func PbPrizeGrades2MdPrizeGrades(pbList []*luckydraw_pb.PrizeGrade) []*model.PrizeGrade {
	if pbList == nil {
		return nil
	}
	data := make([]*model.PrizeGrade, 0)
	for _, v := range pbList {
		data = append(data, PbPrizeGradeToMdPrizeGrade(v))
	}
	return data
}

func MdPrizeGrades2PbPrizeGrades(mdList []*model.PrizeGrade) []*luckydraw_pb.PrizeGrade {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.PrizeGrade, 0)
	for _, v := range mdList {
		data = append(data, MdPrizeGradeToPbPrizeGrade(v))
	}
	return data
}

func PbPrizePoolToMdPrizePool(pb *luckydraw_pb.PrizePool) *model.PrizePool {
	if pb == nil {
		return nil
	}
	return &model.PrizePool{
		Id:         pb.Id,
		Uuid:       pb.Uuid,
		ActivityId: pb.ActivityId,
		PrizeId:    pb.PrizeId,
		GradeId:    pb.GradeId,
		LuckValue:  pb.LuckValue,
		Status:     pb.Status,
		Total:      pb.Total,
		Remainder:  pb.Remainder,
	}
}

func MdPrizePoolToPbPrizePool(md *model.PrizePool) *luckydraw_pb.PrizePool {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.PrizePool{
		Id:         md.Id,
		Uuid:       md.Uuid,
		ActivityId: md.ActivityId,
		PrizeId:    md.PrizeId,
		GradeId:    md.GradeId,
		LuckValue:  md.LuckValue,
		Status:     md.Status,
		Total:      md.Total,
		Remainder:  md.Remainder,
	}
}

func PbPrizePools2MdPrizePools(pbList []*luckydraw_pb.PrizePool) []*model.PrizePool {
	if pbList == nil {
		return nil
	}
	data := make([]*model.PrizePool, 0)
	for _, v := range pbList {
		data = append(data, PbPrizePoolToMdPrizePool(v))
	}
	return data
}

func MdPrizePools2PbPrizePools(mdList []*model.PrizePool) []*luckydraw_pb.PrizePool {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.PrizePool, 0)
	for _, v := range mdList {
		data = append(data, MdPrizePoolToPbPrizePool(v))
	}
	return data
}

func PbProjectToMdProject(pb *luckydraw_pb.Project) *model.Project {
	if pb == nil {
		return nil
	}
	return &model.Project{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Status: pb.Status,
	}
}

func MdProjectToPbProject(md *model.Project) *luckydraw_pb.Project {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.Project{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Status: md.Status,
	}
}

func PbProjects2MdProjects(pbList []*luckydraw_pb.Project) []*model.Project {
	if pbList == nil {
		return nil
	}
	data := make([]*model.Project, 0)
	for _, v := range pbList {
		data = append(data, PbProjectToMdProject(v))
	}
	return data
}

func MdProjects2PbProjects(mdList []*model.Project) []*luckydraw_pb.Project {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.Project, 0)
	for _, v := range mdList {
		data = append(data, MdProjectToPbProject(v))
	}
	return data
}

func PbUserToMdUser(pb *luckydraw_pb.User) *model.User {
	if pb == nil {
		return nil
	}
	return &model.User{
		Id:       pb.Id,
		Uuid:     pb.Uuid,
		Passport: pb.Passport,
		Email:    pb.Email,
		Phone:    pb.Phone,
		Name:     pb.Name,
		Address:  pb.Address,
		Score:    pb.Score,
		Status:   pb.Status,
	}
}

func MdUserToPbUser(md *model.User) *luckydraw_pb.User {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.User{
		Id:       md.Id,
		Uuid:     md.Uuid,
		Passport: md.Passport,
		Email:    md.Email,
		Phone:    md.Phone,
		Name:     md.Name,
		Address:  md.Address,
		Score:    md.Score,
		Status:   md.Status,
	}
}

func PbUsers2MdUsers(pbList []*luckydraw_pb.User) []*model.User {
	if pbList == nil {
		return nil
	}
	data := make([]*model.User, 0)
	for _, v := range pbList {
		data = append(data, PbUserToMdUser(v))
	}
	return data
}

func MdUsers2PbUsers(mdList []*model.User) []*luckydraw_pb.User {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.User, 0)
	for _, v := range mdList {
		data = append(data, MdUserToPbUser(v))
	}
	return data
}

func PbUserPoolToMdUserPool(pb *luckydraw_pb.UserPool) *model.UserPool {
	if pb == nil {
		return nil
	}
	return &model.UserPool{
		Id:         pb.Id,
		Uuid:       pb.Uuid,
		ActivityId: pb.ActivityId,
		UserId:     pb.UserId,
		Status:     pb.Status,
		Score:      pb.Score,
	}
}

func MdUserPoolToPbUserPool(md *model.UserPool) *luckydraw_pb.UserPool {
	if md == nil {
		return nil
	}
	return &luckydraw_pb.UserPool{
		Id:         md.Id,
		Uuid:       md.Uuid,
		ActivityId: md.ActivityId,
		UserId:     md.UserId,
		Status:     md.Status,
		Score:      md.Score,
	}
}

func PbUserPools2MdUserPools(pbList []*luckydraw_pb.UserPool) []*model.UserPool {
	if pbList == nil {
		return nil
	}
	data := make([]*model.UserPool, 0)
	for _, v := range pbList {
		data = append(data, PbUserPoolToMdUserPool(v))
	}
	return data
}

func MdUserPools2PbUserPools(mdList []*model.UserPool) []*luckydraw_pb.UserPool {
	if mdList == nil {
		return nil
	}
	data := make([]*luckydraw_pb.UserPool, 0)
	for _, v := range mdList {
		data = append(data, MdUserPoolToPbUserPool(v))
	}
	return data
}

func ListReqParams(req *luckydraw_pb.ListReq) (pageSize int64, page int64, keyword string) {
	if req != nil {
		pageSize = req.PageSize
		page = req.Page
		keyword = req.Keyword
	}
	return
}
