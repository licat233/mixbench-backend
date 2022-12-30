/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:48:35
 * @Description: licat233@gmail.com
 */
package rbac

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPoliciesLogic {
	return &AddPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPoliciesLogic) AddPolicies(req *types.PoliciesReq) (any, error) {
	rules := make([][]string, 0)
	for _, v := range req.Rules {
		rules = append(rules, utils.StructToSlice(v))
	}
	ok, err := l.svcCtx.CasbinEnforcer.AddPolicies(rules)
	if logState && err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultBoolResp(ok, errors.New("rule already exists"))
}
