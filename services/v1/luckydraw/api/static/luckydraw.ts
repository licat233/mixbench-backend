import webapi from "./gocliRequest"
import * as components from "./luckydrawComponents"
export * from "./luckydrawComponents"

/**
 * @description "抽奖者身份验证"
 * @param req
 */
export function userIdentify(req: components.UserIdentifyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/userIdentify`, req)
}

/**
 * @description "获取活动奖品列表"
 * @param req
 */
export function activityPrizes(req: components.ActivityPrizesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/activityPrizes`, req)
}

/**
 * @description '预留的调试端口'
 * @param params
 */
export function test(params: components.TestReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/test`, params)
}

/**
 * @description "获取自己在某活动的中奖名单"
 * @param req
 */
export function getALuckyList(req: components.GetALuckyListReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/activityLuckyList`, req)
}

/**
 * @description "抽取奖品"
 * @param req
 */
export function luckydraw(req: components.LuckydrawReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/luckydraw`, req)
}

/**
 * @description "查询中奖信息"
 * @param req
 */
export function searchLucky(req: components.SearchLuckyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/searchLucky`, req)
}

/**
 * @description "挑选奖品"
 * @param req
 */
export function selectPrizes(req: components.SelectPrizesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/selectPrizes`, req)
}

/**
 * @description "添加活动 基础API"
 * @param req
 */
export function addActivity(req: components.AddActivityReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/activity`, req)
}

/**
 * @description "更新活动 基础API"
 * @param req
 */
export function putActivity(req: components.PutActivityReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/activity`, req)
}

/**
 * @description "获取活动 基础API"
 * @param req
 */
export function getActivity(req: components.GetActivityReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity`, req)
}

/**
 * @description "删除活动 基础API"
 * @param req
 */
export function delActivity(req: components.DelActivityReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/activity`, req)
}

/**
 * @description "获取活动列表 基础API"
 * @param params
 */
export function getActivityList(params: components.GetActivityListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/list`, params)
}

/**
 * @description "获取活动枚举列表 基础API"
 * @param req
 */
export function getActivityEnums(req: components.GetActivityEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/enums`, req)
}

/**
 * @description 
 * @param params
 */
export function activityDetail(params: components.ActivityDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/detail`, params)
}

/**
 * @description 
 * @param params
 */
export function activityDetails(params: components.GetActivityListReqParams) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/activity/details`, params)
}

/**
 * @description "指定项目的活动"
 * @param params
 */
export function projectActivities(params: components.ProjectActivitiesReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/projectActivities`, params)
}

/**
 * @description "获取某用户参与的活动"
 * @param params
 */
export function userActivities(params: components.UserActivitiesReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/userActivities`, params)
}

/**
 * @description 获取活动模式枚举列表
 */
export function activityPatternEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/patternEnumList`)
}

/**
 * @description "添加中奖名单 基础API"
 * @param req
 */
export function addLucky(req: components.AddLuckyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/lucky`, req)
}

/**
 * @description "更新中奖名单 基础API"
 * @param req
 */
export function putLucky(req: components.PutLuckyReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/lucky`, req)
}

/**
 * @description "获取中奖名单 基础API"
 * @param req
 */
export function getLucky(req: components.GetLuckyReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky`, req)
}

/**
 * @description "删除中奖名单 基础API"
 * @param req
 */
export function delLucky(req: components.DelLuckyReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/lucky`, req)
}

/**
 * @description "获取中奖名单列表 基础API"
 * @param params
 */
export function getLuckyList(params: components.GetLuckyListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky/list`, params)
}

/**
 * @description "获取中奖名单枚举列表 基础API"
 * @param req
 */
export function getLuckyEnums(req: components.GetLuckyEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky/enums`, req)
}

/**
 * @description "获取抽奖单详细列表"
 * @param req
 */
export function luckyDetails(req: components.LuckyDetailsReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/lucky/details`, req)
}

/**
 * @description "编辑抽奖单by prize"
 * @param req
 */
export function putLuckyByPrize(req: components.LuckyPutByPrizeReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/lucky/putLuckyByPrize`, req)
}

/**
 * @description "添加奖品名单 基础API"
 * @param req
 */
export function addPrize(req: components.AddPrizeReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/prize`, req)
}

/**
 * @description "更新奖品名单 基础API"
 * @param req
 */
export function putPrize(req: components.PutPrizeReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/prize`, req)
}

/**
 * @description "获取奖品名单 基础API"
 * @param req
 */
export function getPrize(req: components.GetPrizeReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize`, req)
}

/**
 * @description "删除奖品名单 基础API"
 * @param req
 */
export function delPrize(req: components.DelPrizeReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/prize`, req)
}

/**
 * @description "获取奖品名单列表 基础API"
 * @param params
 */
export function getPrizeList(params: components.GetPrizeListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize/list`, params)
}

/**
 * @description "获取奖品名单枚举列表 基础API"
 * @param req
 */
export function getPrizeEnums(req: components.GetPrizeEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize/enums`, req)
}

/**
 * @description "获取活动的奖品"
 * @param req
 */
export function getActivityPrizes(req: components.GetActivityPrizesReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize/activityPrizes`, req)
}

/**
 * @description "添加奖品级别 基础API"
 * @param req
 */
export function addPrizeGrade(req: components.AddPrizeGradeReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/prizeGrade`, req)
}

/**
 * @description "更新奖品级别 基础API"
 * @param req
 */
export function putPrizeGrade(req: components.PutPrizeGradeReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/prizeGrade`, req)
}

/**
 * @description "获取奖品级别 基础API"
 * @param req
 */
export function getPrizeGrade(req: components.GetPrizeGradeReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade`, req)
}

/**
 * @description "删除奖品级别 基础API"
 * @param req
 */
export function delPrizeGrade(req: components.DelPrizeGradeReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/prizeGrade`, req)
}

/**
 * @description "获取奖品级别列表 基础API"
 * @param params
 */
export function getPrizeGradeList(params: components.GetPrizeGradeListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade/list`, params)
}

/**
 * @description "获取奖品级别枚举列表 基础API"
 * @param req
 */
export function getPrizeGradeEnums(req: components.GetPrizeGradeEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade/enums`, req)
}

/**
 * @description "添加活动奖品池，某活动中的奖品信息 基础API"
 * @param req
 */
export function addPrizePool(req: components.AddPrizePoolReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/prizePool`, req)
}

/**
 * @description "更新活动奖品池，某活动中的奖品信息 基础API"
 * @param req
 */
export function putPrizePool(req: components.PutPrizePoolReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/prizePool`, req)
}

/**
 * @description "获取活动奖品池，某活动中的奖品信息 基础API"
 * @param req
 */
export function getPrizePool(req: components.GetPrizePoolReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool`, req)
}

/**
 * @description "删除活动奖品池，某活动中的奖品信息 基础API"
 * @param req
 */
export function delPrizePool(req: components.DelPrizePoolReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/prizePool`, req)
}

/**
 * @description "获取活动奖品池，某活动中的奖品信息列表 基础API"
 * @param params
 */
export function getPrizePoolList(params: components.GetPrizePoolListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool/list`, params)
}

/**
 * @description "获取活动奖品池，某活动中的奖品信息枚举列表 基础API"
 * @param req
 */
export function getPrizePoolEnums(req: components.GetPrizePoolEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool/enums`, req)
}

/**
 * @description "重置某活动的奖品池概率"
 * @param req
 */
export function resetLuckValue(req: components.ResetLuckValueReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/prizePool/resetLuckValue`, req)
}

/**
 * @description "添加项目 基础API"
 * @param req
 */
export function addProject(req: components.AddProjectReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/project`, req)
}

/**
 * @description "更新项目 基础API"
 * @param req
 */
export function putProject(req: components.PutProjectReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/project`, req)
}

/**
 * @description "获取项目 基础API"
 * @param req
 */
export function getProject(req: components.GetProjectReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project`, req)
}

/**
 * @description "删除项目 基础API"
 * @param req
 */
export function delProject(req: components.DelProjectReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/project`, req)
}

/**
 * @description "获取项目列表 基础API"
 * @param params
 */
export function getProjectList(params: components.GetProjectListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/list`, params)
}

/**
 * @description "获取项目枚举列表 基础API"
 * @param req
 */
export function getProjectEnums(req: components.GetProjectEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/enums`, req)
}

/**
 * @description "获取项目详细列表"
 * @param params
 */
export function projectDetails(params: components.GetProjectListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/projectDetails`, params)
}

/**
 * @description "获取项目详细"
 * @param params
 */
export function projectDetail(params: components.ProjectDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/projectDetail`, params)
}

/**
 * @description "获取项目树形枚举数据 []TreeOptions"
 */
export function projectTreeOptions() {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/projectTreeOptions`)
}

/**
 * @description "添加用户名单 基础API"
 * @param req
 */
export function addUser(req: components.AddUserReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/user`, req)
}

/**
 * @description "更新用户名单 基础API"
 * @param req
 */
export function putUser(req: components.PutUserReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/user`, req)
}

/**
 * @description "获取用户名单 基础API"
 * @param req
 */
export function getUser(req: components.GetUserReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user`, req)
}

/**
 * @description "删除用户名单 基础API"
 * @param req
 */
export function delUser(req: components.DelUserReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/user`, req)
}

/**
 * @description "获取用户名单列表 基础API"
 * @param params
 */
export function getUserList(params: components.GetUserListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/list`, params)
}

/**
 * @description "获取用户名单枚举列表 基础API"
 * @param req
 */
export function getUserEnums(req: components.GetUserEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/enums`, req)
}

/**
 * @description "获取用户详情"
 * @param params
 */
export function userDetail(params: components.UserDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/detail`, params)
}

/**
 * @description "获取用户详情列表"
 * @param params
 */
export function userDetails(params: components.GetUserListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/details`, params)
}

/**
 * @description "添加活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理 基础API"
 * @param req
 */
export function addUserPool(req: components.AddUserPoolReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/userPool`, req)
}

/**
 * @description "更新活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理 基础API"
 * @param req
 */
export function putUserPool(req: components.PutUserPoolReq) {
	return webapi.put<components.BaseResp>(`/v1/api/luckydraw/userPool`, req)
}

/**
 * @description "获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理 基础API"
 * @param req
 */
export function getUserPool(req: components.GetUserPoolReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool`, req)
}

/**
 * @description "删除活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理 基础API"
 * @param req
 */
export function delUserPool(req: components.DelUserPoolReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/luckydraw/userPool`, req)
}

/**
 * @description "获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理列表 基础API"
 * @param params
 */
export function getUserPoolList(params: components.GetUserPoolListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool/list`, params)
}

/**
 * @description "获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理枚举列表 基础API"
 * @param req
 */
export function getUserPoolEnums(req: components.GetUserPoolEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool/enums`, req)
}
