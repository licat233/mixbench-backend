import type * as components from "./luckydrawComponents"
export * from "./luckydrawComponents"

/**
 * @description "抽奖者身份验证"
 * @param req
 */
export function identify(req: components.IdentifyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/identify`, req)
}

/**
 * @description "获取活动奖品列表"
 * @param req
 */
export function prizes(req: components.PrizesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/prizes`, req)
}

/**
 * @description "获取自己在某活动的中奖名单"
 * @param req
 */
export function luckyList(req: components.LuckyListReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/luckyList`, req)
}

/**
 * @description "抽取奖品"
 * @param req
 */
export function goodluck(req: components.GoodluckReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/goodluck`, req)
}

/**
 * @description "查询中奖信息"
 * @param req
 */
export function search(req: components.SearchReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/search`, req)
}

/**
 * @description "挑选奖品"
 * @param req
 */
export function selectPrizes(req: components.SelectPrizesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/luckydraw/selectPrizes`, req)
}
