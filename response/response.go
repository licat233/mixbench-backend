/*
 * @Author: licat
 * @Date: 2023-01-11 11:27:52
 * @LastEditors: licat
 * @LastEditTime: 2023-01-11 14:46:15
 * @Description: licat233@gmail.com
 */
package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	SILENT        = 0 // 不提示错误
	WARN_MESSAGE  = 1 // 警告信息提示
	ERROR_MESSAGE = 2 // 错误信息提示
	NOTIFICATION  = 4 // 通知提示
	REDIRECT      = 9 // 页面跳转
)

// 规范返回格式
type Body struct {
	//响应成功的基本格式
	Status bool `json:"success"` //【必填】响应状态
	// Message string      `json:"message,omitempty"` //【选填】自己加上去的，如果请求成功，该提示不要也罢
	Data interface{} `json:"data,omitempty"` //【选填】响应的业务数据
	//涉及列表请求
	Total    int64 `json:"total,omitempty"`    // 【选填】数据总个数
	PageSize int64 `json:"pageSize,omitempty"` // 【选填】单页数量
	Page     int64 `json:"current,omitempty"`  // 【选填】当前页码，current与antd前端对接
	// TotalPage int64 `json:"totalPage,omitempty"` // 【选填】自己加上去的，总共有多少页，根据前端的pageSize来计算： (Total + PageSize - 1) / PageSize
	//如果是失败响应，附带debug信息
	ErrorCode    int    `json:"errorCode,omitempty"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage,omitempty"` // 【选填】向用户显示消息
	// TraceMessage string `json:"traceMessage,omitempty"` // 【选填】自己加上去的，调试错误信息，请勿在生产环境下使用，可有可无
	ShowType int    `json:"showType,omitempty"` // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId  string `json:"traceId,omitempty"`  // 【选填】方便后端故障排除：唯一的请求ID
	Host     string `json:"host,omitempty"`     // 【选填】方便后端故障排除：当前访问服务器的主机
}

/**
 * @description: 响应体处理
 * @param {http.ResponseWriter} w
 * @param {interface{}} resp
 * @param {error} err 该错误可能为原生 error，或者自定义的 Errorx，以及rpc返回的 status.Error
 * @return {*}
 */
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		//错误处理
		body.Status = false
		body.ErrorCode = 400 //无从得知当前错误类型，该方案暂时摒弃
		body.ErrorMessage = err.Error()

	}

	httpx.OkJson(w, body)
}
