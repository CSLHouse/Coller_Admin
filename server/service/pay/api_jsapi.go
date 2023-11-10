package pay

import (
	"fmt"
	payRequest "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
	payRes "github.com/flipped-aurora/gin-vue-admin/server/model/pay/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

const (
	WechatPayAPIServer       = "https://api.mch.weixin.qq.com"  // 微信支付 API 地址
	WechatPayAPIServerBackup = "https://api2.mch.weixin.qq.com" // 微信支付 API 备份地址
)

// APIResult 微信支付API v3 请求结果
type APIResult struct {
	// 本次请求所使用的 HTTPRequest
	Request *nethttp.Request
	// 本次请求所获得的 HTTPResponse
	Response *nethttp.Response
}

type PayApi struct{}

// CloseOrder 关闭订单
//
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
func (a *PayApi) CloseOrder(req payRequest.CloseOrderRequest) (result *APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in CloseOrderRequest")
	}

	localVarPath := WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}/close"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(utils.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &payRequest.CloseRequest{
		Mchid: req.Mchid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentType := "application/json"

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Prepay JSAPI支付下单
//
// 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易会话标识后再按Native、JSAPI、APP等不同场景生成交易串调起支付。
func (a *PayApi) Prepay(req payRequest.PrepayRequest) (resp *payRes.PrepayResponse, result *APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := WechatPayAPIServer + "/v3/pay/transactions/jsapi"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Setup Content-Type
	localVarHTTPContentType := "application/json"

	// Perform Http Request
	result, err = a.Client.Request(localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(payRes.PrepayResponse)
	err = utils.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrderById 微信支付订单号查询订单
//
// 商户可以通过查询订单接口主动查询订单状态
//func (a *PayApi) QueryOrderById(ctx context.Context, req QueryOrderByIdRequest) (resp *payments.Transaction, result *core.APIResult, err error) {
//	var (
//		localVarHTTPMethod   = nethttp.MethodGet
//		localVarPostBody     interface{}
//		localVarQueryParams  neturl.Values
//		localVarHeaderParams = nethttp.Header{}
//	)
//
//	// Make sure Path Params are properly set
//	if req.TransactionId == nil {
//		return nil, nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderByIdRequest")
//	}
//
//	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/id/{transaction_id}"
//	// Build Path with Path Params
//	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", neturl.PathEscape(core.ParameterToString(*req.TransactionId, "")), -1)
//
//	// Make sure All Required Params are properly set
//	if req.Mchid == nil {
//		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByIdRequest")
//	}
//
//	// Setup Query Params
//	localVarQueryParams = neturl.Values{}
//	localVarQueryParams.Add("mchid", core.ParameterToString(*req.Mchid, ""))
//
//	// Determine the Content-Type Header
//	localVarHTTPContentTypes := []string{}
//	// Setup Content-Type
//	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
//
//	// Perform Http Request
//	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
//	if err != nil {
//		return nil, result, err
//	}
//
//	// Extract payments.Transaction from Http Response
//	resp = new(payments.Transaction)
//	err = core.UnMarshalResponse(result.Response, resp)
//	if err != nil {
//		return nil, result, err
//	}
//	return resp, result, nil
//}
//
//// QueryOrderByOutTradeNo 商户订单号查询订单
////
//// 商户可以通过查询订单接口主动查询订单状态
//func (a *PayApi) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *payments.Transaction, result *core.APIResult, err error) {
//	var (
//		localVarHTTPMethod   = nethttp.MethodGet
//		localVarPostBody     interface{}
//		localVarQueryParams  neturl.Values
//		localVarHeaderParams = nethttp.Header{}
//	)
//
//	// Make sure Path Params are properly set
//	if req.OutTradeNo == nil {
//		return nil, nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in QueryOrderByOutTradeNoRequest")
//	}
//
//	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}"
//	// Build Path with Path Params
//	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)
//
//	// Make sure All Required Params are properly set
//	if req.Mchid == nil {
//		return nil, nil, fmt.Errorf("field `Mchid` is required and must be specified in QueryOrderByOutTradeNoRequest")
//	}
//
//	// Setup Query Params
//	localVarQueryParams = neturl.Values{}
//	localVarQueryParams.Add("mchid", core.ParameterToString(*req.Mchid, ""))
//
//	// Determine the Content-Type Header
//	localVarHTTPContentTypes := []string{}
//	// Setup Content-Type
//	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
//
//	// Perform Http Request
//	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
//	if err != nil {
//		return nil, result, err
//	}
//
//	// Extract payments.Transaction from Http Response
//	resp = new(payments.Transaction)
//	err = core.UnMarshalResponse(result.Response, resp)
//	if err != nil {
//		return nil, result, err
//	}
//	return resp, result, nil
//}
