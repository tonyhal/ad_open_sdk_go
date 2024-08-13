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

// ToolsPromotionRaiseVersionGetV30ApiService ToolsPromotionRaiseVersionGetV30Api service
type ToolsPromotionRaiseVersionGetV30ApiService service

type ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsPromotionRaiseVersionGetV30ApiService
	advertiserId *int64
	promotionId  *int64
	page         *int64
	pageSize     *int64
}

// 广告主ID
func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告ID
func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) PromotionId(promotionId int64) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	r.promotionId = &promotionId
	return r
}

// 页码，默认值：1
func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) Page(page int64) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，允许值：1-100，默认值：10
func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) Execute() (*ToolsPromotionRaiseVersionGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsPromotionRaiseVersionGetGet Method for OpenApiV30ToolsPromotionRaiseVersionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest
*/
func (a *ToolsPromotionRaiseVersionGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest {
	return &ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPromotionRaiseVersionGetV30Response
func (a *ToolsPromotionRaiseVersionGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsPromotionRaiseVersionGetGetRequest) (*ToolsPromotionRaiseVersionGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPromotionRaiseVersionGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/promotion_raise_version/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.promotionId == nil {
		return localVarReturnValue, nil, ReportError("promotionId is required and must be specified")
	}
	if *r.promotionId < 1 {
		return localVarReturnValue, nil, ReportError("promotionId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_id", r.promotionId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
