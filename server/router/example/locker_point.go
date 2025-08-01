package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LockerPointRouter struct {}

// InitLockerPointRouter 初始化 网点管理 路由信息
func (s *LockerPointRouter) InitLockerPointRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	lockerPointRouter := Router.Group("lockerPoint").Use(middleware.OperationRecord())
	lockerPointRouterWithoutRecord := Router.Group("lockerPoint")
	lockerPointRouterWithoutAuth := PublicRouter.Group("lockerPoint")
	{
		lockerPointRouter.POST("createLockerPoint", lockerPointApi.CreateLockerPoint)   // 新建网点管理
		lockerPointRouter.DELETE("deleteLockerPoint", lockerPointApi.DeleteLockerPoint) // 删除网点管理
		lockerPointRouter.DELETE("deleteLockerPointByIds", lockerPointApi.DeleteLockerPointByIds) // 批量删除网点管理
		lockerPointRouter.PUT("updateLockerPoint", lockerPointApi.UpdateLockerPoint)    // 更新网点管理
	}
	{
		lockerPointRouterWithoutRecord.GET("findLockerPoint", lockerPointApi.FindLockerPoint)        // 根据ID获取网点管理
		lockerPointRouterWithoutRecord.GET("getLockerPointList", lockerPointApi.GetLockerPointList)  // 获取网点管理列表
	}
	{
	    lockerPointRouterWithoutAuth.GET("getLockerPointPublic", lockerPointApi.GetLockerPointPublic)  // 网点管理开放接口
	}
}
