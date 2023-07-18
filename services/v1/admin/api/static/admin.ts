import webapi from "./gocliRequest"
import type * as components from "./adminComponents"
export * from "./adminComponents"

/**
 * @description "登录"
 * @param req
 */
export function login(req: components.LoginReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/public/login`, req)
}

/**
 * @description "邮箱登录"
 * @param req
 */
export function loginByEmail(req: components.LoginByEmailReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/public/loginByEmail`, req)
}

/**
 * @description "给登录邮箱发送验证码"
 * @param req
 */
export function loginEmailCaptcha(req: components.LoginEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/public/loginEmailCaptcha`, req)
}

/**
 * @description "给注册邮箱发送验证码"
 * @param req
 */
export function registerEmailCaptcha(req: components.RegisterEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/public/registerEmailCaptcha`, req)
}

/**
 * @description "给绑定邮箱发送验证码"
 * @param req
 */
export function bindEmailCaptcha(req: components.BindEmailCaptchaReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/public/bindEmailCaptcha`, req)
}

/**
 * @description '预留的调试端口'
 * @param params
 */
export function test(params: components.TestReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/public/test`, params)
}

/**
 * @description "密码加密端口"
 * @param params
 */
export function pwdEncryption(params: components.PwdEncryptionReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/public/pwdEncryption`, params)
}

/**
 * @description "密码比较端口"
 * @param params
 */
export function pwdCompare(params: components.PwdCompareReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/public/pwdCompare`, params)
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
 * @description "添加活动 基础API"
 * @param req
 */
export function addCampaign(req: components.AddCampaignReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/campaign`, req)
}

/**
 * @description "更新活动 基础API"
 * @param req
 */
export function putCampaign(req: components.PutCampaignReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/campaign`, req)
}

/**
 * @description "获取活动 基础API"
 * @param req
 */
export function getCampaign(req: components.GetCampaignReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/campaign`, req)
}

/**
 * @description "删除活动 基础API"
 * @param req
 */
export function delCampaign(req: components.DelCampaignReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/campaign`, req)
}

/**
 * @description "获取活动列表 基础API"
 * @param params
 */
export function getCampaignList(params: components.GetCampaignListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/campaign/list`, params)
}

/**
 * @description "获取活动枚举列表 基础API"
 * @param req
 */
export function getCampaignEnums(req: components.GetCampaignEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/campaign/enums`, req)
}

/**
 * @description "添加公司 基础API"
 * @param req
 */
export function addCompany(req: components.AddCompanyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/company`, req)
}

/**
 * @description "更新公司 基础API"
 * @param req
 */
export function putCompany(req: components.PutCompanyReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/company`, req)
}

/**
 * @description "获取公司 基础API"
 * @param req
 */
export function getCompany(req: components.GetCompanyReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/company`, req)
}

/**
 * @description "删除公司 基础API"
 * @param req
 */
export function delCompany(req: components.DelCompanyReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/company`, req)
}

/**
 * @description "获取公司列表 基础API"
 * @param params
 */
export function getCompanyList(params: components.GetCompanyListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/company/list`, params)
}

/**
 * @description "获取公司枚举列表 基础API"
 * @param req
 */
export function getCompanyEnums(req: components.GetCompanyEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/company/enums`, req)
}

/**
 * @description "添加项目 基础API"
 * @param req
 */
export function addProject(req: components.AddProjectReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "更新项目 基础API"
 * @param req
 */
export function putProject(req: components.PutProjectReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "获取项目 基础API"
 * @param req
 */
export function getProject(req: components.GetProjectReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "删除项目 基础API"
 * @param req
 */
export function delProject(req: components.DelProjectReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "获取项目列表 基础API"
 * @param params
 */
export function getProjectList(params: components.GetProjectListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/list`, params)
}

/**
 * @description "获取项目枚举列表 基础API"
 * @param req
 */
export function getProjectEnums(req: components.GetProjectEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/enums`, req)
}

/**
 * @description "添加网站 基础API"
 * @param req
 */
export function addWeb(req: components.AddWebReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/web`, req)
}

/**
 * @description "更新网站 基础API"
 * @param req
 */
export function putWeb(req: components.PutWebReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/web`, req)
}

/**
 * @description "获取网站 基础API"
 * @param req
 */
export function getWeb(req: components.GetWebReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/web`, req)
}

/**
 * @description "删除网站 基础API"
 * @param req
 */
export function delWeb(req: components.DelWebReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/web`, req)
}

/**
 * @description "获取网站列表 基础API"
 * @param params
 */
export function getWebList(params: components.GetWebListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/web/list`, params)
}

/**
 * @description "获取网站枚举列表 基础API"
 * @param req
 */
export function getWebEnums(req: components.GetWebEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/web/enums`, req)
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
