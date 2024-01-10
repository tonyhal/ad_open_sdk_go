/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
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

// PromotionAidGetV30ApiService PromotionAidGetV30Api service
type PromotionAidGetV30ApiService service

type ApiOpenApiV30PromotionAidGetGetRequest struct {
	ctx          context.Context
	ApiService   *PromotionAidGetV30ApiService
	advertiserId *int64
	promotionIds *[]int64
}

// 广告账户id
func (r *ApiOpenApiV30PromotionAidGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30PromotionAidGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 待查询广告ids，一次最多支持10个
func (r *ApiOpenApiV30PromotionAidGetGetRequest) PromotionIds(promotionIds []int64) *ApiOpenApiV30PromotionAidGetGetRequest {
	r.promotionIds = &promotionIds
	return r
}

func (r *ApiOpenApiV30PromotionAidGetGetRequest) Execute() (*PromotionAidGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30PromotionAidGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30PromotionAidGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30PromotionAidGetGetRequest) WithLog(enable bool) *ApiOpenApiV30PromotionAidGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30PromotionAidGetGet Method for OpenApiV30PromotionAidGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30PromotionAidGetGetRequest
*/
func (a *PromotionAidGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30PromotionAidGetGetRequest {
	return &ApiOpenApiV30PromotionAidGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PromotionAidGetV30Response
func (a *PromotionAidGetV30ApiService) getExecute(r *ApiOpenApiV30PromotionAidGetGetRequest) (*PromotionAidGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *PromotionAidGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/promotion/aid/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.promotionIds == nil {
		return localVarReturnValue, nil, ReportError("promotionIds is required and must be specified")
	}
	if len(*r.promotionIds) < 1 {
		return localVarReturnValue, nil, ReportError("promotionIds must have at least 1 elements")
	}
	if len(*r.promotionIds) > 10 {
		return localVarReturnValue, nil, ReportError("promotionIds must have less than 10 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_ids", r.promotionIds)
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
