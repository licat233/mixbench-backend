package tools

import (
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"
	"strings"
)

func NewStatusResponse() *luckydraw_pb.StatusResp {
	return &luckydraw_pb.StatusResp{}
}

func PbActivityToMdActivity(pbActivity *luckydraw_pb.Activity) *model.Activity {
	if pbActivity == nil {
		return nil
	}
	return &model.Activity{
		Id:        pbActivity.Id,
		Uuid:      pbActivity.Uuid,
		Name:      pbActivity.Name,
		ProjectId: pbActivity.ProjectId,
		Status:    pbActivity.Status,
		Pattern:   pbActivity.Pattern,
		MaxCount:  pbActivity.MaxCount,
		MaxPrizes: pbActivity.MaxPrizes,
		StartTime: utils.Int64ToTime(pbActivity.StartTime),
		EndTime:   utils.Int64ToTime(pbActivity.EndTime),
	}
}

func MdActivityToPbActivity(mdActivity *model.Activity) *luckydraw_pb.Activity {
	if mdActivity == nil {
		return nil
	}
	return &luckydraw_pb.Activity{
		Id:        mdActivity.Id,
		Uuid:      mdActivity.Uuid,
		Name:      mdActivity.Name,
		ProjectId: mdActivity.ProjectId,
		Status:    mdActivity.Status,
		Pattern:   mdActivity.Pattern,
		MaxCount:  mdActivity.MaxCount,
		MaxPrizes: mdActivity.MaxPrizes,
		StartTime: mdActivity.StartTime.Unix(),
		EndTime:   mdActivity.EndTime.Unix(),
	}
}

func MdActivitiesToPbActivities(mdActivities []*model.Activity) []*luckydraw_pb.Activity {
	pbActivities := make([]*luckydraw_pb.Activity, 0, len(mdActivities))
	for _, mdActivity := range mdActivities {
		pbActivities = append(pbActivities, MdActivityToPbActivity(mdActivity))
	}
	return pbActivities
}

func PbLuckyToMdLucky(pbLucky *luckydraw_pb.Lucky) *model.Lucky {
	if pbLucky == nil {
		return nil
	}
	return &model.Lucky{
		Id:          pbLucky.Id,
		Uuid:        pbLucky.Uuid,
		UserId:      pbLucky.UserId,
		PrizePoolId: pbLucky.PrizePoolId,
		Status:      pbLucky.Status,
		LuckDate:    utils.Int64ToTime(pbLucky.LuckDate),
		Ip:          pbLucky.Ip,
	}
}

func MdLuckyToPbLucky(mdLucky *model.Lucky) *luckydraw_pb.Lucky {
	if mdLucky == nil {
		return nil
	}
	return &luckydraw_pb.Lucky{
		Id:          mdLucky.Id,
		Uuid:        mdLucky.Uuid,
		UserId:      mdLucky.UserId,
		PrizePoolId: mdLucky.PrizePoolId,
		Status:      mdLucky.Status,
		LuckDate:    mdLucky.LuckDate.Unix(),
		Ip:          mdLucky.Ip,
	}
}

func MdLuckyListToPbLuckyList(mdLuckyList []*model.Lucky) []*luckydraw_pb.Lucky {
	pbLuckyList := make([]*luckydraw_pb.Lucky, 0, len(mdLuckyList))
	for _, mdLucky := range mdLuckyList {
		pbLuckyList = append(pbLuckyList, MdLuckyToPbLucky(mdLucky))
	}
	return pbLuckyList
}

func PbPrizeToMdPrize(pbPrize *luckydraw_pb.Prize) *model.Prize {
	if pbPrize == nil {
		return nil
	}
	return &model.Prize{
		Id:     pbPrize.Id,
		Uuid:   pbPrize.Uuid,
		Name:   pbPrize.Name,
		Image:  pbPrize.Image,
		Price:  pbPrize.Price,
		Status: pbPrize.Status,
	}
}

func MdPrizeToPbPrize(mdPrize *model.Prize) *luckydraw_pb.Prize {
	if mdPrize == nil {
		return nil
	}
	return &luckydraw_pb.Prize{
		Id:     mdPrize.Id,
		Uuid:   mdPrize.Uuid,
		Name:   mdPrize.Name,
		Image:  mdPrize.Image,
		Price:  mdPrize.Price,
		Status: mdPrize.Status,
	}
}

func MdPrizesToPbPrizes(mdPrizeList []*model.Prize) []*luckydraw_pb.Prize {
	pbPrizeList := make([]*luckydraw_pb.Prize, 0, len(mdPrizeList))
	for _, mdPrize := range mdPrizeList {
		pbPrizeList = append(pbPrizeList, MdPrizeToPbPrize(mdPrize))
	}
	return pbPrizeList
}

func PbPrizeGradeToMdPrizeGrade(pbPrizeGrade *luckydraw_pb.PrizeGrade) *model.PrizeGrade {
	if pbPrizeGrade == nil {
		return nil
	}
	return &model.PrizeGrade{
		Id:    pbPrizeGrade.Id,
		Name:  pbPrizeGrade.Name,
		IsWin: pbPrizeGrade.IsWin,
	}
}

func MdPrizeGradeToPbPrizeGrade(mdPrizeGrade *model.PrizeGrade) *luckydraw_pb.PrizeGrade {
	if mdPrizeGrade == nil {
		return nil
	}
	return &luckydraw_pb.PrizeGrade{
		Id:    mdPrizeGrade.Id,
		Name:  mdPrizeGrade.Name,
		IsWin: mdPrizeGrade.IsWin,
	}
}

func MdPrizeGradesToPbPrizeGrades(mdPrizeGrades []*model.PrizeGrade) []*luckydraw_pb.PrizeGrade {
	pbPrizeGrades := make([]*luckydraw_pb.PrizeGrade, 0, len(mdPrizeGrades))
	for _, mdPrizeGrade := range mdPrizeGrades {
		pbPrizeGrades = append(pbPrizeGrades, MdPrizeGradeToPbPrizeGrade(mdPrizeGrade))
	}
	return pbPrizeGrades
}

func PbPrizePoolToMdPrizePool(pbPrizePool *luckydraw_pb.PrizePool) *model.PrizePool {
	if pbPrizePool == nil {
		return nil
	}
	return &model.PrizePool{
		Id:         pbPrizePool.Id,
		Uuid:       pbPrizePool.Uuid,
		ActivityId: pbPrizePool.ActivityId,
		PrizeId:    pbPrizePool.PrizeId,
		GradeId:    pbPrizePool.GradeId,
		LuckValue:  pbPrizePool.LuckValue,
		Status:     pbPrizePool.Status,
		Total:      pbPrizePool.Total,
		Remainder:  pbPrizePool.Remainder,
	}
}

func MdPrizePoolToPbPrizePool(mdPrizePool *model.PrizePool) *luckydraw_pb.PrizePool {
	if mdPrizePool == nil {
		return nil
	}
	return &luckydraw_pb.PrizePool{
		Id:         mdPrizePool.Id,
		Uuid:       mdPrizePool.Uuid,
		ActivityId: mdPrizePool.ActivityId,
		PrizeId:    mdPrizePool.PrizeId,
		GradeId:    mdPrizePool.GradeId,
		LuckValue:  mdPrizePool.LuckValue,
		Status:     mdPrizePool.Status,
		Total:      mdPrizePool.Total,
		Remainder:  mdPrizePool.Remainder,
	}
}

func MdPrizePoolsToPbPrizePools(mdPrizePools []*model.PrizePool) []*luckydraw_pb.PrizePool {
	pbPrizePools := make([]*luckydraw_pb.PrizePool, 0, len(mdPrizePools))
	for _, mdPrizePool := range mdPrizePools {
		pbPrizePools = append(pbPrizePools, MdPrizePoolToPbPrizePool(mdPrizePool))
	}
	return pbPrizePools
}

func PbProjectToMdProject(pbProject *luckydraw_pb.Project) *model.Project {
	if pbProject == nil {
		return nil
	}
	return &model.Project{
		Id:     pbProject.Id,
		Uuid:   pbProject.Uuid,
		Name:   pbProject.Name,
		Status: pbProject.Status,
	}
}

func MdProjectToPbProject(mdProject *model.Project) *luckydraw_pb.Project {
	if mdProject == nil {
		return nil
	}
	return &luckydraw_pb.Project{
		Id:     mdProject.Id,
		Uuid:   mdProject.Uuid,
		Name:   mdProject.Name,
		Status: mdProject.Status,
	}
}

func MdProjectsToPbProjects(mdProjects []*model.Project) []*luckydraw_pb.Project {
	pbProjects := make([]*luckydraw_pb.Project, 0, len(mdProjects))
	for _, mdProject := range mdProjects {
		pbProjects = append(pbProjects, MdProjectToPbProject(mdProject))
	}
	return pbProjects
}

func PbUserToMdUser(pbUser *luckydraw_pb.User) *model.User {
	if pbUser == nil {
		return nil
	}
	return &model.User{
		Id:       pbUser.Id,
		Uuid:     pbUser.Uuid,
		Passport: pbUser.Passport,
		Email:    pbUser.Email,
		Phone:    pbUser.Phone,
		Name:     pbUser.Name,
		Address:  pbUser.Address,
		Score:    pbUser.Score,
		Status:   pbUser.Status,
	}
}

func MdUserToPbUser(mdUser *model.User) *luckydraw_pb.User {
	if mdUser == nil {
		return nil
	}
	return &luckydraw_pb.User{
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

func MdUsersToPbUsers(mdUsers []*model.User) []*luckydraw_pb.User {
	pbUsers := make([]*luckydraw_pb.User, 0, len(mdUsers))
	for _, mdUser := range mdUsers {
		pbUsers = append(pbUsers, MdUserToPbUser(mdUser))
	}
	return pbUsers
}

func PbUserPoolToMdUserPool(pbUserPool *luckydraw_pb.UserPool) *model.UserPool {
	if pbUserPool == nil {
		return nil
	}
	return &model.UserPool{
		Id:         pbUserPool.Id,
		Uuid:       pbUserPool.Uuid,
		ActivityId: pbUserPool.ActivityId,
		UserId:     pbUserPool.UserId,
		Status:     pbUserPool.Status,
		Score:      pbUserPool.Score,
	}
}

func MdUserPoolToPbUserPool(mdUserPool *model.UserPool) *luckydraw_pb.UserPool {
	if mdUserPool == nil {
		return nil
	}
	return &luckydraw_pb.UserPool{
		Id:         mdUserPool.Id,
		Uuid:       mdUserPool.Uuid,
		ActivityId: mdUserPool.ActivityId,
		UserId:     mdUserPool.UserId,
		Status:     mdUserPool.Status,
		Score:      mdUserPool.Score,
	}
}

func MdUserPoolsToPbUserPools(mdUserPoolList []*model.UserPool) []*luckydraw_pb.UserPool {
	pbUserPoolList := make([]*luckydraw_pb.UserPool, 0, len(mdUserPoolList))
	for _, mdUserPool := range mdUserPoolList {
		pbUserPoolList = append(pbUserPoolList, MdUserPoolToPbUserPool(mdUserPool))
	}
	return pbUserPoolList
}

func ConvertLuckyDetails(in []*model.LuckyDetail) []*luckydraw_pb.LuckyDetail {
	var res []*luckydraw_pb.LuckyDetail
	for _, item := range in {
		res = append(res, &luckydraw_pb.LuckyDetail{
			Lucky:      MdLuckyToPbLucky(item.Lucky),
			User:       MdUserToPbUser(item.User),
			PrizePool:  MdPrizePoolToPbPrizePool(item.PrizePool),
			Activity:   MdActivityToPbActivity(item.Activity),
			Prize:      MdPrizeToPbPrize(item.Prize),
			PrizeGrade: MdPrizeGradeToPbPrizeGrade(item.PrizeGrade),
			Project:    MdProjectToPbProject(item.Project),
		})
	}
	return res
}

func NewListReq(in *luckydraw_pb.ListReq) *luckydraw_pb.ListReq {
	if in == nil {
		return &luckydraw_pb.ListReq{
			PageSize: -1,
			Page:     -1,
			Keyword:  "",
		}
	}
	in.Keyword = strings.TrimSpace(in.Keyword)
	return in
}

func MdEnumToPbEnum(enum *model.Enum) *luckydraw_pb.Enum {
	return &luckydraw_pb.Enum{
		Label: enum.Label,
		Value: enum.Value,
	}
}

func MdEnumsToPbEnums(enumList []*model.Enum) []*luckydraw_pb.Enum {
	list := make([]*luckydraw_pb.Enum, len(enumList))
	for index, enum := range enumList {
		list[index] = MdEnumToPbEnum(enum)
	}
	return list
}
