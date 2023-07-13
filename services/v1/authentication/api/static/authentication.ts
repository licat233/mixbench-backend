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
 * @param params
 */
export function auth(params: components.AuthReqParams) {
	return webapi.get<null>(`/v1/api/authentication/auth`, params)
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
export function addIpblack(req: components.AddIpblackReq) {
	return webapi.post<components.BaseResp>(`/v1/api/authentication/ipblack`, req)
}

/**
 * @description "更新IP黑名单 基础API"
 * @param req
 */
export function putIpblack(req: components.PutIpblackReq) {
	return webapi.put<components.BaseResp>(`/v1/api/authentication/ipblack`, req)
}

/**
 * @description "获取IP黑名单 基础API"
 * @param req
 */
export function getIpblack(req: components.GetIpblackReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipblack`, req)
}

/**
 * @description "删除IP黑名单 基础API"
 * @param req
 */
export function delIpblack(req: components.DelIpblackReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/authentication/ipblack`, req)
}

/**
 * @description "获取IP黑名单列表 基础API"
 * @param params
 */
export function getIpblackList(params: components.GetIpblackListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipblack/list`, params)
}

/**
 * @description "获取IP黑名单枚举列表 基础API"
 * @param req
 */
export function getIpblackEnums(req: components.GetIpblackEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/authentication/ipblack/enums`, req)
}
