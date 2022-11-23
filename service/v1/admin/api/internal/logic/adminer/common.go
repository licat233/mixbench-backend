package adminer

import (
	"context"
	"mixbench/common/ctxdata"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin"
)

func SuperAdminerIdentify(ctx context.Context, svcCtx *svc.ServiceContext) (bool, error) {
	uid := ctxdata.GetUidFromCtx(ctx)
	resp, err := svcCtx.AdminRpc.GetAdminer(ctx, &admin.ItemReq{
		Id: uid,
	})
	if err != nil {
		return false, err
	}
	return resp.Adminer.IsSuperAdmin == 1, nil
}
