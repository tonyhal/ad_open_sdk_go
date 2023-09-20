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

// QianchuanToolsEstimateAudienceV10ApiService QianchuanToolsEstimateAudienceV10Api service
type QianchuanToolsEstimateAudienceV10ApiService service

type ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanToolsEstimateAudienceV10ApiService
	advertiserId   *int64
	marketingGoal  *QianchuanToolsEstimateAudienceV10MarketingGoal
	externalAction *QianchuanToolsEstimateAudienceV10ExternalAction
	productId      *int64
	awemeId        *int64
	audience       *QianchuanToolsEstimateAudienceV10Audience
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) MarketingGoal(marketingGoal QianchuanToolsEstimateAudienceV10MarketingGoal) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) ExternalAction(externalAction QianchuanToolsEstimateAudienceV10ExternalAction) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) ProductId(productId int64) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.productId = &productId
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) Audience(audience QianchuanToolsEstimateAudienceV10Audience) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.audience = &audience
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) Execute() (*QianchuanToolsEstimateAudienceV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsEstimateAudienceGet Method for OpenApiV10QianchuanToolsEstimateAudienceGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest
*/
func (a *QianchuanToolsEstimateAudienceV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest {
	return &ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsEstimateAudienceV10Response
func (a *QianchuanToolsEstimateAudienceV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsEstimateAudienceGetRequest) (*QianchuanToolsEstimateAudienceV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanToolsEstimateAudienceV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/estimate_audience/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.externalAction == nil {
		return localVarReturnValue, nil, ReportError("externalAction is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId)
	}
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.audience != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience", r.audience)
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
