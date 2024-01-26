package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"gorm.io/gorm"
	"strings"
)

type OrderService struct{}

func (o *OrderService) GetProductCartById(userId int, id int) (cartItem wechat.CartItem, err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	db.Where("user_id = ? and id = ?", userId, id).First(&cartItem)
	return cartItem, err
}

func (o *OrderService) GetProductCartByIds(userId int, ids []int) (cartItem []wechat.CartItem, err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	db.Debug().Where("user_id = ? and id in ?", userId, ids).First(&cartItem)
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

func (o *OrderService) GetProductOrderListByStatus(searchInfo request.StateInfo) (list []wechat.Order, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	state := searchInfo.State
	db := global.GVA_DB.Model(&wechat.Order{})

	var cmdSN string
	var cmdReceiver string
	var cmdOrderType string
	var cmdCreateTime string

	if len(searchInfo.OrderSn) > 0 {
		cmdSN += fmt.Sprintf("order_sn = %s", strings.TrimSpace(searchInfo.OrderSn))
	}
	if len(searchInfo.ReceiverKeyword) > 0 {
		cmdReceiver += fmt.Sprintf("receiver_keyword = %s", strings.TrimSpace(searchInfo.ReceiverKeyword))
	}

	if searchInfo.OrderType > 0 {
		cmdOrderType += fmt.Sprintf("order_type = %d", searchInfo.OrderType-100)
	}
	if !searchInfo.CreateTime.IsZero() {
		cmdCreateTime += fmt.Sprintf("create_at = %v", searchInfo.CreateTime)
	}

	cmdSearch := ""
	cmds := [4]string{cmdSN, cmdReceiver, cmdOrderType, cmdCreateTime}
	isFirst := true
	for _, cmd := range cmds {
		if len(cmd) > 0 {
			if isFirst {
				cmdSearch += cmd
				isFirst = false
			} else {
				cmdSearch += " and " + cmd
			}
		}
	}

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
	case 0, 3, 4:
		{
			if len(cmdSearch) > 0 {
				cmdSearch = fmt.Sprintf("%s and status = %d", cmdSearch, state)
			}
			err = db.Where(cmdSearch).Count(&total).Error
			if err != nil {
				return list, total, err
			} else {
				err = db.Where(cmdSearch).Limit(limit).Offset(offset).Preload("OrderItemList").Order("id desc").Find(&list).Error
			}
			return list, total, err
		}
	case 1, 2:
		{
			if len(cmdSearch) > 0 {
				cmdSearch = fmt.Sprintf("%s and status = 1 or status = 2", cmdSearch)
			}
			err = db.Where(cmdSearch).Count(&total).Error
			if err != nil {
				return list, total, err
			} else {
				err = db.Where(cmdSearch).Limit(limit).Offset(offset).Preload("OrderItemList").Order("id desc").Find(&list).Error
			}
			return list, total, err
		}
	default:
		return list, total, err
	}
}

func (o *OrderService) UpdateOrderStatus(e *wechatReq.PaySuccessRequest, status int) (err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Debug().Where("id = ?", e.OrderId).Updates(map[string]interface{}{"pay_type": e.PayType, "status": status}).Error
	return err
}

func (o *OrderService) UpdateOrderStatusById(orderId int, status int) (err error) {
	db := global.GVA_DB.Model(&wechat.Order{})
	err = db.Debug().Where("id = ?", orderId).UpdateColumn("status", status).Error
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
