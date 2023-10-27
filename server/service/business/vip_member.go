package business

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
	"strings"
)

type VIPMemberService struct{}

var VIPMemberServiceApp = new(VIPMemberService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaMember
//@description: 创建客户
//@param: e model.ExaMember
//@return: err error

func (exa *VIPMemberService) CreateVIPMember(e *business.VIPMember) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *VIPMemberService) CreateVIPMemberSynchronous(member *business.VIPMember, order *business.VIPOrder, statement *business.VIPStatement, statistics *business.VIPStatistics) (err error) {
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// 创建会员
	err = VIPMemberServiceApp.CreateVIPMember(member)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 订单
	err = VIPOrderServiceApp.CreateVIPOrder(order)
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
//@param: e model.ExaMember
//@return: err error

func (exa *VIPMemberService) DeleteVIPMember(e business.VIPMember) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaMember
//@description: 更新客户
//@param: e *model.ExaMember
//@return: err error

func (exa *VIPMemberService) UpdateVIPMember(e *business.VIPMember) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *VIPMemberService) UpdateVIPMemberSynchronous(member *business.VIPMember, statement *business.VIPStatement, statistics *business.VIPStatistics) (err error) {
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// 更新会员
	err = VIPMemberServiceApp.UpdateVIPMember(member)
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
	//err = global.GVA_DB.Save(e).Error
	//return err
}

// 更新剩余次数remainTimes
func (exa *VIPMemberService) UpdateVIPMemberRemainTimes(e *business.VIPMember, id int, num int) (err error) {
	err = global.GVA_DB.Where("id = ?", id).Update("remainTimes", gorm.Expr("remainTimes - ?", num)).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaMember
//@description: 获取客户信息
//@param: id int
//@return: member model.ExaMember, err error

func (exa *VIPMemberService) GetVIPMember(id int) (member business.VIPMember, err error) {
	err = global.GVA_DB.Where("card_id = ?", id).First(&member).Error
	return
}

func (exa *VIPMemberService) GetVIPMemberWithTelephone(telephone int) (member business.VIPMember, err error) {
	err = global.GVA_DB.Where("telephone = ?", telephone).First(&member).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMemberInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *VIPMemberService) GetVIPMemberInfoList(sysUserAuthorityID int, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.VIPMember{})
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
	var MemberList []business.VIPMember
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return MemberList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("Combo").Where("sys_user_authority_id in ?", dataId).Find(&MemberList).Error
	}
	return MemberList, total, err
}

// 根据卡号、联系方式搜索会员
func (exa *VIPMemberService) SearchVIPMember(sysUserAuthorityID int, searchInfo request.MemberSearchInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var MemberList []business.VIPMember
	cmd := fmt.Sprintf("sys_user_authority_id = %d", sysUserAuthorityID)
	if searchInfo.Telephone >= 1000 {
		cmd += fmt.Sprintf(" and telephone like '%%%d%%'", searchInfo.Telephone)
	}
	if len(searchInfo.MemberName) > 1 {
		cmd += fmt.Sprintf(" and member_name like '%%%s%%'", strings.TrimSpace(searchInfo.MemberName))
	}
	if len(searchInfo.Deadline) >= 10 {
		cmd += fmt.Sprintf(" and deadline > '%s'", strings.TrimSpace(searchInfo.Deadline))
	}
	if searchInfo.State > 0 {
		cmd += fmt.Sprintf(" and state = %d", searchInfo.State)
	}
	if limit > 0 && offset > 0 {
		cmd += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}
	db := global.GVA_DB.Model(&business.VIPMember{})
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return MemberList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Debug().Preload("Combo").Where(cmd).Find(&MemberList).Error
	}
	return MemberList, total, err
}

// 根据卡号、联系方式搜索会员
func (exa *VIPMemberService) SearchVipCard(sysUserAuthorityID int, cardInfo request.CardInfo) (list interface{}, err error) {
	var MemberList []business.VIPMember
	db := global.GVA_DB.Model(&business.VIPMember{})
	cmd := fmt.Sprintf("sys_user_authority_id = %d and telephone like '%%%d%%' or card_id like '%%%d%%'", sysUserAuthorityID, cardInfo.OnlyId, cardInfo.OnlyId)
	err = db.Preload("Combo").Where(cmd).Find(&MemberList).Error
	return MemberList, err
}
