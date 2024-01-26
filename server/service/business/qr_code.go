package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
	"time"
)

type QrCodeService struct{}

func (exa *QrCodeService) GetQrCodeInfoList(info request.PageInfo) (list []business.QrCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.QrCode{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, total, err
}

func (exa *QrCodeService) GetQrCodeById(id int) (qrcode business.QrCode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&qrcode).Error
	return qrcode, err
}

func (exa *QrCodeService) CreateQrCode(e business.QrCode) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *QrCodeService) DeleteQrCodeById(id int) (err error) {
	var qrcode business.QrCode
	err = global.GVA_DB.Where("id = ?", id).Delete(&qrcode).Error
	return err
}

func (exa *QrCodeService) UpdateQrCode(e *business.QrCode) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *QrCodeService) UpdateQrCodeCount(id int) (err error) {
	db := global.GVA_DB.Model(&business.QrCode{})
	err = db.Where("id = ?", id).Update("count", gorm.Expr("count+ ?", 1)).Error
	return err
}

func (exa *QrCodeService) GetExpiredQrCodeList() (list []business.QrCode, err error) {
	db := global.GVA_DB.Model(&business.QrCode{})
	offTimes := time.Now().Add(604800 * time.Minute)
	err = db.Where("update_at < ?", offTimes).Find(&list).Error
	return list, err
}

func (exa *QrCodeService) UpdateExpiredQrCodeState() (err error) {
	db := global.GVA_DB.Model(&business.QrCode{})
	offTimes := time.Now().Add(604800 * time.Minute)
	err = db.Where("updated_at < ?", offTimes).Update("is_expired", 1).Error
	return err
}
