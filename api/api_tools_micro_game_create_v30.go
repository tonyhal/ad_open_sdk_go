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

// ToolsMicroGameCreateV30ApiService ToolsMicroGameCreateV30Api service
type ToolsMicroGameCreateV30ApiService service

type ApiOpenApiV30ToolsMicroGameCreatePostRequest struct {
	ctx                            context.Context
	ApiService                     *ToolsMicroGameCreateV30ApiService
	toolsMicroGameCreateV30Request *ToolsMicroGameCreateV30Request
}

func (r *ApiOpenApiV30ToolsMicroGameCreatePostRequest) ToolsMicroGameCreateV30Request(toolsMicroGameCreateV30Request ToolsMicroGameCreateV30Request) *ApiOpenApiV30ToolsMicroGameCreatePostRequest {
	r.toolsMicroGameCreateV30Request = &toolsMicroGameCreateV30Request
	return r
}

func (r *ApiOpenApiV30ToolsMicroGameCreatePostRequest) Execute() (*ToolsMicroGameCreateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsMicroGameCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsMicroGameCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsMicroGameCreatePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsMicroGameCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsMicroGameCreatePost Method for OpenApiV30ToolsMicroGameCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsMicroGameCreatePostRequest
*/
func (a *ToolsMicroGameCreateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsMicroGameCreatePostRequest {
	return &ApiOpenApiV30ToolsMicroGameCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsMicroGameCreateV30Response
func (a *ToolsMicroGameCreateV30ApiService) postExecute(r *ApiOpenApiV30ToolsMicroGameCreatePostRequest) (*ToolsMicroGameCreateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsMicroGameCreateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/micro_game/create/"

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
	localVarPostBody = r.toolsMicroGameCreateV30Request
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
