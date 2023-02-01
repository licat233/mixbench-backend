import type * as components from "./adminComponents"
export * from "./adminComponents"

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
 * @description "获取活动列表，因为设计数组传参，post方式会比较合适"
 * @param req
 */
export function activityList(req: components.ActivityListReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/activity/list`, req)
}

/**
 * @description 
 * @param params
 */
export function activityDetail(params: components.ActivityDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity/detail`, params)
}

/**
 * @description 
 * @param req
 */
export function activityDetails(req: components.ActivityListReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/activity/details`, req)
}

/**
 * @description "获取活动信息"
 * @param params
 */
export function activityInfo(params: components.ActivityGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity`, params)
}

/**
 * @description "添加活动"
 * @param req
 */
export function activityAdd(req: components.ActivityAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/activity`, req)
}

/**
 * @description "编辑活动"
 * @param req
 */
export function activityPut(req: components.ActivityPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/activity`, req)
}

/**
 * @description "删除活动"
 * @param req
 */
export function activityDel(req: components.ActivityDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/activity`, req)
}

/**
 * @description "枚举列表"
 * @param params
 */
export function activityEnumList(params: components.EnumListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity/enumList`, params)
}

/**
 * @description "指定项目的活动"
 * @param params
 */
export function projectActivity(params: components.ProjectActivityReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity/projectActivity`, params)
}

/**
 * @description "获取某用户参与的活动"
 * @param params
 */
export function userActivities(params: components.UserActivitiesReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity/userActivities`, params)
}

/**
 * @description 获取活动模式枚举列表
 */
export function activityPatternEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/activity/patternEnumList`)
}

/**
 * @description "获取管理员列表"
 * @param params
 */
export function adminerList(params: components.AdminerListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer/list`, params)
}

/**
 * @description "获取管理员信息"
 * @param params
 */
export function adminerInfo(params: components.AdminerGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer`, params)
}

/**
 * @description "添加管理员"
 * @param req
 */
export function adminerAdd(req: components.AdminerAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "编辑管理员"
 * @param req
 */
export function adminerPut(req: components.AdminerPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "删除管理员"
 * @param req
 */
export function adminerDel(req: components.AdminerDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/adminer`, req)
}

/**
 * @description "枚举列表"
 */
export function adminerEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/adminer/enumList`)
}

/**
 * @description "获取抽奖单列表"
 * @param params
 */
export function luckyList(params: components.LuckyListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/lucky/list`, params)
}

/**
 * @description "获取抽奖单详细列表"
 * @param req
 */
export function luckyDetails(req: components.LuckyDetailsReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/lucky/details`, req)
}

/**
 * @description "获取抽奖单信息"
 * @param params
 */
export function luckyInfo(params: components.LuckyGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/lucky`, params)
}

/**
 * @description "添加抽奖单"
 * @param req
 */
export function luckyAdd(req: components.LuckyAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/lucky`, req)
}

/**
 * @description "编辑抽奖单"
 * @param req
 */
export function luckyPut(req: components.LuckyPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/lucky`, req)
}

/**
 * @description "编辑抽奖单by prize"
 * @param req
 */
export function putByPrize(req: components.LuckyPutByPrizeReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/lucky/putByPrize`, req)
}

/**
 * @description "删除抽奖单"
 * @param req
 */
export function luckyDel(req: components.LuckyDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/lucky`, req)
}

/**
 * @description "枚举列表"
 */
export function luckyEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/lucky/enumList`)
}

/**
 * @description "获取奖品列表"
 * @param params
 */
export function prizeList(params: components.PrizeListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prize/list`, params)
}

/**
 * @description "获取奖品信息"
 * @param params
 */
export function prizeInfo(params: components.PrizeGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prize`, params)
}

/**
 * @description "添加奖品"
 * @param req
 */
export function prizeAdd(req: components.PrizeAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/prize`, req)
}

/**
 * @description "编辑奖品"
 * @param req
 */
export function prizePut(req: components.PrizePutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/prize`, req)
}

/**
 * @description "删除奖品"
 * @param req
 */
export function prizeDel(req: components.PrizeDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/prize`, req)
}

/**
 * @description "枚举列表"
 */
export function prizeEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prize/enumList`)
}

/**
 * @description "获取某活动奖品池的奖品"
 * @param params
 */
export function activityPrizes(params: components.ActivityPrizesReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prize/activityPrizes`, params)
}

/**
 * @description "获取奖品级别列表"
 * @param params
 */
export function prizeGradeList(params: components.PrizeGradeListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prizeGrade/list`, params)
}

/**
 * @description "获取奖品级别信息"
 * @param params
 */
export function prizeGradeInfo(params: components.PrizeGradeGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prizeGrade`, params)
}

/**
 * @description "添加奖品级别"
 * @param req
 */
export function prizeGradeAdd(req: components.PrizeGradeAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/prizeGrade`, req)
}

/**
 * @description "编辑奖品级别"
 * @param req
 */
export function prizeGradePut(req: components.PrizeGradePutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/prizeGrade`, req)
}

/**
 * @description "删除奖品级别"
 * @param req
 */
export function prizeGradeDel(req: components.PrizeGradeDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/prizeGrade`, req)
}

/**
 * @description "枚举列表"
 */
export function prizeGradeEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prizeGrade/enumList`)
}

/**
 * @description "获取奖品池列表"
 * @param params
 */
export function prizesPoolList(params: components.PrizePoolListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prizesPool/list`, params)
}

/**
 * @description "获取奖品池信息"
 * @param params
 */
export function prizesPoolInfo(params: components.PrizePoolGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/prizesPool`, params)
}

/**
 * @description "添加奖品池"
 * @param req
 */
export function prizesPoolAdd(req: components.PrizePoolAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/prizesPool`, req)
}

/**
 * @description "编辑奖品池"
 * @param req
 */
export function prizesPoolPut(req: components.PrizePoolPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/prizesPool`, req)
}

/**
 * @description "删除奖品池"
 * @param req
 */
export function prizesPoolDel(req: components.PrizePoolDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/prizesPool`, req)
}

/**
 * @description 重置某活动的奖品池概率
 * @param req
 */
export function resetLuckValue(req: components.ResetLuckValueReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/prizesPool/resetLuckValue`, req)
}

/**
 * @description "获取项目列表"
 * @param params
 */
export function projectList(params: components.ProjectListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/list`, params)
}

/**
 * @description "获取项目详细列表"
 * @param params
 */
export function projectDetails(params: components.ProjectListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/projectDetails`, params)
}

/**
 * @description "获取项目信息"
 * @param params
 */
export function projectInfo(params: components.ProjectGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project`, params)
}

/**
 * @description "添加项目"
 * @param req
 */
export function projectAdd(req: components.ProjectAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "编辑项目"
 * @param req
 */
export function projectPut(req: components.ProjectPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "删除项目"
 * @param req
 */
export function projectDel(req: components.ProjectDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/project`, req)
}

/**
 * @description "枚举列表"
 */
export function projectEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/enumList`)
}

/**
 * @description "获取项目详细"
 * @param params
 */
export function projectDetail(params: components.ProjectDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/projectDetail`, params)
}

/**
 * @description "获取项目树形枚举数据 []TreeOptions"
 */
export function projectTreeOptions() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/project/projectTreeOptions`)
}

/**
 * @description "获取授权策略"
 * @param req
 */
export function getPolicy(req: components.PolicyReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/policy`, req)
}

/**
 * @description "添加授权策略"
 * @param req
 */
export function addPolicy(req: components.PolicyReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/rbac/policy`, req)
}

/**
 * @description "修改授权规则"
 * @param req
 */
export function putPolicy(req: components.PolicyReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/rbac/policy`, req)
}

/**
 * @description "删除授权规则"
 * @param req
 */
export function delPolicy(req: components.PolicyReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/rbac/policy`, req)
}

/**
 * @description "批量获取授权策略"
 * @param req
 */
export function getPolicies(req: components.PoliciesReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/policies`, req)
}

/**
 * @description "批量添加授权策略"
 * @param req
 */
export function addPolicies(req: components.PoliciesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/rbac/policies`, req)
}

/**
 * @description "批量修改授权规则"
 * @param req
 */
export function putPolicies(req: components.PoliciesReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/rbac/policies`, req)
}

/**
 * @description "批量删除授权规则"
 * @param req
 */
export function delPolicies(req: components.PoliciesReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/rbac/policies`, req)
}

/**
 * @description "获取当前策略中显示的操作列表"
 */
export function getAllActions() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/all-actions`)
}

/**
 * @description "获取策略中的所有授权规则"
 */
export function getAllPolicy() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/all-policies`)
}

/**
 * @description "获取当前策略中显示的 object 列表"
 */
export function getAllObjects() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/all-objects`)
}

/**
 * @description "将获得所有与该域相关联的用户"
 * @param req
 */
export function getAllUsersByDomain(req: components.DomainReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/domain-users`, req)
}

/**
 * @description "获取用户在某域内的权限"
 * @param req
 */
export function getPermissionsForUserInDomain(req: components.RoleAndDomainReq) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/rbac/permissions`, req)
}

/**
 * @description "获取用户列表"
 * @param params
 */
export function userList(params: components.UserListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/user/list`, params)
}

/**
 * @description "获取用户详情列表"
 * @param params
 */
export function userDetail(params: components.UserDetailReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/user/detail`, params)
}

/**
 * @description "获取用户详情列表"
 * @param params
 */
export function userDetails(params: components.UserListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/user/details`, params)
}

/**
 * @description "获取用户信息"
 * @param params
 */
export function userInfo(params: components.UserGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/user`, params)
}

/**
 * @description "添加用户"
 * @param req
 */
export function userAdd(req: components.UserAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/user`, req)
}

/**
 * @description "编辑用户"
 * @param req
 */
export function userPut(req: components.UserPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/user`, req)
}

/**
 * @description "删除用户"
 * @param req
 */
export function userDel(req: components.UserDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/user`, req)
}

/**
 * @description "枚举列表"
 */
export function userEnumList() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/user/enumList`)
}

/**
 * @description "获取用户池列表"
 * @param params
 */
export function usersPoolList(params: components.UserPoolListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/usersPool/list`, params)
}

/**
 * @description "获取用户池信息"
 * @param params
 */
export function usersPoolInfo(params: components.UserPoolGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/usersPool`, params)
}

/**
 * @description "添加用户池"
 * @param req
 */
export function usersPoolAdd(req: components.UserPoolAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/usersPool`, req)
}

/**
 * @description "编辑用户池"
 * @param req
 */
export function usersPoolPut(req: components.UserPoolPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/usersPool`, req)
}

/**
 * @description "删除用户池"
 * @param req
 */
export function usersPoolDel(req: components.UserPoolDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/usersPool`, req)
}

/**
 * @description "管理员校验自己的密码"
 * @param req
 */
export function verifyPassword(req: components.VerifyPasswordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/account/verifyPassword`, req)
}

/**
 * @description 管理员修改自己的账号信息，当修改密码时，需要提供验证码
 * @param req
 */
export function putCurrent(req: components.PutCurrentReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/account/putCurrent`, req)
}

/**
 * @description "获取自己的账号信息"
 */
export function currentInfo() {
	return webapi.get<components.BaseResp>(`/v1/api/admin/account/currentInfo`)
}

/**
 * @description "修改自己的密码"
 * @param req
 */
export function putPassword(req: components.PutPasswordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/account/putPassword`, req)
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

/**
 * @description "获取商品列表"
 * @param params
 */
export function goodsList(params: components.GoodsListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goods/list`, params)
}

/**
 * @description "获取商品信息"
 * @param params
 */
export function goodsInfo(params: components.GoodsGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goods`, params)
}

/**
 * @description "添加商品"
 * @param req
 */
export function goodsAdd(req: components.GoodsAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/goods`, req)
}

/**
 * @description "编辑商品"
 * @param req
 */
export function goodsPut(req: components.GoodsPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/goods`, req)
}

/**
 * @description "删除商品"
 * @param req
 */
export function goodsDel(req: components.GoodsDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/goods`, req)
}

/**
 * @description "获取商品分类列表"
 * @param params
 */
export function goodsCategoryList(params: components.GoodsCategoryListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goodsCategory/list`, params)
}

/**
 * @description "获取商品分类信息"
 * @param params
 */
export function goodsCategoryInfo(params: components.GoodsCategoryGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goodsCategory`, params)
}

/**
 * @description "添加商品分类"
 * @param req
 */
export function goodsCategoryAdd(req: components.GoodsCategoryAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/goodsCategory`, req)
}

/**
 * @description "编辑商品分类"
 * @param req
 */
export function goodsCategoryPut(req: components.GoodsCategoryPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/goodsCategory`, req)
}

/**
 * @description "删除商品分类"
 * @param req
 */
export function goodsCategoryDel(req: components.GoodsCategoryDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/goodsCategory`, req)
}

/**
 * @description "获取商品标签列表"
 * @param params
 */
export function goodsTagList(params: components.GoodsTagListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goodsTag/list`, params)
}

/**
 * @description "获取商品标签信息"
 * @param params
 */
export function goodsTagInfo(params: components.GoodsTagGetReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/admin/goodsTag`, params)
}

/**
 * @description "添加商品标签"
 * @param req
 */
export function goodsTagAdd(req: components.GoodsTagAddReq) {
	return webapi.post<components.BaseResp>(`/v1/api/admin/goodsTag`, req)
}

/**
 * @description "编辑商品标签"
 * @param req
 */
export function goodsTagPut(req: components.GoodsTagPutReq) {
	return webapi.put<components.BaseResp>(`/v1/api/admin/goodsTag`, req)
}

/**
 * @description "删除商品标签"
 * @param req
 */
export function goodsTagDel(req: components.GoodsTagDelReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/admin/goodsTag`, req)
}
