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

func pbLuckerToMdLucker(pbLucker *admin_pb.Lucker) *model.Lucker {
	if pbLucker == nil {
		return nil
	}
	return &model.Lucker{
		Id:           pbLucker.Id,
		Uuid:         pbLucker.Uuid,
		UserId:       pbLucker.UserId,
		PrizesPoolId: pbLucker.PrizesPoolId,
		Status:       pbLucker.Status,
		LuckDate:     utils.Int64ToTime(pbLucker.LuckDate),
		Ip:           pbLucker.Ip,
	}
}

func mdLuckerToPbLucker(mdLucker *model.Lucker) *admin_pb.Lucker {
	if mdLucker == nil {
		return nil
	}
	return &admin_pb.Lucker{
		Id:           mdLucker.Id,
		Uuid:         mdLucker.Uuid,
		UserId:       mdLucker.UserId,
		PrizesPoolId: mdLucker.PrizesPoolId,
		Status:       mdLucker.Status,
		LuckDate:     mdLucker.LuckDate.Unix(),
		Ip:           mdLucker.Ip,
	}
}

func mdLuckerListToPbLuckerList(mdLuckerList []*model.Lucker) []*admin_pb.Lucker {
	pbLuckerList := make([]*admin_pb.Lucker, 0, len(mdLuckerList))
	for _, mdLucker := range mdLuckerList {
		pbLuckerList = append(pbLuckerList, mdLuckerToPbLucker(mdLucker))
	}
	return pbLuckerList
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

func pbPrizesPoolToMdPrizesPool(pbPrizesPool *admin_pb.PrizesPool) *model.PrizesPool {
	if pbPrizesPool == nil {
		return nil
	}
	return &model.PrizesPool{
		Id:         pbPrizesPool.Id,
		Uuid:       pbPrizesPool.Uuid,
		ActivityId: pbPrizesPool.ActivityId,
		PrizeId:    pbPrizesPool.PrizeId,
		GradeId:    pbPrizesPool.GradeId,
		LuckValue:  pbPrizesPool.LuckValue,
		Status:     pbPrizesPool.Status,
		Total:      pbPrizesPool.Total,
		Remainder:  pbPrizesPool.Remainder,
	}
}

func mdPrizesPoolToPbPrizesPool(mdPrizesPool *model.PrizesPool) *admin_pb.PrizesPool {
	if mdPrizesPool == nil {
		return nil
	}
	return &admin_pb.PrizesPool{
		Id:         mdPrizesPool.Id,
		Uuid:       mdPrizesPool.Uuid,
		ActivityId: mdPrizesPool.ActivityId,
		PrizeId:    mdPrizesPool.PrizeId,
		GradeId:    mdPrizesPool.GradeId,
		LuckValue:  mdPrizesPool.LuckValue,
		Status:     mdPrizesPool.Status,
		Total:      mdPrizesPool.Total,
		Remainder:  mdPrizesPool.Remainder,
	}
}

func mdPrizesPoolListToPbPrizesPoolList(mdPrizesPoolList []*model.PrizesPool) []*admin_pb.PrizesPool {
	pbPrizesPoolList := make([]*admin_pb.PrizesPool, 0, len(mdPrizesPoolList))
	for _, mdPrizesPool := range mdPrizesPoolList {
		pbPrizesPoolList = append(pbPrizesPoolList, mdPrizesPoolToPbPrizesPool(mdPrizesPool))
	}
	return pbPrizesPoolList
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

func pbUsersPoolToMdUsersPool(pbUsersPool *admin_pb.UsersPool) *model.UsersPool {
	if pbUsersPool == nil {
		return nil
	}
	return &model.UsersPool{
		Id:         pbUsersPool.Id,
		Uuid:       pbUsersPool.Uuid,
		ActivityId: pbUsersPool.ActivityId,
		UserId:     pbUsersPool.UserId,
		Status:     pbUsersPool.Status,
		Score:      pbUsersPool.Score,
	}
}

func mdUsersPoolToPbUsersPool(mdUsersPool *model.UsersPool) *admin_pb.UsersPool {
	if mdUsersPool == nil {
		return nil
	}
	return &admin_pb.UsersPool{
		Id:         mdUsersPool.Id,
		Uuid:       mdUsersPool.Uuid,
		ActivityId: mdUsersPool.ActivityId,
		UserId:     mdUsersPool.UserId,
		Status:     mdUsersPool.Status,
		Score:      mdUsersPool.Score,
	}
}

func mdUsersPoolListToPbUsersPoolList(mdUsersPoolList []*model.UsersPool) []*admin_pb.UsersPool {
	pbUsersPoolList := make([]*admin_pb.UsersPool, 0, len(mdUsersPoolList))
	for _, mdUsersPool := range mdUsersPoolList {
		pbUsersPoolList = append(pbUsersPoolList, mdUsersPoolToPbUsersPool(mdUsersPool))
	}
	return pbUsersPoolList
}

func verifyAdminerStatus(adminer *model.Adminer) error {
	if adminer.Status == 0 {
		return errors.New("this account has been disabled")
	}
	return nil
}

func convertLuckerDetailsList(in []*model.LuckerDetailsInfo) []*admin_pb.LuckerDetail {
	var res []*admin_pb.LuckerDetail
	for _, item := range in {
		res = append(res, &admin_pb.LuckerDetail{
			Lucker:     mdLuckerToPbLucker(item.Lucker),
			User:       mdUserToPbUser(item.User),
			PrizesPool: mdPrizesPoolToPbPrizesPool(item.PrizesPool),
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
