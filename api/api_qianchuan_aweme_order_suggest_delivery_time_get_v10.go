/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService QianchuanAwemeOrderSuggestDeliveryTimeGetV10Api service
type QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService
	advertiserId *int64
	orderId      *int64
	addAmount    *int64
}

// 广告主id
func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 需要追加预算的订单id
func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) OrderId(orderId int64) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	r.orderId = &orderId
	return r
}

// 追加的预算
func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) AddAmount(addAmount int64) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	r.addAmount = &addAmount
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) Execute() (*QianchuanAwemeOrderSuggestDeliveryTimeGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGet Method for OpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest
*/
func (a *QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeOrderSuggestDeliveryTimeGetV10Response
func (a *QianchuanAwemeOrderSuggestDeliveryTimeGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeOrderSuggestDeliveryTimeGetGetRequest) (*QianchuanAwemeOrderSuggestDeliveryTimeGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeOrderSuggestDeliveryTimeGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/order/suggest/delivery_time/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.orderId == nil {
		return localVarReturnValue, nil, ReportError("orderId is required and must be specified")
	}
	if r.addAmount == nil {
		return localVarReturnValue, nil, ReportError("addAmount is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "add_amount", r.addAmount)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
