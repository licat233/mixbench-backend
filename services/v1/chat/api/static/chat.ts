import webapi from "./gocliRequest"
import type * as components from "./chatComponents"
export * from "./chatComponents"

/**
 * @description "游客登录"
 */
export function visitorLogin() {
	return webapi.post<components.BaseResp>(`/v1/api/chat/public/visitorLogin`)
}

/**
 * @description "测试"
 * @param req
 */
export function test(req: components.TestReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/public/test`, req)
}

/**
 * @description "visitor游客获取自己的聊天记录"
 */
export function visitorMessages() {
	return webapi.get<components.BaseResp>(`/v1/api/chat/public/visitor/messages`)
}

/**
 * @description "visitor建立ws连接"
 * @param params
 */
export function visitorServer(params: components.ConnWebsocketReqParams) {
	return webapi.get<null>(`/v1/api/chat/public/ws/visitor`, params)
}

/**
 * @description "kefu建立ws连接"
 * @param params
 */
export function kefuServer(params: components.ConnWebsocketReqParams) {
	return webapi.get<null>(`/v1/api/chat/public/ws/kefu`, params)
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
 * @description "将游客转接给其他客服"
 * @param req
 */
export function transferVisitor(req: components.TransferVisitorReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/transferVisitor`, req)
}

/**
 * @description "获取某个游客的最后一条消息"
 * @param req
 */
export function visitorLastMessage(req: components.VisitorLastMessageReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/visitorLastMessage`, req)
}

/**
 * @description "获取某个游客的全部消息"
 * @param req
 */
export function getVisitorMessages(req: components.GetVisitorMessagesReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/visitorMessages`, req)
}

/**
 * @description "更改某个游客的备注"
 * @param req
 */
export function putVisitorComment(req: components.PutVisitorCommentReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/putVisitorComment`, req)
}

/**
 * @description "添加消息 基础API"
 * @param req
 */
export function addMessage(req: components.AddMessageReq) {
	return webapi.post<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "更新消息 基础API"
 * @param req
 */
export function putMessage(req: components.PutMessageReq) {
	return webapi.put<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "获取消息 基础API"
 * @param params
 */
export function getMessage(params: components.GetMessageReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message`, params)
}

/**
 * @description "删除消息 基础API"
 * @param req
 */
export function delMessage(req: components.DelMessageReq) {
	return webapi.delete<components.BaseResp>(`/v1/api/chat/message`, req)
}

/**
 * @description "获取消息列表 基础API"
 * @param params
 */
export function getMessageList(params: components.GetMessageListReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message/list`, params)
}

/**
 * @description "获取消息枚举列表 基础API"
 * @param req
 */
export function getMessageEnums(req: components.GetMessageEnumsReq) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/message/enums`, req)
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
 * @param params
 */
export function getMessageConfig(params: components.GetMessageConfigReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/messageConfig`, params)
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
 * @param params
 */
export function getVisitor(params: components.GetVisitorReqParams) {
	return webapi.get<components.BaseResp>(`/v1/api/chat/visitor`, params)
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
