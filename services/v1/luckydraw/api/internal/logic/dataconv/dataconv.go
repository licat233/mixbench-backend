package dataconv

import (
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"
)

func PbEnumToApiEnum(in *luckydraw_pb.Enum) *types.Enum {
	if in == nil {
		return nil
	}
	return &types.Enum{
		Label: in.Label,
		Value: in.Value,
	}
}

func PbEnumsToApiEnums(list []*luckydraw_pb.Enum) []*types.Enum {
	res := []*types.Enum{}
	for _, v := range list {
		res = append(res, PbEnumToApiEnum(v))
	}
	return res
}

func PbOptionToApiOption(in *luckydraw_pb.Option) *types.Option {
	if in == nil {
		return nil
	}
	return &types.Option{
		Title: in.Title,
		Value: in.Value,
	}
}

func PbOptionsToApiOptions(list []*luckydraw_pb.Option) []*types.Option {
	res := []*types.Option{}
	for _, v := range list {
		res = append(res, PbOptionToApiOption(v))
	}
	return res
}

func PbTreeOptionToApiTreeOption(in *luckydraw_pb.TreeOption) *types.TreeOption {
	if in == nil {
		return nil
	}
	return &types.TreeOption{
		Title:    in.Title,
		Value:    in.Value,
		Children: PbTreeOptionsToApiTreeOptions(in.Children),
	}
}

func PbTreeOptionsToApiTreeOptions(list []*luckydraw_pb.TreeOption) []types.TreeOption {
	res := []types.TreeOption{}
	for _, v := range list {
		res = append(res, *PbTreeOptionToApiTreeOption(v))
	}
	return res
}

func MdActivityToApiActivity(md *model.Activity) *types.Activity {
	if md == nil {
		return nil
	}
	return &types.Activity{
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

func MdActivities2ApiActivities(mdList []*model.Activity) []*types.Activity {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Activity, 0)
	for _, v := range mdList {
		data = append(data, MdActivityToApiActivity(v))
	}
	return data
}

func PbActivityToApiActivity(pb *luckydraw_pb.Activity) *types.Activity {
	if pb == nil {
		return nil
	}
	return &types.Activity{
		Id:        pb.Id,
		Uuid:      pb.Uuid,
		Name:      pb.Name,
		ProjectId: pb.ProjectId,
		Status:    pb.Status,
		Pattern:   pb.Pattern,
		MaxCount:  pb.MaxCount,
		MaxPrizes: pb.MaxPrizes,
		StartTime: pb.StartTime,
		EndTime:   pb.EndTime,
	}
}

func PbActivities2ApiActivities(pbList []*luckydraw_pb.Activity) []*types.Activity {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Activity, 0)
	for _, v := range pbList {
		data = append(data, PbActivityToApiActivity(v))
	}
	return data
}

func MdLuckyToApiLucky(md *model.Lucky) *types.Lucky {
	if md == nil {
		return nil
	}
	return &types.Lucky{
		Id:          md.Id,
		Uuid:        md.Uuid,
		UserId:      md.UserId,
		PrizePoolId: md.PrizePoolId,
		Status:      md.Status,
		LuckDate:    md.LuckDate.Unix(),
		Ip:          md.Ip,
	}
}

func MdLuckies2ApiLuckies(mdList []*model.Lucky) []*types.Lucky {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Lucky, 0)
	for _, v := range mdList {
		data = append(data, MdLuckyToApiLucky(v))
	}
	return data
}

func PbLuckyToApiLucky(pb *luckydraw_pb.Lucky) *types.Lucky {
	if pb == nil {
		return nil
	}
	return &types.Lucky{
		Id:          pb.Id,
		Uuid:        pb.Uuid,
		UserId:      pb.UserId,
		PrizePoolId: pb.PrizePoolId,
		Status:      pb.Status,
		LuckDate:    pb.LuckDate,
		Ip:          pb.Ip,
	}
}

func PbLuckies2ApiLuckies(pbList []*luckydraw_pb.Lucky) []*types.Lucky {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Lucky, 0)
	for _, v := range pbList {
		data = append(data, PbLuckyToApiLucky(v))
	}
	return data
}

func MdPrizeToApiPrize(md *model.Prize) *types.Prize {
	if md == nil {
		return nil
	}
	return &types.Prize{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Image:  md.Image,
		Price:  md.Price,
		Status: md.Status,
	}
}

func MdPrizes2ApiPrizes(mdList []*model.Prize) []*types.Prize {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Prize, 0)
	for _, v := range mdList {
		data = append(data, MdPrizeToApiPrize(v))
	}
	return data
}

func PbPrizeToApiPrize(pb *luckydraw_pb.Prize) *types.Prize {
	if pb == nil {
		return nil
	}
	return &types.Prize{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Image:  pb.Image,
		Price:  pb.Price,
		Status: pb.Status,
	}
}

func PbPrizes2ApiPrizes(pbList []*luckydraw_pb.Prize) []*types.Prize {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Prize, 0)
	for _, v := range pbList {
		data = append(data, PbPrizeToApiPrize(v))
	}
	return data
}

func MdPrizeGradeToApiPrizeGrade(md *model.PrizeGrade) *types.PrizeGrade {
	if md == nil {
		return nil
	}
	return &types.PrizeGrade{
		Id:    md.Id,
		Uuid:  md.Uuid,
		Name:  md.Name,
		IsWin: md.IsWin,
	}
}

func MdPrizeGrades2ApiPrizeGrades(mdList []*model.PrizeGrade) []*types.PrizeGrade {
	if mdList == nil {
		return nil
	}
	data := make([]*types.PrizeGrade, 0)
	for _, v := range mdList {
		data = append(data, MdPrizeGradeToApiPrizeGrade(v))
	}
	return data
}

func PbPrizeGradeToApiPrizeGrade(pb *luckydraw_pb.PrizeGrade) *types.PrizeGrade {
	if pb == nil {
		return nil
	}
	return &types.PrizeGrade{
		Id:    pb.Id,
		Uuid:  pb.Uuid,
		Name:  pb.Name,
		IsWin: pb.IsWin,
	}
}

func PbPrizeGrades2ApiPrizeGrades(pbList []*luckydraw_pb.PrizeGrade) []*types.PrizeGrade {
	if pbList == nil {
		return nil
	}
	data := make([]*types.PrizeGrade, 0)
	for _, v := range pbList {
		data = append(data, PbPrizeGradeToApiPrizeGrade(v))
	}
	return data
}

func MdPrizePoolToApiPrizePool(md *model.PrizePool) *types.PrizePool {
	if md == nil {
		return nil
	}
	return &types.PrizePool{
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

func MdPrizePools2ApiPrizePools(mdList []*model.PrizePool) []*types.PrizePool {
	if mdList == nil {
		return nil
	}
	data := make([]*types.PrizePool, 0)
	for _, v := range mdList {
		data = append(data, MdPrizePoolToApiPrizePool(v))
	}
	return data
}

func PbPrizePoolToApiPrizePool(pb *luckydraw_pb.PrizePool) *types.PrizePool {
	if pb == nil {
		return nil
	}
	return &types.PrizePool{
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

func PbPrizePools2ApiPrizePools(pbList []*luckydraw_pb.PrizePool) []*types.PrizePool {
	if pbList == nil {
		return nil
	}
	data := make([]*types.PrizePool, 0)
	for _, v := range pbList {
		data = append(data, PbPrizePoolToApiPrizePool(v))
	}
	return data
}

func MdProjectToApiProject(md *model.Project) *types.Project {
	if md == nil {
		return nil
	}
	return &types.Project{
		Id:     md.Id,
		Uuid:   md.Uuid,
		Name:   md.Name,
		Status: md.Status,
	}
}

func MdProjects2ApiProjects(mdList []*model.Project) []*types.Project {
	if mdList == nil {
		return nil
	}
	data := make([]*types.Project, 0)
	for _, v := range mdList {
		data = append(data, MdProjectToApiProject(v))
	}
	return data
}

func PbProjectToApiProject(pb *luckydraw_pb.Project) *types.Project {
	if pb == nil {
		return nil
	}
	return &types.Project{
		Id:     pb.Id,
		Uuid:   pb.Uuid,
		Name:   pb.Name,
		Status: pb.Status,
	}
}

func PbProjects2ApiProjects(pbList []*luckydraw_pb.Project) []*types.Project {
	if pbList == nil {
		return nil
	}
	data := make([]*types.Project, 0)
	for _, v := range pbList {
		data = append(data, PbProjectToApiProject(v))
	}
	return data
}

func MdUserToApiUser(md *model.User) *types.User {
	if md == nil {
		return nil
	}
	return &types.User{
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

func MdUsers2ApiUsers(mdList []*model.User) []*types.User {
	if mdList == nil {
		return nil
	}
	data := make([]*types.User, 0)
	for _, v := range mdList {
		data = append(data, MdUserToApiUser(v))
	}
	return data
}

func PbUserToApiUser(pb *luckydraw_pb.User) *types.User {
	if pb == nil {
		return nil
	}
	return &types.User{
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

func PbUsers2ApiUsers(pbList []*luckydraw_pb.User) []*types.User {
	if pbList == nil {
		return nil
	}
	data := make([]*types.User, 0)
	for _, v := range pbList {
		data = append(data, PbUserToApiUser(v))
	}
	return data
}

func MdUserPoolToApiUserPool(md *model.UserPool) *types.UserPool {
	if md == nil {
		return nil
	}
	return &types.UserPool{
		Id:         md.Id,
		Uuid:       md.Uuid,
		ActivityId: md.ActivityId,
		UserId:     md.UserId,
		Status:     md.Status,
		Score:      md.Score,
	}
}

func MdUserPools2ApiUserPools(mdList []*model.UserPool) []*types.UserPool {
	if mdList == nil {
		return nil
	}
	data := make([]*types.UserPool, 0)
	for _, v := range mdList {
		data = append(data, MdUserPoolToApiUserPool(v))
	}
	return data
}

func PbUserPoolToApiUserPool(pb *luckydraw_pb.UserPool) *types.UserPool {
	if pb == nil {
		return nil
	}
	return &types.UserPool{
		Id:         pb.Id,
		Uuid:       pb.Uuid,
		ActivityId: pb.ActivityId,
		UserId:     pb.UserId,
		Status:     pb.Status,
		Score:      pb.Score,
	}
}

func PbUserPools2ApiUserPools(pbList []*luckydraw_pb.UserPool) []*types.UserPool {
	if pbList == nil {
		return nil
	}
	data := make([]*types.UserPool, 0)
	for _, v := range pbList {
		data = append(data, PbUserPoolToApiUserPool(v))
	}
	return data
}
