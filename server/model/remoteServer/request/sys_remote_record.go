package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysRemoteRecordSearch struct{
    remoteServer.SysRemoteRecord
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
