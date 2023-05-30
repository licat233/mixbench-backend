import webapi from "./gocliRequest"
import * as components from "./adminComponents"
export * from "./adminComponents"

/**
 * @description "鉴权接口"
 */
export function auth() {
	return webapi.get<null>(`/v1/api/admin/auth`)
}

/**
 * @description "登录"
 * @param req
 */
export function login(req: components.LoginReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/login`, req)
}

/**
 * @description "邮箱登录"
 * @param req
 */
export function loginByEmail(req: components.LoginByEmailReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/loginByEmail`, req)
}

/**
 * @description "给登录邮箱发送验证码"
 * @param req
 */
export function loginEmailCaptcha(req: components.LoginEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/loginEmailCaptcha`, req)
}

/**
 * @description "给注册邮箱发送验证码"
 * @param req
 */
export function registerEmailCaptcha(req: components.RegisterEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/registerEmailCaptcha`, req)
}

/**
 * @description "给绑定邮箱发送验证码"
 * @param req
 */
export function bindEmailCaptcha(req: components.BindEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/bindEmailCaptcha`, req)
}

/**
 * @description '预留的调试端口'
 * @param params
 */
export function test(params: components.TestReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/test`, params)
}

/**
 * @description "注销登录"
 */
export function logout() {
	return webapi.post<components.BaseResp>(`/v1/api/admin/logout`)
}

/**
 * @description "获取当前用户信息"
 */
export function currentUser() {
	return webapi.post<components.BaseResp>(`/v1/api/admin/currentUser`)
}

/**
 * @description "添加管理员 基础API"
 * @param req
 */
export function addAdminer(req: components.AddAdminerReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "更新管理员 基础API"
 * @param req
 */
export function putAdminer(req: components.PutAdminerReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "获取管理员 基础API"
 * @param req
 */
export function getAdminer(req: components.GetAdminerReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "删除管理员 基础API"
 * @param req
 */
export function delAdminer(req: components.DelAdminerReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "获取管理员列表 基础API"
 * @param params
 */
export function getAdminerList(params: components.GetAdminerListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer/list`, params)
}

/**
 * @description "获取管理员枚举列表 基础API"
 * @param req
 */
export function getAdminerEnums(req: components.GetAdminerEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer/enums`, req)
}

/**
 * @description "管理员校验自己的密码"
 * @param req
 */
export function verifyPassword(req: components.VerifyPasswordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer/verifyPassword`, req)
}

/**
 * @description "管理员修改自己的账号信息，当修改密码时，需要提供验证码"
 * @param req
 */
export function putCurrent(req: components.PutCurrentReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer/putCurrent`, req)
}

/**
 * @description "获取自己的账号信息"
 */
export function currentInfo() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer/currentInfo`)
}

/**
 * @description "修改自己的密码"
 * @param req
 */
export function putPassword(req: components.PutPasswordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer/putPassword`, req)
}

/**
 * @description "忘记密码，通过邮箱重置密码"
 * @param req
 */
export function forgotPassword(req: components.ForgotPasswordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer/forgotPassword`, req)
}

/**
 * @description "上传头像"
 */
export function uploadAvatar() {
	return webapi.post<components.BaseResp>(`/v1/api/admin/upload/avatar`)
}

/**
 * @description "上传logo"
 */
export function uploadLogo() {
	return webapi.post<components.BaseResp>(`/v1/api/admin/upload/logo`)
}

/**
 * @description "上传图片"
 */
export function uploadImg() {
	return webapi.post<components.BaseResp>(`/v1/api/admin/upload/image`)
}
