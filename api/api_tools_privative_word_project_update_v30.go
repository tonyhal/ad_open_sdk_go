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

// ToolsPrivativeWordProjectUpdateV30ApiService ToolsPrivativeWordProjectUpdateV30Api service
type ToolsPrivativeWordProjectUpdateV30ApiService service

type ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest struct {
	ctx                                       context.Context
	ApiService                                *ToolsPrivativeWordProjectUpdateV30ApiService
	toolsPrivativeWordProjectUpdateV30Request *ToolsPrivativeWordProjectUpdateV30Request
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest) ToolsPrivativeWordProjectUpdateV30Request(toolsPrivativeWordProjectUpdateV30Request ToolsPrivativeWordProjectUpdateV30Request) *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest {
	r.toolsPrivativeWordProjectUpdateV30Request = &toolsPrivativeWordProjectUpdateV30Request
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest) Execute() (*ToolsPrivativeWordProjectUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsPrivativeWordProjectUpdatePost Method for OpenApiV30ToolsPrivativeWordProjectUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest
*/
func (a *ToolsPrivativeWordProjectUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest {
	return &ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPrivativeWordProjectUpdateV30Response
func (a *ToolsPrivativeWordProjectUpdateV30ApiService) postExecute(r *ApiOpenApiV30ToolsPrivativeWordProjectUpdatePostRequest) (*ToolsPrivativeWordProjectUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPrivativeWordProjectUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/privative_word/project/update/"

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
	localVarPostBody = r.toolsPrivativeWordProjectUpdateV30Request
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
