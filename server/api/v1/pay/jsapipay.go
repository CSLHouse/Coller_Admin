package pay

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/flipped-aurora/gin-vue-admin/server/client/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
	address, err := accountService.GetMemberReceiveAddressById(orderReq.MemberReceiveAddressId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var order wechat.Order
	var payReq payRequest.PrepayRequest
	payReq.Appid = utils.String(orderReq.AppId)
	payReq.Mchid = utils.String(consts.MachID)
	orderDescription := ""
	var goodsDetail []payRequest.GoodsDetail

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

		var goodDetail payRequest.GoodsDetail
		goodDetail.MerchantGoodsId = utils.String(product.ProductSN)
		goodDetail.GoodsName = utils.String(product.Name)
		goodDetail.Quantity = utils.Int64(int64(cartItem.Quantity))
		goodDetail.UnitPrice = utils.Int64(int64(cartItem.Price * 100))

		promotionType := product.PromotionType
		var orderItem wechat.OrderItem
		promotionAmount := float32(0)
		// TODO: 各个优惠计算
		if promotionType == 0 {
			order.PromotionInfo = "无优惠"
		} else if promotionType == 1 {
			order.PromotionInfo = "单品优惠"
		} else if promotionType == 2 {
			order.PromotionInfo = "会员优惠"
		} else if promotionType == 3 {
			order.PromotionInfo = "阶梯价格"
		} else if promotionType == 4 {
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

		orderDescription = fmt.Sprintf("%s x%d ", cartItem.ProductName, cartItem.Quantity)

		goodsDetail = append(goodsDetail, goodDetail)
	}
	n, err := snowflake.NewNode(1)
	if err != nil {
		global.GVA_LOG.Error("创建id失败!", zap.Error(err))
	}
	order.UserId = userId
	order.CouponId = orderReq.CouponId
	order.OrderSn = fmt.Sprintf("%d", n.Generate())
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
		//StoreInfo: &payRequest.StoreInfo{
		//	Address:  utils.String("广东省深圳市南山区科技中一道10000号"),
		//	AreaCode: utils.String("440305"),
		//	Id:       utils.String("0001"),
		//	Name:     utils.String("腾讯大厦分店"),
		//},
	}
	res, _, err := jspaymentService.PrepayWithRequestPayment(payReq)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		fmt.Println("支付失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("--[PrepayWithRequestPayment]-res:", *res)
	err = orderService.UpdateOrderPrepayId(order.ID, *res.PrepayId)
	if err != nil {
		global.GVA_LOG.Error("更新订单预支付交易会话标识失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	var data payRes.GenerateOrderResponse
	data.OrderId = order.ID
	data.Payment = *res
	response.OkWithData(data, c)
}

// PrepayWithRequestPayment Jsapi支付下单，并返回调起支付的请求参数
func (p *PayApi) PrepayWithRequestPayment(c *gin.Context) {

	var payReq payRequest.PrepayRequest
	err := c.ShouldBindJSON(&payReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	payReq = payRequest.PrepayRequest{
		Appid:         utils.String("wxfc872e07778f40eb"),
		Mchid:         utils.String("1658099442"),
		Description:   utils.String("Image形象店-深圳腾大-QQ公仔"),
		OutTradeNo:    utils.String("1217752501201407033233368018"),
		TimeExpire:    utils.String(time.Now().Format(time.RFC3339)),
		Attach:        utils.String("自定义数据说明"),
		NotifyUrl:     utils.String("https://www.weixin.qq.com/wxpay/pay.php"),
		GoodsTag:      utils.String("WXG"),
		SupportFapiao: utils.Bool(false),
		Amount: &payRequest.Amount{
			Currency: utils.String("CNY"),
			Total:    utils.Int64(1),
		},
		Payer: &payRequest.Payer{
			Openid: utils.String("oc9uK66ml1M-nV7Sn3DDlf7IKbPQ"),
		},
		Detail: &payRequest.Detail{
			CostPrice: utils.Int64(608800),
			GoodsDetail: []payRequest.GoodsDetail{payRequest.GoodsDetail{
				GoodsName:        utils.String("iPhoneX 256G"),
				MerchantGoodsId:  utils.String("ABC"),
				Quantity:         utils.Int64(1),
				UnitPrice:        utils.Int64(828800),
				WechatpayGoodsId: utils.String("1001"),
			}},
			InvoiceId: utils.String("wx123"),
		},
		SceneInfo: &payRequest.SceneInfo{
			DeviceId:      utils.String("013467007045764"),
			PayerClientIp: utils.String("14.23.150.211"),
			StoreInfo: &payRequest.StoreInfo{
				Address:  utils.String("广东省深圳市南山区科技中一道10000号"),
				AreaCode: utils.String("440305"),
				Id:       utils.String("0001"),
				Name:     utils.String("腾讯大厦分店"),
			},
		},
		SettleInfo: &payRequest.SettleInfo{
			ProfitSharing: utils.Bool(false),
		},
	}
	res, _, err := jspaymentService.PrepayWithRequestPayment(payReq)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		fmt.Println("支付失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("--[PrepayWithRequestPayment]-res:", res)
	response.OkWithMessage("创建成功", c)
}
