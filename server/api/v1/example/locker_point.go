package example

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type LockerPointApi struct {}



// CreateLockerPoint 创建网点管理
// @Tags LockerPoint
// @Summary 创建网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.LockerPoint true "创建网点管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /lockerPoint/createLockerPoint [post]
func (lockerPointApi *LockerPointApi) CreateLockerPoint(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var lockerPoint example.LockerPoint
	err := c.ShouldBindJSON(&lockerPoint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = lockerPointService.CreateLockerPoint(ctx,&lockerPoint)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteLockerPoint 删除网点管理
// @Tags LockerPoint
// @Summary 删除网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.LockerPoint true "删除网点管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /lockerPoint/deleteLockerPoint [delete]
func (lockerPointApi *LockerPointApi) DeleteLockerPoint(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := lockerPointService.DeleteLockerPoint(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLockerPointByIds 批量删除网点管理
// @Tags LockerPoint
// @Summary 批量删除网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /lockerPoint/deleteLockerPointByIds [delete]
func (lockerPointApi *LockerPointApi) DeleteLockerPointByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := lockerPointService.DeleteLockerPointByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLockerPoint 更新网点管理
// @Tags LockerPoint
// @Summary 更新网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.LockerPoint true "更新网点管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /lockerPoint/updateLockerPoint [put]
func (lockerPointApi *LockerPointApi) UpdateLockerPoint(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var lockerPoint example.LockerPoint
	err := c.ShouldBindJSON(&lockerPoint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = lockerPointService.UpdateLockerPoint(ctx,lockerPoint)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLockerPoint 用id查询网点管理
// @Tags LockerPoint
// @Summary 用id查询网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询网点管理"
// @Success 200 {object} response.Response{data=example.LockerPoint,msg=string} "查询成功"
// @Router /lockerPoint/findLockerPoint [get]
func (lockerPointApi *LockerPointApi) FindLockerPoint(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	relockerPoint, err := lockerPointService.GetLockerPoint(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(relockerPoint, c)
}
// GetLockerPointList 分页获取网点管理列表
// @Tags LockerPoint
// @Summary 分页获取网点管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.LockerPointSearch true "分页获取网点管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /lockerPoint/getLockerPointList [get]
func (lockerPointApi *LockerPointApi) GetLockerPointList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo exampleReq.LockerPointSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lockerPointService.GetLockerPointInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetLockerPointPublic 不需要鉴权的网点管理接口
// @Tags LockerPoint
// @Summary 不需要鉴权的网点管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerPoint/getLockerPointPublic [get]
func (lockerPointApi *LockerPointApi) GetLockerPointPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    lockerPointService.GetLockerPointPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的网点管理接口信息",
    }, "获取成功", c)
}
