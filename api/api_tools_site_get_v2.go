/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
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

// ToolsSiteGetV2ApiService ToolsSiteGetV2Api service
type ToolsSiteGetV2ApiService service

type ApiOpenApi2ToolsSiteGetGetRequest struct {
	ctx           context.Context
	ApiService    *ToolsSiteGetV2ApiService
	advertiserId  *int64
	xOrangeCaller *string
	page          *int64
	pageSize      *int64
	status        *ToolsSiteGetV2Status
	filtering     *ToolsSiteGetV2Filtering
}

// 广告主id，范围：1 &lt;&#x3D; advertiser_id &lt;&#x3D; MAX_INT64
func (r *ApiOpenApi2ToolsSiteGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsSiteGetGetRequest) XOrangeCaller(xOrangeCaller string) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.xOrangeCaller = &xOrangeCaller
	return r
}

// 页码，默认值: 1，范围：page &gt;&#x3D; 1
func (r *ApiOpenApi2ToolsSiteGetGetRequest) Page(page int64) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.page = &page
	return r
}

// 页面数据量，默认值：10，范围：20～300
func (r *ApiOpenApi2ToolsSiteGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 建站粗粒度状态，详见【附录-枚举值-建站粗粒度状态】
func (r *ApiOpenApi2ToolsSiteGetGetRequest) Status(status ToolsSiteGetV2Status) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.status = &status
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsSiteGetGetRequest) Filtering(filtering ToolsSiteGetV2Filtering) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsSiteGetGetRequest) Execute() (*ToolsSiteGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSiteGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSiteGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSiteGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSiteGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSiteGetGet Method for OpenApi2ToolsSiteGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSiteGetGetRequest
*/
func (a *ToolsSiteGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSiteGetGetRequest {
	return &ApiOpenApi2ToolsSiteGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSiteGetV2Response
func (a *ToolsSiteGetV2ApiService) getExecute(r *ApiOpenApi2ToolsSiteGetGetRequest) (*ToolsSiteGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsSiteGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/site/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
	if r.xOrangeCaller != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Orange-Caller", r.xOrangeCaller)
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
