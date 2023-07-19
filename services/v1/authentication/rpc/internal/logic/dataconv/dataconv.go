package dataconv

import "mixbench/services/v1/authentication/rpc/authentication_pb"

func ListReqParams(req *authentication_pb.ListReq) (pageSize int64, page int64, keyword string) {
	if req != nil {
		pageSize = req.PageSize
		page = req.Page
		keyword = req.Keyword
	}
	return
}
