/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsClueFormDetailV2ApiService ToolsClueFormDetailV2Api service
type ToolsClueFormDetailV2ApiService service

type ApiOpenApi2ToolsClueFormDetailGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueFormDetailV2ApiService
	advertiserId *int64
	instanceId   *int64
}

func (r *ApiOpenApi2ToolsClueFormDetailGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueFormDetailGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsClueFormDetailGetRequest) InstanceId(instanceId int64) *ApiOpenApi2ToolsClueFormDetailGetRequest {
	r.instanceId = &instanceId
	return r
}

func (r *ApiOpenApi2ToolsClueFormDetailGetRequest) Execute() (*ToolsClueFormDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueFormDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueFormDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueFormDetailGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueFormDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueFormDetailGet Method for OpenApi2ToolsClueFormDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueFormDetailGetRequest
*/
func (a *ToolsClueFormDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueFormDetailGetRequest {
	return &ApiOpenApi2ToolsClueFormDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueFormDetailV2Response
func (a *ToolsClueFormDetailV2ApiService) getExecute(r *ApiOpenApi2ToolsClueFormDetailGetRequest) (*ToolsClueFormDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueFormDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/form/detail/"

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
	if r.instanceId == nil {
		return localVarReturnValue, nil, ReportError("instanceId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "instance_id", r.instanceId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
