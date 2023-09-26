package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRemoteRecordRouter struct {
}

// InitSysRemoteRecordRouter 初始化 远程操作记录 路由信息
func (s *SysRemoteRecordRouter) InitSysRemoteRecordRouter(Router *gin.RouterGroup) {
	remoteRecordRouter := Router.Group("remoteRecord").Use(middleware.OperationRecord())
	remoteRecordRouterWithoutRecord := Router.Group("remoteRecord")
	var remoteRecordApi = v1.ApiGroupApp.RemoteServerApiGroup.SysRemoteRecordApi
	{
		remoteRecordRouter.POST("createSysRemoteRecord", remoteRecordApi.CreateSysRemoteRecord)   // 新建远程操作记录
		remoteRecordRouter.DELETE("deleteSysRemoteRecord", remoteRecordApi.DeleteSysRemoteRecord) // 删除远程操作记录
		remoteRecordRouter.DELETE("deleteSysRemoteRecordByIds", remoteRecordApi.DeleteSysRemoteRecordByIds) // 批量删除远程操作记录
		remoteRecordRouter.PUT("updateSysRemoteRecord", remoteRecordApi.UpdateSysRemoteRecord)    // 更新远程操作记录
	}
	{
		remoteRecordRouterWithoutRecord.GET("findSysRemoteRecord", remoteRecordApi.FindSysRemoteRecord)        // 根据ID获取远程操作记录
		remoteRecordRouterWithoutRecord.GET("getSysRemoteRecordList", remoteRecordApi.GetSysRemoteRecordList)  // 获取远程操作记录列表
	}
}
