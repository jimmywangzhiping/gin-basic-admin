package remoteServer

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysRemotesApi
	SysRemoteRecordApi
}

var (
	sysRemoteService       = service.ServiceGroupApp.RemoteServerServiceGroup.SysRemotesService
	sysRemoteRecordService = service.ServiceGroupApp.RemoteServerServiceGroup.SysRemoteRecordService
)
