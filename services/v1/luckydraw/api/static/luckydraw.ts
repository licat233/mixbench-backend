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
 * @param params
 */
export function getActivity(params: components.GetActivityReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity`, params)
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
 * @param req
 */
export function getActivityList(req: components.GetActivityListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/list`, req)
}

/**
 * @description "获取活动枚举列表 基础API"
 * @param params
 */
export function getActivityEnums(params: components.GetActivityEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/activity/enums`, params)
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
 * @param params
 */
export function getLucky(params: components.GetLuckyReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky`, params)
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
 * @param req
 */
export function getLuckyList(req: components.GetLuckyListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky/list`, req)
}

/**
 * @description "获取中奖名单枚举列表 基础API"
 * @param params
 */
export function getLuckyEnums(params: components.GetLuckyEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/lucky/enums`, params)
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
 * @param params
 */
export function getPrize(params: components.GetPrizeReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize`, params)
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
 * @param req
 */
export function getPrizeList(req: components.GetPrizeListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize/list`, req)
}

/**
 * @description "获取奖品名单枚举列表 基础API"
 * @param params
 */
export function getPrizeEnums(params: components.GetPrizeEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prize/enums`, params)
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
 * @param params
 */
export function getPrizeGrade(params: components.GetPrizeGradeReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade`, params)
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
 * @param req
 */
export function getPrizeGradeList(req: components.GetPrizeGradeListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade/list`, req)
}

/**
 * @description "获取奖品级别枚举列表 基础API"
 * @param params
 */
export function getPrizeGradeEnums(params: components.GetPrizeGradeEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizeGrade/enums`, params)
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
 * @param params
 */
export function getPrizePool(params: components.GetPrizePoolReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool`, params)
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
 * @param req
 */
export function getPrizePoolList(req: components.GetPrizePoolListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool/list`, req)
}

/**
 * @description "获取活动奖品池，某活动中的奖品信息枚举列表 基础API"
 * @param params
 */
export function getPrizePoolEnums(params: components.GetPrizePoolEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/prizePool/enums`, params)
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
 * @param params
 */
export function getProject(params: components.GetProjectReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project`, params)
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
 * @param req
 */
export function getProjectList(req: components.GetProjectListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/list`, req)
}

/**
 * @description "获取项目枚举列表 基础API"
 * @param params
 */
export function getProjectEnums(params: components.GetProjectEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/project/enums`, params)
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
 * @param params
 */
export function getUser(params: components.GetUserReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user`, params)
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
 * @param req
 */
export function getUserList(req: components.GetUserListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/list`, req)
}

/**
 * @description "获取用户名单枚举列表 基础API"
 * @param params
 */
export function getUserEnums(params: components.GetUserEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/user/enums`, params)
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
 * @param params
 */
export function getUserPool(params: components.GetUserPoolReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool`, params)
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
 * @param req
 */
export function getUserPoolList(req: components.GetUserPoolListReq) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool/list`, req)
}

/**
 * @description "获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理枚举列表 基础API"
 * @param params
 */
export function getUserPoolEnums(params: components.GetUserPoolEnumsReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/luckydraw/userPool/enums`, params)
}
