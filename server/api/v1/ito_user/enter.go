package ito_user

import "github.com/flipped-aurora/gin-vue-admin/server/service/ito_user"

type ApiGroup struct {
	UsersApi *UsersApi
	LockerOrdersApi *LockerOrdersApi
	LockerPricingRulesApi *LockerPricingRulesApi
}

var (
	usersApi                  = new(UsersApi)
	lockerOrdersApi           = new(LockerOrdersApi)
	lockerPricingRulesApi     = new(LockerPricingRulesApi)

	usersService              = new(ito_user.UsersService)
	lockerOrdersService       = new(ito_user.LockerOrdersService)
	lockerPricingRulesService = new(ito_user.LockerPricingRulesService)
)
