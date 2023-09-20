/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
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

// QianchuanToolsAllowCouponV10ApiService QianchuanToolsAllowCouponV10Api service
type QianchuanToolsAllowCouponV10ApiService service

type ApiOpenApiV10QianchuanToolsAllowCouponGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanToolsAllowCouponV10ApiService
	advertiserId   *int64
	marketingGoal  *QianchuanToolsAllowCouponV10MarketingGoal
	campaignScene  *QianchuanToolsAllowCouponV10CampaignScene
	marketingScene *QianchuanToolsAllowCouponV10MarketingScene
	awemeIds       *[]int64
	productIds     *[]int64
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) MarketingGoal(marketingGoal QianchuanToolsAllowCouponV10MarketingGoal) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) CampaignScene(campaignScene QianchuanToolsAllowCouponV10CampaignScene) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.campaignScene = &campaignScene
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) MarketingScene(marketingScene QianchuanToolsAllowCouponV10MarketingScene) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.marketingScene = &marketingScene
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) AwemeIds(awemeIds []int64) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.awemeIds = &awemeIds
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) ProductIds(productIds []int64) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.productIds = &productIds
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) Execute() (*QianchuanToolsAllowCouponV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsAllowCouponGet Method for OpenApiV10QianchuanToolsAllowCouponGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsAllowCouponGetRequest
*/
func (a *QianchuanToolsAllowCouponV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest {
	return &ApiOpenApiV10QianchuanToolsAllowCouponGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsAllowCouponV10Response
func (a *QianchuanToolsAllowCouponV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsAllowCouponGetRequest) (*QianchuanToolsAllowCouponV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanToolsAllowCouponV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/allow_coupon/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.campaignScene == nil {
		return localVarReturnValue, nil, ReportError("campaignScene is required and must be specified")
	}
	if r.marketingScene == nil {
		return localVarReturnValue, nil, ReportError("marketingScene is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_scene", r.campaignScene)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_scene", r.marketingScene)
	if r.awemeIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_ids", r.awemeIds)
	}
	if r.productIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_ids", r.productIds)
	}
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
