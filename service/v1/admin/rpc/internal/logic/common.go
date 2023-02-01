package logic

import (
	"errors"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/model"
)

func newStatusResponse() *admin_pb.StatusResp {
	return &admin_pb.StatusResp{}
}

func pbActivityToMdActivity(pbActivity *admin_pb.Activity) *model.Activity {
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

func mdActivityToPbActivity(mdActivity *model.Activity) *admin_pb.Activity {
	if mdActivity == nil {
		return nil
	}
	return &admin_pb.Activity{
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

func mdActivityListToPbActivityList(mdActivityList []*model.Activity) []*admin_pb.Activity {
	pbActivityList := make([]*admin_pb.Activity, 0, len(mdActivityList))
	for _, mdActivity := range mdActivityList {
		pbActivityList = append(pbActivityList, mdActivityToPbActivity(mdActivity))
	}
	return pbActivityList
}

func pbAdminerToMdAdminer(pbAdminer *admin_pb.Adminer) *model.Adminer {
	if pbAdminer == nil {
		return nil
	}
	return &model.Adminer{
		Id:           pbAdminer.Id,
		Uuid:         pbAdminer.Uuid,
		Name:         pbAdminer.Name,
		Avatar:       pbAdminer.Avatar,
		Passport:     pbAdminer.Passport,
		Password:     pbAdminer.Password,
		Email:        pbAdminer.Email,
		Status:       pbAdminer.Status,
		IsSuperAdmin: pbAdminer.IsSuperAdmin,
		LoginCount:   pbAdminer.LoginCount,
		LastLogin:    utils.Int64ToTime(pbAdminer.LastLogin),
	}
}

func mdAdminerToPbAdminer(mdAdminer *model.Adminer) *admin_pb.Adminer {
	if mdAdminer == nil {
		return nil
	}
	return &admin_pb.Adminer{
		Id:           mdAdminer.Id,
		Uuid:         mdAdminer.Uuid,
		Name:         mdAdminer.Name,
		Avatar:       mdAdminer.Avatar,
		Passport:     mdAdminer.Passport,
		Password:     mdAdminer.Password,
		Email:        mdAdminer.Email,
		Status:       mdAdminer.Status,
		IsSuperAdmin: mdAdminer.IsSuperAdmin,
		LoginCount:   mdAdminer.LoginCount,
		LastLogin:    mdAdminer.LastLogin.Unix(),
	}
}

func mdAdminerListToPbAdminerList(mdAdminerList []*model.Adminer) []*admin_pb.Adminer {
	pbAdminerList := make([]*admin_pb.Adminer, 0, len(mdAdminerList))
	for _, mdAdminer := range mdAdminerList {
		pbAdminerList = append(pbAdminerList, mdAdminerToPbAdminer(mdAdminer))
	}
	return pbAdminerList
}

func pbLuckyToMdLucky(pbLucky *admin_pb.Lucky) *model.Lucky {
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

func mdLuckyToPbLucky(mdLucky *model.Lucky) *admin_pb.Lucky {
	if mdLucky == nil {
		return nil
	}
	return &admin_pb.Lucky{
		Id:          mdLucky.Id,
		Uuid:        mdLucky.Uuid,
		UserId:      mdLucky.UserId,
		PrizePoolId: mdLucky.PrizePoolId,
		Status:      mdLucky.Status,
		LuckDate:    mdLucky.LuckDate.Unix(),
		Ip:          mdLucky.Ip,
	}
}

func mdLuckyListToPbLuckyList(mdLuckyList []*model.Lucky) []*admin_pb.Lucky {
	pbLuckyList := make([]*admin_pb.Lucky, 0, len(mdLuckyList))
	for _, mdLucky := range mdLuckyList {
		pbLuckyList = append(pbLuckyList, mdLuckyToPbLucky(mdLucky))
	}
	return pbLuckyList
}

func pbPrizeToMdPrize(pbPrize *admin_pb.Prize) *model.Prize {
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

func mdPrizeToPbPrize(mdPrize *model.Prize) *admin_pb.Prize {
	if mdPrize == nil {
		return nil
	}
	return &admin_pb.Prize{
		Id:     mdPrize.Id,
		Uuid:   mdPrize.Uuid,
		Name:   mdPrize.Name,
		Image:  mdPrize.Image,
		Price:  mdPrize.Price,
		Status: mdPrize.Status,
	}
}

func mdPrizeListToPbPrizeList(mdPrizeList []*model.Prize) []*admin_pb.Prize {
	pbPrizeList := make([]*admin_pb.Prize, 0, len(mdPrizeList))
	for _, mdPrize := range mdPrizeList {
		pbPrizeList = append(pbPrizeList, mdPrizeToPbPrize(mdPrize))
	}
	return pbPrizeList
}

func pbPrizeGradeToMdPrizeGrade(pbPrizeGrade *admin_pb.PrizeGrade) *model.PrizeGrade {
	if pbPrizeGrade == nil {
		return nil
	}
	return &model.PrizeGrade{
		Id:    pbPrizeGrade.Id,
		Name:  pbPrizeGrade.Name,
		IsWin: pbPrizeGrade.IsWin,
	}
}

func mdPrizeGradeToPbPrizeGrade(mdPrizeGrade *model.PrizeGrade) *admin_pb.PrizeGrade {
	if mdPrizeGrade == nil {
		return nil
	}
	return &admin_pb.PrizeGrade{
		Id:    mdPrizeGrade.Id,
		Name:  mdPrizeGrade.Name,
		IsWin: mdPrizeGrade.IsWin,
	}
}

func mdPrizeGradeListToPbPrizeGradeList(mdPrizeGradeList []*model.PrizeGrade) []*admin_pb.PrizeGrade {
	pbPrizeGradeList := make([]*admin_pb.PrizeGrade, 0, len(mdPrizeGradeList))
	for _, mdPrizeGrade := range mdPrizeGradeList {
		pbPrizeGradeList = append(pbPrizeGradeList, mdPrizeGradeToPbPrizeGrade(mdPrizeGrade))
	}
	return pbPrizeGradeList
}

func pbPrizePoolToMdPrizePool(pbPrizePool *admin_pb.PrizePool) *model.PrizePool {
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

func mdPrizePoolToPbPrizePool(mdPrizePool *model.PrizePool) *admin_pb.PrizePool {
	if mdPrizePool == nil {
		return nil
	}
	return &admin_pb.PrizePool{
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

func mdPrizePoolListToPbPrizePoolList(mdPrizePoolList []*model.PrizePool) []*admin_pb.PrizePool {
	pbPrizePoolList := make([]*admin_pb.PrizePool, 0, len(mdPrizePoolList))
	for _, mdPrizePool := range mdPrizePoolList {
		pbPrizePoolList = append(pbPrizePoolList, mdPrizePoolToPbPrizePool(mdPrizePool))
	}
	return pbPrizePoolList
}

func pbProjectToMdProject(pbProject *admin_pb.Project) *model.Project {
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

func mdProjectToPbProject(mdProject *model.Project) *admin_pb.Project {
	if mdProject == nil {
		return nil
	}
	return &admin_pb.Project{
		Id:     mdProject.Id,
		Uuid:   mdProject.Uuid,
		Name:   mdProject.Name,
		Status: mdProject.Status,
	}
}

func mdProjectListToPbProjectList(mdProjectList []*model.Project) []*admin_pb.Project {
	pbProjectList := make([]*admin_pb.Project, 0, len(mdProjectList))
	for _, mdProject := range mdProjectList {
		pbProjectList = append(pbProjectList, mdProjectToPbProject(mdProject))
	}
	return pbProjectList
}

func pbUserToMdUser(pbUser *admin_pb.User) *model.User {
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

func mdUserToPbUser(mdUser *model.User) *admin_pb.User {
	if mdUser == nil {
		return nil
	}
	return &admin_pb.User{
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

func mdUserListToPbUserList(mdUserList []*model.User) []*admin_pb.User {
	pbUserList := make([]*admin_pb.User, 0, len(mdUserList))
	for _, mdUser := range mdUserList {
		pbUserList = append(pbUserList, mdUserToPbUser(mdUser))
	}
	return pbUserList
}

func pbUserPoolToMdUserPool(pbUserPool *admin_pb.UserPool) *model.UserPool {
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

func mdUserPoolToPbUserPool(mdUserPool *model.UserPool) *admin_pb.UserPool {
	if mdUserPool == nil {
		return nil
	}
	return &admin_pb.UserPool{
		Id:         mdUserPool.Id,
		Uuid:       mdUserPool.Uuid,
		ActivityId: mdUserPool.ActivityId,
		UserId:     mdUserPool.UserId,
		Status:     mdUserPool.Status,
		Score:      mdUserPool.Score,
	}
}

func mdUserPoolListToPbUserPoolList(mdUserPoolList []*model.UserPool) []*admin_pb.UserPool {
	pbUserPoolList := make([]*admin_pb.UserPool, 0, len(mdUserPoolList))
	for _, mdUserPool := range mdUserPoolList {
		pbUserPoolList = append(pbUserPoolList, mdUserPoolToPbUserPool(mdUserPool))
	}
	return pbUserPoolList
}

func verifyAdminerStatus(adminer *model.Adminer) error {
	if adminer.Status == 0 {
		return errors.New("this account has been disabled")
	}
	return nil
}

func convertLuckyDetailsList(in []*model.LuckyDetailsInfo) []*admin_pb.LuckyDetail {
	var res []*admin_pb.LuckyDetail
	for _, item := range in {
		res = append(res, &admin_pb.LuckyDetail{
			Lucky:      mdLuckyToPbLucky(item.Lucky),
			User:       mdUserToPbUser(item.User),
			PrizePool:  mdPrizePoolToPbPrizePool(item.PrizePool),
			Activity:   mdActivityToPbActivity(item.Activity),
			Prize:      mdPrizeToPbPrize(item.Prize),
			PrizeGrade: mdPrizeGradeToPbPrizeGrade(item.PrizeGrade),
			Project:    mdProjectToPbProject(item.Project),
		})
	}
	return res
}

func newListReq(in *admin_pb.ListReq) *admin_pb.ListReq {
	if in == nil {
		return &admin_pb.ListReq{
			PageSize: -1,
			Page:     -1,
			Keyword:  "",
		}
	}
	return in
}

func mdEnumToPbEnum(enum *model.Enum) *admin_pb.Enum {
	return &admin_pb.Enum{
		Label: enum.Label,
		Value: enum.Value,
	}
}

func mdEnumListToPbEnumList(enumList []*model.Enum) []*admin_pb.Enum {
	list := make([]*admin_pb.Enum, len(enumList))
	for index, enum := range enumList {
		list[index] = mdEnumToPbEnum(enum)
	}
	return list
}
