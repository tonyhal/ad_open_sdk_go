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

// ToolsPrivativeWordPromotionAddV30ApiService ToolsPrivativeWordPromotionAddV30Api service
type ToolsPrivativeWordPromotionAddV30ApiService service

type ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest struct {
	ctx                                      context.Context
	ApiService                               *ToolsPrivativeWordPromotionAddV30ApiService
	toolsPrivativeWordPromotionAddV30Request *ToolsPrivativeWordPromotionAddV30Request
}

func (r *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest) ToolsPrivativeWordPromotionAddV30Request(toolsPrivativeWordPromotionAddV30Request ToolsPrivativeWordPromotionAddV30Request) *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest {
	r.toolsPrivativeWordPromotionAddV30Request = &toolsPrivativeWordPromotionAddV30Request
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest) Execute() (*ToolsPrivativeWordPromotionAddV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsPrivativeWordPromotionAddPost Method for OpenApiV30ToolsPrivativeWordPromotionAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest
*/
func (a *ToolsPrivativeWordPromotionAddV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest {
	return &ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPrivativeWordPromotionAddV30Response
func (a *ToolsPrivativeWordPromotionAddV30ApiService) postExecute(r *ApiOpenApiV30ToolsPrivativeWordPromotionAddPostRequest) (*ToolsPrivativeWordPromotionAddV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPrivativeWordPromotionAddV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/privative_word/promotion/add/"

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
	localVarPostBody = r.toolsPrivativeWordPromotionAddV30Request
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
