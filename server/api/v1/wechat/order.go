package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/client/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	payRequest "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
	payRes "github.com/flipped-aurora/gin-vue-admin/server/model/pay/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	wechatRes "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderApi struct{}

func (e *OrderApi) GenerateConfirmOrder(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	fmt.Println(reqIds)
	cartItemList := make([]wechatRes.CartPromotionItem, 0)
	var totalAmount float32 = 0
	var promotionAmount float32 = 0
	for _, id := range reqIds.Ids {
		cartItem, err := orderService.GetProductCartById(userId, id)
		if err != nil {
			global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			response.FailWithMessage("获取购物车物品失败", c)
			return
		}
		var cartPromotionItem wechatRes.CartPromotionItem
		cartPromotionItem.ID = cartItem.ID
		cartPromotionItem.CreatedAt = cartItem.CreatedAt
		cartPromotionItem.ProductId = cartItem.ProductId
		cartPromotionItem.ProductSkuId = cartItem.ProductSkuId
		cartPromotionItem.UserId = cartItem.UserId
		cartPromotionItem.Quantity = cartItem.Quantity
		cartPromotionItem.Price = cartItem.Price
		cartPromotionItem.ProductPic = cartItem.ProductPic
		cartPromotionItem.ProductName = cartItem.ProductName
		cartPromotionItem.ProductSubTitle = cartItem.ProductSubTitle
		cartPromotionItem.ProductSkuCode = cartItem.ProductSkuCode
		cartPromotionItem.MemberNickname = cartItem.MemberNickname
		cartPromotionItem.DeleteStatus = cartItem.DeleteStatus
		cartPromotionItem.ProductCategoryId = cartItem.ProductCategoryId
		cartPromotionItem.ProductBrand = cartItem.ProductBrand
		cartPromotionItem.ProductSn = cartItem.ProductSn
		cartPromotionItem.ProductAttr = cartItem.ProductAttr

		product, err := wechatService.GetProductByID(cartItem.ProductId)
		promotion := product.PromotionType
		// TODO: 各个优惠计算
		if promotion == 0 {
			cartPromotionItem.PromotionMessage = "无优惠"
		} else if promotion == 1 {
			cartPromotionItem.PromotionMessage = "单品优惠"
		} else if promotion == 2 {
			cartPromotionItem.PromotionMessage = "会员优惠"
		} else if promotion == 3 {
			cartPromotionItem.PromotionMessage = "阶梯价格"
		} else if promotion == 4 {
			fullReductionList, err := wechatService.GetProductFullReductionByProductId(cartItem.ProductId)
			if err != nil {
				global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			}
			reductionStr := ""
			cartPromotionItem.ReduceAmount = 0
			for _, reduction := range fullReductionList {
				if cartPromotionItem.Price >= reduction.FullPrice && cartPromotionItem.ReduceAmount < reduction.ReducePrice {
					cartPromotionItem.ReduceAmount = reduction.ReducePrice
				}
				reductionStr += fmt.Sprintf("满%.2f元，减%.2f元,", reduction.FullPrice, reduction.ReducePrice)
			}
			cartPromotionItem.PromotionMessage = fmt.Sprintf("满减优惠：%s", reductionStr)
		}
		skuStock, err := wechatService.GetProductSKUStockById(cartItem.ProductSkuId)
		if err != nil {
			global.GVA_LOG.Error("获取SKU库存失败!", zap.Error(err))
		}
		cartPromotionItem.RealStock = skuStock.Stock
		cartItemList = append(cartItemList, cartPromotionItem)
		totalAmount += cartPromotionItem.Price
		promotionAmount += cartPromotionItem.ReduceAmount
	}
	address, err := accountService.GetMemberReceiveAddressList(userId)

	var order wechatRes.GenerateOrderResModel
	order.CartPromotionItemList = cartItemList
	order.MemberReceiveAddressList = address
	order.CalcAmount.TotalAmount = totalAmount
	order.CalcAmount.PromotionAmount = promotionAmount
	order.CalcAmount.PayAmount = totalAmount - promotionAmount
	response.OkWithData(order, c)
}

func (e *OrderApi) GenerateOrder(c *gin.Context) {
	var orderReq wechatReq.OrderCreateRequest
	err := c.ShouldBindJSON(&orderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	address, err := accountService.GetMemberReceiveAddressById(orderReq.MemberReceiveAddressId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var order wechat.Order
	order.TotalAmount = 0
	order.PayAmount = 0
	order.PromotionAmount = 0
	for _, id := range orderReq.CartIds {
		cartItem, err := orderService.GetProductCartById(userId, id)
		if err != nil {
			global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			response.FailWithMessage("获取购物车物品失败", c)
			return
		}
		product, err := wechatService.GetProductByID(cartItem.ProductId)
		promotion := product.PromotionType
		var orderItem wechat.OrderItem
		promotionAmount := float32(0)
		// TODO: 各个优惠计算
		if promotion == 0 {
			order.PromotionInfo = "无优惠"
		} else if promotion == 1 {
			order.PromotionInfo = "单品优惠"
		} else if promotion == 2 {
			order.PromotionInfo = "会员优惠"
		} else if promotion == 3 {
			order.PromotionInfo = "阶梯价格"
		} else if promotion == 4 {
			fullReductionList, err := wechatService.GetProductFullReductionByProductId(cartItem.ProductId)
			if err != nil {
				global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			}
			reductionStr := ""
			for _, reduction := range fullReductionList {
				if cartItem.Price >= reduction.FullPrice && promotionAmount < reduction.ReducePrice {
					promotionAmount = reduction.ReducePrice
				}
				reductionStr += fmt.Sprintf("满%.2f元，减%.2f元,", reduction.FullPrice, reduction.ReducePrice)
			}
			order.PromotionAmount += promotionAmount
			order.PromotionInfo = fmt.Sprintf("满减优惠：%s", reductionStr)
			orderItem.PromotionAmount = promotionAmount
			orderItem.PromotionName = fmt.Sprintf("满减优惠：%s", reductionStr)
		}
		// 计算优惠前总金额
		order.TotalAmount += order.TotalAmount + cartItem.Price*float32(cartItem.Quantity)
		// 该商品经过优惠后的实际金额
		realAmount := cartItem.Price*float32(cartItem.Quantity) - promotionAmount

		orderItem.OrderId = order.ID
		orderItem.ProductId = cartItem.ProductId
		orderItem.ProductSkuId = cartItem.ProductSkuId
		orderItem.UserId = cartItem.UserId
		orderItem.Quantity = cartItem.Quantity
		orderItem.Price = cartItem.Price
		orderItem.ProductPic = cartItem.ProductPic
		orderItem.ProductName = cartItem.ProductName
		orderItem.ProductSubTitle = cartItem.ProductSubTitle
		orderItem.ProductSkuCode = cartItem.ProductSkuCode
		orderItem.MemberNickname = cartItem.MemberNickname
		orderItem.DeleteStatus = cartItem.DeleteStatus
		orderItem.ProductCategoryId = cartItem.ProductCategoryId
		orderItem.ProductBrand = cartItem.ProductBrand
		orderItem.ProductSn = cartItem.ProductSn
		orderItem.ProductAttr = cartItem.ProductAttr
		orderItem.CouponAmount = 0
		orderItem.IntegrationAmount = 0
		orderItem.RealAmount = realAmount
		orderItem.GiftIntegration = 0
		orderItem.GiftGrowth = 0
		order.OrderItemList = append(order.OrderItemList, &orderItem)
	}

	order.UserId = userId
	order.CouponId = orderReq.CouponId
	order.OrderSn = ""
	order.Username = utils.GetUserName(c)
	order.FreightAmount = 0
	order.PayAmount += order.TotalAmount - order.PromotionAmount
	order.IntegrationAmount = float32(orderReq.UseIntegration / 1000) // 1000积分抵1元
	order.CouponAmount = 0
	order.DiscountAmount = 0
	order.PayType = orderReq.PayType
	order.SourceType = 1
	order.Status = 0
	order.OrderType = 0
	//order.DeliveryCompany = ""
	//order.DeliverySn = ""
	order.AutoConfirmDay = 7
	order.Integration = 0
	order.Growth = 0
	//order.BillType = 0
	//order.BillHeader = ""
	//order.BillContent = ""
	//order.BillReceiverPhone = ""
	//order.BillReceiverEmail = ""
	order.ReceiverPhone = address.PhoneNumber
	order.ReceiverName = address.Name
	order.ReceiverPostCode = address.PostCode
	order.ReceiverProvince = address.Province
	order.ReceiverCity = address.City
	order.ReceiverRegion = address.Region
	order.ReceiverDetailAddress = address.DetailAddress
	order.Note = orderReq.Note
	order.ConfirmStatus = 0
	order.DeleteStatus = 0
	order.UseIntegration = orderReq.UseIntegration
	order.PaymentTime = ""
	//order.DeliveryTime = ""
	//order.ReceiveTime = ""
	//order.CommentTime = ""
	//order.ModifyTime = ""
	err = orderService.CreateOrder(&order)
	if err != nil {
		global.GVA_LOG.Error("创建订单数据失败!", zap.Error(err))
		response.FailWithMessage("创建订单数据失败", c)
		return
	}

	response.OkWithData(order, c)
}

func (e *OrderApi) GetOrderDetail(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	order, err := orderService.GetProductOrderById(reqId.ID)
	if err != nil {
		response.FailWithMessage("获取订单数据失败", c)
		return
	}

	var queryReq payRequest.QueryOrderByOutTradeNoRequest
	queryReq.Mchid = utils.String(consts.MachID)
	queryReq.OutTradeNo = utils.String(order.OrderSn)
	_, res, _, err := jspaymentService.QueryOrderByOutTradeNo(queryReq, order.PrepayId)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		fmt.Println("支付失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("--[GetOrderDetail]-res:", *res)
	var data payRes.GenerateOrderDetailResponse
	data.Order = order
	data.Payment = *res
	response.OkWithData(data, c)
}

func (e *OrderApi) GetOrderList(c *gin.Context) {
	var stateInfo request.StateInfo
	err := c.ShouldBindQuery(&stateInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(stateInfo, utils.StateInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	orderList, total, err := orderService.GetProductOrderListByStatus(stateInfo)
	if err != nil {
		global.GVA_LOG.Error("获取订单数据失败!", zap.Error(err))
		response.FailWithMessage("获取订单数据失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     orderList,
		Total:    total,
		Page:     stateInfo.Page,
		PageSize: stateInfo.PageSize,
	}, "获取成功", c)
}

func (e *OrderApi) PaySuccess(c *gin.Context) {
	var paySuccess wechatReq.PaySuccessRequest
	err := c.ShouldBindJSON(&paySuccess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrderStatus(&paySuccess)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}
	response.OkWithMessage("支付成功", c)
}

func (e *OrderApi) CancelOrder(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	outTrade, err := orderService.CancelOrder(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}
	err = jspaymentService.CloseOrder(outTrade)
	if err != nil {
		global.GVA_LOG.Error("删除订单数据失败!", zap.Error(err))
		response.FailWithMessage("删除订单数据失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
