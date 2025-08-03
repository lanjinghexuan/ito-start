package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(ito_user.Users{}, example.LockerPoint{}, ito_user.LockerOrders{}, ito_user.LockerPricingRules{})
	if err != nil {
		return err
	}
	return nil
}
