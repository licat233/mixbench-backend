/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:03
 * @LastEditors: licat
 * @LastEditTime: 2023-01-11 09:47:24
 * @Description: licat233@gmail.com
 */
package respx

/** 与前端约定的数据规范
interface ErrorInfoStructure {
	success: boolean; // 请求状态
	data?: any; // 返回数据
	total?: number; // 数据总个数
	pageSize: number //单页数量
	current?: number; // 当前页码
	page?: number; // 【自己加上去的】总共有多少页，根据前端的pageSize来计算
}
**/

/**
* 前端要求的返回格式规范
  const result = {
   data: dataSource,
   total: tableListDataSource.length,
   success: true,
   pageSize: finalPageSize,
   current: parseInt(`${params.currentPage}`, 10) || 1,
 };
*/

// 规范返回格式
type BaseResp struct {
	//响应成功的基本格式
	Status  bool        `json:"success"`           //【响应状态】
	Message string      `json:"message,omitempty"` //【给予的提示信息】
	Data    interface{} `json:"data,omitempty"`    //【选填】响应的业务数据
	//涉及列表请求
	Total     int64 `json:"total,omitempty"`     // 【选填】数据总个数
	PageSize  int64 `json:"pageSize,omitempty"`  // 【选填】单页数量
	Page      int64 `json:"current,omitempty"`   // 【选填】当前页码，current与antd前端对接
	TotalPage int64 `json:"totalPage,omitempty"` // 【自己加上去的】总共有多少页，根据前端的pageSize来计算
	//如果是失败响应，附带debug信息
	ErrorCode    int    `json:"errorCode,omitempty"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage,omitempty"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage,omitempty"` // 【选填】自己加上去的，调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType,omitempty"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId,omitempty"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host,omitempty"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}
