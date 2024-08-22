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

// ToolsAipThirdSiteUpdateV2ApiService ToolsAipThirdSiteUpdateV2Api service
type ToolsAipThirdSiteUpdateV2ApiService service

type ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest struct {
	ctx                              context.Context
	ApiService                       *ToolsAipThirdSiteUpdateV2ApiService
	toolsAipThirdSiteUpdateV2Request *ToolsAipThirdSiteUpdateV2Request
}

func (r *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest) ToolsAipThirdSiteUpdateV2Request(toolsAipThirdSiteUpdateV2Request ToolsAipThirdSiteUpdateV2Request) *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest {
	r.toolsAipThirdSiteUpdateV2Request = &toolsAipThirdSiteUpdateV2Request
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest) Execute() (*ToolsAipThirdSiteUpdateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest) WithLog(enable bool) *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAipThirdSiteUpdatePost Method for OpenApi2ToolsAipThirdSiteUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest
*/
func (a *ToolsAipThirdSiteUpdateV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest {
	return &ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAipThirdSiteUpdateV2Response
func (a *ToolsAipThirdSiteUpdateV2ApiService) postExecute(r *ApiOpenApi2ToolsAipThirdSiteUpdatePostRequest) (*ToolsAipThirdSiteUpdateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAipThirdSiteUpdateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aip_third_site/update/"

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
	localVarPostBody = r.toolsAipThirdSiteUpdateV2Request
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
