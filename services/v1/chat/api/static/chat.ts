import webapi from "./gocliRequest"
import * as components from "./chatComponents"
export * from "./chatComponents"

/**
 * @description "游客登录"
 * @param req
 */
export function visitorLogin(req: components.VisitorLoginReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/visitorLogin`, req)
}

/**
 * @description "客服登录"
 * @param req
 */
export function kefuLogin(req: components.KefuLoginReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/kefuLogin`, req)
}

/**
 * @description "测试"
 * @param req
 */
export function test(req: components.TestReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/test`, req)
}

/**
 * @description "visitor建立ws连接"
 */
export function visitorServer() {
	return webapi.get<null>(`/v1/api/chat/wsVisitor`)
}

/**
 * @description "保存消息配置"
 * @param req
 */
export function saveMsgConf(req: components.SaveMsgConfReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/saveMsgConf`, req)
}

/**
 * @description "获取某客服的游客列表"
 * @param params
 */
export function visitorList(params: components.VisitorListReqParams) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/visitorList`, params)
}

/**
 * @description "kefu建立ws连接"
 */
export function kefuServer() {
	return webapi.get<null>(`/v1/api/chat/wsKefu`)
}

/**
 * @description "添加游客 基础API"
 * @param req
 */
export function addVisitor(req: components.AddVisitorReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/visitor`, req)
}

/**
 * @description "更新游客 基础API"
 * @param req
 */
export function putVisitor(req: components.PutVisitorReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/visitor`, req)
}

/**
 * @description "获取游客 基础API"
 * @param req
 */
export function getVisitor(req: components.GetVisitorReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/visitor`, req)
}

/**
 * @description "删除游客 基础API"
 * @param req
 */
export function delVisitor(req: components.DelVisitorReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/visitor`, req)
}

/**
 * @description "获取游客列表 基础API"
 * @param params
 */
export function getVisitorList(params: components.GetVisitorListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/visitor/list`, params)
}

/**
 * @description "获取游客枚举列表 基础API"
 * @param req
 */
export function getVisitorEnums(req: components.GetVisitorEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/visitor/enums`, req)
}

/**
 * @description "添加客服 基础API"
 * @param req
 */
export function addKefu(req: components.AddKefuReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/kefu`, req)
}

/**
 * @description "更新客服 基础API"
 * @param req
 */
export function putKefu(req: components.PutKefuReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/kefu`, req)
}

/**
 * @description "获取客服 基础API"
 * @param req
 */
export function getKefu(req: components.GetKefuReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/kefu`, req)
}

/**
 * @description "删除客服 基础API"
 * @param req
 */
export function delKefu(req: components.DelKefuReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/kefu`, req)
}

/**
 * @description "获取客服列表 基础API"
 * @param params
 */
export function getKefuList(params: components.GetKefuListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/kefu/list`, params)
}

/**
 * @description "获取客服枚举列表 基础API"
 * @param req
 */
export function getKefuEnums(req: components.GetKefuEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/kefu/enums`, req)
}

/**
 * @description "添加聊天记录 基础API"
 * @param req
 */
export function addChatRecord(req: components.AddChatRecordReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/chatRecord`, req)
}

/**
 * @description "更新聊天记录 基础API"
 * @param req
 */
export function putChatRecord(req: components.PutChatRecordReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/chatRecord`, req)
}

/**
 * @description "获取聊天记录 基础API"
 * @param req
 */
export function getChatRecord(req: components.GetChatRecordReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/chatRecord`, req)
}

/**
 * @description "删除聊天记录 基础API"
 * @param req
 */
export function delChatRecord(req: components.DelChatRecordReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/chatRecord`, req)
}

/**
 * @description "获取聊天记录列表 基础API"
 * @param params
 */
export function getChatRecordList(params: components.GetChatRecordListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/chatRecord/list`, params)
}

/**
 * @description "获取聊天记录枚举列表 基础API"
 * @param req
 */
export function getChatRecordEnums(req: components.GetChatRecordEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/chatRecord/enums`, req)
}

/**
 * @description "添加消息, 主要用于缓存未发送成功的消息 基础API"
 * @param req
 */
export function addMessage(req: components.AddMessageReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "更新消息, 主要用于缓存未发送成功的消息 基础API"
 * @param req
 */
export function putMessage(req: components.PutMessageReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "获取消息, 主要用于缓存未发送成功的消息 基础API"
 * @param req
 */
export function getMessage(req: components.GetMessageReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "删除消息, 主要用于缓存未发送成功的消息 基础API"
 * @param req
 */
export function delMessage(req: components.DelMessageReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "获取消息, 主要用于缓存未发送成功的消息列表 基础API"
 * @param params
 */
export function getMessageList(params: components.GetMessageListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message/list`, params)
}

/**
 * @description "获取消息, 主要用于缓存未发送成功的消息枚举列表 基础API"
 * @param req
 */
export function getMessageEnums(req: components.GetMessageEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message/enums`, req)
}

/**
 * @description "添加IP黑名单 基础API"
 * @param req
 */
export function addIpblack(req: components.AddIpblackReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/ipblack`, req)
}

/**
 * @description "更新IP黑名单 基础API"
 * @param req
 */
export function putIpblack(req: components.PutIpblackReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/ipblack`, req)
}

/**
 * @description "获取IP黑名单 基础API"
 * @param req
 */
export function getIpblack(req: components.GetIpblackReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/ipblack`, req)
}

/**
 * @description "删除IP黑名单 基础API"
 * @param req
 */
export function delIpblack(req: components.DelIpblackReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/ipblack`, req)
}

/**
 * @description "获取IP黑名单列表 基础API"
 * @param params
 */
export function getIpblackList(params: components.GetIpblackListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/ipblack/list`, params)
}

/**
 * @description "获取IP黑名单枚举列表 基础API"
 * @param req
 */
export function getIpblackEnums(req: components.GetIpblackEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/ipblack/enums`, req)
}

/**
 * @description "添加消息配置 基础API"
 * @param req
 */
export function addMessageConfig(req: components.AddMessageConfigReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/messageConfig`, req)
}

/**
 * @description "更新消息配置 基础API"
 * @param req
 */
export function putMessageConfig(req: components.PutMessageConfigReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/messageConfig`, req)
}

/**
 * @description "获取消息配置 基础API"
 * @param req
 */
export function getMessageConfig(req: components.GetMessageConfigReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/messageConfig`, req)
}

/**
 * @description "删除消息配置 基础API"
 * @param req
 */
export function delMessageConfig(req: components.DelMessageConfigReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/messageConfig`, req)
}

/**
 * @description "获取消息配置列表 基础API"
 * @param params
 */
export function getMessageConfigList(params: components.GetMessageConfigListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/messageConfig/list`, params)
}

/**
 * @description "获取消息配置枚举列表 基础API"
 * @param req
 */
export function getMessageConfigEnums(req: components.GetMessageConfigEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/messageConfig/enums`, req)
}
