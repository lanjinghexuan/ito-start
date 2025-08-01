import service from '@/utils/request'
// @Tags LockerOrders
// @Summary 创建lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerOrders true "创建lockerOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /lockerOrders/createLockerOrders [post]
export const createLockerOrders = (data) => {
  return service({
    url: '/lockerOrders/createLockerOrders',
    method: 'post',
    data
  })
}

// @Tags LockerOrders
// @Summary 删除lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerOrders true "删除lockerOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerOrders/deleteLockerOrders [delete]
export const deleteLockerOrders = (params) => {
  return service({
    url: '/lockerOrders/deleteLockerOrders',
    method: 'delete',
    params
  })
}

// @Tags LockerOrders
// @Summary 批量删除lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除lockerOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerOrders/deleteLockerOrders [delete]
export const deleteLockerOrdersByIds = (params) => {
  return service({
    url: '/lockerOrders/deleteLockerOrdersByIds',
    method: 'delete',
    params
  })
}

// @Tags LockerOrders
// @Summary 更新lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerOrders true "更新lockerOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lockerOrders/updateLockerOrders [put]
export const updateLockerOrders = (data) => {
  return service({
    url: '/lockerOrders/updateLockerOrders',
    method: 'put',
    data
  })
}

// @Tags LockerOrders
// @Summary 用id查询lockerOrders表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.LockerOrders true "用id查询lockerOrders表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lockerOrders/findLockerOrders [get]
export const findLockerOrders = (params) => {
  return service({
    url: '/lockerOrders/findLockerOrders',
    method: 'get',
    params
  })
}

// @Tags LockerOrders
// @Summary 分页获取lockerOrders表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取lockerOrders表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lockerOrders/getLockerOrdersList [get]
export const getLockerOrdersList = (params) => {
  return service({
    url: '/lockerOrders/getLockerOrdersList',
    method: 'get',
    params
  })
}

// @Tags LockerOrders
// @Summary 不需要鉴权的lockerOrders表接口
// @Accept application/json
// @Produce application/json
// @Param data query ito_userReq.LockerOrdersSearch true "分页获取lockerOrders表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerOrders/getLockerOrdersPublic [get]
export const getLockerOrdersPublic = () => {
  return service({
    url: '/lockerOrders/getLockerOrdersPublic',
    method: 'get',
  })
}
