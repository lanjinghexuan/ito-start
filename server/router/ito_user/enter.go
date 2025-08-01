package ito_user

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	ito_user_api "github.com/flipped-aurora/gin-vue-admin/server/api/v1/ito_user"
)

type RouterGroup struct {
	UsersRouter
	LockerOrdersRouter
	LockerPricingRulesRouter
}

var (
	usersApi              *ito_user_api.UsersApi
	lockerOrdersApi       *ito_user_api.LockerOrdersApi
	lockerPricingRulesApi *ito_user_api.LockerPricingRulesApi
)

func Init() {
	usersApi = api.ApiGroupApp.Ito_userApiGroup.UsersApi
	lockerOrdersApi = api.ApiGroupApp.Ito_userApiGroup.LockerOrdersApi
	lockerPricingRulesApi = api.ApiGroupApp.Ito_userApiGroup.LockerPricingRulesApi
}
