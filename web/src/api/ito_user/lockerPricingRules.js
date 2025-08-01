import service from '@/utils/request'
// @Tags LockerPricingRules
// @Summary 创建lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPricingRules true "创建lockerPricingRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /lockerPricingRules/createLockerPricingRules [post]
export const createLockerPricingRules = (data) => {
  return service({
    url: '/lockerPricingRules/createLockerPricingRules',
    method: 'post',
    data
  })
}

// @Tags LockerPricingRules
// @Summary 删除lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPricingRules true "删除lockerPricingRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerPricingRules/deleteLockerPricingRules [delete]
export const deleteLockerPricingRules = (params) => {
  return service({
    url: '/lockerPricingRules/deleteLockerPricingRules',
    method: 'delete',
    params
  })
}

// @Tags LockerPricingRules
// @Summary 批量删除lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除lockerPricingRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerPricingRules/deleteLockerPricingRules [delete]
export const deleteLockerPricingRulesByIds = (params) => {
  return service({
    url: '/lockerPricingRules/deleteLockerPricingRulesByIds',
    method: 'delete',
    params
  })
}

// @Tags LockerPricingRules
// @Summary 更新lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPricingRules true "更新lockerPricingRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lockerPricingRules/updateLockerPricingRules [put]
export const updateLockerPricingRules = (data) => {
  return service({
    url: '/lockerPricingRules/updateLockerPricingRules',
    method: 'put',
    data
  })
}

// @Tags LockerPricingRules
// @Summary 用id查询lockerPricingRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.LockerPricingRules true "用id查询lockerPricingRules表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lockerPricingRules/findLockerPricingRules [get]
export const findLockerPricingRules = (params) => {
  return service({
    url: '/lockerPricingRules/findLockerPricingRules',
    method: 'get',
    params
  })
}

// @Tags LockerPricingRules
// @Summary 分页获取lockerPricingRules表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取lockerPricingRules表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lockerPricingRules/getLockerPricingRulesList [get]
export const getLockerPricingRulesList = (params) => {
  return service({
    url: '/lockerPricingRules/getLockerPricingRulesList',
    method: 'get',
    params
  })
}

// @Tags LockerPricingRules
// @Summary 不需要鉴权的lockerPricingRules表接口
// @Accept application/json
// @Produce application/json
// @Param data query ito_userReq.LockerPricingRulesSearch true "分页获取lockerPricingRules表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerPricingRules/getLockerPricingRulesPublic [get]
export const getLockerPricingRulesPublic = () => {
  return service({
    url: '/lockerPricingRules/getLockerPricingRulesPublic',
    method: 'get',
  })
}
