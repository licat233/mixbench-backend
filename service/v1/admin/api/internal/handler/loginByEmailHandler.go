/*
 * @Author: licat
 * @Date: 2023-02-18 19:01:16
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 11:22:22
 * @Description: licat233@gmail.com
 */
package handler

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginByEmailLogic(r.Context(), svcCtx, r)
		resp, err := l.LoginByEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
