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

// AdUpdateStatusV2ApiService AdUpdateStatusV2Api service
type AdUpdateStatusV2ApiService service

type ApiOpenApi2AdUpdateStatusPostRequest struct {
	ctx                     context.Context
	ApiService              *AdUpdateStatusV2ApiService
	adUpdateStatusV2Request *AdUpdateStatusV2Request
}

func (r *ApiOpenApi2AdUpdateStatusPostRequest) AdUpdateStatusV2Request(adUpdateStatusV2Request AdUpdateStatusV2Request) *ApiOpenApi2AdUpdateStatusPostRequest {
	r.adUpdateStatusV2Request = &adUpdateStatusV2Request
	return r
}

func (r *ApiOpenApi2AdUpdateStatusPostRequest) Execute() (*AdUpdateStatusV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AdUpdateStatusPostRequest) AccessToken(accessToken string) *ApiOpenApi2AdUpdateStatusPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdUpdateStatusPostRequest) WithLog(enable bool) *ApiOpenApi2AdUpdateStatusPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdUpdateStatusPost Method for OpenApi2AdUpdateStatusPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdUpdateStatusPostRequest
*/
func (a *AdUpdateStatusV2ApiService) Post(ctx context.Context) *ApiOpenApi2AdUpdateStatusPostRequest {
	return &ApiOpenApi2AdUpdateStatusPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdUpdateStatusV2Response
func (a *AdUpdateStatusV2ApiService) postExecute(r *ApiOpenApi2AdUpdateStatusPostRequest) (*AdUpdateStatusV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdUpdateStatusV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/ad/update/status/"

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
	localVarPostBody = r.adUpdateStatusV2Request
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
