package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(ito_user.Users{}, ito_user.LockerOrders{}, ito_user.LockerPricingRules{}, ito_user.LockerOrders{})
	if err != nil {
		return err
	}
	return nil
}
