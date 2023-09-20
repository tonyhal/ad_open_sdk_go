/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// ToolsTaskRaiseOptimizationIdsGetV2ApiService ToolsTaskRaiseOptimizationIdsGetV2Api service
type ToolsTaskRaiseOptimizationIdsGetV2ApiService service

type ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest struct {
	ctx             context.Context
	ApiService      *ToolsTaskRaiseOptimizationIdsGetV2ApiService
	advertiserId    *int64
	platformVersion *ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion
}

func (r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) PlatformVersion(platformVersion ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion) *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest {
	r.platformVersion = &platformVersion
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) Execute() (*ToolsTaskRaiseOptimizationIdsGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsTaskRaiseOptimizationIdsGetGet Method for OpenApi2ToolsTaskRaiseOptimizationIdsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest
*/
func (a *ToolsTaskRaiseOptimizationIdsGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest {
	return &ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsTaskRaiseOptimizationIdsGetV2Response
func (a *ToolsTaskRaiseOptimizationIdsGetV2ApiService) getExecute(r *ApiOpenApi2ToolsTaskRaiseOptimizationIdsGetGetRequest) (*ToolsTaskRaiseOptimizationIdsGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsTaskRaiseOptimizationIdsGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/task_raise/optimization_ids/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.platformVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform_version", r.platformVersion)
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
