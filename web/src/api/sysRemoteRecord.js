import service from '@/utils/request'

// @Tags SysRemoteRecord
// @Summary 创建远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemoteRecord true "创建远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /remoteRecord/createSysRemoteRecord [post]
export const createSysRemoteRecord = (data) => {
  return service({
    url: '/remoteRecord/createSysRemoteRecord',
    method: 'post',
    data
  })
}

// @Tags SysRemoteRecord
// @Summary 删除远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemoteRecord true "删除远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remoteRecord/deleteSysRemoteRecord [delete]
export const deleteSysRemoteRecord = (data) => {
  return service({
    url: '/remoteRecord/deleteSysRemoteRecord',
    method: 'delete',
    data
  })
}

// @Tags SysRemoteRecord
// @Summary 批量删除远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remoteRecord/deleteSysRemoteRecord [delete]
export const deleteSysRemoteRecordByIds = (data) => {
  return service({
    url: '/remoteRecord/deleteSysRemoteRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags SysRemoteRecord
// @Summary 更新远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysRemoteRecord true "更新远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remoteRecord/updateSysRemoteRecord [put]
export const updateSysRemoteRecord = (data) => {
  return service({
    url: '/remoteRecord/updateSysRemoteRecord',
    method: 'put',
    data
  })
}

// @Tags SysRemoteRecord
// @Summary 用id查询远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysRemoteRecord true "用id查询远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remoteRecord/findSysRemoteRecord [get]
export const findSysRemoteRecord = (params) => {
  return service({
    url: '/remoteRecord/findSysRemoteRecord',
    method: 'get',
    params
  })
}

// @Tags SysRemoteRecord
// @Summary 分页获取远程操作记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取远程操作记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remoteRecord/getSysRemoteRecordList [get]
export const getSysRemoteRecordList = (params) => {
  return service({
    url: '/remoteRecord/getSysRemoteRecordList',
    method: 'get',
    params
  })
}
