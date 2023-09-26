package remoteServer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer"
	remoteServerReq "github.com/flipped-aurora/gin-vue-admin/server/model/remoteServer/request"
	"gorm.io/gorm"
)

type SysRemoteRecordService struct {
}

// CreateSysRemoteRecord 创建远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) CreateSysRemoteRecord(remoteRecord *remoteServer.SysRemoteRecord) (err error) {
	err = global.GVA_DB.Create(remoteRecord).Error
	return err
}

// DeleteSysRemoteRecord 删除远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) DeleteSysRemoteRecord(remoteRecord remoteServer.SysRemoteRecord) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&remoteServer.SysRemoteRecord{}).Where("id = ?", remoteRecord.ID).Update("deleted_by", remoteRecord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&remoteRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSysRemoteRecordByIds 批量删除远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) DeleteSysRemoteRecordByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&remoteServer.SysRemoteRecord{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&remoteServer.SysRemoteRecord{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSysRemoteRecord 更新远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) UpdateSysRemoteRecord(remoteRecord remoteServer.SysRemoteRecord) (err error) {
	err = global.GVA_DB.Save(&remoteRecord).Error
	return err
}

// GetSysRemoteRecord 根据id获取远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) GetSysRemoteRecord(id uint) (remoteRecord remoteServer.SysRemoteRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&remoteRecord).Error
	return
}

// GetSysRemoteRecordInfoList 分页获取远程操作记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (remoteRecordService *SysRemoteRecordService) GetSysRemoteRecordInfoList(info remoteServerReq.SysRemoteRecordSearch) (list []remoteServer.SysRemoteRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&remoteServer.SysRemoteRecord{})
	var remoteRecords []remoteServer.SysRemoteRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&remoteRecords).Error
	return remoteRecords, total, err
}
