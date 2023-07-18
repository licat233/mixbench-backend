import webapi from "./gocliRequest"
import type * as components from "./authenticationComponents"
export * from "./authenticationComponents"

/**
 * @description "预留的调试端口"
 * @param params
 */
export function test(params: components.TestReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/test`, params)
}

/**
 * @description "jwt鉴权接口"
 */
export function auth() {
	return webapi.get<null>(`/v1/api/authentication/auth`)
}

/**
 * @description "获取数学验证码"
 * @param req
 */
export function getMathCaptcha(req: components.CaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/captcha/math`, req)
}

/**
 * @description "验证码图片"
 */
export function showCaptchaImage() {
	return webapi.get<null>(`/v1/api/authentication/captcha/show`)
}

/**
 * @description "验证码校验"
 * @param req
 */
export function verifyCaptcha(req: components.ProcessCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/captcha/process`, req)
}

/**
 * @description "向邮箱发送验证码"
 * @param req
 */
export function sendEmailCaptcha(req: components.EmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/captcha/email`, req)
}

/**
 * @description "邮箱验证码校验"
 * @param req
 */
export function verifyEmailCaptcha(req: components.EmailCaptchaVerifyReq) {
	return webapi.post<components.EmailCaptchaResp>(`/v1/api/authentication/captcha/email/process`, req)
}

/**
 * @description "添加IP黑名单 基础API"
 * @param req
 */
export function addIpBlacklist(req: components.AddIpBlacklistReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/ipBlacklist`, req)
}

/**
 * @description "更新IP黑名单 基础API"
 * @param req
 */
export function putIpBlacklist(req: components.PutIpBlacklistReq) {
	return webapi.put<components.BaseResp>(`/v1/api/authentication/ipBlacklist`, req)
}

/**
 * @description "获取IP黑名单 基础API"
 * @param req
 */
export function getIpBlacklist(req: components.GetIpBlacklistReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipBlacklist`, req)
}

/**
 * @description "删除IP黑名单 基础API"
 * @param req
 */
export function delIpBlacklist(req: components.DelIpBlacklistReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/authentication/ipBlacklist`, req)
}

/**
 * @description "获取IP黑名单列表 基础API"
 * @param params
 */
export function getIpBlacklistList(params: components.GetIpBlacklistListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipBlacklist/list`, params)
}

/**
 * @description "获取IP黑名单枚举列表 基础API"
 * @param req
 */
export function getIpBlacklistEnums(req: components.GetIpBlacklistEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipBlacklist/enums`, req)
}

/**
 * @description "添加token黑名单 基础API"
 * @param req
 */
export function addTokenBlacklist(req: components.AddTokenBlacklistReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/tokenBlacklist`, req)
}

/**
 * @description "更新token黑名单 基础API"
 * @param req
 */
export function putTokenBlacklist(req: components.PutTokenBlacklistReq) {
	return webapi.put<components.BaseResp>(`/v1/api/authentication/tokenBlacklist`, req)
}

/**
 * @description "获取token黑名单 基础API"
 * @param req
 */
export function getTokenBlacklist(req: components.GetTokenBlacklistReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/tokenBlacklist`, req)
}

/**
 * @description "删除token黑名单 基础API"
 * @param req
 */
export function delTokenBlacklist(req: components.DelTokenBlacklistReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/authentication/tokenBlacklist`, req)
}

/**
 * @description "获取token黑名单列表 基础API"
 * @param params
 */
export function getTokenBlacklistList(params: components.GetTokenBlacklistListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/tokenBlacklist/list`, params)
}

/**
 * @description "获取token黑名单枚举列表 基础API"
 * @param req
 */
export function getTokenBlacklistEnums(req: components.GetTokenBlacklistEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/tokenBlacklist/enums`, req)
}
