package pay

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	wechatApi "github.com/flipped-aurora/gin-vue-admin/server/api/v1/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/client/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	payRequest "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
	payRes "github.com/flipped-aurora/gin-vue-admin/server/model/pay/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type PayApi struct{}

func (e *PayApi) GenerateOrder(c *gin.Context) {
	var orderReq wechatReq.OrderCreateRequest
	err := c.ShouldBindJSON(&orderReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)

	var order wechat.Order
	var payReq payRequest.PrepayRequest
	payReq.Appid = utils.String(orderReq.AppId)
	payReq.Mchid = utils.String(consts.MachID)
	orderDescription := ""
	var goodsDetail []payRequest.GoodsDetail

	order.TotalAmount = 0
	order.PayAmount = 0
	order.PromotionAmount = 0

	if orderReq.BuyType > 0 {
		productCartList, err := orderService.GetProductCartByIds(userId, orderReq.Ids)
		if err != nil || len(productCartList) < 1 {
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
			var goodDetail payRequest.GoodsDetail
			goodDetail.MerchantGoodsId = utils.String(product.ProductSN)
			goodDetail.GoodsName = utils.String(product.Name)
			goodDetail.Quantity = utils.Int64(int64(cartItem.Quantity))
			goodDetail.UnitPrice = utils.Int64(int64(cartItem.Price * 100))

			var orderItem wechat.OrderItem
			promotionAmount := float32(0)
			_, promotionMessage, reduceAmount := wechatApi.CalculateProductPromotionPrice(product, nil)

			order.PromotionAmount += reduceAmount
			order.PromotionInfo = promotionMessage
			orderItem.PromotionAmount = reduceAmount
			orderItem.PromotionName = promotionMessage

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

			orderDescription = fmt.Sprintf("%s x%d ", cartItem.ProductName, cartItem.Quantity)

			goodsDetail = append(goodsDetail, goodDetail)
		}
	}

	//address, err := accountService.GetMemberReceiveAddressById(orderReq.MemberReceiveAddressId)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}

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
	order.Integration = 0
	order.Growth = 0
	//order.ReceiverPhone = address.PhoneNumber
	//order.ReceiverName = address.Name
	//order.ReceiverPostCode = address.PostCode
	//order.ReceiverProvince = address.Province
	//order.ReceiverCity = address.City
	//order.ReceiverRegion = address.Region
	//order.ReceiverDetailAddress = address.DetailAddress
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

	payReq.Description = utils.String(orderDescription)
	payReq.OutTradeNo = utils.String(order.OrderSn)
	payReq.TimeExpire = utils.String(time.Now().Format(time.RFC3339))
	payReq.NotifyUrl = utils.String("https://www.weixin.qq.com/wxpay/pay.php")
	payReq.GoodsTag = utils.String("WXG")
	payReq.SettleInfo = &payRequest.SettleInfo{
		ProfitSharing: utils.Bool(false),
	}
	payReq.SupportFapiao = utils.Bool(false)
	payReq.Amount = &payRequest.Amount{
		Currency: utils.String("CNY"),
		Total:    utils.Int64(int64(order.PayAmount * 100)),
	}
	payReq.Payer = &payRequest.Payer{
		Openid: utils.String(orderReq.OpenId),
	}
	payReq.Detail = &payRequest.Detail{
		CostPrice:   utils.Int64(int64(order.TotalAmount * 100)),
		GoodsDetail: goodsDetail,
		//InvoiceId:   utils.String("wx123"),
	}
	payReq.SceneInfo = &payRequest.SceneInfo{
		DeviceId:      utils.String("013467007045764"),
		PayerClientIp: utils.String(orderReq.IP),
		StoreInfo: &payRequest.StoreInfo{
			Address:  utils.String("河南省商丘市中骏雍景台27栋113商铺"),
			AreaCode: utils.String("476100"),
			Id:       utils.String("0001"),
			Name:     utils.String("猪迪克星动乐园"),
		},
	}
	res, _, err := jspaymentService.PrepayWithRequestPayment(payReq)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		fmt.Println("支付失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	//err = orderService.UpdateOrderStatusById(order.ID, 0)
	//if err != nil {
	//	global.GVA_LOG.Error("更新订单支付状态失败!", zap.Error(err))
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	err = orderService.UpdateOrderPrepayId(order.ID, *res.PrepayId)
	if err != nil {
		global.GVA_LOG.Error("更新订单预支付交易会话标识失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var data payRes.GenerateOrderResponse
	data.OrderId = order.ID
	data.Payment = &res
	response.OkWithData(data, c)
}

func (e *PayApi) GetOrderDetail(c *gin.Context) {
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
	var data payRes.GenerateOrderDetailResponse
	data.Order = order
	data.Payment = res
	response.OkWithData(data, c)
}

func (e *PayApi) OrderNotify(c *gin.Context) {
	//request := notify.Request{}
	//c.ShouldBind(&request)
	//mapstructure.Decode(c.Params, &request)
	//if request.EventType == "TRANSACTION.SUCCESS" {
	//	//plaintext, err := wepay.DecryptAES256GCM(
	//	//	aesKey, request.Resource.AssociatedData, request.Resource.Nonce, request.Resource.Ciphertext,
	//	//)
	//	plaintext, err := utils.DecryptAES256GCM(
	//		consts.MchAPIv3Key, request.Resource.AssociatedData, request.Resource.Nonce, request.Resource.Ciphertext,
	//	)
	//	if err != nil {
	//		fmt.Println(err)
	//		zap.S().Error("DecryptAES256GCM err" + err.Error())
	//	}
	//	transaction := payments.Transaction{}
	//	json.Unmarshal([]byte(plaintext), &transaction)
	//	go func() {
	//		// 执行service层代码
	//	}()
	//	tmp := make(map[string]interface{})
	//	tmp["code"] = "SUCCESS"
	//	tmp["message"] = "成功"
	//	tmpJson, _ := json.Marshal(tmp)
	//	c.Writer.Write(tmpJson)
	//} else {
	//	tmp := make(map[string]interface{})
	//	tmp["code"] = "500"
	//	tmp["message"] = "失败"
	//	tmpJson, _ := json.Marshal(tmp)
	//	c.Writer.Write(tmpJson)
	//}
}
