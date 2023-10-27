package business

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type VIPConsumeService struct{}

var VIPConsumeServiceApp = new(VIPConsumeService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaConsume
//@description: 创建客户
//@param: e model.ExaConsume
//@return: err error

func (exa *VIPConsumeService) CreateVIPConsume(e *business.ConsumeRecord) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *VIPConsumeService) CreateVIPConsumeSynchronous(consumeRecord *business.ConsumeRecord, statement *business.VIPStatement, statistics *business.VIPStatistics) (err error) {
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// 创建消费记录
	err = VIPConsumeServiceApp.CreateVIPConsume(consumeRecord)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 流水
	err = VIPOrderServiceApp.CreateVIPStatement(statement)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 统计
	err = VIPOrderServiceApp.CreateVIPStatistics(statistics)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.ExaConsume
//@return: err error

func (exa *VIPConsumeService) DeleteVIPConsume(e business.ConsumeRecord) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaConsume
//@description: 更新客户
//@param: e *model.ExaConsume
//@return: err error

func (exa *VIPConsumeService) UpdateVIPConsume(e *business.ConsumeRecord) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaConsume
//@description: 获取客户信息
//@param: id int
//@return: member model.ExaConsume, err error

func (exa *VIPConsumeService) GetVIPConsume(id int) (member business.ConsumeRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&member).Error
	return
}

func (exa *VIPConsumeService) GetVIPConsumeInfoList(sysUserAuthorityID int, searchInfo request.ConsumeSearchInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	var ConsumeList []business.ConsumeRecord
	cmd := fmt.Sprintf("sys_user_authority_id = %d", sysUserAuthorityID)
	if searchInfo.Telephone >= 1000 {
		cmd += fmt.Sprintf(" and telephone like '%%%d%%'", searchInfo.Telephone)
	}
	if searchInfo.State > 0 {
		cmd += fmt.Sprintf(" and state = %d", searchInfo.State)
	}
	if limit > 0 && offset > 0 {
		cmd += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}
	db := global.GVA_DB.Model(&business.ConsumeRecord{})
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return ConsumeList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("Member").Preload("Member.Combo").Where(cmd).Find(&ConsumeList).Error
	}
	return ConsumeList, total, err
}
