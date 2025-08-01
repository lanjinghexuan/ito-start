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

type LockerPricingRulesApi struct {}



// CreateLockerPricingRules 创建lockerPricingRules表
// @Tags LockerPricingRules
// @Summary 创建lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerPricingRules true "创建lockerPricingRules表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /lockerPricingRules/createLockerPricingRules [post]
func (lockerPricingRulesApi *LockerPricingRulesApi) CreateLockerPricingRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var lockerPricingRules ito_user.LockerPricingRules
	err := c.ShouldBindJSON(&lockerPricingRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    lockerPricingRules.CreatedBy = utils.GetUserID(c)
	err = lockerPricingRulesService.CreateLockerPricingRules(ctx,&lockerPricingRules)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteLockerPricingRules 删除lockerPricingRules表
// @Tags LockerPricingRules
// @Summary 删除lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerPricingRules true "删除lockerPricingRules表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /lockerPricingRules/deleteLockerPricingRules [delete]
func (lockerPricingRulesApi *LockerPricingRulesApi) DeleteLockerPricingRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := lockerPricingRulesService.DeleteLockerPricingRules(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteLockerPricingRulesByIds 批量删除lockerPricingRules表
// @Tags LockerPricingRules
// @Summary 批量删除lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /lockerPricingRules/deleteLockerPricingRulesByIds [delete]
func (lockerPricingRulesApi *LockerPricingRulesApi) DeleteLockerPricingRulesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := lockerPricingRulesService.DeleteLockerPricingRulesByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateLockerPricingRules 更新lockerPricingRules表
// @Tags LockerPricingRules
// @Summary 更新lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ito_user.LockerPricingRules true "更新lockerPricingRules表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /lockerPricingRules/updateLockerPricingRules [put]
func (lockerPricingRulesApi *LockerPricingRulesApi) UpdateLockerPricingRules(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var lockerPricingRules ito_user.LockerPricingRules
	err := c.ShouldBindJSON(&lockerPricingRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    lockerPricingRules.UpdatedBy = utils.GetUserID(c)
	err = lockerPricingRulesService.UpdateLockerPricingRules(ctx,lockerPricingRules)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindLockerPricingRules 用id查询lockerPricingRules表
// @Tags LockerPricingRules
// @Summary 用id查询lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询lockerPricingRules表"
// @Success 200 {object} response.Response{data=ito_user.LockerPricingRules,msg=string} "查询成功"
// @Router /lockerPricingRules/findLockerPricingRules [get]
func (lockerPricingRulesApi *LockerPricingRulesApi) FindLockerPricingRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	relockerPricingRules, err := lockerPricingRulesService.GetLockerPricingRules(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(relockerPricingRules, c)
}
// GetLockerPricingRulesList 分页获取lockerPricingRules表列表
// @Tags LockerPricingRules
// @Summary 分页获取lockerPricingRules表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query ito_userReq.LockerPricingRulesSearch true "分页获取lockerPricingRules表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /lockerPricingRules/getLockerPricingRulesList [get]
func (lockerPricingRulesApi *LockerPricingRulesApi) GetLockerPricingRulesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo ito_userReq.LockerPricingRulesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lockerPricingRulesService.GetLockerPricingRulesInfoList(ctx,pageInfo)
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

// GetLockerPricingRulesPublic 不需要鉴权的lockerPricingRules表接口
// @Tags LockerPricingRules
// @Summary 不需要鉴权的lockerPricingRules表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerPricingRules/getLockerPricingRulesPublic [get]
func (lockerPricingRulesApi *LockerPricingRulesApi) GetLockerPricingRulesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    lockerPricingRulesService.GetLockerPricingRulesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的lockerPricingRules表接口信息",
    }, "获取成功", c)
}
