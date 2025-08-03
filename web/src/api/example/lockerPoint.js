import service from '@/utils/request'
// @Tags LockerPoint
// @Summary 创建网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPoint true "创建网点管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /lockerPoint/createLockerPoint [post]
export const createLockerPoint = (data) => {
  return service({
    url: '/lockerPoint/createLockerPoint',
    method: 'post',
    data
  })
}

// @Tags LockerPoint
// @Summary 删除网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPoint true "删除网点管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerPoint/deleteLockerPoint [delete]
export const deleteLockerPoint = (params) => {
  return service({
    url: '/lockerPoint/deleteLockerPoint',
    method: 'delete',
    params
  })
}

// @Tags LockerPoint
// @Summary 批量删除网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网点管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lockerPoint/deleteLockerPoint [delete]
export const deleteLockerPointByIds = (params) => {
  return service({
    url: '/lockerPoint/deleteLockerPointByIds',
    method: 'delete',
    params
  })
}

// @Tags LockerPoint
// @Summary 更新网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.LockerPoint true "更新网点管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lockerPoint/updateLockerPoint [put]
export const updateLockerPoint = (data) => {
  return service({
    url: '/lockerPoint/updateLockerPoint',
    method: 'put',
    data
  })
}

// @Tags LockerPoint
// @Summary 用id查询网点管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.LockerPoint true "用id查询网点管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lockerPoint/findLockerPoint [get]
export const findLockerPoint = (params) => {
  return service({
    url: '/lockerPoint/findLockerPoint',
    method: 'get',
    params
  })
}

// @Tags LockerPoint
// @Summary 分页获取网点管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取网点管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lockerPoint/getLockerPointList [get]
export const getLockerPointList = (params) => {
  return service({
    url: '/lockerPoint/getLockerPointList',
    method: 'get',
    params
  })
}

// @Tags LockerPoint
// @Summary 不需要鉴权的网点管理接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.LockerPointSearch true "分页获取网点管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /lockerPoint/getLockerPointPublic [get]
export const getLockerPointPublic = () => {
  return service({
    url: '/lockerPoint/getLockerPointPublic',
    method: 'get',
  })
}
