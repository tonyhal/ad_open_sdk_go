/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// ToolsBlueFlowPackageListV30ApiService ToolsBlueFlowPackageListV30Api service
type ToolsBlueFlowPackageListV30ApiService service

type ApiOpenApiV30ToolsBlueFlowPackageListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsBlueFlowPackageListV30ApiService
	advertiserId *int64
	filtering    *ToolsBlueFlowPackageListV30Filtering
}

func (r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤器
func (r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) Filtering(filtering ToolsBlueFlowPackageListV30Filtering) *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) Execute() (*ToolsBlueFlowPackageListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBlueFlowPackageListGet Method for OpenApiV30ToolsBlueFlowPackageListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBlueFlowPackageListGetRequest
*/
func (a *ToolsBlueFlowPackageListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest {
	return &ApiOpenApiV30ToolsBlueFlowPackageListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBlueFlowPackageListV30Response
func (a *ToolsBlueFlowPackageListV30ApiService) getExecute(r *ApiOpenApiV30ToolsBlueFlowPackageListGetRequest) (*ToolsBlueFlowPackageListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsBlueFlowPackageListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/blue_flow_package/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
