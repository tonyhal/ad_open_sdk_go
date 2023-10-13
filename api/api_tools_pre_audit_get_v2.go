/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
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

// ToolsPreAuditGetV2ApiService ToolsPreAuditGetV2Api service
type ToolsPreAuditGetV2ApiService service

type ApiOpenApi2ToolsPreAuditGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsPreAuditGetV2ApiService
	advertiserId *int64
	filter       *ToolsPreAuditGetV2Filter
	page         *int32
	pageSize     *int32
}

// 广告主ID
func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) Filter(filter ToolsPreAuditGetV2Filter) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	r.filter = &filter
	return r
}

// 页码，默认值&#x60;1&#x60;，范围：page &gt;&#x3D; 1
func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) Page(page int32) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值 &#x60;20&#x60;，范围：&#x60;1-50&#x60;
func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) Execute() (*ToolsPreAuditGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPreAuditGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPreAuditGetGet Method for OpenApi2ToolsPreAuditGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPreAuditGetGetRequest
*/
func (a *ToolsPreAuditGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPreAuditGetGetRequest {
	return &ApiOpenApi2ToolsPreAuditGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPreAuditGetV2Response
func (a *ToolsPreAuditGetV2ApiService) getExecute(r *ApiOpenApi2ToolsPreAuditGetGetRequest) (*ToolsPreAuditGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsPreAuditGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/pre_audit/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
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
