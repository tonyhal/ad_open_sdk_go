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

// ToolsBpAssetManagementShareV30ApiService ToolsBpAssetManagementShareV30Api service
type ToolsBpAssetManagementShareV30ApiService service

type ApiOpenApiV30ToolsBpAssetManagementSharePostRequest struct {
	ctx                                   context.Context
	ApiService                            *ToolsBpAssetManagementShareV30ApiService
	toolsBpAssetManagementShareV30Request *ToolsBpAssetManagementShareV30Request
}

func (r *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest) ToolsBpAssetManagementShareV30Request(toolsBpAssetManagementShareV30Request ToolsBpAssetManagementShareV30Request) *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest {
	r.toolsBpAssetManagementShareV30Request = &toolsBpAssetManagementShareV30Request
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest) Execute() (*ToolsBpAssetManagementShareV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBpAssetManagementSharePost Method for OpenApiV30ToolsBpAssetManagementSharePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBpAssetManagementSharePostRequest
*/
func (a *ToolsBpAssetManagementShareV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest {
	return &ApiOpenApiV30ToolsBpAssetManagementSharePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBpAssetManagementShareV30Response
func (a *ToolsBpAssetManagementShareV30ApiService) postExecute(r *ApiOpenApiV30ToolsBpAssetManagementSharePostRequest) (*ToolsBpAssetManagementShareV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsBpAssetManagementShareV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/bp_asset_management/share/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.toolsBpAssetManagementShareV30Request
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
