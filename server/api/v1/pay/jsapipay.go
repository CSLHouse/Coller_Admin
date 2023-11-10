package pay

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	payRequest "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
	payRes "github.com/flipped-aurora/gin-vue-admin/server/model/pay/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PayApi struct{}

// PrepayWithRequestPayment Jsapi支付下单，并返回调起支付的请求参数
func (p *PayApi) PrepayWithRequestPayment(c *gin.Context, req payRequest.PrepayRequest) (resp *payRes.PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {

	ctx := context.Background()
	prepayResp, result, err := jsapipayService.Prepay(ctx, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(payRes.PrepayWithRequestPaymentResponse)
	resp.PrepayId = prepayResp.PrepayId
	resp.SignType = core.String("RSA")
	resp.Appid = req.Appid
	resp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, nil, fmt.Errorf("generate request for payment err:%s", err.Error())
	}
	resp.NonceStr = core.String(nonce)
	resp.Package = core.String("prepay_id=" + *prepayResp.PrepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *resp.Appid, *resp.TimeStamp, *resp.NonceStr, *resp.Package)
	signatureResult, err := jsapipayService.Client.Sign(ctx, message)
	if err != nil {
		return nil, nil, fmt.Errorf("generate sign for payment err:%s", err.Error())
	}
	resp.PaySign = core.String(signatureResult.Signature)
	return resp, result, nil
}
