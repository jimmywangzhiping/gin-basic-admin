package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    remoteServerReq "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SysRemoteRecordApi struct {
}

var remoteRecordService = service.ServiceGroupApp.RemoteServerServiceGroup.SysRemoteRecordService


// CreateSysRemoteRecord 创建远程操作记录
// @Tags SysRemoteRecord
// @Summary 创建远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemoteRecord true "创建远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /remoteRecord/createSysRemoteRecord [post]
func (remoteRecordApi *SysRemoteRecordApi) CreateSysRemoteRecord(c *gin.Context) {
	var remoteRecord remoteServer.SysRemoteRecord
	err := c.ShouldBindJSON(&remoteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    remoteRecord.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "RemoteId":{utils.NotEmpty()},
        "Command":{utils.NotEmpty()},
    }
	if err := utils.Verify(remoteRecord, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := remoteRecordService.CreateSysRemoteRecord(&remoteRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysRemoteRecord 删除远程操作记录
// @Tags SysRemoteRecord
// @Summary 删除远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemoteRecord true "删除远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remoteRecord/deleteSysRemoteRecord [delete]
func (remoteRecordApi *SysRemoteRecordApi) DeleteSysRemoteRecord(c *gin.Context) {
	var remoteRecord remoteServer.SysRemoteRecord
	err := c.ShouldBindJSON(&remoteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    remoteRecord.DeletedBy = utils.GetUserID(c)
	if err := remoteRecordService.DeleteSysRemoteRecord(remoteRecord); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysRemoteRecordByIds 批量删除远程操作记录
// @Tags SysRemoteRecord
// @Summary 批量删除远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /remoteRecord/deleteSysRemoteRecordByIds [delete]
func (remoteRecordApi *SysRemoteRecordApi) DeleteSysRemoteRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := remoteRecordService.DeleteSysRemoteRecordByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysRemoteRecord 更新远程操作记录
// @Tags SysRemoteRecord
// @Summary 更新远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body remoteServer.SysRemoteRecord true "更新远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remoteRecord/updateSysRemoteRecord [put]
func (remoteRecordApi *SysRemoteRecordApi) UpdateSysRemoteRecord(c *gin.Context) {
	var remoteRecord remoteServer.SysRemoteRecord
	err := c.ShouldBindJSON(&remoteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    remoteRecord.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "RemoteId":{utils.NotEmpty()},
          "Command":{utils.NotEmpty()},
      }
    if err := utils.Verify(remoteRecord, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := remoteRecordService.UpdateSysRemoteRecord(remoteRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysRemoteRecord 用id查询远程操作记录
// @Tags SysRemoteRecord
// @Summary 用id查询远程操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query remoteServer.SysRemoteRecord true "用id查询远程操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remoteRecord/findSysRemoteRecord [get]
func (remoteRecordApi *SysRemoteRecordApi) FindSysRemoteRecord(c *gin.Context) {
	var remoteRecord remoteServer.SysRemoteRecord
	err := c.ShouldBindQuery(&remoteRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reremoteRecord, err := remoteRecordService.GetSysRemoteRecord(remoteRecord.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reremoteRecord": reremoteRecord}, c)
	}
}

// GetSysRemoteRecordList 分页获取远程操作记录列表
// @Tags SysRemoteRecord
// @Summary 分页获取远程操作记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query remoteServerReq.SysRemoteRecordSearch true "分页获取远程操作记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remoteRecord/getSysRemoteRecordList [get]
func (remoteRecordApi *SysRemoteRecordApi) GetSysRemoteRecordList(c *gin.Context) {
	var pageInfo remoteServerReq.SysRemoteRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := remoteRecordService.GetSysRemoteRecordInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
