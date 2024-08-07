/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// ToolsClueContactLogListV2ApiService ToolsClueContactLogListV2Api service
type ToolsClueContactLogListV2ApiService service

type ApiOpenApi2ToolsClueContactLogListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueContactLogListV2ApiService
	advertiserId *int64
	clueId       *int64
	filter       *ToolsClueContactLogListV2Filter
}

// 广告主ID
func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 线索ID
func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) ClueId(clueId int64) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	r.clueId = &clueId
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) Filter(filter ToolsClueContactLogListV2Filter) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) Execute() (*ToolsClueContactLogListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueContactLogListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueContactLogListGet Method for OpenApi2ToolsClueContactLogListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueContactLogListGetRequest
*/
func (a *ToolsClueContactLogListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueContactLogListGetRequest {
	return &ApiOpenApi2ToolsClueContactLogListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueContactLogListV2Response
func (a *ToolsClueContactLogListV2ApiService) getExecute(r *ApiOpenApi2ToolsClueContactLogListGetRequest) (*ToolsClueContactLogListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueContactLogListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/contact_log/list/"

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
	if r.clueId == nil {
		return localVarReturnValue, nil, ReportError("clueId is required and must be specified")
	}
	if *r.clueId < 1 {
		return localVarReturnValue, nil, ReportError("clueId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "clue_id", r.clueId)
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
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
