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

// ToolsMicroGameUpdateV30ApiService ToolsMicroGameUpdateV30Api service
type ToolsMicroGameUpdateV30ApiService service

type ApiOpenApiV30ToolsMicroGameUpdatePostRequest struct {
	ctx                            context.Context
	ApiService                     *ToolsMicroGameUpdateV30ApiService
	toolsMicroGameUpdateV30Request *ToolsMicroGameUpdateV30Request
}

func (r *ApiOpenApiV30ToolsMicroGameUpdatePostRequest) ToolsMicroGameUpdateV30Request(toolsMicroGameUpdateV30Request ToolsMicroGameUpdateV30Request) *ApiOpenApiV30ToolsMicroGameUpdatePostRequest {
	r.toolsMicroGameUpdateV30Request = &toolsMicroGameUpdateV30Request
	return r
}

func (r *ApiOpenApiV30ToolsMicroGameUpdatePostRequest) Execute() (*ToolsMicroGameUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsMicroGameUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsMicroGameUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsMicroGameUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsMicroGameUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsMicroGameUpdatePost Method for OpenApiV30ToolsMicroGameUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsMicroGameUpdatePostRequest
*/
func (a *ToolsMicroGameUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsMicroGameUpdatePostRequest {
	return &ApiOpenApiV30ToolsMicroGameUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsMicroGameUpdateV30Response
func (a *ToolsMicroGameUpdateV30ApiService) postExecute(r *ApiOpenApiV30ToolsMicroGameUpdatePostRequest) (*ToolsMicroGameUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsMicroGameUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/micro_game/update/"

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
	localVarPostBody = r.toolsMicroGameUpdateV30Request
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
