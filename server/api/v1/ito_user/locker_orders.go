package ito_user

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ito_user"
    ito_userReq "github.com/flipped-aurora/gin-vue-admin/server/model/ito_user/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type LockerOrdersApi struct {}



// CreateLockerOrders 创建lockerOrders表
// @Tags LockerOrders
// @Summary 创建lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerOrders true "创建lockerOrders表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /lockerOrders/createLockerOrders [post]
func (lockerOrdersApi *LockerOrdersApi) CreateLockerOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var lockerOrders ito_user.LockerOrders
	err := c.ShouldBindJSON(&lockerOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    lockerOrders.CreatedBy = utils.GetUserID(c)
	err = lockerOrdersService.CreateLockerOrders(ctx,&lockerOrders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteLockerOrders 删除lockerOrders表
// @Tags LockerOrders
// @Summary 删除lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerOrders true "删除lockerOrders表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /lockerOrders/deleteLockerOrders [delete]
func (lockerOrdersApi *LockerOrdersApi) DeleteLockerOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := lockerOrdersService.DeleteLockerOrders(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLockerOrdersByIds 批量删除lockerOrders表
// @Tags LockerOrders
// @Summary 批量删除lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /lockerOrders/deleteLockerOrdersByIds [delete]
func (lockerOrdersApi *LockerOrdersApi) DeleteLockerOrdersByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := lockerOrdersService.DeleteLockerOrdersByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLockerOrders 更新lockerOrders表
// @Tags LockerOrders
// @Summary 更新lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerOrders true "更新lockerOrders表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /lockerOrders/updateLockerOrders [put]
func (lockerOrdersApi *LockerOrdersApi) UpdateLockerOrders(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var lockerOrders ito_user.LockerOrders
	err := c.ShouldBindJSON(&lockerOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    lockerOrders.UpdatedBy = utils.GetUserID(c)
	err = lockerOrdersService.UpdateLockerOrders(ctx,lockerOrders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLockerOrders 用id查询lockerOrders表
// @Tags LockerOrders
// @Summary 用id查询lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询lockerOrders表"
// @Success 200 {object} response.Response{data=ito_user.LockerOrders,msg=string} "查询成功"
// @Router /lockerOrders/findLockerOrders [get]
func (lockerOrdersApi *LockerOrdersApi) FindLockerOrders(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	relockerOrders, err := lockerOrdersService.GetLockerOrders(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(relockerOrders, c)
}
// GetLockerOrdersList 分页获取lockerOrders表列表
// @Tags LockerOrders
// @Summary 分页获取lockerOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ito_userReq.LockerOrdersSearch true "分页获取lockerOrders表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /lockerOrders/getLockerOrdersList [get]
func (lockerOrdersApi *LockerOrdersApi) GetLockerOrdersList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ito_userReq.LockerOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lockerOrdersService.GetLockerOrdersInfoList(ctx,pageInfo)
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

// GetLockerOrdersPublic 不需要鉴权的lockerOrders表接口
// @Tags LockerOrders
// @Summary 不需要鉴权的lockerOrders表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerOrders/getLockerOrdersPublic [get]
func (lockerOrdersApi *LockerOrdersApi) GetLockerOrdersPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    lockerOrdersService.GetLockerOrdersPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的lockerOrders表接口信息",
    }, "获取成功", c)
}
