package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CampaignModel = (*customCampaignModel)(nil)

type (
	// CampaignModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCampaignModel.
	CampaignModel interface {
		campaignModel  // extended interface by gozero
		campaign_model // extended interface by genzero
	}

	customCampaignModel struct {
		*defaultCampaignModel
	}
)

// NewCampaignModel returns a model for the database table.
func NewCampaignModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CampaignModel {
	cacheCampaignIdPrefix = "cache:admin:campaign:id:" // modifying cache id prefix by genzero
	return &customCampaignModel{
		defaultCampaignModel: newCampaignModel(conn, c, opts...),
	}
}
