package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
)

type SysRemotesSearch struct {
	remoteServer.SysRemotes
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type SysRemoteExcute struct {
	Id  uint   `json:"id" form:"id"`
	Cmd string `json:"cmd" form:"cmd"`
}
