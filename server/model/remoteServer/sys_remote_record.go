// 自动生成模板SysRemoteRecord
package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 远程操作记录 结构体  SysRemoteRecord
type SysRemoteRecord struct {
      global.GVA_MODEL
      RemoteId  *int `json:"remoteId" form:"remoteId" gorm:"column:remoteId;comment:远程服务器Id;"`  //远程服务器Id 
      Command  string `json:"command" form:"command" gorm:"column:cmd;comment:命令;"`  //命令名称 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:执行命令结果(0:失败;1:成功);"`  //执行命令结果 
      Message  string `json:"message" form:"message" gorm:"column:msg;comment:命令返回消息;"`  //消息 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 远程操作记录 SysRemoteRecord自定义表名 sys_remote_record
func (SysRemoteRecord) TableName() string {
  return "sys_remote_record"
}

