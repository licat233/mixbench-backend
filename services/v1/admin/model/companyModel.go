package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CompanyModel = (*customCompanyModel)(nil)

type (
	// CompanyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCompanyModel.
	CompanyModel interface {
		companyModel  // extended interface by gozero
		company_model // extended interface by genzero
	}

	customCompanyModel struct {
		*defaultCompanyModel
	}
)

// NewCompanyModel returns a model for the database table.
func NewCompanyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CompanyModel {
	cacheCompanyIdPrefix = "cache:admin:company:id:" // modifying cache id prefix by genzero
	return &customCompanyModel{
		defaultCompanyModel: newCompanyModel(conn, c, opts...),
	}
}
