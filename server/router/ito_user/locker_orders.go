package ito_user

import (
	"github.com/gin-gonic/gin"
)

type LockerOrdersRouter struct{}

// InitLockerOrdersRouter 初始化 lockerOrders表 路由信息
func (s *LockerOrdersRouter) InitLockerOrdersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//lockerOrdersRouter := Router.Group("lockerOrders").Use(middleware.OperationRecord())
	//lockerOrdersRouterWithoutRecord := Router.Group("lockerOrders")
	//lockerOrdersRouterWithoutAuth := PublicRouter.Group("lockerOrders")
	//{
	//	lockerOrdersRouter.POST("createLockerOrders", lockerOrdersApi.CreateLockerOrders)   // 新建lockerOrders表
	//	lockerOrdersRouter.DELETE("deleteLockerOrders", lockerOrdersApi.DeleteLockerOrders) // 删除lockerOrders表
	//	lockerOrdersRouter.DELETE("deleteLockerOrdersByIds", lockerOrdersApi.DeleteLockerOrdersByIds) // 批量删除lockerOrders表
	//	lockerOrdersRouter.PUT("updateLockerOrders", lockerOrdersApi.UpdateLockerOrders)    // 更新lockerOrders表
	//}
	//{
	//	lockerOrdersRouterWithoutRecord.GET("findLockerOrders", lockerOrdersApi.FindLockerOrders)        // 根据ID获取lockerOrders表
	//	lockerOrdersRouterWithoutRecord.GET("getLockerOrdersList", lockerOrdersApi.GetLockerOrdersList)  // 获取lockerOrders表列表
	//}
	//{
	//    lockerOrdersRouterWithoutAuth.GET("getLockerOrdersPublic", lockerOrdersApi.GetLockerOrdersPublic)  // lockerOrders表开放接口
	//}
}
