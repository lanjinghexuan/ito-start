package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ito_user"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

func Init() {
	ApiGroupApp.Ito_userApiGroup = new(ito_user.ApiGroup)
	ApiGroupApp.Ito_userApiGroup.UsersApi = new(ito_user.UsersApi)
	ApiGroupApp.Ito_userApiGroup.LockerOrdersApi = new(ito_user.LockerOrdersApi)
	ApiGroupApp.Ito_userApiGroup.LockerPricingRulesApi = new(ito_user.LockerPricingRulesApi)
}

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	Ito_userApiGroup *ito_user.ApiGroup
}
