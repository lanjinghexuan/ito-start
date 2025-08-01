package ito_user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LockerPricingRulesRouter struct {}

// InitLockerPricingRulesRouter 初始化 lockerPricingRules表 路由信息
func (s *LockerPricingRulesRouter) InitLockerPricingRulesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	lockerPricingRulesRouter := Router.Group("lockerPricingRules").Use(middleware.OperationRecord())
	lockerPricingRulesRouterWithoutRecord := Router.Group("lockerPricingRules")
	lockerPricingRulesRouterWithoutAuth := PublicRouter.Group("lockerPricingRules")
	{
		lockerPricingRulesRouter.POST("createLockerPricingRules", lockerPricingRulesApi.CreateLockerPricingRules)   // 新建lockerPricingRules表
		lockerPricingRulesRouter.DELETE("deleteLockerPricingRules", lockerPricingRulesApi.DeleteLockerPricingRules) // 删除lockerPricingRules表
		lockerPricingRulesRouter.DELETE("deleteLockerPricingRulesByIds", lockerPricingRulesApi.DeleteLockerPricingRulesByIds) // 批量删除lockerPricingRules表
		lockerPricingRulesRouter.PUT("updateLockerPricingRules", lockerPricingRulesApi.UpdateLockerPricingRules)    // 更新lockerPricingRules表
	}
	{
		lockerPricingRulesRouterWithoutRecord.GET("findLockerPricingRules", lockerPricingRulesApi.FindLockerPricingRules)        // 根据ID获取lockerPricingRules表
		lockerPricingRulesRouterWithoutRecord.GET("getLockerPricingRulesList", lockerPricingRulesApi.GetLockerPricingRulesList)  // 获取lockerPricingRules表列表
	}
	{
	    lockerPricingRulesRouterWithoutAuth.GET("getLockerPricingRulesPublic", lockerPricingRulesApi.GetLockerPricingRulesPublic)  // lockerPricingRules表开放接口
	}
}
