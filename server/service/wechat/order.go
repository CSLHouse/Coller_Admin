package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"gorm.io/gorm"
)

type OrderService struct{}

func (o *OrderService) GetProductCartById(userId int, id int) (cartItem wechat.CartItem, err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	db.Debug().Where("user_id = ? and id = ?", userId, id).First(&cartItem)
	return cartItem, err
}

func (o *OrderService) CreateOrder(e *wechat.Order) (err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Create(&e).Error
	return err
}

func (o *OrderService) CreateOrderItem(e wechat.OrderItem) (err error) {
	db := global.GVA_DB.Model(&wechat.OrderItem{})
	err = db.Create(&e).Error

	return err
}

func (o *OrderService) GetProductOrderById(id int) (order wechat.Order, err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	db.Debug().Where("id = ?", id).Preload("OrderItemList").First(&order)
	return order, err
}

func (o *OrderService) GetProductOrderList(info request.PageInfo) (list []wechat.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("OrderItemList").Find(&list).Error
	}
	return list, total, err
}

func (o *OrderService) GetProductOrderListByStatus(info request.StateInfo) (list []wechat.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	state := info.State
	db := global.GVA_DB.Model(&wechat.Order{})

	switch state {
	case -1:
		{
			err = db.Count(&total).Error
			if err != nil {
				return list, total, err
			} else {
				err = db.Limit(limit).Offset(offset).Preload("OrderItemList").Order("id desc").Find(&list).Error
			}
			return list, total, err
		}
	case 0, 1, 2, 3, 4:
		{
			err = db.Where("status = ?", state).Count(&total).Error
			if err != nil {
				return list, total, err
			} else {
				err = db.Where("status = ?", state).Limit(limit).Offset(offset).Preload("OrderItemList").Order("id desc").Find(&list).Error
			}
			return list, total, err
		}
	default:
		return list, total, err
	}
}

func (o *OrderService) UpdateOrderStatus(e *wechatReq.PaySuccessRequest) (err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Debug().Where("id = ?", e.OrderId).Updates(map[string]interface{}{"pay_type": e.PayType, "status": 1}).Error
	return err
}

func (o *OrderService) UpdateOrderPrepayId(id int, prepayId string) (err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Debug().Where("id = ?", id).UpdateColumn("prepay_id", prepayId).Error
	return err
}

func (o *OrderService) CancelOrder(id int) (outTrade string, err error) {
	var order wechat.Order
	db := global.GVA_DB
	if err := db.Preload("OrderItemList").First(&order, id).Error; err != nil {
		return "", err
	}
	// 执行关联删除操作
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&order).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return order.OrderSn, err
	}
	return order.OrderSn, nil
}
