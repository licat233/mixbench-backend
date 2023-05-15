/*
 * @Author: licat
 * @Date: 2023-02-18 19:01:16
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:24:20
 * @Description: licat233@gmail.com
 */
package handler

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx, r)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
