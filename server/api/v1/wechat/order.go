package wechat

import (
	"cooller/server/global"
	"cooller/server/model/common/request"
	"cooller/server/model/common/response"
	payRes "cooller/server/model/pay/response"
	"cooller/server/model/wechat"
	wechatReq "cooller/server/model/wechat/request"
	wechatRes "cooller/server/model/wechat/response"
	"cooller/server/utils"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type OrderApi struct{}

// GenerateConfirmOrder 生成确认单信息
func (e *OrderApi) GenerateConfirmOrder(c *gin.Context) {
	var reqIds request.IdsTagReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId <= 0 {
		response.FailWithMessage("Not get userId!", c)
		return
	}
	cartItemList := make([]*wechatRes.CartPromotionItem, 0)
	var totalAmount float32 = 0
	var promotionAmount float32 = 0
	pickUp := 0

	if reqIds.Tag == 1 {
		productCartList, err := orderService.GetProductTmpCartByIds(userId, reqIds.Ids)
		if err != nil {
			global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			response.FailWithMessage("获取购物车物品失败", c)
			return
		}
		for _, cart := range productCartList {
			var cartPromotionItem wechatRes.CartPromotionItem
			cartPromotionItem.ID = cart.ID
			cartPromotionItem.CreatedAt = cart.CreatedAt
			cartPromotionItem.Quantity = cart.Quantity
			cartPromotionItem.ProductPic = cart.SkuStock.Pic
			cartPromotionItem.ProductName = cart.Product.Name
			product, err := wechatService.GetProductByID(cart.ProductId)
			promotionProduct, promotionMessage, reduceAmount := CalculateProductPromotionPrice(product, nil)
			cartPromotionItem.ReduceAmount = reduceAmount
			cartPromotionItem.Price = promotionProduct.Price
			cartPromotionItem.PromotionMessage = promotionMessage
			skuStock, err := wechatService.GetProductSKUStockById(cart.SkuStockId)
			if err != nil {
				global.GVA_LOG.Error("获取SKU库存失败!", zap.Error(err))
			}
			cartPromotionItem.RealStock = skuStock.Stock
			cartItemList = append(cartItemList, &cartPromotionItem)
			totalAmount += cartPromotionItem.Price
			promotionAmount += cartPromotionItem.ReduceAmount
			if promotionProduct.SelfPickup > 0 {
				pickUp = 1
			}
		}
	} else if reqIds.Tag == 2 {
		productCartList, err := orderService.GetProductCartByIds(userId, reqIds.Ids)
		if err != nil {
			global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
			response.FailWithMessage("获取购物车物品失败", c)
			return
		}
		for _, cart := range productCartList {
			var cartPromotionItem wechatRes.CartPromotionItem
			cartPromotionItem.ID = cart.ID
			cartPromotionItem.CreatedAt = cart.CreatedAt
			cartPromotionItem.Quantity = cart.Quantity
			cartPromotionItem.ProductPic = cart.SkuStock.Pic
			cartPromotionItem.ProductName = cart.Product.Name
			product, err := wechatService.GetProductByID(cart.ProductId)
			promotionProduct, promotionMessage, reduceAmount := CalculateProductPromotionPrice(product, nil)
			cartPromotionItem.ReduceAmount = reduceAmount
			cartPromotionItem.Price = promotionProduct.Price
			cartPromotionItem.PromotionMessage = promotionMessage
			skuStock, err := wechatService.GetProductSKUStockById(cart.SkuStockId)
			if err != nil {
				global.GVA_LOG.Error("获取SKU库存失败!", zap.Error(err))
			}
			cartPromotionItem.RealStock = skuStock.Stock
			cartItemList = append(cartItemList, &cartPromotionItem)
			totalAmount += cartPromotionItem.Price
			promotionAmount += cartPromotionItem.ReduceAmount
			if promotionProduct.SelfPickup > 0 {
				pickUp = 1
			}
		}
	}

	address, err := accountService.GetMemberReceiveAddressList(userId)

	var order wechatRes.GenerateOrderResModel
	order.CartPromotionItemList = cartItemList
	order.MemberReceiveAddressList = address
	order.CalcAmount.TotalAmount = totalAmount
	order.CalcAmount.PromotionAmount = promotionAmount
	order.CalcAmount.PayAmount = totalAmount - promotionAmount
	order.PickupType = pickUp
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

	productCartList, err := orderService.GetProductCartByIds(userId, orderReq.Ids)
	if err != nil {
		global.GVA_LOG.Error("获取购物车物品失败!", zap.Error(err))
		response.FailWithMessage("获取购物车物品失败", c)
		return
	}

	for _, cartItem := range productCartList {
		product, err := wechatService.GetProductByID(cartItem.ProductId)
		if err != nil {
			global.GVA_LOG.Error("[GenerateOrder]获取物品失败!", zap.Error(err))
			response.FailWithMessage("[GenerateOrder]获取物品失败", c)
			return
		}
		_, promotionMessage, reduceAmount := CalculateProductPromotionPrice(product, nil)

		var orderItem wechat.OrderItem
		order.PromotionAmount += reduceAmount
		order.PromotionInfo = fmt.Sprintf("满减优惠：%s", promotionMessage)
		orderItem.PromotionAmount = reduceAmount
		orderItem.PromotionName = fmt.Sprintf("满减优惠：%s", promotionMessage)
		// 计算优惠前总金额
		order.TotalAmount += order.TotalAmount + cartItem.Product.Price*float32(cartItem.Quantity)
		// 该商品经过优惠后的实际金额
		realAmount := cartItem.Product.Price*float32(cartItem.Quantity) - reduceAmount

		//orderItem.OrderId = order.ID
		orderItem.ProductId = cartItem.ProductId
		orderItem.ProductSkuId = cartItem.SkuStock.SkuCode
		orderItem.UserId = cartItem.UserId
		orderItem.Quantity = cartItem.Quantity
		orderItem.Price = cartItem.Product.Price
		orderItem.ProductPic = cartItem.Product.Pic
		orderItem.ProductName = cartItem.Product.Name
		orderItem.ProductSubTitle = cartItem.Product.SubTitle
		orderItem.ProductSkuCode = ""
		orderItem.MemberNickname = utils.GetUserName(c)
		orderItem.DeleteStatus = 0
		orderItem.ProductCategoryId = cartItem.Product.ProductCategoryId
		orderItem.ProductBrand = cartItem.Product.BrandName
		orderItem.ProductSn = cartItem.Product.ProductSN
		orderItem.ProductAttr = cartItem.SkuStock.SpData
		orderItem.CouponAmount = 0
		orderItem.IntegrationAmount = 0
		orderItem.RealAmount = realAmount
		orderItem.GiftIntegration = 0
		orderItem.GiftGrowth = 0
		order.OrderItemList = append(order.OrderItemList, &orderItem)
	}
	n, err := snowflake.NewNode(1)
	if err != nil {
		global.GVA_LOG.Error("创建id失败!", zap.Error(err))
	}
	order.UserId = userId
	order.CouponId = orderReq.CouponId
	order.OrderSn = fmt.Sprintf("%d", n.Generate())
	userName := utils.GetNickName(c)
	if len(userName) < 1 {
		userName = utils.GetTelephone(c)
	}
	order.UserName = userName
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
	order.Integration = 100
	order.Growth = 100
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
	order.PaymentTime = time.Now()
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

	var data payRes.GenerateOrderDetailResponse
	data.Order = order
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
	//TODO: 检查是否是秒杀活动商品，如果是，检查是否超出付款时间
	var paySuccess wechatReq.PaySuccessRequest
	err := c.ShouldBindJSON(&paySuccess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrderStatus(&paySuccess, 1)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}
	response.OkWithMessage("支付成功", c)
}

// CancelOrders 取消订单 不付费版
func (e *OrderApi) CancelOrders(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId <= 0 {
		response.FailWithMessage("Not get userId!", c)
		return
	}
	err = orderService.UpdateOrdersStatus(reqIds.Ids, 5)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CloseOrders 关闭订单
func (e *OrderApi) CloseOrders(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId <= 0 {
		response.FailWithMessage("Not get userId!", c)
		return
	}

	err = orderService.UpdateManyOrderStatus(reqIds.Ids, 4)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}

	response.OkWithMessage("关闭成功", c)
}

// DeleteOrders 删除订单
func (e *OrderApi) DeleteOrders(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId <= 0 {
		response.FailWithMessage("Not get userId!", c)
		return
	}
	_, err = orderService.DeleteManyOrder(reqIds.Ids)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (e *OrderApi) UpdateOrderReceiverInfo(c *gin.Context) {
	var address wechatReq.OrderReceiveAddress
	err := c.ShouldBindJSON(&address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrderReceiverInfo(&address)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *OrderApi) UpdateOrderMoneyInfo(c *gin.Context) {
	var info wechatReq.OrderMoneyInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrderMoneyInfo(&info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *OrderApi) UpdateOrderNote(c *gin.Context) {
	var info wechatReq.OrderNoteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrderNoteInfo(&info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *OrderApi) UpdateOrderCompletedStatus(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindJSON(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = orderService.UpdateOrdersStatus(reqIds.Ids, 3)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *OrderApi) GetOrderSetting(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	setting, err := orderService.GetOrderSetting(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("更新订单数据失败!", zap.Error(err))
		response.FailWithMessage("更新订单数据失败", c)
		return
	}

	response.OkWithData(setting, c)
}

func (e *OrderApi) UpdateOrderSetting(c *gin.Context) {
	var home wechat.OrderSetting
	err := c.ShouldBindJSON(&home)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.UpdateOrderSetting(&home)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *OrderApi) GetProductCartList(c *gin.Context) {
	cartList, err := wechatService.GetProductCartList()
	if err != nil {
		global.GVA_LOG.Error("获取商品sku库存失败!", zap.Error(err))
		response.FailWithMessage("获取sku库存失败", c)
		return
	}
	response.OkWithData(cartList, c)
}

// CreateProductCart 创建商品购物车
func (e *OrderApi) CreateProductCart(c *gin.Context) {
	var cart wechat.CartItem
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cart.UserId = utils.GetUserID(c)
	err = wechatService.CreateProductCart(&cart)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateProductCartQuantity 更新商品购物车数量
func (e *OrderApi) UpdateProductCartQuantity(c *gin.Context) {
	var quantityInfo request.QuantityInfo
	err := c.ShouldBindQuery(&quantityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = wechatService.UpdateProductCartQuantity(userId, quantityInfo.ID, quantityInfo.Quantity)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteProductCartById 删除商品购物车
func (e *OrderApi) DeleteProductCartById(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = wechatService.DeleteProductCartById(userId, reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

func (e *OrderApi) DeleteProductCartByIds(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindQuery(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = wechatService.DeleteProductCartByIds(userId, reqIds.Ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// ClearProductCart 清空商品购物车
func (e *OrderApi) ClearProductCart(c *gin.Context) {
	userId := utils.GetUserID(c)
	err := wechatService.ClearProductCartUserId(userId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// CreateProductTmpCart 创建商品购物车
func (e *OrderApi) CreateProductTmpCart(c *gin.Context) {
	var cart wechat.CartTmpItem
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cart.UserId = utils.GetUserID(c)
	err = wechatService.CreateProductTmpCart(&cart)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	data := make(map[string]int)
	data["id"] = cart.ID
	response.OkWithData(data, c)
}
