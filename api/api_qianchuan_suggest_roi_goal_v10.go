/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
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

// QianchuanSuggestRoiGoalV10ApiService QianchuanSuggestRoiGoalV10Api service
type QianchuanSuggestRoiGoalV10ApiService service

type ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest struct {
	ctx                context.Context
	ApiService         *QianchuanSuggestRoiGoalV10ApiService
	advertiserId       *int64
	awemeId            *int64
	marketingScene     *QianchuanSuggestRoiGoalV10MarketingScene
	marketingGoal      *QianchuanSuggestRoiGoalV10MarketingGoal
	productId          *int64
	productNewOpen     *bool
	externalAction     *QianchuanSuggestRoiGoalV10ExternalAction
	campaignScene      *QianchuanSuggestRoiGoalV10CampaignScene
	deepExternalAction *QianchuanSuggestRoiGoalV10DeepExternalAction
	deepBidType        *QianchuanSuggestRoiGoalV10DeepBidType
}

// 广告主ID
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 抖音id
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.awemeId = &awemeId
	return r
}

// 营销场景
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) MarketingScene(marketingScene QianchuanSuggestRoiGoalV10MarketingScene) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.marketingScene = &marketingScene
	return r
}

// MarGoal
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) MarketingGoal(marketingGoal QianchuanSuggestRoiGoalV10MarketingGoal) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

// 商品ID
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.productId = &productId
	return r
}

// 是否开启新品加速
func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) ProductNewOpen(productNewOpen bool) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.productNewOpen = &productNewOpen
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) ExternalAction(externalAction QianchuanSuggestRoiGoalV10ExternalAction) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) CampaignScene(campaignScene QianchuanSuggestRoiGoalV10CampaignScene) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.campaignScene = &campaignScene
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) DeepExternalAction(deepExternalAction QianchuanSuggestRoiGoalV10DeepExternalAction) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.deepExternalAction = &deepExternalAction
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) DeepBidType(deepBidType QianchuanSuggestRoiGoalV10DeepBidType) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.deepBidType = &deepBidType
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) Execute() (*QianchuanSuggestRoiGoalV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanSuggestRoiGoalGet Method for OpenApiV10QianchuanSuggestRoiGoalGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest
*/
func (a *QianchuanSuggestRoiGoalV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest {
	return &ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanSuggestRoiGoalV10Response
func (a *QianchuanSuggestRoiGoalV10ApiService) getExecute(r *ApiOpenApiV10QianchuanSuggestRoiGoalGetRequest) (*QianchuanSuggestRoiGoalV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanSuggestRoiGoalV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/suggest/roi/goal/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.marketingScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_scene", r.marketingScene)
	}
	if r.marketingGoal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
	}
	if r.productNewOpen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_new_open", r.productNewOpen)
	}
	if r.externalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	}
	if r.campaignScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_scene", r.campaignScene)
	}
	if r.deepExternalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_external_action", r.deepExternalAction)
	}
	if r.deepBidType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deep_bid_type", r.deepBidType)
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
