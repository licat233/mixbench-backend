package rbacx

import (
	"database/sql"
	"mixbench/common/utils"
	"path"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewCasbinEnforcer(db *sql.DB) *casbin.SyncedEnforcer {
	if db == nil {
		return nil
	}
	casbinFile := path.Join(utils.ProjectRoot, "deploy/casbin/rbac_with_domains_model.conf")
	adapter, err := sqladapter.NewAdapter(db, "mysql", "casbin_rule")
	if err != nil {
		logx.Errorf("新建Casbin Adapter失败:%v", err)
		return nil
	}

	syncedEnforcer, err := casbin.NewSyncedEnforcer(casbinFile, adapter)
	if err != nil {
		logx.Errorf("新建Casbin Enforcer失败:%v", err)
		return nil
	}

	// Load the policy from DB.
	if err = syncedEnforcer.LoadPolicy(); err != nil {
		logx.Errorf("加载Casbin Policy失败:%v", err)
		return nil
	}

	return syncedEnforcer
}
