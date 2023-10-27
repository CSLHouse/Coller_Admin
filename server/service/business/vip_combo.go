package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type VIPComboService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateVIPCombo
//@description: 创建套餐
//@param: e model.VIPCombo
//@return: err error

func (exa *VIPComboService) CreateVIPCombo(e business.VIPCombo) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVIPCombo
//@description: 删除套餐
//@param: e model.VIPCombo
//@return: err error

func (exa *VIPComboService) DeleteVIPCombo(e business.VIPCombo) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateVIPCombo
//@description: 更新套餐
//@param: e *model.VIPCombo
//@return: err error

func (exa *VIPComboService) UpdateVIPCombo(e *business.VIPCombo) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVIPCombo
//@description: 获取套餐信息
//@param: id int
//@return: customer model.VIPCombo, err error

func (exa *VIPComboService) GetVIPCombo(id int) (customer business.VIPCombo, err error) {
	err = global.GVA_DB.Where("combo_id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVIPComboInfoList
//@description: 获取套餐列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *VIPComboService) GetVIPComboInfoList(sysUserAuthorityID int, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.VIPCombo{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []int
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var ComboList []business.VIPCombo
	err = db.Where("authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return ComboList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("authority_id in ?", dataId).Find(&ComboList).Error
	}
	return ComboList, total, err
}

func (exa *VIPComboService) GetAllVIPComboInfoList(sysUserAuthorityID int) (list interface{}, err error) {
	db := global.GVA_DB.Model(&business.VIPCombo{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []int
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var ComboList []business.VIPCombo
	err = db.Where("authority_id in ?", dataId).Find(&ComboList).Error
	return ComboList, err
}
